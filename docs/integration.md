# Integration and Data Flow

This document describes how the four repositories (Gitopedia, Knowledgebase, Researcher, and Website) coordinate with each other. We also clarify the usage of ULIDs for unique identification and how releases are orchestrated across the system.

## Cross-Repository Workflow

The system's operation involves a pipeline that moves from content creation to publication:

1. **Content Update in Gitopedia** – Articles are introduced via a Draft PR: the Researcher stages new or updated Markdown files under a branch-local `_incoming/` directory and opens a Draft PR labeled to trigger the Copilot Organizer. The Copilot Organizer refactors the branch by relocating files into appropriate `Compendium/` categories, validating front matter and authority references, and preparing summarized source websites for Knowledgebase ingestion. When complete, it marks the PR ready. The final (ready) PR merge constitutes the content update for downstream triggers.

2. **Trigger Knowledgebase Build** – The Gitopedia repository notifies the Knowledgebase repository that content has changed. This can be achieved with a GitHub Action that triggers a repository dispatch event or a `workflow_run`. The payload can include the Git commit SHA of Gitopedia to ensure Knowledgebase knows exactly what version to pull.

3. **Knowledgebase Processing** – The Knowledgebase workflow runs:
   - It pulls the latest Gitopedia content (at the specific commit SHA).
   - Regenerates `meta.json` files and the `index.sqlite` full-text search index.
   - Publishes the updated artifacts (committing meta files, creating/updating a release with the SQLite file, and/or uploading it to S3 as configured).

4. **Trigger Website Rebuild** – Upon successful completion, the Knowledgebase pipeline triggers the Website repository to update:
   - The trigger can be another repository dispatch (for example, Knowledgebase's Action calls the GitHub API to send an event to the Website repo, including perhaps the Knowledgebase release version or a URL to the new index).
   - Alternatively, the Website could poll or be scheduled to check for a new index release. A direct trigger is preferred for immediacy.

5. **Website Deployment** – The Website workflow runs:
   - It pulls content from Gitopedia (ideally the same commit that was used for the Knowledgebase build, to keep in sync).
   - Downloads or accesses the latest `index.sqlite` (if needed for packaging with the search lambda).
   - Rebuilds the static site and deploys it to S3/CloudFront.
   - Updates the Search Lambda with the new index (either by redeploying the function or ensuring the function will fetch the new index).

6. **User Facing** – The new content is live on the site and searchable via the updated index.

## CI Automation and GitHub Access

### GitHub App for Cross-Repository Dispatch

For cross-repository workflow triggers (e.g., `repository_dispatch` events), we use a GitHub App (`gitopedia-bot`) with the following permissions:

- **Contents**: Read and Write (for future tag/release operations)
- **Issues**: Read and Write (for Researcher agent operations)
- **Pull requests**: Read and Write (for Researcher and Copilot Organizer operations)
- **Actions**: Write (to trigger workflows via `repository_dispatch`)

**Setup Steps:**

1. **Install the App**: Install `gitopedia-bot` on the three repositories:
   - `gitopedia/gitopedia`
   - `gitopedia/knowledge-base`
   - `gitopedia/website`

2. **Generate Installation Token**: Use the GitHub App's App ID and private key to generate an installation access token. This token is used in CI workflows to authenticate as the app.

3. **Store Token as Secret**: Store the installation access token as an organization secret `KB_DISPATCH_TOKEN` (or similar) that workflows can use to authenticate when calling the GitHub API for `repository_dispatch` events.

**Usage in Workflows:**

Workflows use the installation token to authenticate GitHub API calls:
```yaml
- name: Trigger Knowledgebase
  env:
    GITHUB_TOKEN: ${{ secrets.KB_DISPATCH_TOKEN }}
  run: |
    gh api repos/gitopedia/knowledge-base/dispatches \
      -X POST \
      -f event_type=content-updated \
      -f client_payload='{"ref":"main","sha":"${{ github.sha }}"}'
```

### SSH Deploy Keys for Git Operations

To enable fully automated git operations (creating branches, committing, pushing), CI workflows are granted GitHub access over SSH:

- A dedicated deploy key (SSH keypair) is created for CI. The public key is added as a Deploy Key on the relevant repositories with write access; the private key is stored as an encrypted secret in the CI environment.
- CI jobs load the private key using `ssh-agent` and configure `known_hosts` for GitHub to prevent MITM prompts.
- With SSH access, workflows can:
  - Create and update branches, Draft PRs, and PR labels to trigger the Copilot Organizer
  - Open GitHub Issues for reporting automation findings or required follow-ups
  - Push commits made by automation (e.g., Copilot Organizer refactors)

### Researcher Integration

The Researcher agent operates somewhat independently, but is integrated via GitHub:

- It monitors Gitopedia issues for content requests.
- When it creates a Draft PR with staged content under `_incoming/`, the Copilot Organizer performs automated refactoring and preparation. After Copilot completion and PR readiness, that triggers the same chain as above. The Researcher agent is the exclusive source of new article content; all articles are created by the automated agent, not by human contributors.
- We may add a label or note in the PR to indicate it was AI-generated for transparency, but the pipeline doesn't change.

### Error Handling and Sync

Each step in the pipeline should ideally wait for the previous one to succeed:

- If Knowledgebase fails to build (e.g., parse error in new Markdown), it should report back (perhaps opening an issue or commenting on the commit/PR) and not trigger the website update. The error can be addressed (the content fixed) and the pipeline re-run.
- The Website build should only use an index that corresponds to the content it's building. By passing the Git commit and/or a release identifier through the dispatch events, we ensure consistency. For example, Knowledgebase could include `content_sha: <Gitopedia commit>` in the dispatch to Website; the Website action then checks out that commit of Gitopedia to build pages, and uses the index file from the Knowledgebase release tagged with that SHA.
- This prevents a rare race condition where multiple content updates in quick succession might otherwise cause mismatch (e.g., building site with newer content but older index or vice versa).

## ULIDs for Unified Identification

We use ULIDs as a way to uniquely and consistently identify content across the system:

- When the Researcher creates a new article, it generates a ULID and places it in the article's front matter (`id` field). This ULID becomes the primary key for that article in the Knowledgebase index and meta.
- ULIDs are 128-bit lexicographically-sortable unique IDs that include a time component[11]. This means they roughly sort by creation time. We leverage this property in various ways (for example, we could sort articles by ID to get a chronological list without needing a separate timestamp).
- The Knowledgebase uses the ULID as the stable identifier for each article record. For instance, the `index.sqlite` might have a table of articles keyed by ULID, and FTS results return the ULID of the matching article.
- The Website can then use this ULID to cross-reference the content. One approach: the site's pages are organized by human-friendly slugs, but we can maintain a mapping of ULID to slug. The Knowledgebase could output a JSON map or include the slug in meta, so that given a ULID we can find the corresponding page.
- For example, a search result returns `id = 01GX...` and `title = "Quantum Computing Overview"`. The site can map that to `/articles/quantum-computing-overview` (either by a convention of slugifying the title or via a lookup table generated at build time from `meta.json`).
- ULIDs ensure that even if titles change or slugs are updated, we have a permanent reference. If a title changes, the article's ULID in front matter stays the same, so Knowledgebase can recognize it as an existing entry and update its title. This decouples identity from title or filename.
- Also, because ULIDs include timestamp, it's easy to visually (or programmatically) detect roughly when an article was added from its ID. This can be used for analytics or ordering.

### Ensuring Uniqueness

The chance of ULID collisions is practically zero, but we will still enforce:

- The Researcher will use a well-tested ULID library to generate IDs.
- The Knowledgebase can double-check that no two articles share the same ULID (if a duplicate is found, log an error or assign a new one if possible, though this situation is unlikely unless someone copied an existing id by mistake).

## Release Coordination and Versioning

Coordinating releases means ensuring that at any given tagged release or deployment, all components are aligned:

- **Git SHAs**: We plan to capture the Git commit hashes in the workflow as content moves through. For example, Knowledgebase might store the Gitopedia commit hash that it indexed (perhaps in a text file or the release name). The Website could embed the Gitopedia commit hash in the deployed site (for instance, in a comment in the HTML or an `/about` page) to trace exactly what content version it's showing. This traceability is important for debugging or for formal version releases.

- **Index Versions**: Each `index.sqlite` could be versioned (even if implicitly by content hash or timestamp). If using GitHub Releases in Knowledgebase, each release could be tagged with an incrementing number or a date (e.g., `index-v2025-11-15`). The Website's deployment action can ensure it grabs the correct version (for example, the latest).

- **Docs Tags**: When we reach major milestones (like MVP, Beta, v1.0), we will tag the repositories to mark those points. For instance:
  - Tag the Gitopedia repo at a commit that forms the content of v1.0 (e.g., `content-v1.0`).
  - Tag the Knowledgebase repo at the corresponding index release (e.g., `index-v1.0`).
  - Tag the Website repo for the code that was used in v1.0 (e.g., `site-v1.0`).
  - These tags allow us to later reproduce that exact state if needed, and serve as reference points in documentation.
  - The term "docs tags" might also refer to tagging versions of documentation. Since documentation (like these design docs) will evolve, we might tag the main project repository when these docs are updated for a new phase or release. This way, if someone wants to read the docs relevant to version 1.0, they can check out the v1.0 tag of the docs.

In practice, much of our deployment will be continuous (new articles go live continuously). Formal versioning is more for major releases or if we package the entire Gitopedia for offline use. Nonetheless, implementing basic version tagging is part of being production-ready.

### Maintaining Sync

It's critical that the search index and the website content remain in sync. We handle this by the sequential trigger design above. There should never be a scenario where the website is deployed without the matching index or vice versa. To safeguard:

- The Website could perform a sanity check on startup: e.g., load a small manifest from Knowledgebase that contains the last content update ID and compare with what it built. If mismatch, it could warn or even fetch new data.
- Similarly, the search Lambda could include an identifier for the index (like a version number) and the website could include that in its search requests to ensure it's querying the right version. (This is an advanced check and usually not necessary if the deployment pipeline is robust.)

### Automated Release Process

When ready for a formal release (say, moving from Phase 1 MVP to Phase 2 Beta):

- We ensure the system is in a stable state (all pipelines green).
- Create tags on each repo as needed (or use a single script to tag all with the same version label).
- Possibly generate a changelog or release notes summarizing what content and features were added since the last tag.
- Continue development beyond that tag for the next iteration.

## Conclusion

Integrating all these pieces creates a cohesive system where each component knows when to act and how to find the data it needs: the Git repositories communicate through Actions, unique IDs tie content together, and version tags mark milestones. This ensures that Gitopedia can continuously evolve with confidence that all parts will remain in lockstep.
