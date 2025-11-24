# Roadmap

This roadmap outlines the development of Gitopedia in stages from a pre-MVP setup to a production-ready system. Each phase is defined with specific objectives and deliverables. The phases are incremental: each builds upon the previous, gradually adding functionality and stability. TODO items are listed under each phase as actionable steps required to achieve the phase's goals.

## Phase 0: Pre-MVP Setup (Project Bootstrapping)

**Goal:** Establish the foundational structure for all components of the system. In this phase, we set up the repositories and basic pipelines, but the end-to-end functionality may be stubbed or manual. The focus is on scaffolding the project so development can proceed in an organized way.

**Deliverables:** All four repositories initialized with minimal code/configuration, basic documentation (this `docs/` directory), and placeholder processes for content flow.

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

**Goal:** Achieve a working end-to-end system where a new article can be created, processed, and displayed on the website with a functional search. The MVP doesn't need to be polished or scalable, but it should demonstrate the core loop: content -> index -> website. The Researcher should be able to add at least one article automatically (even if using a simple approach or minimal AI logic).

**Deliverables:**
- A simple end-to-end demo of Gitopedia: for example, a "Hello World" article created via a research issue, appearing on the live site, searchable by its title or content.
- Basic implementations for content parsing, indexing, static site generation, and a search query endpoint.

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
- **DONE:** Test the end-to-end flow for MVP:
  - **DONE:** Create a dummy "New Article Request" issue on Gitopedia (with a simple prompt).
  - **DONE:** Run the Researcher agent to process this issue (generating an article and PR).
  - **DONE:** Merge the PR into Gitopedia.
  - **DONE:** Verify the Knowledgebase action runs and updates the index.
  - **DONE:** Manually deploy or update the Search API with the new index (for MVP, this might be a manual step if full automation isn't done).
  - **DONE:** Rebuild the Website and deploy the updated static site.
  - **DONE:** Open the website and confirm the new article is visible and searchable.
- **DONE:** Fix any issues uncovered in testing. Addressed minor issues like regex parsing in search UI and incorrect roadmap status updates.

## Phase 2: Enhanced Capabilities and Beta Release

**Goal:** Build on the MVP to add important features, improve quality, and increase the system's knowledge base. In this phase, the focus is on refining functionality, increasing content, and ensuring the system can handle more use cases. We consider this a "Beta" stage where the system is usable by early adopters, but not yet fully polished.

**Deliverables:**
- Several dozen high-quality articles in Gitopedia, many generated via the Researcher agent.
- Improved search features (accuracy, performance, possibly filtering or categorization).
- More automation and robustness in the pipelines (fewer manual steps).
- Internal evaluations or testing by a small set of users.

### Key Improvements and Tasks

- **TODO:** Enrich article metadata and structure. Define additional front-matter fields as needed (e.g. `author`, `tags`/`categories`, `summary`). Ensure the Knowledgebase captures these in `meta.json`. This will allow enhancements like category pages or filtering search results by tag.
- **TODO:** Implement facet validators in CI to ensure: facet values resolve to authority IDs, required facets are present (≥1 topic, at least one of people/orgs/places/date_refs), and authority ID format is correct. Implement canonical title enforcement in CI to prevent duplicate titles.
- **TODO:** Implement authority list management (creation, merging, alias resolution) in the Researcher agent. The agent should check authority lists when creating articles, create new authority entries when needed (subject to review), and suggest merges if similar entities already exist.
- **TODO:** Create redirect mechanism for renamed or merged articles to maintain link integrity. This can be stored in a `redirects.json` file or handled by the Knowledgebase.
- **TODO:** Improve the Researcher agent's output quality. Integrate better prompt templates and multi-step research (e.g. have the agent gather multiple sources before writing). Ensure that the agent includes citations to external sources in the Markdown (perhaps as footnotes or reference list) to maintain credibility. Implement error handling for the agent (e.g. if the API call fails or the content is incomplete, the agent should log an issue or retry).
- **TODO:** Expand content coverage. Come up with a list of target articles to add (possibly through new GitHub issues that the Researcher will pick up). This will both test the system and start populating Gitopedia with real content. For example, create 10–20 research issues on diverse topics and allow the Researcher to process them. Review the PRs and refine the Researcher's prompts or code based on any deficiencies in the outputs.
- **TODO:** Enhance the search functionality. If not already using it, consider upgrading to SQLite FTS5 for better relevance ranking. Implement search result ranking tuning (e.g. boosting results where query appears in title). Add a basic caching mechanism for frequent searches (if needed) or pagination for results if the content grows. Also, consider adding the ability to search within article metadata (like tags or titles only).
- **TODO:** Improve the Website user experience. Organize articles on the site (e.g. add a homepage that lists categories or recent additions, create a navigation menu). Possibly generate an "All articles" index page directly from Knowledgebase metadata (which could include an automatically generated Table of Contents). Ensure that each article page has a consistent layout (maybe a template with a sidebar for related articles or a section showing the article's metadata like last updated date and references).
- **TODO:** Automate deployment steps. By this phase, aim to reduce manual intervention:
  - The Knowledgebase workflow should automatically publish the updated `index.sqlite` (for example, by uploading it to an S3 bucket or updating a GitHub release asset) whenever content changes.
  - The Website should have a continuous deployment pipeline (e.g. GitHub Action or CI job) that triggers on changes to content or index. This may involve the website repo pulling the latest Gitopedia content and knowledgebase index, rebuilding, and syncing to S3/CloudFront.
  - The Search Lambda deployment should also be automated (e.g. use AWS SAM or Serverless framework in CI to deploy the new index and code after Knowledgebase updates).
- **TODO:** Implement version and sync checks:
  - Consider adding the Gitopedia commit hash into Knowledgebase's SQLite (e.g., as a value in a metadata table). The search API could expose it (for debugging).
  - Add the content commit hash into the website build (maybe as a meta tag or in the footer of pages).
  - Write a small script or use existing actions to tag multiple repos. (Since multi-repo tagging isn't atomic, we do it manually in each, but we can automate it via API calls.)
  - Document the release tagging conventions in `releases.md` (for example, using semantic versioning for code changes and date-based for content snapshots, or a unified scheme).
- **TODO:** Conduct testing and gather feedback. Run integration tests for scenarios like "new article added -> searchable on site" to catch any regressions. If possible, have a few users test the system and report issues or suggestions. Use this feedback to fix bugs (e.g. broken links, formatting issues) and improve documentation (e.g. clearer instructions for contributors).
- **TODO:** Performance tuning for Beta. With more content, monitor the size of the `index.sqlite` and search query performance. Ensure the Lambda has enough memory/CPU to handle queries quickly. If search latency grows, consider optimizations like indexing only relevant fields, or in worst case, plan for an external search service in the future (not needed yet, but keep in mind).
- **TODO:** Security review. By Phase 2, the system might be accessible to more users. Review the setup for any security concerns: ensure the search API is not vulnerable to injection (validate queries), the S3 bucket with the site is properly configured for public read (and not listable), and that secrets (API keys, tokens) are secure in CI. Also, make sure the Researcher agent doesn't accidentally leak secrets or get stuck on malicious input from issues.
- **TODO:** Update continuous integration pipelines to include version information in build outputs. For instance, the website could display the software version number somewhere (in a footer or about page) for debugging purposes. This helps with troubleshooting and ensures users can identify which version of the system they are interacting with.

## Phase 3: Production-Ready Release

**Goal:** Finalize the system for production use. This phase emphasizes hardening, scalability, and maintainability. The aim is to have a reliable, automated platform that can continuously grow with minimal manual oversight, suitable for real users.

**Deliverables:**
- Version 1.0 of Gitopedia – a fully functional AI-curated encyclopedia online.
- Stable and monitored infrastructure for the website and search.
- Documentation for future contributors and maintainers.
- A backlog of future improvements (post-1.0) identified.

### Key Tasks

- **TODO:** Refine and enforce content standards. Develop guidelines or automated checks for article quality (e.g. a CI step that lints Markdown, checks for broken links or missing metadata). Possibly integrate a content review process: even if AI generates articles, having a human or a validation step (maybe the Researcher agent itself double-checking facts) could be implemented for critical content.
- **TODO:** Scale up the knowledge base. Increase the number of articles significantly (hundreds or more), either through focused content additions or by enabling community contributions. Ensure the system performs with the larger volume (test search with lots of data, test the static site build time with many pages).
- **TODO:** Robust monitoring and alerts. Set up monitoring for the production site (uptime monitoring for the static site and the search API). Use CloudWatch or an external service to track Lambda usage, errors, and performance. Configure alerts (email/Slack/webhook) for failures in the pipeline (e.g. Knowledgebase index build fails, or the Researcher encounters an error so it doesn't silently stop working).
- **TODO:** Implement analytics and feedback loops. Add analytics to the website to understand how users search and what they read (e.g. Google Analytics or a privacy-friendly alternative). This can inform which areas of content to expand. Provide a feedback channel on the site (even a simple link to the GitHub issues for content requests or corrections).
- **TODO:** Finalize versioning and release processes. By now, adopt a consistent version scheme across repos (see [releases.md](releases.md)). Plan the first official release (v1.0) criteria: define what content and features it includes, and thus what the tags will represent. Leading up to that, use pre-release version tags (like v0.9-beta) if needed to test the release process. Tag a release (e.g. v1.0) for the combined system once it meets the production criteria. Document how future updates will be handled – for instance, how often the Knowledgebase index is updated and deployed, how hotfixes to the site or content are managed, etc.
- **TODO:** Create a script or use existing tools to automatically tag all relevant repos with a single version identifier when doing a coordinated release. This could be a simple manual process documented for maintainers (e.g., "upon release, tag Gitopedia, Knowledgebase, Website repos with the same tag"). Alternatively, automate this as part of the CI/CD pipeline to ensure consistency across repositories.
- **TODO:** Establish a process to revisit the versioning strategy periodically. As the project grows, we might need to adjust (for example, if content updates become very frequent, we may automate monthly content release labels; or if the researcher becomes decoupled and has its own development path, it might need independent versioning). This should be part of regular project maintenance and planning cycles.
- **TODO:** Conduct a security and privacy audit before release. This includes ensuring no sensitive data is exposed, verifying compliance with any terms of service for APIs used by the Researcher, and making sure user data (if any collected) is minimal and protected.
- **TODO:** Improve maintainability. Refactor code where necessary (e.g. if some scripts from Phase 1/2 are hacky, clean them up now). Write additional unit tests for critical components (e.g. test the Knowledgebase parsing on various markdown edge cases, test the Search API with tricky queries). Update documentation to reflect the current state (the docs should match the final implementation details).
- **TODO:** Plan for post-launch. Identify features or improvements that were deferred and log them (e.g. semantic search using embeddings, multi-language support, richer front-end features, etc.). Having a backlog ensures the project can continue to evolve after the initial production release.
- **TODO:** Celebrate the release 🎉. Ensure all contributors (human or AI) are credited appropriately in the project documentation. Announce the availability of Gitopedia to a broader audience, inviting feedback and contributions for the ongoing growth of the encyclopedia.
