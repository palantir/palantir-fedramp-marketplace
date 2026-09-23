# Maintainer Information

Every `data/<name>.yaml` file generates `generated/json/<name>.json`. The JSON is validated against the vendored FedRAMP schemas, then rendered with `TEMPLATE.md` into `generated/markdown/<name>.md`. Commit both the YAML source and generated outputs.

## Updating Information

To update the information, edit the YAML in the data directory, the template, or both, then regenerate the outputs:

```bash
./godelw generate
```

Do not include `$schema` or `metadata.lastUpdated` in YAML. Generation adds `$schema` using the schema selected in `validate.go`. It sets `metadata.lastUpdated` to the current UTC time (RFC3339, whole seconds) when a package's data changes, and preserves it otherwise. YAML comments, formatting, and key order do not count as data changes. Template-only edits regenerate Markdown without changing the JSON timestamp. Each package is handled independently, and repeated generation leaves unchanged files untouched.

`--check` verifies both generated formats without writing files. Generation also removes JSON and Markdown whose YAML source was deleted. The `generated/json` and `generated/markdown` directories are reserved for generated files.

Run all validation and tests:

```bash
./godelw verify --apply=false
```

Optionally, you can add the `fix-nits` label to your PR and the nitbot will render and push the updated markdowns.

## Commands

```bash
./godelw generate              # Regenerate JSON and Markdown
./godelw verify --apply=false  # Verify generated output and run tests
./godelw test                  # Run the Go tests
go run . --check               # Read-only JSON and Markdown check
go run . --validate            # Validate generated JSON
```

## Source files

- `TEMPLATE.md` is the human-readable Markdown template for the data.
- `data/*.yaml` contains the maintained package information, without `$schema` or `metadata.lastUpdated`.
- `generated/json/*.json` and `generated/markdown/*.md` are generated. Do not edit them directly.
- `source.go` converts YAML into validated JSON and manages timestamps.
- `generate.go` renders the template with the JSON data.
- `validate.go` validates the JSON against the vendored FedRAMP schemas in `schemas/`.
- `schemas/README.md` records the upstream schema revision and update instructions.
- `vendor/` contains Go dependencies so generation and validation do not download modules. Refresh it with `go mod vendor` after updating dependencies.
- `generate_test.go` and `validate_test.go` test public-information markdown synchronization and schema validation.
