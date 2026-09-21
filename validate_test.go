package main

import (
	"errors"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

func TestSchemaValidation(t *testing.T) {
	paths, err := dataFiles()
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		return
	}
	schema, err := loadSchema()
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name   string
		mutate func(map[string]any)
	}{
		{"valid", nil},
		{"required", func(data map[string]any) {
			delete(data, "serviceIdentification")
		}},
		{"type", func(data map[string]any) {
			data["contactInformation"] = "not an array"
		}},
		{"enum", func(data map[string]any) {
			data["serviceProperties"].(map[string]any)["deploymentModel"] = "Private Cloud"
		}},
		{"date", func(data map[string]any) {
			data["certifiedServices"].([]any)[0].(map[string]any)["dateAvailable"] = "not-a-date"
		}},
		{"url", func(data map[string]any) {
			data["serviceIdentification"].(map[string]any)["website"] = "not a valid URI"
		}},
		{"common-schema-reference", func(data map[string]any) {
			data["serviceIdentification"].(map[string]any)["logo"] = "https://example.com/logo.txt"
		}},
	}
	for _, path := range paths {
		for _, tc := range cases {
			t.Run(path+"/"+tc.name, func(t *testing.T) {
				data, err := loadDocument(path)
				if err != nil {
					t.Fatal(err)
				}
				if tc.mutate != nil {
					tc.mutate(data)
				}
				err = schema.Validate(data)
				if tc.mutate == nil {
					if err != nil {
						t.Fatal(err)
					}
				} else {
					var validationError *jsonschema.ValidationError
					if !errors.As(err, &validationError) {
						t.Fatalf("expected validation error, got %v", err)
					}
				}
			})
		}
	}
}
