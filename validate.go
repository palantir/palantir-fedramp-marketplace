package main

import (
	"embed"
	"encoding/json"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

const schemaName = "fedramp-certification-package-overview-schema-2026-06-24.json"
const commonSchemaName = "fedramp-common-definitions-schema-2026-06-24.json"

//go:embed schemas/*.json
var schemaFiles embed.FS

func loadSchema() (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat()
	for _, name := range []string{commonSchemaName, schemaName} {
		content, err := schemaFiles.ReadFile("schemas/" + name)
		if err != nil {
			return nil, err
		}
		var doc any
		if err := json.Unmarshal(content, &doc); err != nil {
			return nil, err
		}
		if err := compiler.AddResource("https://fedramp.gov/schemas/"+name, doc); err != nil {
			return nil, err
		}
	}
	return compiler.Compile("https://fedramp.gov/schemas/" + schemaName)
}

func validateDocument(path string) error {
	data, err := loadDocument(path)
	if err != nil {
		return err
	}
	schema, err := loadSchema()
	if err != nil {
		return err
	}
	return schema.Validate(data)
}
