# Roadmap

This roadmap outlines the development of Gitopedia in stages from a pre-MVP setup to a production-ready system. Each phase is defined with specific objectives and deliverables. The phases are incremental: each builds upon the previous, gradually adding functionality and stability. TODO items are listed under each phase as actionable steps required to achieve the phase's goals.

## Phase 0: Pre-MVP Setup (Project Bootstrapping)

**Goal:** Establish the foundational structure for all components of the system. In this phase, we set up the repositories and basic pipelines, but the end-to-end functionality may be stubbed or manual. The focus is on scaffolding the project so development can proceed in an organized way.

**Deliverables:** All four repositories initialized with minimal code/configuration, basic documentation (this `docs/` directory), and placeholder processes for content flow.

### Key Tasks

- **DONE:** Create the Gitopedia repository structure. Set up directories for articles (`Compendium/` folder with category subdirectories) and add a sample Markdown article. Set up the `authority/` directory structure for controlled vocabularies (people, orgs, places, topics). Include a README explaining the article format and how the Researcher agent creates articles. Decide on an article format (e.g. use a [YAML front-matter](https://docs.github.com/en/contributing/writing-for-github-docs/using-yaml-frontmatter) with fields like `id`, `title`, `slug`, etc., where `id` will be a ULID for each article).
- **TODO:** Create documentation in README.md of Gitopedia that summarizes the article format (including facets and authority lists), and how the Researcher agent creates articles via automated PRs. Note that all articles are created by the Researcher agent, not by human contributors.
- **TODO:** Provide ULID generation utilities for the Researcher agent:
  - Ensure the Researcher agent uses a consistent ULID library to generate IDs (for example, `ulid-py` for Python).
  - Possibly create a GitHub Actions chatbot for Gitopedia that, on a new PR without an `id` in front matter, leaves a comment with a suggested ULID or even automatically pushes a commit adding one.
- **TODO:** Provision self-hosted infrastructure for Researcher dependencies:
  - Deploy a self-hosted SearXNG instance (Docker), configure engines/timeouts/rate-limits, and expose an internal `SEARXNG_URL`.
  - Deploy a self-hosted DeepSeek model behind an OpenAI-compatible API (e.g., vLLM gateway). Expose an internal `OPENAI_BASE_URL` and set default `DEEPSEEK_MODEL`.
  - Provision a stable runtime for the Researcher (VM/container) with private network access to SearXNG and DeepSeek endpoints.
- **TODO:** Initialize the Knowledgebase repository. Define a basic schema for metadata (e.g. what fields a `meta.json` should contain for each article). Stub out a script (or notebook) that can read a Markdown file and produce a simple JSON and/or add an entry to a SQLite database. No full functionality yet, but lay the groundwork (create the repository, add dependencies like SQLite libraries in a `requirements.txt`).
- **TODO:** Initialize the Researcher agent repository. Set up a Python project (e.g. a `main.py` that prints a greeting to confirm the environment works). Outline the steps this agent will perform (perhaps in the README or comments). Ensure the repository has access to necessary credentials (store placeholders for GitHub token, etc., to be filled in later) and can be run locally.
- **TODO:** Initialize the Website repository with a new Next.js app. Configure `next.config.js` for static export (set `output: 'export'`). Create a placeholder homepage that confirms the app builds and deploys (for example, a page that says "Gitopedia coming soon"). Verify that `npm run build && next export` produces static files. Set up the deployment target (e.g. an S3 bucket or vercel setup) - this can be manual initially.
- **TODO:** Set up basic CI/CD workflows (GitHub Actions or similar) for cross-repo integration. For example, configure an Action in Gitopedia that triggers (via repository dispatch or a scheduled job) a workflow in Knowledgebase. In Phase 0, these can just echo messages or create dummy files to ensure the wiring is in place.
- **TODO:** Decide on the unique ID strategy for articles. Plan to use ULIDs for each article's stable identifier (for chronological sorting and uniqueness). In Phase 0, implement a small utility or note to generate ULIDs (e.g. integrate a ULID library in the Researcher or a script) and include an example ULID in the sample article's metadata. Set up validation for ULIDs: a pre-commit hook or CI step should verify that any new article has a valid-looking ULID in the `id` field.
- **TODO:** Document the development setup for each component. Ensure that another developer (or AI agent) can pull each repo, run the basic setup, and understand how the pieces will connect. This includes updating this documentation and repository READMEs as needed.

## Phase 1: MVP (Minimum Viable Product)

**Goal:** Achieve a working end-to-end system where a new article can be created, processed, and displayed on the website with a functional search. The MVP doesn't need to be polished or scalable, but it should demonstrate the core loop: content -> index -> website. The Researcher should be able to add at least one article automatically (even if using a simple approach or minimal AI logic).

**Deliverables:**
- A simple end-to-end demo of Gitopedia: for example, a "Hello World" article created via a research issue, appearing on the live site, searchable by its title or content.
- Basic implementations for content parsing, indexing, static site generation, and a search query endpoint.

### Key Tasks

- **DONE:** Implement content ingestion in Gitopedia. Define contribution guidelines for articles (e.g. how sources should be cited in Markdown). Ensure each article has a ULID in its front-matter (from Phase 0 plan). Create a couple of sample articles to use for testing the pipeline (these can be brief).
- **DONE:** Add Draft PR staging with `_incoming/` and Copilot Organizer. Researcher opens Draft PRs with articles and summarized sources staged under `_incoming/`. A Copilot Organizer GitHub Action/agent refactors the branch (moves articles to `Compendium/`, validates front matter and facets, prepares summarized sources for Knowledgebase), pushes commits, and marks the PR ready.
- **TODO:** Implement CI pipeline for Researcher container:
  - Multi-stage Docker build producing a distroless runtime image.
  - Run unit/integration tests inside CI.
  - Authenticate to ECR and push images tagged with commit SHA and `latest`.
  - Surface image URI/tag as an artifact for deployment systems (e.g., consumed by ArgoCD).
- **TODO:** Enable CI GitHub SSH automation:
  - Create a deploy key; add public key to repos with write access, store private key as CI secret.
  - Configure `ssh-agent` and `known_hosts` in CI to safely perform git operations.
  - Automate creating Issues, Draft PRs, applying labels to trigger the Copilot Organizer, and pushing commits generated by automation.
- **TODO:** Integrate self-hosted services and SDKs in the Researcher:
  - Add a SearXNG client (pointed at `SEARXNG_URL`, optional `SEARXNG_API_KEY`) for web search.
  - Use the Python OpenAI SDK for both OpenAI cloud and self-hosted DeepSeek via `OPENAI_BASE_URL`.
  - Implement provider selection (prefer DeepSeek self-hosted, fallback to OpenAI), with timeouts and retries.
  - Add configuration and documentation for `SEARXNG_URL`, `OPENAI_BASE_URL`, `OPENAI_API_KEY`, `OPENAI_MODEL`/`DEEPSEEK_MODEL`.
- **DONE:** Implement the Knowledgebase indexing process. Develop a Python script (or use an existing tool) to parse all Markdown files in Gitopedia. For each article, generate its metadata (`meta.json`), including fields like `id` (ULID), `title`, perhaps an excerpt or list of references. Aggregate these into a searchable SQLite database (using SQLite FTS5 or FTS4 for full-text search on article content). Verify that running this script locally produces an `index.sqlite` with the sample articles.
- **DONE:** Set up an automated workflow for Knowledgebase. Configure a GitHub Action that triggers on new content in Gitopedia (e.g. push events on main) to run the indexing script. For MVP, this action can simply run tests or produce an artifact (it might not yet deploy the SQLite file, but should at least confirm the process). Ensure the workflow can ingest new summarized sources organized by the Copilot phase.
- **DONE:** Implement cross-repo triggers using GitHub Actions:
  - **In Gitopedia**: Create a workflow (on push to main) that triggers Knowledgebase's workflow. This can use `actions/github-script` or a curl to GitHub API to send a `repository_dispatch` to Knowledgebase. (Requires a GitHub token with repo scope stored as a secret.)
  - **In Knowledgebase**: Add a workflow to receive the dispatch (trigger on: `repository_dispatch` with a specific event type, e.g., `content-updated`). This workflow runs the index build. On completion (success), use a similar method to dispatch to Website. Include payload data like `gitopedia_sha` and perhaps knowledgebase's new index version (or a URL to fetch it).
  - **In Website**: Have a workflow on: `repository_dispatch` (e.g., event type `rebuild-site`) that runs the build and deploy. It should use the `gitopedia_sha` from the payload to ensure it builds the matching content version. It should also retrieve the new `index.sqlite` (maybe provided as an artifact URL or downloaded from S3) for packaging with the search lambda or for verification.
  - Secure these triggers by scoping them to specific event types and secrets. (Only accept dispatches from our own workflows, using a shared secret if needed to verify authenticity.)
- **TODO:** Develop the Search Lambda (Search API) in the Website (or a dedicated directory). Use Python or Node.js to create a function that can load the `index.sqlite` file and execute full-text queries. For MVP, you can bundle the SQLite from Knowledgebase manually (e.g. attach it in the lambda package). Implement a simple query handler: accept a query string and return top N article IDs/titles. (Tip: Use SQLite FTS query with rank and snippet to get a summary – this is supported as shown in similar projects). Specific implementation steps:
  - Write the Lambda function code to open the SQLite database and perform a parameterized search query. (Ensure to use parameter binding to prevent any possibility of SQL injection, even though input is likely limited and sanitized by SQLite FTS.)
  - Test the function locally (it can be run with a test event JSON). Use a small test SQLite with known entries to verify that queries return expected results.
  - Set up infrastructure as code for deployment:
    - Define an API Gateway route integration for the Lambda. This can be done with AWS SAM template or Serverless YAML. For example, define a `GET /search` route that triggers the lambda, and enable CORS to the known origins (localhost for dev, and production domain).
    - Automate deployment via GitHub Actions: e.g., use `aws cloudformation package/deploy` commands or the Serverless framework CLI to push updates.
    - Store necessary AWS credentials (or use GitHub OIDC to assume a deploy role) in the repository secrets so CI can deploy the Lambda and upload the static site.
  - After deploying, test end-to-end: from a browser (or Postman) call the `/search` API with a sample query and see that it returns results from the latest content.
- **TODO:** Integrate search on the Website frontend. Create a search input field on the site (perhaps on the homepage or a dedicated search page). When a user submits a query, use JavaScript (fetch API) to call the Search API endpoint. Display the list of results (at least title and snippet, linking to the article pages). For MVP, this can be a very basic UI.
- **TODO:** Implement static page generation for Website. Write a script or Next.js page that iterates over all Markdown articles from Gitopedia (which might be pulled in as a submodule or fetched via GitHub API) and generates a static page for each. This likely involves using a Markdown processing library (like remark or Next.js MDX). Ensure that the sample articles from Gitopedia appear as pages on build. Also, generate an index page or sidebar listing all articles (even a simple list of titles linking to pages). Specific implementation steps:
  - Set up Git access or API access for the build process to retrieve Markdown files from Gitopedia. Possibly store a GitHub token in CI to fetch private repo content (if the repository is private).
  - Write scripts or use Next.js API routes to load content during `getStaticProps`. For simplicity, consider checking out the Gitopedia repository as part of the build step so all markdown files are present locally.
  - Use a Markdown parsing pipeline (`gray-matter` for front matter, `remark` for markdown to React or HTML). Test that an example Markdown file is correctly rendered as a page.
  - Create page templates: one for article pages (which takes content and maybe a list of references to display), and one for index or home page.
  - Ensure that links between articles (if any) are properly handled (we might use absolute URLs or generate Next.js `<Link>` components for internal links).
  - Validate that after export, the files are in the expected locations (Next.js by default might output each page as an `index.html` in a folder named after the slug).
- **DONE:** Implement the index generation script for category `index.md` files and integrate it into GitHub Actions. This includes: parsing Markdown files for titles, handling edge cases (e.g., an article missing a title in front matter should fallback to the first markdown heading), committing the updated category indexes, and ensuring the action has appropriate permissions.
- **TODO:** Connect the Researcher agent minimal functionality. Implement the logic for the Researcher to fetch a "research request" (for MVP, this could be a specific GitHub issue or a hard-coded task). Integrate an API call to an LLM (if available, e.g. OpenAI API) to generate content for a given prompt. If external API access is not available in the development environment, simulate this by using a placeholder text or a locally stored answer. Focus on the flow: the Researcher should take a prompt (issue title/description), produce a Markdown draft (with at least a title, some content, and maybe a fake citation), and create a Pull Request to the Gitopedia repo. Use the GitHub API or `gh` CLI for creating the PR.
- **TODO:** Test the end-to-end flow for MVP:
  - Create a dummy "New Article Request" issue on Gitopedia (with a simple prompt).
  - Run the Researcher agent to process this issue (generating an article and PR).
  - Merge the PR into Gitopedia.
  - Verify the Knowledgebase action runs and updates the index.
  - Manually deploy or update the Search API with the new index (for MVP, this might be a manual step if full automation isn't done).
  - Rebuild the Website and deploy the updated static site.
  - Open the website and confirm the new article is visible and searchable.
- **TODO:** Fix any issues uncovered in testing. Likely areas: Markdown formatting errors, search queries not returning expected results, deployment misconfigurations (e.g., CORS issues between the static site and search API). Address these to ensure the MVP demonstration is successful.

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
