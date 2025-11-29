# Gitopedia Custom Agents

This directory contains documentation and reference copies of GitHub Copilot custom agents used in the Gitopedia project.

## Available Agents

### Encyclopaedist

**File:** [`encyclopaedist.agent.md`](./encyclopaedist.agent.md)

The Encyclopaedist is a content organization specialist that processes Draft Pull Requests created by the Researcher agent. It:

- **Organizes articles** from `_incoming/` into the proper `Compendium/<Category>/` structure
- **Validates front matter** (ULID, title, slug, tags)
- **Flags authority references** that need attention
- **Leaves sources untouched** for Knowledge Base ingestion
- **Removes debug artifacts** (`_debug/` directories)
- **Finalizes PRs** by marking them ready for review

## Deployment

The actual agent files that GitHub Copilot uses are stored in the private `.github-private` repository under `agents/`. The files in this directory are reference copies for documentation purposes.

To deploy or update an agent:

1. Edit the reference copy here
2. Copy to `.github-private/agents/<agent-name>.agent.md`
3. The agent becomes available for use in PR comments

## Invoking Agents

### Via GitHub CLI (Recommended for Automation)

The Researcher agent can invoke the Encyclopaedist after creating a Draft PR:

```bash
# Ensure gh is authenticated on the host machine
gh auth status

# Invoke Copilot with context
gh copilot explain "Process this Draft PR using the Encyclopaedist agent to organize articles"
```

**Note:** The host machine running the Researcher must have:
- `gh` CLI installed
- Valid GitHub authentication (`gh auth login`)
- Access to the target repository

### Via PR Comments (Interactive)

In any PR, you can invoke Copilot with context:

```
@copilot Please organize these articles following the Encyclopaedist guidelines
```

Copilot will use the agent instructions from `.github-private/agents/` when processing requests.

## Agent Development

When creating new agents:

1. Use the `.agent.md` extension for GitHub to recognize the file
2. Include YAML front matter with `name` and `description`
3. Be specific about what the agent should and should NOT do
4. Include example workflows
5. Document commit message conventions

