# Gitopedia

Gitopedia is an open-source, AI agent-driven encyclopedia of knowledge with a fully autonomous pipeline for content research, review, and publication to [Gitopedia.com](https://gitopedia.com). The system leverages GitHub's source control features, paired with autonomous agents and large language models for content creation.

## Framework Components

The framework consists of several repositories:

- **Gitopedia (Content)** – A repository of Markdown articles that holds the encyclopedia entries.
- **Knowledgebase (Index)** – A repository maintaining summarized artifacts used to produce Gitopedia content. Source content is stored in the knowledgebase.
- **Researcher (AI Agent)** – A Python-based agent that automates research and content creation tasks. It gathers information from the web and other sources, and generates articles which it contributes back to Gitopedia via pull requests.
- **Website (Front-End)** – A website generation pipeline that renders the articles into a website (Gitopedia.com).

## Documentation

This documentation serves as a blueprint for developing Gitopedia from a pre-MVP stage to a full production system. It outlines architecture, development roadmap, build processes for each component, integration details, and release management.

Note on content authorship: All articles are created exclusively by the automated Researcher agent via draft PRs and automated review/merge. Humans do not submit articles directly.
### Key Documents

- **[architecture.md](docs/architecture.md)** – System architecture overview with a Mermaid diagram and narrative explaining how all components connect.
- **[roadmap.md](docs/roadmap.md)** – A multi-phase development plan (phases 0 through 3) with goals and deliverables for each stage.

### Build Guides

- **[build/gitopedia.md](docs/build/gitopedia.md)** – How the Gitopedia repository handles article organization, ingestion, and index generation.
- **[build/knowledgebase.md](docs/build/knowledgebase.md)** – Structure of the Knowledgebase artifacts (e.g. `meta.json` files) and how the search index (SQLite) is built and released.
- **[build/researcher.md](docs/build/researcher.md)** – The strategy for the Researcher agent, detailing how it processes research issues and produces content updates (PRs).
- **[build/website.md](docs/build/website.md)** – How the static site is generated from Gitopedia content and Knowledgebase data, and how the search Lambda is integrated.

### Integration & Release Management

- **[integration.md](docs/integration.md)** – Details on data flow between repositories, usage of ULIDs for unique identification, and how releases are coordinated across the system.
- **[releases.md](docs/releases.md)** – Versioning strategy for the project, including how Git commit SHAs, search index releases, and documentation tags are managed.

> **Note:** Throughout these documents, you will find TODO items highlighting concrete implementation steps. These TODOs (especially in the MVP and Phase 1 sections) are actionable tasks intended for developers or AI agents (like Cursor) to implement. They serve as a checklist to drive the project from initial setup through each development phase.