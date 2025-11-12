# Release Versioning Strategy

This document describes how we handle versioning and releases for the Gitopedia project. With multiple repositories and a continuously updated knowledge base, it's important to define how to label and reference specific states of the system.

## Source Control and Git SHAs

Every commit in each repository is identified by a Git SHA (hash). In our integration process, we propagate these SHAs to ensure traceability:

- When Knowledgebase builds an index, it records the Gitopedia commit SHA from which the content was taken. This could be stored in the Knowledgebase release notes or in the SQLite database (e.g., in a metadata table).
- When the Website is built, it knows the Gitopedia content SHA (passed from Knowledgebase) and can embed it or log it. For example, we might include a JSON file in the deployed site that contains:

```json
{
  "gitopedia_commit": "<sha>",
  "built_at": "<timestamp>"
}
```

- This way, if someone finds an issue on the site or in search results, we can pinpoint exactly which content version and code version they are using.

In general, during active development (Phase 0-2), we rely on commit SHAs and automated triggers rather than manual version bumps. This continuous deployment approach means every commit could potentially be deployed.

## Semantic Versioning for Code

For the code repositories (Website and Researcher, and to some extent Knowledgebase if we treat its build scripts as a product), we will adopt semantic versioning once we approach a stable release:

- **Website code** (Next.js app and search lambda): Will be versioned with tags like v1.0.0, v1.1.0, etc., once we exit the MVP stage. Before v1.0, we might use 0.x versions or simply track via commit until a formal release.
- **Researcher code**: Similarly, if the researcher agent becomes a persistent service or a CLI tool, we version its releases. For example, researcher-v1.0.0 could correspond to the first production-ready iteration. This might be less visible to end users, but it's important for development to know which version of the agent is running (especially if multiple improvements are made over time).
- **Knowledgebase code**: The Knowledgebase repository mostly contains scripts. We can tag versions of the Knowledgebase code (like knowledgebase-v1.0.0) to mark significant changes (e.g., a change in the index schema or build process). This is mainly for developer reference, as users don't directly interact with knowledgebase code.

Semantic versioning (MAJOR.MINOR.PATCH) means:

- Increment **MAJOR** when making incompatible API/format changes (e.g., overhauling how search works or changing the structure of meta.json in a way that requires coordinated changes in the website).
- Increment **MINOR** when adding functionality in a backwards-compatible manner (e.g., adding a new field to meta.json that older code will simply ignore if not present).
- Increment **PATCH** for backwards-compatible bug fixes.

## Versioning Content Releases

Content in Gitopedia is continuously growing, and doesn't have "versions" in the traditional sense (like software does). However, it is useful to mark snapshots of the content:

- We may designate certain snapshots as editions or releases of the knowledge base (for example, "2025 Edition" if we want to do a yearly archival snapshot).
- To do this, we can tag the Gitopedia repository at certain commits. For instance, when we reach 100 articles or when we hit a significant milestone, we tag that commit as `content-v1.0` or `edition-2025-12` (for December 2025 edition). This tag can be used to generate a downloadable package or simply to mark progress.
- The Knowledgebase can also produce a static artifact for that snapshot (the index.sqlite corresponding to that content, maybe a dump of all meta.json). We could attach these to a GitHub Release named "Gitopedia 2025 Edition" for anyone who wants the offline data.

It's important to note that content updates don't follow semantic versioning (since adding a new article isn't a "breaking change" in the same sense; it's just more data). Therefore, content releases might use a date-based versioning or incremental edition numbers.

For example:

- `content-v1.0` might correspond to the MVP content set (maybe a small set of seed articles).
- `content-v2.0` might correspond to the Beta content set (significant expansion).
- Or we might simply use dates: `content-2025-12-01` to indicate a snapshot on Dec 1, 2025.

We will decide on a consistent approach as we approach a stable release:

- For now, using dates in tags for content snapshots is straightforward and unambiguous.
- The code and system as a whole can still have an overarching version (like v1.0 for the first public release) that implies a certain content state as well.

## SQLite Index Versioning

Each build of the Knowledgebase's index.sqlite should be tied to the content commit (as mentioned, the SHA or content tag can serve as its version).

- When Knowledgebase publishes an index, it can name the file or release with a suffix of the content tag or short SHA. For example: `index-01GX....sqlite` where the prefix is the ULID or short commit of the content, or `index-v1.0.sqlite` if aligning with a content release.
- Internally, the SQLite file could have a table containing version info:

```sql
CREATE TABLE metadata (key TEXT, value TEXT);
INSERT INTO metadata (key, value) VALUES ('gitopedia_commit', '<sha>');
INSERT INTO metadata (key, value) VALUES ('build_time', '<timestamp>');
```

This ensures anyone inspecting the DB can know its origin.

- If we publish the index for external use, we will attach a README or metadata with it describing which content it corresponds to.

Since the search index is derived data, we don't version it independently from content. It's essentially a "build artifact" of a given content version. The Knowledgebase repository's code version is separate from the index version:

- Knowledgebase code v1.1 might generate the same content index as v1.0 if no content changed or if changes are backward-compatible.
- But if Knowledgebase code v2.0 introduced a new indexing method (say we switch to a different search engine), then that index might be considered a new format – in which case we would definitely increment the major version of knowledgebase code, and possibly mark indexes differently (e.g., include an index format version in the metadata).

## Container Images (ECR) Versioning

The Researcher is packaged as a container and published to Amazon ECR:

- Images are tagged with the short Git SHA (e.g., `researcher:<sha>`) and `latest`.
- CI builds, tests, and pushes on changes to the Researcher codebase. The image URI and tag are surfaced as a build artifact for downstream deployment systems (e.g., ArgoCD).
- For release milestones, we additionally tag images with the semantic version (e.g., `researcher:v1.0.0`) to align with code tags.
- Provenance can be tracked by embedding the Git SHA and build timestamp into image labels and/or an `/app/version.json` artifact inside the image.

## Documentation Versioning (Project Docs)

The documentation in `/docs` (like this file) will be updated as the project evolves. We treat this documentation as part of the main repository (likely Gitopedia repo or a central "Gitopedia project" repo). To keep documentation in sync with code:

- Whenever we cut a major release (v1.0, v2.0, etc.), we should snapshot the docs. This could be done by tagging the commit of the docs with the release tag. For example, the same tag v1.0.0 can be applied to the commit that contains the finalized docs for 1.0.
- If we publish docs on a website (e.g., via GitHub Pages or another site), we might maintain versioned sections (like a drop-down for docs version). Given this is an internal project, that might be overkill, but tagging ensures we can always retrieve the state of these design docs at any point in time.
- These docs are meant to guide development, so they will likely be continuously refined rather than kept static for old versions. Tagging them is mainly for historical reference if needed.

## Conclusion

In summary, our versioning strategy is:

- Use Git tags and GitHub Releases to mark significant milestones for each repo and for the system as a whole.
- Propagate version identifiers (especially commit hashes) through the build pipeline to maintain transparency of what content/code is deployed.
- Adopt semantic versioning for code when appropriate, and use date or milestone-based versioning for content.
- Maintain documentation alongside development, and tag it with releases.

By carefully versioning each part of the system, we ensure that we can reproduce past states, coordinate multi-repo changes, and communicate clearly about what "version" of Gitopedia someone is using at any given time.

## References

1. [Adding Search to My Static Blog Using AWS Lambda and SQLite — » Henry J Schmale's Blog](https://www.henryschmale.org/2021/07/09/blog-search.html)
2. [AI vs. Machine Learning: How Do They Differ? | Google Cloud](https://cloud.google.com/learn/artificial-intelligence-vs-machine-learning)
3. [Deploying and Securing a Static NextJS Site on AWS | by Adriana Ito | Medium](https://medium.com/@ito.dri/deploying-and-securing-a-static-nextjs-site-on-aws-9e2c9756cfbe)
4. [Identifiers 101: Understanding and Implementing UUIDs and ULIDs - DEV Community](https://dev.to/siddhantkcode/identifiers-101-understanding-and-implementing-uuids-and-ulids-2kc6)
