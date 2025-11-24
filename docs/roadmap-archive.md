# Roadmap Archive: Pre-MVP and Phase 1

This document archives the completed phases of the Gitopedia development roadmap.

## Phase 0: Pre-MVP Setup (Project Bootstrapping)

**Goal:** Establish the foundational structure for all components of the system.

**Status:** COMPLETED

### Key Tasks

- **DONE:** Create the Gitopedia repository structure. Set up directories for articles (`Compendium/` folder with category subdirectories) and add a sample Markdown article. Set up the `authority/` directory structure for controlled vocabularies (people, orgs, places, topics). Include a README explaining the article format and how the Researcher agent creates articles. Decide on an article format (e.g. use a [YAML front-matter](https://docs.github.com/en/contributing/writing-for-github-docs/using-yaml-frontmatter) with fields like `id`, `title`, `slug`, etc., where `id` will be a ULID for each article).
- **DONE:** Create documentation in README.md of Gitopedia that summarizes the article format (including facets and authority lists), and how the Researcher agent creates articles via automated PRs. Note that all articles are created by the Researcher agent, not by human contributors.
- **DONE:** Provide ULID generation utilities for the Researcher agent:
  - Created `scripts/generate_ulid.py` helper script for ULID generation.
  - Set up validation in CI and pre-commit hooks to verify ULID format in article front matter.
- **DONE:** Provision self-hosted infrastructure for Researcher dependencies:
  - **DONE:** Provision a stable runtime for the Researcher (VM/container) with private network access to SearXNG and DeepSeek endpoints.
  - **DONE:** Deploy a self-hosted DeepSeek model behind an OpenAI-compatible API (e.g., vLLM gateway). Expose an internal `OPENAI_BASE_URL` and set default `DEEPSEEK_MODEL`.
  - **SKIPPED:** Deploy a self-hosted SearXNG instance (Docker). We are using DuckDuckGo directly for MVP to avoid complexity; rate limits are managed internally.
  - **Note:** Infrastructure setup is documented in `researcher/infra/README.md` using Docker Compose.
- **DONE:** Initialize the Knowledgebase repository. Defined schema for metadata, created `scripts/build_index.py` (Python) to parse Markdown files and generate SQLite FTS5 index. Added dependencies in `requirements.txt`.
- **DONE:** Initialize the Researcher agent repository. Set up Go project structure (`go.mod`) with README outlining the agent's workflow.
- **DONE:** Initialize the Website repository with Next.js app. Configured `next.config.js` for static export. Created static page generation that reads from Gitopedia content. Set up AWS S3 and CloudFront deployment via CDK.
- **DONE:** Set up basic CI/CD workflows for cross-repo integration. Implemented `repository_dispatch` triggers from Gitopedia to Knowledgebase. Knowledgebase workflow builds index and uploads to S3. Website workflow builds and deploys to S3/CloudFront.
- **DONE:** Decide on the unique ID strategy for articles. Using ULIDs for each article's stable identifier. Implemented `scripts/generate_ulid.py` and validation in CI/pre-commit hooks.
- **DONE:** Document the development setup for each component. Created comprehensive documentation in `docs/` directory and README files in each repository.

## Phase 1: MVP (Minimum Viable Product)

**Goal:** Achieve a working end-to-end system where a new article can be created, processed, and displayed on the website with a functional search.

**Status:** COMPLETED

### Key Tasks

- **DONE:** Implement content ingestion in Gitopedia. Define contribution guidelines for articles (e.g. how sources should be cited in Markdown). Ensure each article has a ULID in its front-matter (from Phase 0 plan). Create a couple of sample articles to use for testing the pipeline (these can be brief).
- **DONE:** Add Draft PR staging with `_incoming/` and Copilot Organizer. Researcher opens Draft PRs with articles and summarized sources staged under `_incoming/`. A Copilot Organizer GitHub Action/agent refactors the branch (moves articles to `Compendium/`, validates front matter and facets, prepares summarized sources for Knowledgebase), pushes commits, and marks the PR ready.
- **SKIPPED:** Implement CI pipeline for Researcher container. We are running the Researcher locally as a Go binary for the MVP. Containerization and CI publishing are deferred.
- **SKIPPED:** Enable CI GitHub SSH automation. The Researcher uses the GitHub API (via `google/go-github`) for all operations (branches, files, PRs), so SSH keys are not required.
- **DONE:** Integrate self-hosted services and SDKs in the Researcher:
  - **SKIPPED:** Add a SearXNG client. Using `duckduckgo-search` (direct) for simplicity.
  - **DONE:** Use the Go OpenAI SDK (`sashabaranov/go-openai`) for both OpenAI cloud and self-hosted DeepSeek via `OPENAI_BASE_URL`.
  - **DONE:** Add configuration via `.env` file for `OPENAI_BASE_URL`, `OPENAI_API_KEY`, `OPENAI_MODEL`.
- **DONE:** Implement the Knowledgebase indexing process. Develop a Python script (or use an existing tool) to parse all Markdown files in Gitopedia. For each article, generate its metadata (`meta.json`), including fields like `id` (ULID), `title`, perhaps an excerpt or list of references. Aggregate these into a searchable SQLite database (using SQLite FTS5 or FTS4 for full-text search on article content). Verify that running this script locally produces an `index.sqlite` with the sample articles.
- **DONE:** Set up an automated workflow for Knowledgebase. Configure a GitHub Action that triggers on new content in Gitopedia (e.g. push events on main) to run the indexing script. For MVP, this action can simply run tests or produce an artifact (it might not yet deploy the SQLite file, but should at least confirm the process). Ensure the workflow can ingest new summarized sources organized by the Copilot phase.
- **DONE:** Implement cross-repo triggers using GitHub Actions:
  - **DONE - In Gitopedia**: Workflow triggers Knowledgebase via `repository_dispatch` on push to main, using a GitHub App (`gitopedia-bot`) token generated automatically in the workflow (no manual token rotation required).
  - **DONE - In Knowledgebase**: Workflow builds index and uploads to S3 with OIDC authentication, listens for `repository_dispatch` events (`content-updated`), and dispatches a `rebuild-site` event to the Website repo after successful index upload.
  - **DONE - In Website**: Workflow builds and deploys static site to S3/CloudFront with OIDC authentication, and listens for `repository_dispatch` events (`rebuild-site`) to rebuild the site when Knowledgebase completes.
  - **DONE**: AWS infrastructure configured with IAM OIDC roles for secure CI/CD access.
- **DONE:** Develop the Search Lambda (Search API) in the Website (or a dedicated directory). Implemented a Python Lambda in `website/search-api/app.py` that downloads `index.sqlite` from the Knowledgebase S3 bucket, executes SQLite FTS queries, and returns JSON results. The Solus CDK `GitopediaStack` defines the Lambda, grants it read access to the KB index bucket, and exposes it via an API Gateway `GET /search` endpoint with CORS for `gitopedia.org`. Website CI builds a deployable Lambda zip artifact.
- **DONE:** Integrate search on the Website frontend. Added a basic search UI on the homepage (`website/pages/index.jsx`) that calls the Search API endpoint (configured via `NEXT_PUBLIC_SEARCH_API_URL`), displays results with titles and snippets, and links to article pages. For MVP, this is a simple form and result list.
- **DONE:** Implement static page generation for Website. Created Next.js pages that:
  - Checkout Gitopedia repository during CI build to access Markdown files.
  - Use `lib/content.js` to parse Markdown with `gray-matter` and `remark` for HTML conversion.
  - Generate static pages via `getStaticPaths` and `getStaticProps` for all articles.
  - Created homepage (`pages/index.jsx`) listing all articles and article pages (`pages/[...slug].jsx`) for individual articles.
  - Handle front matter serialization (Date objects to ISO strings) for JSON compatibility.
  - Deploy to S3 and CloudFront via GitHub Actions with OIDC authentication.
- **DONE:** Implement the index generation script for category `index.md` files and integrate it into GitHub Actions. Created `scripts/build_category_indexes.py` that:
  - Scans `Compendium/` directory structure for articles.
  - Extracts titles from front matter (with fallback to first heading).
  - Generates/updates `index.md` in each category directory.
  - Integrated into CI workflow with auto-commit on PRs and pushes to main.
  - Added to pre-commit hooks for local validation.
- **DONE:** Set up AWS infrastructure via CDK:
  - Created `GitopediaStack` with S3 buckets for website and knowledgebase index.
  - Configured CloudFront distribution for website with custom domain support.
  - Set up ECR repository for Researcher container images.
  - Created IAM OIDC roles for GitHub Actions (Website, Knowledgebase, Researcher CI roles).
  - Configured SSL certificate stack for `gitopedia.org` domain.
  - Deployed infrastructure and configured domain DNS (Route53) for `gitopedia.org`.
- **DONE:** Connect the Researcher agent minimal functionality. Implemented a Go-based agent in `gitopedia/researcher` that:
  - Fetches "research request" issues via GitHub API.
  - Performs searches using DuckDuckGo (simulated via API or scraping).
  - Calls an LLM (OpenAI/DeepSeek) to draft Markdown content with citations.
  - Creates a branch, stages content in `_incoming/`, and opens a Draft PR.
  - Validated with unit tests and a local `Makefile` workflow.
- **DONE:** Handle Git operations. Implemented `internal/github` package using `google/go-github` to manage branches, files, and PRs via API calls (avoiding local git shell-outs for simpler containerization).
- **DONE:** Integrate search. Implemented `internal/search` package using DuckDuckGo.
- **DONE:** Integrate LLM. Implemented `internal/llm` package using `sashabaranov/go-openai` SDK, configurable for self-hosted DeepSeek via `OPENAI_BASE_URL`.

