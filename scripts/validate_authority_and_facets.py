#!/usr/bin/env python3
import json
import re
import sys
from pathlib import Path

try:
    import yaml  # type: ignore
except Exception:
    yaml = None

REPO_ROOT = Path(__file__).resolve().parents[1]
AUTHORITY_DIR = REPO_ROOT / "authority"
COMPENDIUM_DIR = REPO_ROOT / "Compendium"

ULID_RE = re.compile(r"^[0-9A-HJKMNP-TV-Z]{26}$")


def error(msg: str) -> None:
    print(f"ERROR: {msg}", file=sys.stderr)


def load_authority_files() -> tuple[bool, dict[str, dict]]:
    ok = True
    entities: dict[str, dict] = {}
    if not AUTHORITY_DIR.exists():
        # Not critical if not present yet (early bootstrapping), return ok
        return True, entities
    for filename in ["people.json", "orgs.json", "places.json", "topics.json"]:
        path = AUTHORITY_DIR / filename
        if not path.exists():
            # Allow empty authority during bootstrapping
            continue
        try:
            data = json.loads(path.read_text(encoding="utf-8"))
            if not isinstance(data, list):
                error(f"{path} must be a JSON array")
                ok = False
                continue
            for entry in data:
                if not isinstance(entry, dict):
                    error(f"{path} contains a non-object entry")
                    ok = False
                    continue
                ent_id = entry.get("id")
                if not ent_id or not isinstance(ent_id, str):
                    error(f"{path} entry missing string 'id'")
                    ok = False
                    continue
                if ent_id in entities:
                    error(f"Duplicate authority id '{ent_id}' across authority files")
                    ok = False
                else:
                    entities[ent_id] = entry
        except json.JSONDecodeError as e:
            error(f"{path} is not valid JSON: {e}")
            ok = False
    return ok, entities


def parse_front_matter(md_path: Path) -> dict | None:
    text = md_path.read_text(encoding="utf-8", errors="ignore")
    if not text.startswith("---"):
        return None
    parts = text.split("\n")
    # find closing '---'
    try:
        end_idx = parts[1:].index("---") + 1
    except ValueError:
        return None
    fm_text = "\n".join(parts[1:end_idx])
    if yaml is None:
        # If PyYAML isn't available, skip deep parsing (handled in CI with dependency)
        return {}
    try:
        data = yaml.safe_load(fm_text) or {}
        if not isinstance(data, dict):
            return {}
        return data
    except Exception:
        return {}


def validate_ulid(val: str) -> bool:
    return bool(ULID_RE.match(val.upper()))


def validate_articles() -> bool:
    ok = True
    if not COMPENDIUM_DIR.exists():
        # nothing to validate yet
        return True
    seen_ids: set[str] = set()
    for md_path in COMPENDIUM_DIR.rglob("*.md"):
        if md_path.name.lower() == "index.md":
            continue
        fm = parse_front_matter(md_path)
        if fm is None:
            error(f"{md_path}: missing or malformed YAML front matter")
            ok = False
            continue
        # if PyYAML not available locally, we won't block; CI will install and re-run
        if fm == {} and yaml is None:
            continue
        # Required fields: id, title, slug
        for field in ("id", "title", "slug"):
            if field not in fm or not isinstance(fm.get(field), str) or not fm.get(field):
                error(f"{md_path}: missing required front matter field '{field}'")
                ok = False
        art_id = fm.get("id")
        if isinstance(art_id, str):
            if not validate_ulid(art_id):
                error(f"{md_path}: 'id' is not a valid ULID format")
                ok = False
            if art_id in seen_ids:
                error(f"{md_path}: duplicate article id '{art_id}'")
                ok = False
            seen_ids.add(art_id)
        # Basic slug sanity
        slug = fm.get("slug")
        if isinstance(slug, str) and (" " in slug or slug != slug.strip()):
            error(f"{md_path}: 'slug' should be a URL-friendly string without spaces")
            ok = False
    return ok


def main() -> int:
    a_ok, _ = load_authority_files()
    c_ok = validate_articles()
    return 0 if (a_ok and c_ok) else 1


if __name__ == "__main__":
    sys.exit(main())


