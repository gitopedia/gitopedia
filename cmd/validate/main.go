package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gitopedia/gitopedia/pkg/frontmatter"
)

var ulidRegex = regexp.MustCompile(`^[0-9A-HJKMNP-TV-Z]{26}$`)
var linkRegex = regexp.MustCompile(`\[.*?\]\((.*?)\)`)

type AuthorityEntry struct {
	ID string `json:"id"`
}

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
	// Assume we run from repo root
	authorityDir := filepath.Join(repoRoot, "authority")
	compendiumDir := filepath.Join(repoRoot, "Compendium")

	authorityIDs, err := loadAuthorities(authorityDir)
	if err != nil {
		// Warn but don't fail if bootstrapping? Script said:
		// "Not critical if not present yet"
		log.Printf("Warning loading authorities: %v", err)
		authorityIDs = make(map[string]bool)
	}

	return validateArticles(compendiumDir, authorityIDs)
}

func loadAuthorities(dir string) (map[string]bool, error) {
	ids := make(map[string]bool)
	files := []string{"people.json", "orgs.json", "places.json", "topics.json"}

	for _, fname := range files {
		path := filepath.Join(dir, fname)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", fname, err)
		}

		var entries []AuthorityEntry
		if err := json.Unmarshal(data, &entries); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", fname, err)
		}

		for _, e := range entries {
			if ids[e.ID] {
				return nil, fmt.Errorf("duplicate authority ID: %s", e.ID)
			}
			ids[e.ID] = true
		}
	}
	return ids, nil
}

func validateArticles(dir string, authIDs map[string]bool) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}

	seenIDs := make(map[string]string)
	seenTitles := make(map[string]string)
	hasError := false

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			return nil
		}
		if strings.ToLower(d.Name()) == "index.md" {
			return nil
		}

		if err := validateFile(path, dir, authIDs, seenIDs, seenTitles); err != nil {
			log.Printf("Error in %s: %v", path, err)
			hasError = true
		}
		return nil
	})

	if err != nil {
		return err
	}
	if hasError {
		return fmt.Errorf("validation failed")
	}
	return nil
}

func validateFile(path, rootDir string, authIDs map[string]bool, seenIDs, seenTitles map[string]string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fm, body, err := frontmatter.Parse(content)
	if err != nil {
		return fmt.Errorf("front matter error: %w", err)
	}

	// Required fields
	if fm.ID == "" {
		return fmt.Errorf("missing 'id'")
	}
	if fm.Title == "" {
		return fmt.Errorf("missing 'title'")
	}
	if fm.Slug == "" {
		return fmt.Errorf("missing 'slug'")
	}

	// ULID
	if !ulidRegex.MatchString(strings.ToUpper(fm.ID)) {
		return fmt.Errorf("invalid ULID: %s", fm.ID)
	}
	if existing, ok := seenIDs[fm.ID]; ok {
		return fmt.Errorf("duplicate ID %s (also in %s)", fm.ID, existing)
	}
	seenIDs[fm.ID] = path

	// Title
	normTitle := strings.ToLower(strings.TrimSpace(fm.Title))
	if existing, ok := seenTitles[normTitle]; ok {
		return fmt.Errorf("duplicate title '%s' (also in %s)", fm.Title, existing)
	}
	seenTitles[normTitle] = path

	// Slug
	if strings.Contains(fm.Slug, " ") || fm.Slug != strings.TrimSpace(fm.Slug) {
		return fmt.Errorf("invalid slug '%s'", fm.Slug)
	}

	// Tags
	if len(fm.Tags) == 0 {
		return fmt.Errorf("missing 'tags'")
	}

	// Facets check against Authority
	checkFacets := func(name string, values []string) error {
		for _, v := range values {
			if strings.Contains(v, ":") {
				if !authIDs[v] && len(authIDs) > 0 {
					return fmt.Errorf("authority ID '%s' in '%s' not found", v, name)
				}
			}
		}
		return nil
	}

	if err := checkFacets("tags", fm.Tags); err != nil {
		return err
	}
	if err := checkFacets("people", fm.People); err != nil {
		return err
	}
	if err := checkFacets("orgs", fm.Orgs); err != nil {
		return err
	}
	if err := checkFacets("places", fm.Places); err != nil {
		return err
	}

	// Validate Links
	matches := linkRegex.FindAllStringSubmatch(body, -1)
	for _, m := range matches {
		link := strings.TrimSpace(m[1])
		if link == "" || strings.HasPrefix(link, "#") || strings.HasPrefix(link, "mailto:") {
			continue
		}

		if strings.HasPrefix(link, "http") {
			if _, err := url.Parse(link); err != nil {
				return fmt.Errorf("invalid URL: %s", link)
			}
		} else {
			// Internal link
			var target string
			if strings.HasPrefix(link, "/") {
				// Absolute path relative to Compendium root
				target = filepath.Join(rootDir, link)
			} else {
				// Relative to current file
				target = filepath.Join(filepath.Dir(path), link)
			}

			// Decode URL encoding (e.g. %20 -> space)
			if u, err := url.QueryUnescape(target); err == nil {
				target = u
			}

			// Check existence
			if _, err := os.Stat(target); os.IsNotExist(err) {
				// Try with .md extension
				if _, err := os.Stat(target + ".md"); os.IsNotExist(err) {
					// Try index.md inside directory
					if _, err := os.Stat(filepath.Join(target, "index.md")); os.IsNotExist(err) {
						// Log warning instead of erroring for now, to avoid breaking build on existing content
						// Or strict? User said "enforce".
						// But newly generated content might have issues?
						// Let's be strict for "Pipeline Robustness".
						return fmt.Errorf("broken internal link: %s", link)
					}
				}
			}
		}
	}

	return nil
}
