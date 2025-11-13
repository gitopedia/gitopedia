#!/usr/bin/env python3
"""
Generate a GitHub App installation access token.

Usage:
    python3 scripts/generate-app-token.py <APP_ID> <PRIVATE_KEY_PATH> <INSTALLATION_ID>

Example:
    python3 scripts/generate-app-token.py 123456 /path/to/private-key.pem 98765432
"""

import sys
import jwt
import time
import requests
from pathlib import Path


def generate_jwt(app_id: str, private_key_path: str) -> str:
    """Generate a JWT token for GitHub App authentication."""
    private_key = Path(private_key_path).read_text()
    
    now = int(time.time())
    payload = {
        "iat": now - 60,  # Issued at (1 minute ago to account for clock skew)
        "exp": now + 600,  # Expires in 10 minutes
        "iss": app_id  # Issuer (App ID)
    }
    
    token = jwt.encode(payload, private_key, algorithm="RS256")
    return token


def get_installation_token(jwt_token: str, installation_id: str) -> str:
    """Get an installation access token using a JWT token."""
    url = f"https://api.github.com/app/installations/{installation_id}/access_tokens"
    headers = {
        "Authorization": f"Bearer {jwt_token}",
        "Accept": "application/vnd.github.v3+json",
        "X-GitHub-Api-Version": "2022-11-28"
    }
    
    response = requests.post(url, headers=headers)
    response.raise_for_status()
    
    data = response.json()
    return data["token"]


def main():
    if len(sys.argv) != 4:
        print(__doc__, file=sys.stderr)
        sys.exit(1)
    
    app_id = sys.argv[1]
    private_key_path = sys.argv[2]
    installation_id = sys.argv[3]
    
    try:
        # Generate JWT
        jwt_token = generate_jwt(app_id, private_key_path)
        
        # Get installation token
        installation_token = get_installation_token(jwt_token, installation_id)
        
        print(installation_token)
        
    except FileNotFoundError:
        print(f"Error: Private key file not found: {private_key_path}", file=sys.stderr)
        sys.exit(1)
    except requests.exceptions.HTTPError as e:
        print(f"Error: Failed to get installation token: {e}", file=sys.stderr)
        if hasattr(e.response, 'text'):
            print(f"Response: {e.response.text}", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()

