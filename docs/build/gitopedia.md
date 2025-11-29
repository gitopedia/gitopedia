# Gitopedia Repository – Content Organization and Ingestion

The Gitopedia repository is the central knowledge base containing all Markdown articles (encyclopedia entries). This document explains how articles are organized, how new content is added (ingestion process), and how the repository maintains an index of articles.

## Repository Structure

All articles are stored as Markdown (`.md`) files in a designated content directory (`Compendium/` folder at the root of the repository). Articles are organized into nested category subdirectories that mirror broad domains for human navigation. Each article is a standalone Markdown file. A possible layout is:

```
/ (repository root)
├── Compendium/
│   ├── Science/
│   │   ├── Physics/
│   │   │   ├── quantum-mechanics.md
│   │   │   └── ...
│   │   ├── Chemistry/
│   │   │   └── ...
│   │   └── ...
│   ├── History/
│   │   ├── WWII/
│   │   │   ├── d-day.md
│   │   │   └── ...
│   │   └── ...
│   ├── Technology/
│   │   ├── AI/
│   │   │   ├── OpenAI.md
│   │   │   └── ...
│   │   └── ...
│   └── ... (other categories)
├── authority/
│   ├── people.json
│   ├── orgs.json
│   ├── places.json
│   └── topics.json
├── README.md (repository introduction)
└── ... (other files like license, contribution guide, etc.)
```

- **Category Organization:** Articles are organized into nested category subdirectories within `Compendium/`. Categories mirror broad domains (e.g., `Science/Physics/`, `History/WWII/`, `Technology/AI/`, etc.) to facilitate human navigation. The Researcher agent determines the appropriate category path based on the article's topic and facets.
- **Article Filenames:** We use human-readable slugs (usually based on the article title or topic). For example, an article titled "OpenAI" might be stored as `Compendium/Technology/AI/OpenAI.md`. Use singular form for classes (e.g., "Planet" not "Planets"). For disambiguation, use parentheses: `Mercury (planet).md` vs `Mercury (element).md`. Filenames should be lowercase and descriptive. Avoid spaces or special characters.
- **Stable IDs:** Every article is assigned a unique identifier (ULID) that is included within the file's front matter. Articles use readable slugs for filenames; artifacts use ULIDs for stable references. The ULID is used by other components (Knowledgebase, integration scripts) to reference the article reliably, even if the filename or title changes.

## Article Format and Front Matter

Each Markdown file may begin with a [YAML front matter](https://docs.github.com/en/contributing/writing-for-github-docs/using-yaml-frontmatter) section providing metadata. The expected front matter fields include:

```yaml
---
id: 01HABCD1234XYZ...    # ULID for the article
title: "OpenAI"
slug: "openai"            # Human-readable slug
created: 2025-11-12
author: "Gitopedia Bot"
tags: ["artificial intelligence", "LLM"]
summary: "A research organization focused on developing safe artificial intelligence."
---
```

- **`id`:** A ULID string that uniquely identifies the article. This is generated when the article is first created (by the Researcher agent). Once assigned, this ID should remain with the article even if the title or content changes.
- **`title`:** The human-friendly canonical title of the article. One canonical title per concept is enforced with CI checks. (The article's Markdown content should also start with an H1 heading of the same title for consistency.)
- **`slug`:** A human-readable slug used in URLs and filenames. Should match the filename (without extension and path).
- **`created`** (optional): The date the article was created (or first added). Could be used for sorting or indicating freshness.
- **`author`** (optional): The creator of the article (e.g., "Gitopedia Bot").
- **`tags`** (optional): An array of tags or categories. This can help group articles by topic. (Tags might be single words or phrases like "AI" or "History".)
- **`summary`** (optional): A concise summary of the article content.

After the front matter, the article content should begin with a top-level heading (e.g. `# OpenAI`) followed by the body of the article.

### Citations and References

In the article content, it's encouraged to cite sources for factual statements. A consistent citation format should be used. One approach (especially for AI-generated content) is to include reference links in brackets, e.g. 【source†】 style references. These can later be converted to footnotes or reference lists when rendering on the site. For now, authors can include such references inline; the static site builder or Knowledgebase may post-process them into a more readable form on the website.

## Facets and Metadata

The Knowledgebase generates structured metadata for each article in `meta.json` files. These metadata files include **facets** that categorize and link articles to authority-controlled entities:

```json
{
  "id": "01HABCD1234XYZ567890123456",
  "title": "OpenAI",
  "slug": "openai",
  "index": {
    "topics": ["Artificial intelligence", "Large language models"],
    "people": ["Sam Altman"],
    "orgs": ["OpenAI"],
    "places": ["United States"],
    "date_refs": ["2023", "2024"],
    "entities": ["GPT-4", "GPT-5"]
  }
}
```

### Facet Types

- **`topics[]`:** Array of topic identifiers that describe the subject matter of the article.
- **`people[]`:** Array of person identifiers mentioned or central to the article.
- **`orgs[]`:** Array of organization identifiers mentioned or central to the article.
- **`places[]`:** Array of location identifiers mentioned or central to the article.
- **`date_refs[]`:** Array of date references (years, decades, etc.) relevant to the article.
- **`entities[]`:** Array of other named entities (products, concepts, etc.) mentioned in the article.

All facet values should resolve to authority IDs (see Authority Lists below). CI validators ensure facet values resolve to authority entries or suggest merges for new values.

### Category Lint Requirements

Each article must meet minimum facet requirements:
- **Required:** At least one `topic` must be specified.
- **Required:** At least one of `people`, `orgs`, `places`, or `date_refs` must be specified.

CI checks enforce these requirements on all pull requests.

## Authority Lists

Authority lists provide controlled vocabularies for facets, ensuring consistency and enabling disambiguation. Authority files are stored in the `/authority/` directory:

- `/authority/people.json` - Controlled list of people
- `/authority/orgs.json` - Controlled list of organizations
- `/authority/places.json` - Controlled list of places
- `/authority/topics.json` - Controlled list of topics

Each authority file contains an array of entries with the following structure:

```json
{
  "id": "org:openai",
  "label": "OpenAI",
  "aliases": ["Open AI"]
}
```

- **`id`:** A stable identifier for the entity (e.g., `org:openai`, `person:altman-sam`, `place:united-states`).
- **`label`:** The canonical name for the entity.
- **`aliases[]`:** Array of alternative names or spellings that should resolve to this entity.

When the Researcher agent creates articles, it should reference entities by their authority IDs in the facets. If a new entity is encountered, the agent should either:
1. Check if an alias exists in the authority lists
2. Create a new authority entry (subject to review)
3. Suggest a merge if a similar entity already exists

## Example Article Structure

**Path:** `Compendium/Technology/AI/OpenAI.md`

**Front Matter:**
```yaml
---
id: 01HABCD1234XYZ567890123456
title: "OpenAI"
slug: "openai"
created: 2025-11-12
tags: ["artificial intelligence", "LLM", "technology"]
---

# OpenAI

OpenAI is an artificial intelligence research organization...
```

**Corresponding Knowledgebase `meta.json`:**
```json
{
  "id": "01HABCD1234XYZ567890123456",
  "title": "OpenAI",
  "slug": "openai",
  "index": {
    "topics": ["topic:artificial-intelligence", "topic:large-language-models"],
    "people": ["person:altman-sam"],
    "orgs": ["org:openai"],
    "places": ["place:united-states"],
    "date_refs": ["2023", "2024"],
    "entities": ["GPT-4", "GPT-5"]
  }
}
```

**Authority Entry Example (`authority/orgs.json`):**
```json
{
  "id": "org:openai",
  "label": "OpenAI",
  "aliases": ["Open AI"]
}
```

## Governance and Automation

### Naming Rules

- **Singular for classes:** Use singular form for entity types (e.g., "Planet" not "Planets").
- **Disambiguators in parentheses:** When multiple entities share a name, use parentheses to disambiguate:
  - `Mercury (planet).md` vs `Mercury (element).md`
  - `Washington (state).md` vs `Washington (city).md`

### Canonical Titles and Redirects

- **One canonical title per concept:** Each concept should have exactly one canonical article title. CI checks enforce this rule.
- **Redirects:** If an article is renamed or merged, redirects should be created to maintain link integrity. Redirects can be stored in a `redirects.json` file or handled by the Knowledgebase.

### Facet Validators

CI checks ensure:
- Facet values resolve to authority IDs (or suggest merges for similar entities).
- Required facets are present (at least one topic, and at least one of people/orgs/places/date_refs).
- Authority IDs follow the correct format (e.g., `org:`, `person:`, `place:`, `topic:`).

### Category Lint

Each article must have:
- **≥1 topic** in the facets
- **At least one of:** `people`, `orgs`, `places`, or `date_refs`

CI will fail PRs that don't meet these requirements.

## Relations (Future Enhancement)

For richer knowledge graphs, a `relations` table can be added to store relationships between entities:

```json
{
  "id": "relation:001",
  "subject_id": "org:openai",
  "predicate": "founded_by",
  "object_id": "person:altman-sam"
}
```

This enables querying relationships like "organizations founded by X" or "people associated with Y".

## Adding and Ingesting Articles

All new articles in Gitopedia are created exclusively by the Researcher agent. Humans do not submit articles directly; content is generated automatically by the AI agent.

### Staging via `_incoming/` and Copilot Organizer

New or updated articles are first staged by the Researcher agent into a branch-local `_incoming/` directory at the root of the branch. The Researcher also stages summarized source materials under `_incoming/sources/`. The Researcher then opens a Draft Pull Request and triggers a Custom Copilot Agent ([Encyclopaedist](../agents/README.md)) that:

- Analyzes the `_incoming/` contents (articles only, not sources)
- Refactors and moves articles into the appropriate `Compendium/` category paths
- Flags any authority entries that need review (does not create them automatically)
- Ensures front matter (ULID, title, slug, tags) is valid and consistent
- **Leaves `_incoming/sources/` untouched** – these are temporary staging for the Knowledgebase
- Pushes commits back to the PR branch and transitions the PR from Draft when complete

**Source materials flow:** After the PR merges, the Knowledgebase indexer ingests sources from `_incoming/sources/` into its SQLite database, then deletes them from gitopedia. This keeps the repository focused on articles while making sources searchable via the Knowledgebase.

This ensures all organization and compliance is fully automated before merge.

### Via Pull Requests

The Researcher agent adds or updates articles through pull requests. The agent submits a Draft PR that contains new or updated Markdown files staged under `_incoming/`. Each PR should ideally contain one article addition/update at a time, including:

- The article Markdown file(s) staged under `_incoming/` (the Copilot Organizer will relocate them into the appropriate `Compendium/` subdirectories).
- Updates to authority lists if new entities are introduced.
- (Optional) an update to `index.md` (the table of contents) – if not, an automated process will update it.

When a PR is merged into the main branch, it triggers the ingestion pipeline:

1. **Knowledgebase Update:** A GitHub Action in Gitopedia (or a webhook) notifies the Knowledgebase repository that new content is available. The Knowledgebase will:
   - Pull the latest articles from `Compendium/` and update its metadata and index
   - Ingest sources from `_incoming/sources/` into its SQLite `sources` table
   - Delete `_incoming/sources/` from gitopedia after successful ingestion
   - See [knowledgebase.md](knowledgebase.md) for details.
2. **Website Rebuild:** After the Knowledgebase processes the update (or in parallel), the Website repository is prompted to rebuild the static site so the new article becomes available online.
3. **Search Index Refresh:** The updated search index (SQLite) from Knowledgebase is deployed to the Search API, ensuring both articles and sources are searchable.

(The integration between these steps is automated; see [integration.md](../integration.md) for configuration of cross-repo triggers.)

### Via the Researcher Agent

The Researcher (automation) creates new articles by opening PRs. It monitors for requests (for example, a GitHub Issue titled "Article Request: X") and, when it has generated the content for X, it will:

- Create a new branch on the Gitopedia repo and stage generated Markdown file(s) under the branch-local `_incoming/` directory (each with front matter including a new ULID, title, and slug).
- Extract facets from the content and reference appropriate authority IDs.
- Create or update authority entries if new entities are encountered.
- Stage summarized website sources under `_incoming/sources/` (these will be ingested into the Knowledgebase after the PR merges).
- Commit and push the branch, then open a Draft pull request (labeled to trigger the Copilot Organizer). The PR description might reference the original issue and any notes about sources used.
- This PR follows the same review and merge process as above.

Automated CI checks validate Researcher-generated PRs for quality – checking for formatting, obvious errors, missing metadata, or incorrect facet assignments. PRs that pass all checks are automatically merged.

## File Organization and Naming Conventions

- All article files reside under the `Compendium/` directory, organized into nested category subdirectories. Categories mirror broad domains (e.g., `Science/Physics/`, `History/WWII/`, `Technology/AI/`, etc.) that help organize and navigate the content. The Researcher agent determines the appropriate category path when creating articles.
- Filenames should roughly reflect the article title. Use singular form for classes. For disambiguation, use parentheses. For example:
  - Title: "OpenAI" → Filename: `Compendium/Technology/AI/OpenAI.md`
  - Title: "Mercury (planet)" → Filename: `Compendium/Science/Astronomy/Mercury (planet).md`
- If an article's title changes, it's preferred to rename the file to match (and create redirects if necessary). The `id` ULID in the front matter ensures that the Knowledgebase can still recognize it as the same article even if the file name changes.
- Non-article files:
  - `README.md`: Introduction and contribution guide for the repository.
  - Category `index.md` files: Generated indexes for each category directory (see below).

## Category Index Files

Each category directory may contain an `index.md` file that serves as a table of contents for articles within that category. These indexes help viewers on GitHub navigate articles within specific categories.

### Automatic Generation

Instead of maintaining these lists by hand (which can become error-prone as articles proliferate), we will automate their generation:

- A script (e.g., `scripts/build_index.py`) will scan each category directory in `Compendium/` for `.md` files.
- For each article in a category, it can extract the title (from the front matter or first heading) and perhaps the first line or a specified description field if available.
- The script will then generate or update an `index.md` file in each category directory with a list of articles in that category. For example, `Compendium/Technology/AI/index.md` might contain:

```markdown
# AI Articles

- [OpenAI](OpenAI.md) – *Artificial intelligence research organization.*
- [Machine Learning](machine-learning.md) – *Computational approach to learning from data.*
```

- Articles are listed alphabetically within each category index.

### Automation

We will integrate this script into the development workflow:

- **Pre-merge check:** A GitHub Action runs on pull requests to Gitopedia, executing the index generation script and checking if any category `index.md` files need an update. If updates are needed, the action automatically commits them to the PR branch.
- **Post-merge automation:** After a PR is merged, an automated action updates category index files on the main branch. This action runs the script and pushes commits updating the relevant category indexes automatically.

## Article Creation Workflow Summary

All articles are created by the Researcher agent. The workflow is as follows:

1. **Article Generation:** The Researcher agent generates content in Markdown, including the YAML front matter with a new ULID, title, and slug. The agent ensures the text follows style guidelines (neutral tone, uses Markdown syntax for formatting, includes citations for claims). The agent uses singular form for classes and parentheses for disambiguation.
2. **Draft PR & Staging:** The Researcher agent creates a branch and stages new/updated articles under `_incoming/` and summarized sources under `_incoming/sources/`, then opens a Draft PR labeled to trigger the Copilot Organizer.
3. **Copilot Organizer (Encyclopaedist):** An automated Copilot agent organizes the branch by relocating articles into `Compendium/` categories and validating front matter. It leaves sources in `_incoming/sources/` for the Knowledgebase to ingest post-merge. It pushes commits and marks the PR ready when complete.
4. **Continuous Integration Checks:** On the PR, automated checks will run (once set up). These include:
   - Markdown linting
   - Front matter validation (ULID, title, slug present)
   - Facet validation (required facets present, values resolve to authority IDs)
   - Category lint (≥1 topic, at least one of people/orgs/places/date_refs)
   - Canonical title check (no duplicate titles)
   - Category index generation script to ensure category `index.md` files are up-to-date
5. **Automated Review and Merge:** Automated CI checks validate the PR. If all checks pass, the PR is automatically merged. If issues are detected, the Researcher agent can be notified to apply fixes and update the PR.
6. **Post-Merge Automation:** The merge triggers the Knowledgebase update (indexing) and then the Website rebuild, as described earlier. Within a short time, the new article becomes available on the live site and in search results.

By following this structure, the Gitopedia repository will remain organized and easy to navigate, even as it scales to many articles. The combination of automated indexing, consistent metadata, authority-controlled facets, and governance rules will facilitate seamless integration with the other system components (Knowledgebase indexing and Website generation).
