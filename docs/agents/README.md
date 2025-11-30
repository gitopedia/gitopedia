# Gitopedia Agent Documentation

This directory contains documentation for the article organization logic used in the Gitopedia project.

## Automated Article Organization

The Researcher agent now includes built-in article organization, replacing the need for a separate Copilot agent. After creating a Draft PR with articles in `_incoming/`, the Researcher automatically:

- **Categorizes articles** using LLM-based analysis of tags and content
- **Moves articles** from `_incoming/` to `Compendium/<Category>/` paths
- **Validates front matter** (ULID, title, slug, tags)
- **Removes debug artifacts** (`_debug/` directories)
- **Marks PRs ready** for review

## Reference: Encyclopaedist Agent

**File:** [`encyclopaedist.agent.md`](./encyclopaedist.agent.md)

This file contains the original Encyclopaedist agent definition. While no longer used automatically (the Researcher handles organization directly), it serves as:

- Documentation of the organization logic and rules
- Reference for category inference guidelines
- Template for future Copilot agent development

## Manual Invocation (Optional)

If you need to manually organize articles using Copilot Chat in your IDE:

1. Open the PR in your IDE
2. Use Copilot Chat with the Encyclopaedist agent instructions
3. Follow the guidelines in `encyclopaedist.agent.md`

This is optional - the Researcher handles organization automatically.

