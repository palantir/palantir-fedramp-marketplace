package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadmeMatchesTemplateAndJSON(t *testing.T) {
	if err := run(true, false); err != nil {
		t.Fatal(err)
	}
}

func TestNoData(t *testing.T) {
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
	for _, state := range []string{"missing", "empty"} {
		t.Run(state, func(t *testing.T) {
			if state == "empty" {
				if err := os.Mkdir("data", 0755); err != nil {
					t.Fatal(err)
				}
			}
			for _, mode := range []struct{ check, validate bool }{
				{false, false}, {true, false}, {false, true},
			} {
				if err := run(mode.check, mode.validate); err != nil {
					t.Fatal(err)
				}
			}
			if _, err := os.Stat("package-information"); !os.IsNotExist(err) {
				t.Fatal("no-data run created output directory")
			}
		})
	}
	if err := os.WriteFile("data/invalid.json", []byte("{"), 0644); err != nil {
		t.Fatal(err)
	}
	for _, validate := range []bool{false, true} {
		if err := run(true, validate); err == nil {
			t.Fatal("existing invalid JSON should fail verification")
		}
	}
}

func TestReadmeCheckAndUpdate(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "data"), 0755); err != nil {
		t.Fatal(err)
	}
	for name, content := range map[string]string{
		"data/first.json":  `{"title":"Example"}`,
		"TEMPLATE.md":      "# {{.title}}\n",
		"data/second.json": `{"title":"Second"}`,
		"README.md":        "Maintained by hand\n",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(root); err != nil {
			t.Error(err)
		}
	})
	if err := run(true, false); err == nil {
		t.Fatal("missing output should fail check")
	}
	if _, err := os.Stat("package-information"); !os.IsNotExist(err) {
		t.Fatal("check created output directory")
	}
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	second, err := os.ReadFile("package-information/second.md")
	if err != nil || string(second) != "# Second\n" {
		t.Fatalf("second output: %q, %v", second, err)
	}
	output := "package-information/first.md"
	if err := os.WriteFile(output, []byte("stale\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := run(true, false); err == nil || !strings.Contains(err.Error(), "-stale") {
		t.Fatalf("expected stale README error and diff, got %v", err)
	}
	current, err := os.ReadFile(output)
	if err != nil || string(current) != "stale\n" {
		t.Fatalf("check mode changed the README: %q, %v", current, err)
	}
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	if err := run(true, false); err != nil {
		t.Fatal(err)
	}
	current, err = os.ReadFile(output)
	if err != nil || string(current) != "# Example\n" {
		t.Fatalf("unexpected output: %q, %v", current, err)
	}
	manual, err := os.ReadFile("README.md")
	if err != nil || string(manual) != "Maintained by hand\n" {
		t.Fatalf("generation changed README: %q, %v", manual, err)
	}
	before, err := os.Stat(output)
	if err != nil {
		t.Fatal(err)
	}
	if err := run(false, false); err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(output)
	if err != nil {
		t.Fatal(err)
	}
	if !before.ModTime().Equal(after.ModTime()) {
		t.Fatal("unchanged README was rewritten")
	}
}
