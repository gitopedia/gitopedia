#!/usr/bin/env python3
import os
from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parents[1]
COMPENDIUM_DIR = REPO_ROOT / "Compendium"

def is_markdown_article(path: Path) -> bool:
    if not path.is_file():
        return False
    if path.name.lower() == "index.md":
        return False
    return path.suffix.lower() == ".md"

def generate_index_for_directory(category_dir: Path) -> None:
    articles = sorted(
        [p for p in category_dir.iterdir() if is_markdown_article(p)],
        key=lambda p: p.name.lower(),
    )
    if not articles:
        return

    # Title from directory name
    title = f"{category_dir.name} Articles"
    lines = [f"# {title}", ""]
    for article in articles:
        lines.append(f"- [{article.stem}]({article.name})")
    lines.append("")

    index_path = category_dir / "index.md"
    index_path.write_text("\n".join(lines), encoding="utf-8")

def walk_compendium_and_generate_indexes() -> None:
    if not COMPENDIUM_DIR.exists():
        return
    for root, dirs, files in os.walk(COMPENDIUM_DIR):
        # Skip hidden or staging directories if any appear
        dirs[:] = [d for d in dirs if not d.startswith(".") and d != "_incoming"]
        category_dir = Path(root)
        generate_index_for_directory(category_dir)

if __name__ == "__main__":
    walk_compendium_and_generate_indexes()


