package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGenerationPipeline(t *testing.T) {
	source, err := os.ReadFile("data/pfcs-package.yaml")
	if err != nil {
		t.Fatal(err)
	}
	template, err := os.ReadFile("TEMPLATE.md")
	if err != nil {
		t.Fatal(err)
	}
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(root); err != nil {
			t.Error(err)
		}
	})
	write := func(path string, content []byte) {
		t.Helper()
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, content, 0644); err != nil {
			t.Fatal(err)
		}
	}
	read := func(path string) []byte {
		t.Helper()
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		return content
	}
	path := "data/example.yaml"
	write(path, source)
	write("data/other.yaml", source)
	write("TEMPLATE.md", template)
	if err := run(true, false); err == nil {
		t.Fatal("missing JSON should fail check")
	}
	if _, err := os.Stat("generated"); !os.IsNotExist(err) {
		t.Fatal("check wrote output")
	}
	first := time.Date(2026, 9, 22, 12, 0, 0, 0, time.UTC)
	if err := updateJSON(path, false, first); err != nil {
		t.Fatal(err)
	}
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	if err := run(true, false); err != nil {
		t.Fatal(err)
	}
	if err := run(false, true); err != nil {
		t.Fatal(err)
	}
	jsonFile, markdown := jsonPath(path), outputPath(jsonPath(path))
	initialJSON, initialMarkdown := read(jsonFile), read(markdown)
	otherJSON := read("generated/json/other.json")
	if !bytes.Contains(initialMarkdown, []byte(first.Format(time.RFC3339))) {
		t.Fatal("Markdown missing generated timestamp")
	}
	before, err := os.Stat(jsonFile)
	if err != nil {
		t.Fatal(err)
	}
	write(path, append([]byte("# A comment\n\n"), source...))
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(jsonFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(initialJSON, read(jsonFile)) || !before.ModTime().Equal(after.ModTime()) {
		t.Fatal("format-only change rewrote JSON")
	}
	write("TEMPLATE.md", append(template, []byte("\nTemplate update\n")...))
	if err := run(true, false); err == nil {
		t.Fatal("template edit should fail check")
	}
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(initialJSON, read(jsonFile)) {
		t.Fatal("template edit changed timestamp")
	}
	changed := bytes.Replace(source, []byte("PFCS ISSO"), []byte("Updated PFCS ISSO"), 1)
	write(path, changed)
	if err := run(true, false); err == nil {
		t.Fatal("source edit should fail check")
	}
	if !bytes.Equal(initialJSON, read(jsonFile)) {
		t.Fatal("check modified JSON")
	}
	second := first.Add(time.Hour)
	if err := updateJSON(path, false, second); err != nil {
		t.Fatal(err)
	}
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(read(markdown), []byte(second.Format(time.RFC3339))) {
		t.Fatal("Markdown has stale timestamp")
	}
	updated := read(jsonFile)
	if !bytes.Equal(otherJSON, read("generated/json/other.json")) {
		t.Fatal("editing one source changed another package")
	}
	if err := updateJSON(path, false, second.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(updated, read(jsonFile)) {
		t.Fatal("repeated run changed JSON")
	}
	invalid := bytes.Replace(changed, []byte("deploymentModel: Public Cloud"), []byte("deploymentModel: Invalid"), 1)
	write(path, invalid)
	if err := run(false, false); err == nil {
		t.Fatal("invalid source passed schema validation")
	}
	if !bytes.Equal(updated, read(jsonFile)) {
		t.Fatal("invalid source overwrote JSON")
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err := run(true, false); err == nil {
		t.Fatal("orphan outputs passed check")
	}
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{jsonFile, markdown} {
		if _, err := os.Stat(p); !os.IsNotExist(err) {
			t.Fatalf("orphan remains: %s", p)
		}
	}
}

func TestSourceYAML(t *testing.T) {
	for _, tc := range []struct {
		name, content string
		valid         bool
	}{
		{"date", "metadata: {}\ndate: 2026-09-22\n", true},
		{"duplicate", "metadata: {}\nname: first\nname: second\n", false},
		{"multiple documents", "metadata: {}\n---\nmetadata: {}\n", false},
		{"managed timestamp", "metadata:\n  lastUpdated: 2026-09-22T12:00:00Z\n", false},
		{"missing metadata", "name: example\n", false},
		{"invalid", "metadata: [", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "source.yaml")
			if err := os.WriteFile(path, []byte(tc.content), 0644); err != nil {
				t.Fatal(err)
			}
			doc, err := loadSource(path)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v, error=%v", tc.valid, err)
			}
			if tc.name == "date" && doc["date"] != "2026-09-22" {
				t.Fatalf("date changed: %v", doc["date"])
			}
			if tc.name == "managed timestamp" && !strings.Contains(err.Error(), "generated") {
				t.Fatal(err)
			}
		})
	}
}
