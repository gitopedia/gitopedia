# Researcher Agent – Automated Content Creation

The Researcher is a Go-based AI agent designed to autonomously generate new Gitopedia articles. It monitors requests (for example, issues labeled "article request" in the Gitopedia repository), conducts research on the given topics, and produces well-structured Markdown articles complete with sources. Finally, it creates a pull request to add the articles to the Gitopedia repository.

## Strategy and Workflow

### Triggering Tasks

The Researcher looks for tasks in the form of GitHub issues or other triggers. The primary mode envisioned is:

- When a GitHub Issue is created in the Gitopedia repo requesting articles (using a specific template or label, e.g., "research needed"), the Researcher will pick it up. Each issue may contain a request for a range of topics, not just a single topic, allowing the Researcher to generate multiple articles from one issue.
- Alternatively, a maintainer could trigger the Researcher via a workflow dispatch or a scheduled job (for periodic content generation).

### Web Search & Retrieval

- It performs "Deep Research" by scraping web pages using a headless browser (Chrome/Chromedp).
- It extracts readable text from HTML content to provide rich context to the LLM.
- It avoids relying solely on search snippets, ensuring better accuracy.

### Writing the Article

The Researcher uses an AI language model (via OpenAI-compatible API) to compose the article. This involves:

- Formulating a prompt for the LLM that includes the article outline, the key points to cover (possibly gleaned from the sources), and instructions to produce Markdown with proper headings and maybe a references section.
- For example, the prompt might say: "Write a comprehensive encyclopedia article about 'Quantum Computing'. Include an introduction, key concepts, and a conclusion. Use Markdown formatting. When stating facts or statistics, reference credible sources in the format 【source】."
- The agent may generate the article in sections, especially if using retrieval-augmented generation: it could feed the LLM one source at a time and ask for summaries, then compile those into the final article to ensure accuracy.

### Including Citations

A critical aspect of the Researcher's output is citing sources. After gathering facts from various references, the agent should annotate the article with citations. The format can be the bracketed reference style used throughout Gitopedia (e.g., 【1†】 or 【source†Lx-Ly】). The agent will also compile a reference list at the end of the article or ensure the inline citations link to the sources. (One simple approach is to include the full URL in the parentheses of the reference in Markdown, or use reference footnotes.)

### Creating a Pull Request (Draft with `_incoming/`)

Once the article content is generated for all topics in the issue:

- For each topic, the Researcher assigns a unique ULID for the article and inserts it into the front matter (as id).
- It saves each article as a Markdown file under a branch-local staging directory `_incoming/` at the root of the branch (e.g., `research/issue-42-topics/_incoming/<slug>.md`). The filename is derived from the title (the agent should slugify the title to create the filename).
- It includes summarized website sources used during research alongside the articles in the branch (e.g., under `_incoming/sources/`), so new sources can be ingested into Knowledgebase.
- Using the GitHub API (via a library like PyGithub or through direct HTTP requests), the agent will:
  - Create a new git branch (e.g., `research/issue-42-topics`) on the Gitopedia repository.
  - Add all new article files and summarized sources under `_incoming/`.
  - Commit the changes with a message like `Stage articles in _incoming: Quantum Computing, Machine Learning, Neural Networks`.
  - Open a Draft Pull Request on Gitopedia's repository (labeled to trigger the Copilot Organizer), with a title referencing the articles (and perhaps "via Researcher bot") and a description that closes the original issue (e.g., "Closes #42. This PR drafts new articles on Quantum Computing, Machine Learning, and Neural Networks.").
  - The Copilot Organizer will refactor the branch (move files from `_incoming/` into proper `Compendium/` categories, validate front matter and facets, and prepare summarized sources for Knowledgebase ingestion), push commits back to the PR branch, and mark the PR ready when complete.

### Feedback Loop (Automated)

Automated CI and the Copilot Organizer perform validation and refactoring. If fixes are required, the Researcher can update the PR branch (or the Copilot Organizer can re-run after changes). Once the PR is ready and all checks pass, it is automatically merged. After merge, the normal content pipeline proceeds (Knowledgebase indexing, Website deployment), making the new article live.

## Implementation Details

### Environment and Tools

The Researcher agent runs as a Go binary, either locally or in a Docker container.

- **Containerized Execution**: Recommended for "Deep Research" capabilities (headless browser).
- **Local Execution**: Can run on a developer's machine if `chromium` is installed.

### Self-hosted Services and LLM Access

The Researcher supports both self-hosted and cloud providers:

- **Web Search**: Custom HTML scraping and headless browser fetching (via `chromedp`). No external API keys required for basic operation.
- **LLM via DeepSeek (self-hosted)**: The primary LLM is a self-hosted DeepSeek model served by Ollama.
- **LLM via OpenAI API (cloud)**: Can fallback to OpenAI by configuring `OPENAI_BASE_URL`.

Configuration is handled via `config/base.env` (defaults) and `.env` (overrides) files, or environment variables.

### Runtime

The Researcher is built as a standalone Go binary:

- **Build**: `go build -o researcher main.go`
- **Run**: `./researcher` (or via Docker Compose)
- **Configuration**: `config/base.env` contains all default settings; `.env` file only needs to override specific values (e.g., `GITHUB_TOKEN` or `GITHUB_APP_*`).

### Key Components of the Code

1. **Issue Fetching**: Uses GitHub API (PAT or App Auth) to find open issues with "research category" label.
2. **Topic Expansion**: Asks LLM to suggest topics based on the category.
3. **Deep Research**:
   - Performs web search (HTML scraping).
   - Fetches top results using headless browser (`chromedp`) to handle JS-heavy sites.
   - Extracts readable text.
4. **Drafting & Citations**: Uses LLM to write article with Markdown footnotes `[^1]` linking to sources.
5. **Pull Request**: Creates a PR with the new article and summarized sources in `_incoming/`.

### Example Pseudocode Flow

```python
issues = github_api.search_issues(repo="gitopedia", label="research", state="open")
for issue in issues:
    topics = parse_topics_from_issue(issue)  # e.g. ["Quantum Computing", "Machine Learning", "Neural Networks"]
    articles = []
    for topic in topics:
        info = perform_web_search(topic)
        outline = draft_outline(topic, info)
        article_md = call_language_model(topic, outline, info)
        article_md = ensure_citations(article_md, info)
        ulid = generate_ulid()  # unique ID for the article
        article_md = add_front_matter(article_md, ulid, topic)
        filename = make_filename(topic)
        articles.append((filename, article_md))
    branch_name = f"research/issue-{issue.number}"
    create_branch_and_commit(articles, branch=branch_name)
    pr = open_pull_request(branch_name, title=f"Add articles: {', '.join(topics)}", body=f"Auto-generated by Researcher\n\nCloses #{issue.number}")
```

*(The above is simplified. In reality, the `perform_web_search` and `call_language_model` would involve multiple steps and use external APIs.)*

### Configuration

The agent requires configuration such as:

- `SEARXNG_URL` and optional `SEARXNG_API_KEY` for web search aggregation.
- `OPENAI_API_KEY` for OpenAI cloud access.
- `OPENAI_BASE_URL` to point the OpenAI SDK at a self-hosted, OpenAI-compatible DeepSeek endpoint (leave default for OpenAI cloud).
- `OPENAI_MODEL` and/or `DEEPSEEK_MODEL` to select default models for generation.
- `GITHUB_TOKEN` with permissions to push to Gitopedia.
- Search API keys (e.g., Bing API key) to fetch web results.
- `AWS_ACCOUNT_ID`, `AWS_REGION`, and ECR repository URI for image publishing (CI).

These will be provided via environment variables or secret management in the CI workflow.

### Setup Notes (Self-hosted)

- Provision a self-hosted SearXNG instance (Docker is recommended) and expose it internally to the Researcher runtime. Lock down engines, set sane timeouts, and configure a rate limit.
- Provision a self-hosted DeepSeek model server with an OpenAI-compatible API (e.g., via vLLM or a compatible gateway). Expose it internally and set `OPENAI_BASE_URL` to that endpoint in the Researcher environment. Allocate sufficient GPU/CPU and configure context/window to match article sizes.
- Ensure both services are reachable from the Researcher runtime (firewall rules, VPC, or Docker network as appropriate).
- Create an ECR repository and grant CI permissions to authenticate and push images. Use OIDC or scoped AWS credentials in CI to log in to ECR.

### GitHub CLI Requirements

The Researcher relies on the host machine's GitHub CLI (`gh`) for invoking the Encyclopaedist agent after creating Draft PRs:

1. **Install gh CLI**: https://cli.github.com/
2. **Authenticate**: Run `gh auth login` and complete the OAuth flow
3. **Verify**: Run `gh auth status` to confirm authentication

The Researcher uses the authenticated `gh` session to:
- Create and manage PRs
- Invoke the Encyclopaedist Copilot agent for article organization

**Note:** The `gh` CLI must be installed and authenticated on the machine running the Researcher. Docker deployments should mount the host's `~/.config/gh` directory or use `gh auth token` for non-interactive authentication.

```bash
# Non-interactive authentication (for CI/containers)
echo "$GH_TOKEN" | gh auth login --with-token
```

### Encyclopaedist Integration

The Researcher uses a non-blocking workflow:

1. **Before starting research:**
   - Check all open PRs for any that are ready to merge
   - Merge ready PRs and close their tracking issues
   
2. **Select a research task:**
   - Filter out issues that already have open PRs
   - Pick a random issue from the remaining available issues
   
3. **After completing research:**
   - Create Draft PR with articles in `_incoming/`
   - **Invoke the Encyclopaedist** by posting a `@copilot` comment
   - Check again for ready PRs and merge them
   - **Continue immediately** - does not wait for the new PR

The Encyclopaedist agent (triggered asynchronously) handles:
- Moving articles from `_incoming/` to proper `Compendium/<Category>/` paths
- Validating front matter (ULID, title, slug, tags)
- Flagging authority reference issues
- Removing debug artifacts
- Marking the PR ready for review

On the next Researcher run, the now-ready PR will be merged automatically.

See [../agents/README.md](../agents/README.md) for full agent documentation.

## Future Improvements

- **Better Source Integration**: In Phase 2, the Researcher could use the Knowledgebase's index to find if Gitopedia already has related content to avoid redundancy or to link articles together.
- **Multi-step Reasoning**: Use a chain-of-thought prompting where the agent first lists what it wants to find out, searches for each sub-question, then composes the article. This can improve factual accuracy.
- **Plagiarism Check**: Incorporate a step to ensure the generated text isn't copying large verbatim sections from sources (except for necessary quotes). The agent should rephrase information in its own words.
- **Updating Existing Articles**: The Researcher could also be used to periodically scan articles that might need updating (for example, if an article's sources are outdated, or a "update request" label is applied). It can then gather new info and suggest edits via PR.
- **Multiple Agents or Parallelism**: For scaling, we might run multiple Researcher jobs in parallel if there are many requests. Each would handle a different issue (to be careful, coordinate by locking issues when one agent starts working on them).
- **Human-in-the-loop**: Provide a mechanism for a human expert to review or augment the prompt or outline before the AI writes the article, to improve quality.

## Conclusion

With the Researcher agent in place, the Gitopedia project grows its content autonomously. The agent handles research, drafting, staging, and submission. Automated CI merges successful PRs and triggers downstream indexing and site updates without manual intervention.
