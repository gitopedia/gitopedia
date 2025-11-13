# GitHub App Setup for Cross-Repository Automation

This guide documents the setup of the `gitopedia-bot` GitHub App for cross-repository workflow triggers.

## Overview

The `gitopedia-bot` GitHub App is used to trigger workflows across repositories via `repository_dispatch` events. This enables the automated pipeline:
- Gitopedia → Knowledgebase (on content updates)
- Knowledgebase → Website (on index updates)

## App Permissions

The app requires the following repository permissions:

- **Contents**: Read and Write
  - Future: Create tags/releases for versioning
- **Issues**: Read and Write
  - Researcher agent: Search/read issues, comment, close, create new issues
- **Pull requests**: Read and Write
  - Researcher agent: Create Draft PRs, apply labels
  - Copilot Organizer: Read PRs, update status, apply/remove labels
- **Actions**: Write
  - Trigger workflows via `repository_dispatch` events

## Installation Steps

### 1. Install the App on Repositories

Install `gitopedia-bot` on all three repositories:
- `gitopedia/gitopedia`
- `gitopedia/knowledge-base`
- `gitopedia/website`

**Via GitHub CLI:**
```bash
# Get installation ID (after installing via UI)
gh api orgs/gitopedia/installations

# Install on repositories (if not already done via UI)
gh api \
  -X POST \
  /app/installations/{installation_id}/access_tokens \
  -f repositories[]=gitopedia \
  -f repositories[]=knowledge-base \
  -f repositories[]=website
```

**Via GitHub UI:**
1. Go to: https://github.com/organizations/gitopedia/settings/apps
2. Click on `gitopedia-bot`
3. Click "Configure" → "Repository access"
4. Select "Only select repositories" and choose the three repositories
5. Click "Save"

### 2. Generate Installation Access Token

The installation access token is used by CI workflows to authenticate as the app when calling the GitHub API.

**Using a script (recommended):**

Create a script `scripts/generate-app-token.sh`:

```bash
#!/bin/bash
# Generate GitHub App installation token
# Usage: ./generate-app-token.sh <APP_ID> <PRIVATE_KEY_PATH> <INSTALLATION_ID>

APP_ID=$1
PRIVATE_KEY_PATH=$2
INSTALLATION_ID=$3

if [ -z "$APP_ID" ] || [ -z "$PRIVATE_KEY_PATH" ] || [ -z "$INSTALLATION_ID" ]; then
  echo "Usage: $0 <APP_ID> <PRIVATE_KEY_PATH> <INSTALLATION_ID>"
  exit 1
fi

# Generate JWT
JWT=$(python3 <<EOF
import jwt
import time
from pathlib import Path

app_id = "${APP_ID}"
private_key = Path("${PRIVATE_KEY_PATH}").read_text()

now = int(time.time())
payload = {
    "iat": now - 60,
    "exp": now + 600,
    "iss": app_id
}

token = jwt.encode(payload, private_key, algorithm="RS256")
print(token)
EOF
)

# Get installation token
INSTALLATION_TOKEN=$(curl -s -X POST \
  -H "Authorization: Bearer ${JWT}" \
  -H "Accept: application/vnd.github.v3+json" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  "https://api.github.com/app/installations/${INSTALLATION_ID}/access_tokens" \
  | jq -r '.token')

echo "Installation token:"
echo "${INSTALLATION_TOKEN}"
```

**Manual steps:**

1. Generate a JWT token for the app (valid for 10 minutes):
   ```python
   import jwt
   import time
   from pathlib import Path
   
   app_id = "YOUR_APP_ID"
   private_key = Path("path/to/private-key.pem").read_text()
   
   now = int(time.time())
   payload = {
       "iat": now - 60,
       "exp": now + 600,
       "iss": app_id
   }
   
   token = jwt.encode(payload, private_key, algorithm="RS256")
   print(token)
   ```

2. Use the JWT to get an installation token:
   ```bash
   curl -X POST \
     -H "Authorization: Bearer <JWT_TOKEN>" \
     -H "Accept: application/vnd.github.v3+json" \
     "https://api.github.com/app/installations/<INSTALLATION_ID>/access_tokens"
   ```

3. Extract the `token` field from the response.

### 3. Store App Credentials as Organization Secrets

Instead of storing short-lived installation tokens, we store the **GitHub App credentials** as organization secrets and let workflows generate a fresh installation token on every run.

Create the following **organization secrets** (once):

- `GITOPEDIA_APP_ID` – the numeric App ID of `gitopedia-bot`
- `GITOPEDIA_APP_PRIVATE_KEY` – the full PEM contents of the app's private key

**Via GitHub UI:**
1. Go to: https://github.com/organizations/gitopedia/settings/secrets/actions
2. Click "New organization secret"
3. Name: `GITOPEDIA_APP_ID`
4. Value: the App ID (e.g. `123456`)
5. Repository access: "All repositories" (or at least the three Gitopedia repos)
6. Repeat for `GITOPEDIA_APP_PRIVATE_KEY`, pasting the PEM contents of the private key

**Note:** Installation tokens expire after 1 hour, but the App ID and private key are long‑lived. Workflows will derive fresh tokens automatically using these secrets, so no manual regeneration is required.

## Usage in Workflows

Workflows use the GitHub App credentials to generate an installation token and then trigger `repository_dispatch` events. For example, in the Gitopedia repo:

```yaml
jobs:
  dispatch-kb:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Generate GitHub App token
        id: app-token
        uses: tibdex/github-app-token@v2
        with:
          app_id: ${{ secrets.GITOPEDIA_APP_ID }}
          private_key: ${{ secrets.GITOPEDIA_APP_PRIVATE_KEY }}

      - name: Dispatch to Knowledgebase (repository_dispatch)
        env:
          KB_TOKEN: ${{ steps.app-token.outputs.token }}
        run: |
          payload=$(jq -n --arg sha "${GITHUB_SHA}" '{event_type:"content-updated", client_payload:{gitopedia_sha:$sha}}')
          curl -s -X POST \
            -H "Accept: application/vnd.github+json" \
            -H "Authorization: token ${KB_TOKEN}" \
            https://api.github.com/repos/gitopedia/knowledge-base/dispatches \
            -d "${payload}"
```

## Troubleshooting

### Token Expired
If you get `401 Unauthorized`, the installation token may have expired. Regenerate it using the steps above.

### App Not Installed
If you get `404 Not Found`, verify the app is installed on the target repository:
```bash
gh api repos/gitopedia/knowledge-base
```

### Insufficient Permissions
If you get `403 Forbidden`, check that the app has the required permissions (Actions: Write) and is installed on the repository.

## Future Enhancements

- **Auto-regenerate tokens in workflows**: Store the app's private key and App ID as organization secrets, and add a workflow step to generate fresh tokens automatically (valid for 1 hour). This eliminates the need for manual token regeneration.

