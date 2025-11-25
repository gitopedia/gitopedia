package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

func TestIntegration(t *testing.T) {
	// 1. Setup temp environment
	tmpDir, err := os.MkdirTemp("", "gitopedia-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	compendiumDir := filepath.Join(tmpDir, "Compendium")
	if err := os.MkdirAll(compendiumDir, 0755); err != nil {
		t.Fatal(err)
	}

	// 2. Create a dummy article
	articleContent := `---
id: 01H1V5X8J9K0L1M2N3P4Q5R6S7
title: "Integration Test Article"
slug: "integration-test"
tags: ["Test"]
summary: "A test article."
created: 2025-01-01
author: "Tester"
---

This is an integration test.
[Internal Link](other.md)
`
	if err := os.WriteFile(filepath.Join(compendiumDir, "test.md"), []byte(articleContent), 0644); err != nil {
		t.Fatal(err)
	}

	// 3. Run Indexer
	cwd, _ := os.Getwd()
	var repoRoot string
	if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
		repoRoot = cwd
	} else {
		repoRoot = filepath.Dir(cwd)
	}

	indexerPath := filepath.Join(repoRoot, "../knowledge-base/cmd/indexer/main.go")
	if _, err := os.Stat(indexerPath); os.IsNotExist(err) {
		t.Skipf("Indexer not found at %s, skipping integration test", indexerPath)
	}
	absIndexerPath, _ := filepath.Abs(indexerPath)

	// Output dir for index
	outDir := filepath.Join(tmpDir, "out")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Run indexer
	// We must run from knowledge-base root so it finds go.mod
	kbRoot := filepath.Dir(filepath.Dir(filepath.Dir(absIndexerPath)))
	cmd := exec.Command("go", "run", absIndexerPath)
	cmd.Dir = kbRoot
	cmd.Env = append(os.Environ(), fmt.Sprintf("GITOPEDIA_DIR=%s", compendiumDir))
	
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Indexer failed: %v\nOutput:\n%s", err, string(out))
	}

	// 4. Verify Index
	// Indexer writes to kbRoot/out/index.sqlite
	dbPath := filepath.Join(kbRoot, "out", "index.sqlite")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	var title string
	err = db.QueryRow("SELECT title FROM articles WHERE id = '01H1V5X8J9K0L1M2N3P4Q5R6S7'").Scan(&title)
	if err != nil {
		t.Fatalf("Failed to query article: %v", err)
	}

	if title != "Integration Test Article" {
		t.Errorf("Expected title 'Integration Test Article', got '%s'", title)
	}

	// Verify FTS
	var count int
	err = db.QueryRow("SELECT count(*) FROM article_index WHERE article_index MATCH 'integration'").Scan(&count)
	if err != nil {
		t.Fatalf("FTS query failed: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 match for 'integration', got %d", count)
	}

	log.Println("Integration test passed!")
}

