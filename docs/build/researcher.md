# Researcher Agent – Automated Content Creation

The Researcher is a Python-based AI agent designed to autonomously generate new Gitopedia articles. It monitors requests (for example, issues labeled "article request" in the Gitopedia repository), conducts research on the given topics, and produces well-structured Markdown articles complete with sources. Finally, it creates a pull request to add the articles to the Gitopedia repository.

## Strategy and Workflow

### Triggering Tasks

The Researcher looks for tasks in the form of GitHub issues or other triggers. The primary mode envisioned is:

- When a GitHub Issue is created in the Gitopedia repo requesting articles (using a specific template or label, e.g., "research needed"), the Researcher will pick it up. Each issue may contain a request for a range of topics, not just a single topic, allowing the Researcher to generate multiple articles from one issue.
- Alternatively, a maintainer could trigger the Researcher via a workflow dispatch or a scheduled job (for periodic content generation).

### Information Gathering

Upon receiving a set of topics from an issue, the Researcher gathers background information for each topic:

- It can perform web searches (using an API like Bing Web Search or Google Custom Search) to find relevant reliable sources.
- It may also query the existing Knowledgebase (to avoid duplicating content or to find related articles for context).
- In future enhancements, the agent might use more sophisticated tools (for example, a vector search on existing content, or accessing specific databases).

### Writing the Article

The Researcher uses an AI language model to compose the article. This typically involves:

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

The Researcher agent can run as a GitHub Action in the Researcher repository or as an external service:

- **As a GitHub Action**: We can set up a workflow that triggers on a schedule (e.g., daily) or on repository dispatch. This workflow would run the Researcher script inside a runner that has internet access. We will store necessary secrets (API keys for search/LLM and a GitHub PAT for committing) as encrypted secrets in the repo.
- **As an external long-running bot**: Alternatively, one could run the Researcher on a server or as a GitHub App that listens to events. For simplicity, using GitHub Actions is a good starting point.

### Self-hosted Services and LLM Access

The Researcher supports both self-hosted and cloud providers:

- **Web Search via SearXNG (self-hosted)**: The agent queries a self-hosted SearXNG instance for aggregation of search results across multiple engines. The SearXNG endpoint and optional API key are provided via environment variables. The client enforces timeouts and rate limits.
- **LLM via DeepSeek (self-hosted, OpenAI-compatible API)**: The primary LLM can be a self-hosted DeepSeek model exposed behind an OpenAI-compatible REST API (e.g., served by vLLM or another gateway). The agent uses the Python OpenAI SDK with a custom `base_url` to talk to this endpoint.
- **LLM via OpenAI API (cloud)**: The agent also has access to the OpenAI API and includes the Python OpenAI SDK in its toolkit. It can fall back to OpenAI when the self-hosted model is unavailable or when specific models/features are needed.

Recommended provider order:
1) DeepSeek (self-hosted) for primary generation
2) OpenAI (cloud) as fallback for reliability/specific tasks

### Containerization, Registry, and Runtime

The Researcher is packaged and deployed as a container:

- **Docker Image (multi-stage, distroless runtime)**:
  - Build stage uses a Python base image to install dependencies and prepare artifacts (e.g., wheels/venv).
  - Final stage uses a distroless Python runtime image for minimal attack surface and size.
  - The entrypoint runs the Researcher process (CLI/service) with configuration via environment variables.
- **Amazon ECR**:
  - An ECR repository is created to host the Researcher images.
  - CI pipeline authenticates to ECR, builds the image, runs tests, tags with commit SHA and `latest`, and pushes.
- **Deployment**:
  - Deployment is managed by an external GitOps system (e.g., ArgoCD). The container image reference (ECR URI + tag) is updated as part of the delivery pipeline. This document references that approach at a high level without detailing the ArgoCD configuration.

### Key Components of the Code

1. **Issue Fetching**: Use the GitHub API to search for open issues in Gitopedia with the designated label (e.g., "research"). Pick one (or iterate through all) to process.
2. **Topic Parsing**: Extract the range of topics and any additional instructions from the issue. The issue template might have fields like "Topics:" (which may contain a list or range of topics) and "Specific questions to answer:" which the agent can use to guide research. The Researcher should parse the issue to identify all topics requested, which could be specified as a list, a range, or described in natural language.
3. **Web Search & Retrieval**: For each topic in the range, perform a web search. This can be done via API. Parse the top results and retrieve content (this might involve sending HTTP requests to those pages and extracting text). The agent should filter out unreliable sources; prioritize academic articles, reputable websites, etc.
4. **Drafting Content**: Construct a draft outline. For example: Introduction, Subtopics (each as a section), perhaps a "References" section.
5. **Calling LLM**: Use an AI model to flesh out the draft. This could be done in one go (one prompt that includes outline and some sourced info) or iteratively (section by section). Provide the model with the relevant facts and ask it to incorporate them with citations. (The Researcher might include the URLs or reference names in the prompt so the model can place them accordingly).
6. **Post-processing**: The generated text may need cleaning. The agent should ensure:
   - The Markdown syntax is correct (all headings start with `#`, lists and code blocks are properly formatted).
   - Citations placeholders match actual references. It might replace raw URLs in the output with the 【†】 style and compile those into footnotes.
   - The front matter (id, title, tags) is prepended.
7. **Error Handling**: If the LLM output is not satisfactory (too short, or missing key info), the agent can retry with a refined prompt. If it fails multiple times, the agent might leave a comment on the issue stating it couldn't complete the task for human intervention.

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

## Future Improvements

- **Better Source Integration**: In Phase 2, the Researcher could use the Knowledgebase's index to find if Gitopedia already has related content to avoid redundancy or to link articles together.
- **Multi-step Reasoning**: Use a chain-of-thought prompting where the agent first lists what it wants to find out, searches for each sub-question, then composes the article. This can improve factual accuracy.
- **Plagiarism Check**: Incorporate a step to ensure the generated text isn't copying large verbatim sections from sources (except for necessary quotes). The agent should rephrase information in its own words.
- **Updating Existing Articles**: The Researcher could also be used to periodically scan articles that might need updating (for example, if an article's sources are outdated, or a "update request" label is applied). It can then gather new info and suggest edits via PR.
- **Multiple Agents or Parallelism**: For scaling, we might run multiple Researcher jobs in parallel if there are many requests. Each would handle a different issue (to be careful, coordinate by locking issues when one agent starts working on them).
- **Human-in-the-loop**: Provide a mechanism for a human expert to review or augment the prompt or outline before the AI writes the article, to improve quality.

## TODO Summary for Researcher Implementation

- **TODO**: Set up the Researcher repository with the necessary dependencies (OpenAI API client, requests for web scraping, PyGithub, etc.). Write a `researcher.py` script that encapsulates the logic to take an issue and produce a PR.
- **TODO**: Implement issue retrieval and filtering. The script should be able to parse issues to extract a range of topics (which may be specified as a list, range, or in natural language). It should safely ignore or defer issues that are too broad or lack enough information, possibly by commenting on them asking for clarification.
- **TODO**: Integrate a search step. For MVP, using a simple approach like calling the Wikipedia API for each topic or using one general web search query per topic and picking a couple of top results might suffice. Parse those results (maybe using an HTML to text extractor).
- **TODO**: Integrate the LLM. Start with a known API (if available in the environment). Design the prompt carefully to get Markdown output. Test with a small example (e.g., ask it to write short articles on a few simple topics) and adjust the prompt formatting.
- **TODO**: Handle Git operations. Use a GitHub personal access token (stored as `GH_TOKEN`) to authenticate git or GitHub API operations. Ensure the token has repo access to Gitopedia. Use a high-level library or shell out to git. Steps: clone Gitopedia (or use GitHub API to create file via POST – but cloning and pushing might be easier for multiple files), create a new branch, add all article files for the topics in the issue, commit, push, then use GitHub API to open PR. This requires careful error handling (e.g., if branch name already exists, choose another).
- **TODO**: Testing the full loop. Simulate a run: create a dummy issue in a test environment with multiple topics, run the Researcher (perhaps locally first), and see that it creates a PR with all articles. Check the content quality. Refine as needed.
- **TODO**: Schedule or trigger the Researcher. For now, set up a GitHub Action in the Researcher repo to run daily and process at most N issues per run (to avoid spamming PRs). Mark issues as done or add a comment when a PR is opened so the agent doesn't duplicate work on them. Note that each issue may generate multiple articles, so consider this when setting limits.
- **TODO**: Resource management. Ensure that the Researcher's run has appropriate timeouts and limits (so it doesn't hang if a website is not responding, or if the LLM call takes too long). Use try/except and timeouts for network calls.
- **TODO**: Logging and Monitoring. The Researcher should log its actions (which issue it's working on, what sources it found, etc.) either to the console (for the Action logs) or back to the issue (posting a comment like "Researcher is working on this issue..."). This helps in debugging if it fails or produces suboptimal results.

## Conclusion

With the Researcher agent in place, the Gitopedia project grows its content autonomously. The agent handles research, drafting, staging, organization (via Copilot), validation, and submission. Automated CI merges successful PRs and triggers downstream indexing and site updates without manual intervention.
