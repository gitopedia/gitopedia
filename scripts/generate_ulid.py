#!/usr/bin/env python3
"""
Generate a ULID for new articles.

Usage:
  python scripts/generate_ulid.py

Notes:
- Prefers the 'ulid-py' package if installed:
    pip install ulid-py
- As an alternative, you can also run:
    npx ulid
"""
import sys

try:
    import ulid  # type: ignore
except Exception:
    ulid = None


def main() -> int:
    if ulid is None:
        sys.stderr.write(
            "ulid-py is not installed.\n"
            "Install with: pip install ulid-py\n"
            "Or use Node alternative: npx ulid\n"
        )
        return 1
    u = ulid.new()
    print(u.str)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())


