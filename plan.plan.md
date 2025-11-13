<!-- e024a905-d96a-48c3-b36a-a00b18d4229d ede5b878-144f-4859-9351-194cde7a4907 -->
# Complete Cross-Repo CI/CD Dispatch Chain

## Overview

Currently, the cross-repo dispatch chain is partially implemented but not functional. This plan will complete the automation so that:

1. Gitopedia content updates → trigger Knowledge-base index build
2. Knowledge-base index build success → trigger Website rebuild
3. All workflows respond to `repository_dispatch` events

## Implementation Steps

### Step 1: Configure GitHub App for Cross-Repo Dispatch

**Status**: GitHub App `gitopedia-bot` created with required permissions

**File**: Organization secrets (via GitHub UI or CLI)

**Action**: Set up authentication for `repository_dispatch` calls using GitHub App:

1. **Install the App**: Install `gitopedia-bot` on all three repositories:
   - `gitopedia/gitopedia`
   - `gitopedia/knowledge-base`
   - `gitopedia/website`

2. **Store App Credentials as Secrets**: Instead of storing short-lived installation tokens, store the GitHub App credentials as organization secrets:
   - `GITOPEDIA_APP_ID` – the numeric App ID
   - `GITOPEDIA_APP_PRIVATE_KEY` – the PEM contents of the private key

   These are long‑lived and only need to be set once.

3. **Generate Tokens in Workflows**: Use a GitHub Action (e.g. `tibdex/github-app-token@v2`) in each workflow to generate a fresh installation token on every run:
   ```yaml
   - name: Generate GitHub App token
     id: app-token
     uses: tibdex/github-app-token@v2
     with:
       app_id: ${{ secrets.GITOPEDIA_APP_ID }}
       private_key: ${{ secrets.GITOPEDIA_APP_PRIVATE_KEY }}
   ```

   Subsequent steps use `steps.app-token.outputs.token` as the `Authorization: token` when calling the GitHub API.

**Result**: Installation tokens are short‑lived and automatically regenerated per run; no manual rotation or PATs are required.

**Documentation**: See `docs/setup/github-app.md` for detailed setup instructions.

**App Permissions Required**:
- **Contents**: Read and Write (for future tag/release operations)
- **Issues**: Read and Write (for Researcher agent operations)
- **Pull requests**: Read and Write (for Researcher and Copilot Organizer operations)
- **Actions**: Write (to trigger workflows via `repository_dispatch`)

### Step 2: Add repository_dispatch Trigger to Knowledge-base Workflow

**File**: `knowledge-base/.github/workflows/build-index.yml`

**Changes**:

- Add `repository_dispatch` to the `on:` section with event type `content-updated`
- Update the `deploy` job condition to also run on `repository_dispatch` events
- Extract `gitopedia_sha` from `github.event.client_payload.gitopedia_sha` if available
- Use the SHA to checkout the specific Gitopedia commit (currently uses latest from main)

**Code location**: Lines 3-6 (trigger section), line 44 (deploy condition)

### Step 3: Add Dispatch from Knowledge-base to Website

**File**: `knowledge-base/.github/workflows/build-index.yml`

**Changes**:

- Add a new job `dispatch-website` that runs after successful `deploy` job
- Use the same pattern as Gitopedia dispatch (curl to GitHub API)
- Dispatch event type: `rebuild-site`
- Payload: Include `gitopedia_sha` and optionally `kb_index_version` or S3 path
- Use `KB_DISPATCH_TOKEN` secret (same token can be reused)

**Code location**: After line 61 (new job after deploy)

### Step 4: Add repository_dispatch Trigger to Website Workflow

**File**: `website/.github/workflows/site-build.yml`

**Changes**:

- Add `repository_dispatch` to the `on:` section with event type `rebuild-site`
- Update the `deploy` job condition to also run on `repository_dispatch` events
- Extract `gitopedia_sha` from `github.event.client_payload.gitopedia_sha` if available
- Use the SHA to checkout the specific Gitopedia commit

**Code location**: Lines 3-6 (trigger section), line 43 (deploy condition)

### Step 5: Update Gitopedia Dispatch Workflow

**File**: `gitopedia/.github/workflows/content-dispatch.yml`

**Changes**:

- Remove the "placeholder" from the name
- Optionally: Make the dispatch step fail if token is missing (instead of silently skipping) to surface configuration issues
- Or keep the skip behavior but add a warning

**Code location**: Line 1 (name), lines 17-20 (token check)

### Step 6: Test the Complete Chain

**Action**: Manual testing

- Make a test commit to Gitopedia main branch
- Verify Knowledge-base workflow triggers and builds index
- Verify Website workflow triggers and rebuilds site
- Check that all three workflows complete successfully

## Files to Modify

1. `knowledge-base/.github/workflows/build-index.yml`

- Add `repository_dispatch` trigger
- Add dispatch job to Website
- Update deploy condition

2. `website/.github/workflows/site-build.yml`

- Add `repository_dispatch` trigger
- Update deploy condition

3. `gitopedia/.github/workflows/content-dispatch.yml`

- Update name (remove "placeholder")
- Optionally improve error handling

## Dependencies

- GitHub App `gitopedia-bot` installed on all three repositories
- Organization secrets `GITOPEDIA_APP_ID` and `GITOPEDIA_APP_PRIVATE_KEY` configured
- All three repositories must be in the same GitHub organization
- Workflows must have appropriate permissions (already configured via OIDC for AWS)

## Testing Strategy

1. Test Gitopedia → Knowledge-base dispatch
2. Test Knowledge-base → Website dispatch
3. Test end-to-end: Gitopedia push → Knowledge-base build → Website deploy
4. Verify S3 uploads and CloudFront invalidations occur
5. Check that specific commit SHAs are used when provided in payload

_Last tested: pending automated chain run triggered by a commit to gitopedia/main._

### To-dos

- [x] Install GitHub App `gitopedia-bot` on all three repositories (gitopedia, knowledge-base, website)
- [x] Create org secrets `GITOPEDIA_APP_ID` and `GITOPEDIA_APP_PRIVATE_KEY`
- [x] Add repository_dispatch trigger to knowledge-base/.github/workflows/build-index.yml with event type content-updated
- [x] Add dispatch-website job to knowledge-base workflow that triggers Website rebuild after successful index deploy
- [x] Add repository_dispatch trigger to website/.github/workflows/site-build.yml with event type rebuild-site
- [x] Update gitopedia/.github/workflows/content-dispatch.yml to remove 'placeholder' from name and improve error handling
- [ ] Test complete dispatch chain: Gitopedia push → Knowledge-base build → Website deploy

