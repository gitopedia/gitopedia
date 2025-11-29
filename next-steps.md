# Next Steps: Knowledge Base Integration

This document outlines the remaining work needed to complete the source materials integration between Gitopedia and the Knowledge Base.

## Current State

- **Researcher**: Creates Draft PRs with articles in `_incoming/` and sources in `_incoming/sources/`
- **Encyclopaedist**: Copilot agent that organizes articles into `Compendium/`, leaves sources untouched
- **KB Indexer**: Currently indexes articles only; does not handle sources

## Required KB Indexer Enhancements

### 1. Add Sources Table and FTS Index

Update `cmd/indexer/main.go` to create the sources schema:

```sql
CREATE TABLE sources (
    id TEXT PRIMARY KEY,           -- ULID from front matter
    url TEXT NOT NULL,             -- Original source URL
    title TEXT,                    -- Source page title
    related_article TEXT,          -- Slug of the article this source was used for
    summary TEXT,                  -- Brief description
    content TEXT,                  -- Full summarized content
    model TEXT,                    -- LLM model used for summarization
    language TEXT,                 -- Detected language
    created TEXT                   -- ISO date when source was captured
);

CREATE VIRTUAL TABLE source_index USING fts5(
    content, 
    title, 
    summary, 
    id UNINDEXED, 
    url UNINDEXED
);
```

### 2. Ingest Sources from `_incoming/sources/`

Add logic to scan and parse source files. Each source file has this front matter:

```yaml
---
id: 01JXXXXXXXXXXXXXXXXXX
slug: "article-slug--domain-com-1"
title: "Source: Page Title"
url: "https://example.com/page"
type: source
related_article: "article-slug"
created: 2025-11-29
tags: ["Source"]
summary: "Summarized source material for Topic"
model: "deepseek-r1"        # optional
language: "en"              # optional
---

<summarized content>
```

**Implementation steps:**
1. Walk `_incoming/sources/` directory in gitopedia
2. Parse each `.md` file's front matter
3. De-duplicate by URL (skip if URL already exists in database)
4. Insert into `sources` table and `source_index` FTS

### 3. Cleanup `_incoming/sources/` After Ingestion

After successful source ingestion, the KB indexer should:

1. Delete all files in `_incoming/sources/` from gitopedia
2. Commit the deletion with message: `chore: cleanup ingested sources`
3. Push to gitopedia main branch

This requires the KB indexer to have write access to gitopedia (via GitHub App or deploy key).

### 4. Update GitHub Actions Workflow

The KB build workflow needs to:

1. Checkout gitopedia at the triggering commit SHA
2. Run the indexer (articles + sources)
3. Push source cleanup commit back to gitopedia
4. Upload `index.sqlite` to S3

Example workflow addition:

```yaml
- name: Cleanup ingested sources
  env:
    GITHUB_TOKEN: ${{ secrets.GITOPEDIA_WRITE_TOKEN }}
  run: |
    cd gitopedia
    if [ -d "_incoming/sources" ] && [ "$(ls -A _incoming/sources)" ]; then
      rm -rf _incoming/sources/*
      git add _incoming/sources
      git commit -m "chore: cleanup ingested sources" || true
      git push
    fi
```

## Search API Updates (Optional)

If we want to expose source search separately from article search:

1. Add `/search/sources` endpoint to the Search Lambda
2. Query `source_index` FTS table
3. Return source metadata including `url` and `related_article`

## Testing Checklist

- [ ] KB indexer creates `sources` table on fresh database
- [ ] KB indexer parses source front matter correctly
- [ ] KB indexer de-duplicates sources by URL
- [ ] KB indexer inserts sources into FTS index
- [ ] KB indexer deletes `_incoming/sources/` after ingestion
- [ ] KB indexer commits and pushes cleanup to gitopedia
- [ ] Search API can query sources (if implemented)

## Future: Researcher Enhancements

### Internal Article Linking

Once the Knowledge Base is operational, enhance the Researcher to add internal links to other Gitopedia articles:

1. **Query existing articles** before generating new content
   - Fetch article titles/slugs from KB index or gitopedia `Compendium/`
   - Build a lookup of topic → article path

2. **Prompt LLM to include links** during article generation
   - Provide list of related existing articles in the prompt context
   - Instruct LLM to naturally link to related articles using `[Topic Name](/path/to/article)`

3. **Post-generation link enrichment** (alternative approach)
   - After article is generated, scan for entity mentions
   - Match against existing article titles/aliases
   - Insert links where appropriate

**Depends on:** KB indexer being functional to query existing article inventory

### Source Reference Validation

The Researcher already generates a `## References` section with footnotes. Future enhancement:

- Validate that all `[^N]` citations in the article body have corresponding entries in References
- Warn if References section is missing or empty
- Ensure source URLs are valid and accessible

## Dependencies

- GitHub App or deploy key with write access to gitopedia (for cleanup commits)
- Update KB workflow to checkout gitopedia with write permissions
- Ensure S3 bucket permissions for index upload remain intact

