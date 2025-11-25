package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gitopedia/gitopedia/pkg/frontmatter"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	repoRoot, err := os.Getwd()
	if err != nil {
		return err
	}
	compendiumDir := filepath.Join(repoRoot, "Compendium")

	return walkAndGenerate(compendiumDir)
}

func walkAndGenerate(root string) error {
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return nil
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip hidden folders or _incoming
			base := filepath.Base(path)
			if strings.HasPrefix(base, ".") || base == "_incoming" {
				return filepath.SkipDir
			}

			if path != root {
				if err := generateIndex(path); err != nil {
					log.Printf("Error generating index for %s: %v", path, err)
				}
			}
		}
		return nil
	})
	return err
}

func generateIndex(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var articles []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if strings.ToLower(name) == "index.md" {
			continue
		}
		if strings.ToLower(filepath.Ext(name)) != ".md" {
			continue
		}
		articles = append(articles, name)
	}

	if len(articles) == 0 {
		return nil
	}

	sort.Strings(articles)

	dirName := filepath.Base(dir)
	title := fmt.Sprintf("%s Articles", dirName)
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s\n\n", title))

	for _, fname := range articles {
		// Try to get title from front matter
		fpath := filepath.Join(dir, fname)
		linkTitle := strings.TrimSuffix(fname, filepath.Ext(fname))

		content, err := os.ReadFile(fpath)
		if err == nil {
			fm, _, err := frontmatter.Parse(content)
			if err == nil && fm.Title != "" {
				linkTitle = fm.Title
			}
		}

		sb.WriteString(fmt.Sprintf("- [%s](%s)\n", linkTitle, fname))
	}
	sb.WriteString("\n")

	indexPath := filepath.Join(dir, "index.md")
	return os.WriteFile(indexPath, []byte(sb.String()), 0644)
}
