# Knowledgebase Repository – Metadata and Indexing

The Knowledgebase repository is responsible for transforming the unstructured Markdown articles from Gitopedia into structured data and search indices. It serves as the "brains" that enable efficient querying of the content. This includes generating machine-readable metadata (JSON) for each article and building a full-text search index (SQLite database) that the website's search feature can use.

## Source Materials Ingestion

In addition to Markdown articles, the pipeline ingests summarized website sources produced during research. The flow is:

1. **Researcher** stages sources under `_incoming/sources/` in Gitopedia alongside articles
2. **Encyclopaedist** (Copilot Organizer) leaves sources untouched – it only organizes articles
3. **PR merges** with sources still in `_incoming/sources/`
4. **KB Indexer** ingests sources into SQLite (not as files) and cleans up gitopedia

When the KB indexer runs after a PR merge, sources are:

- Identified and de-duplicated against existing sources in the SQLite database
- Parsed and inserted into the `sources` table with full-text search support
- Linked to their related articles via the `related_article` field
- **Deleted from `_incoming/sources/`** in gitopedia after successful ingestion

This keeps gitopedia focused on articles (no source file bloat) while making sources searchable via the Knowledgebase SQLite index.

## Metadata Artifacts (meta.json)

For each article in Gitopedia, the Knowledgebase produces a JSON metadata file (often referred to as a meta artifact). These JSON files extract key information from the articles, making it easy to query and integrate into other systems.

### Metadata Contents

A typical `meta.json` for an article includes:

- **`id`:** The ULID of the article (as defined in the article's front matter).
- **`title`:** The title of the article.
- **`slug`:** The article's filename or a URL-friendly slug (for linking purposes).
- **`created` or `last_updated`:** Timestamps for when the article was created and/or last updated (could be derived from Git history or front matter).
- **`tags`:** List of tags/categories (if the article provides them).
- **`excerpt`:** A short excerpt or summary of the article. This might be the first few sentences of the content or a dedicated summary field from front matter if available.
- **`index`:** Facets object containing:
  - **`topics[]`:** Array of topic identifiers
  - **`people[]`:** Array of person identifiers
  - **`orgs[]`:** Array of organization identifiers
  - **`places[]`:** Array of location identifiers
  - **`date_refs[]`:** Array of date references
  - **`entities[]`:** Array of other named entities
- **`references`** (optional): A list of reference identifiers or URLs that the article cites (if we choose to extract these for quick reference).
- **`content_length`** (optional): Word count or character count of the article, which might be useful for certain analyses.

All `meta.json` files can be stored in a directory (e.g., `meta/`) within the Knowledgebase repository. Each file could be named with the article's ULID or slug. For example: `meta/01HABCD1234XYZ.json` containing the metadata for the article with that ID.

These JSON artifacts allow other tools or future features (like an API endpoint or data analysis jobs) to quickly access article metadata without parsing the full Markdown every time.

### Generation Process

The Knowledgebase build script will:

1. Read each Markdown file from the Gitopedia repository (the script will likely clone or pull the latest Gitopedia content as part of its run).
2. Parse the front matter YAML (to get `id`, `title`, `tags`, etc.) and the content.
3. Extract facets from the content and resolve them to authority IDs from the `authority/` directory in Gitopedia.
4. Possibly compute or extract an excerpt (e.g., first paragraph) for use in summaries.
5. Construct a JSON object for the article and output it to `meta/<id>.json`.

We can use a library like PyYAML (for parsing front matter) and Python's `json` module to output the files. This step ensures that we have structured data ready for indexing and other uses.

## Search Index (SQLite) Build

The Knowledgebase also builds the search index which is a SQLite database enabling full-text search across all articles. We use SQLite's [Full-Text Search (FTS)](https://www.sqlite.org/fts5.html) engine to create a fast search index of the article content.

### Index Schema

A simple approach is to create an FTS5 virtual table that contains the text of all articles (and perhaps titles) and maps back to article IDs. For example, the database might be structured as:

- A table `articles` for basic info (`id`, `title`, maybe `path`).
- An FTS5 virtual table `article_index` that indexes the article text (and possibly the title). The ULID and title can be stored as unindexed fields for retrieval.

**Pseudo-SQL schema:**

```sql
-- Articles table and FTS index
CREATE TABLE articles (
    id TEXT PRIMARY KEY,
    title TEXT,
    path TEXT,
    author TEXT,
    summary TEXT,
    tags TEXT,           -- JSON array
    meta_json TEXT       -- Full metadata as JSON
);

CREATE VIRTUAL TABLE article_index USING fts5(content, title, summary, tags, id UNINDEXED);

-- Sources table and FTS index
CREATE TABLE sources (
    id TEXT PRIMARY KEY,           -- ULID
    url TEXT NOT NULL,             -- Original source URL
    title TEXT,                    -- Source page title
    related_article TEXT,          -- Slug of the article this source was used for
    summary TEXT,                  -- Brief description
    content TEXT,                  -- Full summarized content
    model TEXT,                    -- LLM model used for summarization (optional)
    language TEXT,                 -- Detected language (optional)
    created TEXT                   -- ISO date when source was captured
);

CREATE VIRTUAL TABLE source_index USING fts5(content, title, summary, id UNINDEXED, url UNINDEXED);
```

We insert each article's full Markdown (or rendered text) into the `article_index` and each source's summarized content into the `source_index`. SQLite will then allow querying both tables with full-text search queries.

> **Note:** We may choose to index the Markdown content as raw text or strip out Markdown syntax for the index. It's often beneficial to remove formatting and just index the plain text of articles.

After inserting all articles, we can query the FTS table for matches. The index supports features like relevance ranking and snippet generation. For example, using SQLite's FTS, one can query:

```sql
SELECT id, snippet(article_index) as snippet, rank(matchinfo(article_index)) as score
FROM article_index
WHERE article_index MATCH 'machine learning'
ORDER BY score;
```

This would return matching articles with a snippet of context.

### FTS Version

We will use FTS5 if available, as it provides enhancements over FTS3/FTS4 (better query capabilities, etc.). Most modern SQLite versions include FTS5 by default. (If the deployment environment has an older SQLite, we might fall back to FTS4. For example, an AWS Lambda Python runtime historically had FTS3/4. We will verify the version and adjust accordingly.)

### Building the Index

The Knowledgebase build process (e.g., the Go indexer in `cmd/indexer/`) will:

1. **Index Articles:**
   - Gather all articles from `Compendium/` in gitopedia
   - Parse front matter and content for each article
   - Insert into `articles` table and `article_index` FTS

2. **Ingest Sources:**
   - Scan `_incoming/sources/` in gitopedia for source files
   - Parse each source's front matter (`id`, `url`, `title`, `related_article`, etc.)
   - De-duplicate against existing sources in the database (by URL)
   - Insert new sources into `sources` table and `source_index` FTS

3. **Cleanup Gitopedia:**
   - After successful source ingestion, delete `_incoming/sources/*` from gitopedia
   - Commit and push the cleanup to keep gitopedia focused on articles only

4. **Finalize:**
   - Save and close the database file
   - Upload to S3 for the Search Lambda

Once built, `index.sqlite` will contain all searchable articles and sources. The file is deployed for use by the search service.

### Releasing the Index

After building, the Knowledgebase repository uploads `index.sqlite` directly to the dedicated Knowledgebase index S3 bucket created by the Solus CDK `GitopediaStack`:

- The CI workflow copies `out/index.sqlite` to `s3://<GITOPEDIA_KB_INDEX_BUCKET>/index.sqlite`.
- The Search Lambda (defined in the same CDK stack) is configured with:
  - `INDEX_BUCKET` = this S3 bucket name
  - `INDEX_KEY` = `index.sqlite`

On cold start, the Search Lambda downloads the latest `index.sqlite` from this bucket into `/tmp` and serves queries from it. This keeps GitHub repos focused on source/metadata while S3 holds the heavy index file.

## Repository Structure

The Knowledgebase repository itself can store the `meta.json` files and possibly a copy of the `index.sqlite` (though we might `.gitignore` the SQLite to avoid bloating the repo). Typically:

- `meta/` folder with JSONs (these can be committed to the repo to version the metadata).
- Optionally, a lightweight representation of the index (like an exported CSV of keywords or just rely on SQLite).
- The SQLite index file might be too large to keep in git; we will treat it as a build output, not source. That means it will be generated on demand and released but not versioned in git history.

## Build Automation

### Triggering the Build

The Knowledgebase repository will have a GitHub Action workflow that triggers on an event from Gitopedia:

- We can use a repository dispatch or a workflow trigger so that whenever content changes in Gitopedia (Phase 1 setup), it starts the Knowledgebase build.
- The workflow will checkout the Gitopedia repo (perhaps using the commit SHA that triggered it to ensure consistency), checkout Knowledgebase (if meta files will be committed), run the Python scripts.
- Post-build, handle outputs: if using releases/artifacts, use `actions/upload-artifact` or GitHub API calls to upload `index.sqlite`. If committing meta, push the changes.
- After that, it can either commit the new meta files to Knowledgebase (if we are storing them in git) and push, or skip committing if we choose to keep meta purely as build artifacts. Committing meta has the advantage of versioning the structured data and enabling diff reviews of what changed in an article's metadata.

### Testing the Index

We should include tests or verification steps:

- For example, after building the index, run a few sample queries against it within the workflow to ensure it returns expected results (for known article titles or keywords).
- Also verify that the count of articles in the index matches the count of Markdown files, etc.

## Versioning

Manage versioning of the index:

- Implement a version number or timestamp in the Knowledgebase that corresponds to the content state. For instance, embed the Git commit SHA of Gitopedia (or a hash of content) into a `version.txt` or into the SQLite as a table. This can help the Website or Search Lambda verify it's using the correct index for the content version it has.
- In practice, if our automation is tight, the website will always fetch the latest index after content updates, so content and index should stay in sync. But tracking the version (e.g., Knowledgebase could tag releases like `content-<git-short-sha>`) adds safety.

By maintaining the Knowledgebase as described, we ensure that any updates to the Gitopedia content are quickly translated into an updated search index and metadata store. This separation of concerns means the Gitopedia repo remains focused on content, while Knowledgebase handles the heavy processing required for search and data retrieval. The output of Knowledgebase (the SQLite index and JSON metadata) is then utilized by the Website and other services.

For implementation tasks and TODOs, see the [roadmap.md](../../roadmap.md) file which outlines all development tasks organized by phase.
