# Architecture

This document outlines how the Gitopedia system components interact to form a complete knowledge platform. The architecture has evolved to support multi-phase article generation with semantic search capabilities via a vector database.

## System Overview

```mermaid
flowchart TB
    subgraph User["👤 User"]
        Browser["Web Browser"]
    end

    subgraph Website["🌐 Website (Next.js)"]
        StaticSite["Static Site<br/>(S3 + CloudFront)"]
        SearchAPI["Search Lambda<br/>(Go + SQLite FTS)"]
    end

    subgraph Gitopedia["📚 Gitopedia Repo"]
        Compendium["Compendium/<br/>(Markdown Articles)"]
        IncomingSources["_incoming/sources/<br/>(Pending Sources)"]
        Authority["authority/<br/>(Entity JSON)"]
    end

    subgraph KnowledgeBase["🧠 Knowledge-Base"]
        Indexer["Indexer<br/>(Go CLI)"]
        IngestPipeline["Ingest Pipeline<br/>(CI Workflow)"]
        SQLiteDB["SQLite DB<br/>(FTS5 Index)"]
        Qdrant["Qdrant<br/>(Vector DB)"]
        KBAPI["KB API Server<br/>(Go HTTP)"]
    end

    subgraph Researcher["🔬 Researcher Agent"]
        Agent["Agent<br/>(Go)"]
        LLMClient["LLM Client<br/>(Multi-Model)"]
        SearchClient["Web Search<br/>(DuckDuckGo)"]
        KBClient["KB Client"]
    end

    subgraph Infrastructure["⚙️ Infrastructure"]
        Ollama["Ollama<br/>(LLM Server)"]
        Embeddings["nomic-embed-text<br/>(768d Embeddings)"]
    end

    Browser -->|views| StaticSite
    Browser -->|search query| SearchAPI
    SearchAPI -->|FTS results| Browser

    Compendium -->|markdown| StaticSite
    Compendium -->|articles| Indexer
    IncomingSources -->|sources| IngestPipeline

    Indexer -->|index| SQLiteDB
    IngestPipeline -->|store| SQLiteDB
    IngestPipeline -->|embeddings| Qdrant
    IngestPipeline -->|delete sources| IncomingSources

    SQLiteDB -->|FTS data| SearchAPI
    Qdrant -->|vector search| KBAPI
    SQLiteDB -->|metadata| KBAPI

    Agent -->|new PRs| Gitopedia
    Agent -->|web research| SearchClient
    Agent -->|generate content| LLMClient
    Agent -->|semantic search| KBClient
    KBClient -->|query| KBAPI

    LLMClient -->|inference| Ollama
    IngestPipeline -->|generate| Embeddings
    Embeddings -->|API| Ollama
```

## Component Details

### Gitopedia Repository (Content)

The central content repository containing all encyclopedia articles as Markdown files with YAML frontmatter.

**Directory Structure:**
```
Compendium/
├── Science/
│   ├── Physics/
│   │   ├── quantum-mechanics.md
│   │   └── index.md
│   └── index.md
├── _incoming/
│   └── sources/           # Pending source summaries (cleared by KB ingest)
├── _debug/                # Debug output (thinking traces, raw sources)
└── authority/
    ├── people.json        # Entity authority file
    ├── organizations.json
    └── ...
```

**Article Frontmatter:**
```yaml
---
id: 01KBCVQXJS3QK3JCRGTWBFH2A6  # ULID
title: Quantum Mechanics
author: Gitopedia Researcher
summary: An overview of quantum mechanics...
tags: [physics, quantum, science]
created: 2025-12-03T04:06:38Z    # UTC datetime
model: qwen3:32b                  # LLM used for generation
researcher_version: 0.3.5        # Researcher version
---
```

### Knowledge-Base Repository

Maintains the search index and vector embeddings for semantic search.

```mermaid
flowchart LR
    subgraph Input
        Articles["Articles<br/>(Compendium/*.md)"]
        Sources["Sources<br/>(_incoming/sources/*.md)"]
    end

    subgraph Processing
        Indexer["cmd/indexer"]
        Ingest["cmd/ingest"]
        Embed["Ollama<br/>nomic-embed-text"]
    end

    subgraph Storage
        SQLite["SQLite<br/>(FTS5)"]
        Qdrant["Qdrant<br/>(768d vectors)"]
    end

    subgraph API
        Server["cmd/server<br/>:8081"]
    end

    Articles --> Indexer
    Indexer --> SQLite

    Sources --> Ingest
    Ingest --> Embed
    Embed --> Qdrant
    Ingest --> SQLite

    SQLite --> Server
    Qdrant --> Server
```

**Database Schema:**

```sql
-- Articles table with FTS5
CREATE TABLE articles (
    id TEXT PRIMARY KEY,           -- ULID
    title TEXT NOT NULL,
    path TEXT NOT NULL,            -- Relative path in Compendium
    author TEXT,
    summary TEXT,
    tags TEXT,                     -- JSON array
    meta_json TEXT                 -- Full frontmatter as JSON
);
CREATE VIRTUAL TABLE articles_fts USING fts5(title, summary, content);

-- Sources table
CREATE TABLE sources (
    id TEXT PRIMARY KEY,           -- ULID
    url TEXT UNIQUE,
    title TEXT,
    topic TEXT,                    -- Related article topic
    summary TEXT,
    language TEXT,
    model TEXT,                    -- LLM used for summarization
    created_at TEXT,
    tags TEXT                      -- JSON array
);
CREATE VIRTUAL TABLE sources_fts USING fts5(title, summary, content);
```

**Vector Collections (Qdrant):**
- `sources`: 768-dimension vectors with payload (url, title, topic, summary)
- `articles`: 768-dimension vectors with payload (title, path, category)

**ULID to UUID Conversion:**
Qdrant requires UUID format for point IDs. The knowledge-base automatically converts ULIDs to UUID format:
```
ULID:  01KBCVQXJS3QK3JCRGTWBFH2A6
UUID:  019ad9bb-f659-1de6-3933-10d716f88946
```

### Researcher Agent

The AI-powered agent that creates encyclopedia content automatically.

```mermaid
flowchart TB
    subgraph Triggers
        Issues["GitHub Issues<br/>(research-request label)"]
        Schedule["Scheduled Run"]
    end

    subgraph Research["Research Phase"]
        WebSearch["DuckDuckGo Search"]
        Fetch["Headless Chrome<br/>Page Fetch"]
        Summarize["LLM Summarize<br/>(qwen3:14b)"]
    end

    subgraph Generation["Generation Phase"]
        EntityExtract["Entity Extraction<br/>(qwen3:14b + thinking)"]
        ArticleGen["Article Generation<br/>(qwen3:32b + thinking)"]
        AddRefs["Add References<br/>(qwen3:14b)"]
    end

    subgraph Output
        PR["Draft PR"]
        Sources["Source Summaries<br/>(_incoming/sources/)"]
        Article["Article<br/>(Compendium/)"]
    end

    Issues --> Research
    Schedule --> Research

    WebSearch --> Fetch
    Fetch --> Summarize
    Summarize --> EntityExtract
    EntityExtract --> ArticleGen
    ArticleGen --> AddRefs

    AddRefs --> PR
    Summarize --> Sources
    AddRefs --> Article
```

**Multi-Model Configuration:**
The researcher uses different models for different task complexities:

| Task | Model | Thinking Mode |
|------|-------|---------------|
| Topic Suggestion | qwen3:8b | No |
| JSON Conversion | qwen3:8b | No |
| Source Summarization | qwen3:14b | Yes |
| Entity Extraction | qwen3:14b | Yes |
| Article Generation | qwen3:32b | Yes |
| Reference Addition | qwen3:14b | No |

**Two-Step Article Generation:**
1. **Content Generation**: LLM generates article without references
2. **Citation Addition**: Separate LLM call adds footnote references `[^N]` based only on provided sources

This prevents hallucinated references by ensuring citations only come from actual sources.

### Website

Next.js static site with serverless search.

```mermaid
flowchart LR
    subgraph Build["Build Time"]
        Markdown["Compendium/*.md"]
        GrayMatter["gray-matter<br/>(parse frontmatter)"]
        Remark["remark-html<br/>(render markdown)"]
        NextBuild["next build<br/>(static export)"]
    end

    subgraph Deploy["Deployment"]
        S3["S3 Bucket"]
        CloudFront["CloudFront CDN"]
    end

    subgraph Runtime["Runtime"]
        Lambda["Search Lambda<br/>(Go)"]
        SQLite["SQLite Index"]
    end

    Markdown --> GrayMatter
    GrayMatter --> Remark
    Remark --> NextBuild
    NextBuild --> S3
    S3 --> CloudFront

    Lambda --> SQLite
```

**Displayed Metadata:**
- Article creation datetime (UTC)
- LLM model used for generation
- Researcher version
- "Work in Progress" disclaimer

## Data Flow

### Content Creation Pipeline

```mermaid
sequenceDiagram
    participant User
    participant GitHub as GitHub Issues
    participant Researcher
    participant Ollama
    participant Gitopedia
    participant KB as Knowledge-Base
    participant Website

    User->>GitHub: Create issue (research-request)
    
    loop For each topic
        Researcher->>Ollama: Suggest topics (qwen3:8b)
        Researcher->>Researcher: Web search + fetch
        Researcher->>Ollama: Summarize sources (qwen3:14b)
        Researcher->>Ollama: Extract entities (qwen3:14b)
        Researcher->>Ollama: Generate article (qwen3:32b)
        Researcher->>Ollama: Add references (qwen3:14b)
        Researcher->>Gitopedia: Create PR with article + sources
    end
    
    Gitopedia->>Gitopedia: Merge PR
    Gitopedia->>KB: Dispatch content-updated
    KB->>KB: Run ingest workflow
    KB->>Ollama: Generate embeddings
    KB->>KB: Store in SQLite + Qdrant
    KB->>Gitopedia: Delete _incoming/sources/
    KB->>Website: Trigger rebuild
    Website->>Website: Build static site
    Website->>User: Updated content live
```

### Source Ingestion Pipeline

```mermaid
sequenceDiagram
    participant Gitopedia
    participant KBWorkflow as KB Ingest Workflow
    participant Ollama
    participant SQLite
    participant Qdrant

    Gitopedia->>KBWorkflow: repository_dispatch (content-updated)
    KBWorkflow->>KBWorkflow: Checkout gitopedia
    KBWorkflow->>KBWorkflow: Check for sources in _incoming/sources/
    
    alt Has sources
        KBWorkflow->>Ollama: Pull nomic-embed-text
        loop For each source
            KBWorkflow->>KBWorkflow: Parse frontmatter
            KBWorkflow->>SQLite: Upsert source metadata
            KBWorkflow->>Ollama: Generate embedding
            KBWorkflow->>Qdrant: Upsert vector (ULID→UUID)
            KBWorkflow->>KBWorkflow: Delete source file
        end
        KBWorkflow->>Gitopedia: Commit + push deletions
    end
```

## Versioning

The researcher agent uses semantic versioning stored in `researcher/VERSION`:

```
0.3.5
```

**Version Updates:**
- Git `post-commit` hook auto-increments patch version
- Version appears in article frontmatter (`researcher_version`)
- Version displayed on website footer and article list

**Changelog:**
Changes are documented in `researcher/CHANGELOG.md` following Keep a Changelog format.

## Infrastructure

### Local Development (Docker Compose)

```yaml
services:
  ollama:
    image: ollama/ollama:latest
    ports: ["11434:11434"]
    # GPU support enabled
    
  qdrant:
    image: qdrant/qdrant:latest
    ports: ["6333:6333", "6334:6334"]
    
  researcher:
    build: .
    depends_on: [ollama, qdrant]
```

### CI/CD Workflows

| Repository | Workflow | Trigger | Action |
|------------|----------|---------|--------|
| gitopedia | content-dispatch | push to main | Dispatch to knowledge-base |
| knowledge-base | ingest | repository_dispatch | Ingest sources, generate embeddings |
| knowledge-base | build-index | push to main | Rebuild article index |
| website | site-build | push to main | Build and deploy static site |

## Summary

The Gitopedia architecture enables fully automated encyclopedia content creation:

1. **Research**: AI agent searches web, summarizes sources
2. **Generate**: Multi-model LLM pipeline creates articles with proper citations
3. **Store**: Sources ingested into knowledge-base with vector embeddings
4. **Index**: Full-text and semantic search indexes maintained
5. **Publish**: Static site automatically rebuilt and deployed

All content flows through Git, ensuring version control and auditability. The use of ULIDs provides stable cross-repository references, while the vector database enables future semantic search and RAG capabilities.
