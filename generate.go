//go:generate go run .

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

var templatePath = "TEMPLATE.md"

func dataFiles() ([]string, error) {
	return filepath.Glob("data/*.json")
}

func outputPath(path string) string {
	return filepath.Join("package-information", strings.TrimSuffix(filepath.Base(path), ".json")+".md")
}

func main() {
	check := flag.Bool("check", false, "check generated Markdown without modifying it")
	validate := flag.Bool("validate", false, "validate the public information JSON")
	flag.Parse()
	err := run(*check, *validate)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(check, validate bool) error {
	paths, err := dataFiles()
	if err != nil {
		return err
	}
	for _, path := range paths {
		if validate {
			err = validateDocument(path)
		} else {
			err = updateReadme(path, check)
		}
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
	}
	if validate {
		fmt.Println("Valid JSON")
	}
	return nil
}

func loadDocument(path string) (map[string]any, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	err = json.Unmarshal(content, &data)
	return data, err
}

func renderReadme(path string) ([]byte, error) {
	data, err := loadDocument(path)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(filepath.Base(templatePath)).Option("missingkey=error").Funcs(template.FuncMap{
		"join": func(values []any, separator string) string {
			parts := make([]string, len(values))
			for i, value := range values {
				parts[i] = fmt.Sprint(value)
			}
			return strings.Join(parts, separator)
		},
		"sortServices": sortServices,
	}).ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	return output.Bytes(), err
}

func sortServices(values []any) ([]map[string]any, error) {
	services := make([]map[string]any, len(values))
	for i, value := range values {
		service, ok := value.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("certifiedServices[%d] must be an object", i)
		}
		if _, ok := service["serviceName"].(string); !ok {
			return nil, fmt.Errorf("certifiedServices[%d].serviceName must be a string", i)
		}
		services[i] = service
	}
	sort.SliceStable(services, func(i, j int) bool {
		return strings.ToLower(services[i]["serviceName"].(string)) < strings.ToLower(services[j]["serviceName"].(string))
	})
	return services, nil
}

func updateReadme(path string, check bool) error {
	output := outputPath(path)
	expected, err := renderReadme(path)
	if err != nil {
		return err
	}
	current, err := os.ReadFile(output)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if bytes.Equal(current, expected) {
		fmt.Printf("%s is current.\n", output)
		return nil
	}
	if check {
		return fmt.Errorf("%s is out of date. Run `./godelw generate`.\n%s", output, readmeDiff(output, current, expected))
	}
	if err := os.MkdirAll("package-information", 0755); err != nil {
		return err
	}
	if err := os.WriteFile(output, expected, 0644); err != nil {
		return err
	}
	fmt.Printf("Updated %s.\n", output)
	return nil
}

// readmeDiff shows one unified hunk spanning all changed lines.
func readmeDiff(path string, current, expected []byte) string {
	before := strings.SplitAfter(string(current), "\n")
	after := strings.SplitAfter(string(expected), "\n")
	start := 0
	for start < len(before) && start < len(after) && before[start] == after[start] {
		start++
	}
	endBefore, endAfter := len(before), len(after)
	for endBefore > start && endAfter > start && before[endBefore-1] == after[endAfter-1] {
		endBefore--
		endAfter--
	}
	var diff strings.Builder
	fmt.Fprintf(&diff, "--- %s\n+++ %s (generated)\n@@ -%d,%d +%d,%d @@\n", path, path, start+1, endBefore-start, start+1, endAfter-start)
	for _, line := range before[start:endBefore] {
		fmt.Fprintf(&diff, "-%s", line)
	}
	for _, line := range after[start:endAfter] {
		fmt.Fprintf(&diff, "+%s", line)
	}
	return diff.String()
}
