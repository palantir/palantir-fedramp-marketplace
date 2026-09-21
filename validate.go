package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

const schemaName = "fedramp-certification-package-overview-schema-2026-06-24.json"
const commonSchemaName = "fedramp-common-definitions-schema-2026-06-24.json"
const schemaBaseURL = "https://raw.githubusercontent.com/FedRAMP/schemas/main/"

func loadSchema() (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat()
	client := &http.Client{Timeout: 30 * time.Second}
	for _, name := range []string{commonSchemaName, schemaName} {
		doc, err := fetchSchema(client, schemaBaseURL+name)
		if err != nil {
			return nil, err
		}
		if err := compiler.AddResource("https://fedramp.gov/schemas/"+name, doc); err != nil {
			return nil, err
		}
	}
	return compiler.Compile("https://fedramp.gov/schemas/" + schemaName)
}

func fetchSchema(client *http.Client, url string) (any, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %s: %s", url, response.Status)
	}
	var schema any
	err = json.NewDecoder(response.Body).Decode(&schema)
	return schema, err
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
