# Roadmap

This roadmap outlines the development of Gitopedia in stages from a pre-MVP setup to a production-ready system. Each phase is defined with specific objectives and deliverables. The phases are incremental: each builds upon the previous, gradually adding functionality and stability. TODO items are listed under each phase as actionable steps required to achieve the phase's goals.

**Note:** Phase 0 (Pre-MVP) and Phase 1 (MVP) are complete. See [roadmap-archive.md](roadmap-archive.md) for the history of completed tasks.

## Phase 2: Enhanced Capabilities and Beta Release

**Goal:** Build on the MVP to add important features, improve quality, and increase the system's knowledge base. In this phase, the focus is on refining functionality, increasing content, and ensuring the system can handle more use cases. We consider this a "Beta" stage where the system is usable by early adopters, but not yet fully polished.

**Deliverables:**
- Several dozen high-quality articles in Gitopedia, many generated via the Researcher agent.
- Improved search features (accuracy, performance, possibly filtering or categorization).
- More automation and robustness in the pipelines (fewer manual steps).
- Internal evaluations or testing by a small set of users.

### Key Improvements and Tasks

- **DONE:** Enrich article metadata and structure. Define additional front-matter fields as needed (e.g. `author`, `tags`/`categories`, `summary`). Ensure the Knowledgebase captures these in `meta.json`. This will allow enhancements like category pages or filtering search results by tag.
- **DONE:** Implement facet validators in CI to ensure: facet values resolve to authority IDs, required facets are present (≥1 topic, at least one of people/orgs/places/date_refs), and authority ID format is correct. Implement canonical title enforcement in CI to prevent duplicate titles.
- **DONE:** Implement authority list management (creation, merging, alias resolution) in the Researcher agent. The agent should check authority lists when creating articles, create new authority entries when needed (subject to review), and suggest merges if similar entities already exist.
- **SKIPPED:** Create redirect mechanism for renamed or merged articles. Decision: Strict renaming policy enforced. Renaming an article requires refactoring links; no automated redirects.
- **DONE:** Improve the Researcher agent's output quality. Integrate better prompt templates and multi-step research (e.g. have the agent gather multiple sources before writing). Ensure that the agent includes citations to external sources in the Markdown (perhaps as footnotes or reference list) to maintain credibility. Implement error handling for the agent (e.g. if the API call fails or the content is incomplete, the agent should log an issue or retry).
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
