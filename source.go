package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"go.yaml.in/yaml/v3"
)

func jsonPath(source string) string {
	return filepath.Join("generated/json", strings.TrimSuffix(filepath.Base(source), ".yaml")+".json")
}

func loadSource(path string) (map[string]any, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	var node yaml.Node
	if err := decoder.Decode(&node); err != nil {
		return nil, err
	}
	var extra yaml.Node
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, fmt.Errorf("expected exactly one YAML document")
	}
	// JSON has no date type. Preserve YAML date scalars as their original strings.
	var normalizeDates func(*yaml.Node)
	normalizeDates = func(n *yaml.Node) {
		if n.Tag == "!!timestamp" {
			n.Tag = "!!str"
		}
		for _, child := range n.Content {
			normalizeDates(child)
		}
	}
	normalizeDates(&node)
	var source map[string]any
	if err := node.Decode(&source); err != nil {
		return nil, err
	}
	// Normalize numeric and container types to the same representation as JSON.
	encoded, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	if err := json.Unmarshal(encoded, &data); err != nil {
		return nil, err
	}
	metadata, ok := data["metadata"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("metadata must be an object")
	}
	if _, exists := metadata["lastUpdated"]; exists {
		return nil, fmt.Errorf("metadata.lastUpdated is generated; remove it from the YAML source")
	}
	if _, exists := data["$schema"]; exists {
		return nil, fmt.Errorf("$schema is generated; remove it from the YAML source")
	}
	return data, nil
}

func updateJSON(source string, check bool, now time.Time) error {
	data, err := loadSource(source)
	if err != nil {
		return err
	}
	data["$schema"] = "https://fedramp.gov/schemas/" + schemaName
	output := jsonPath(source)
	current, err := os.ReadFile(output)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	timestamp := now.UTC().Format(time.RFC3339Nano)
	var previous map[string]any
	if json.Unmarshal(current, &previous) == nil {
		if metadata, ok := previous["metadata"].(map[string]any); ok {
			value, _ := metadata["lastUpdated"].(string)
			delete(metadata, "lastUpdated")
			if _, err := time.Parse(time.RFC3339, value); err == nil && reflect.DeepEqual(previous, data) {
				timestamp = value
			}
		}
	}
	data["metadata"].(map[string]any)["lastUpdated"] = timestamp
	schema, err := loadSchema()
	if err != nil {
		return err
	}
	if err := schema.Validate(data); err != nil {
		return err
	}
	expected, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	expected = append(expected, '\n')
	if bytes.Equal(current, expected) {
		return nil
	}
	if check {
		return fmt.Errorf("%s is out of date. Run `./godelw generate`.", output)
	}
	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(output, expected, 0644); err != nil {
		return err
	}
	fmt.Printf("Updated %s.\n", output)
	return nil
}

// These directories contain generated artifacts only. Remove output for deleted sources.
func removeOrphans(sources []string, check bool) error {
	expected := make(map[string]bool)
	for _, source := range sources {
		expected[jsonPath(source)] = true
		expected[outputPath(jsonPath(source))] = true
	}
	for _, pattern := range []string{"generated/json/*.json", "generated/markdown/*.md"} {
		paths, err := filepath.Glob(pattern)
		if err != nil {
			return err
		}
		for _, path := range paths {
			if expected[path] {
				continue
			}
			if check {
				return fmt.Errorf("%s has no YAML source. Run `./godelw generate`.", path)
			}
			if err := os.Remove(path); err != nil {
				return err
			}
			fmt.Printf("Removed %s.\n", path)
		}
	}
	return nil
}
