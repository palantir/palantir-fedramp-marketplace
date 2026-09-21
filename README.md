# Palantir FedRAMP Package Information

## Package Information

| Title | Information | Data |
|---|---|---|
| Palantir Federal Cloud Service (PFCS) | [Markdown](package-information/pfcs-package.md) | [JSON](data/pfcs-package.json) |
| Palantir Federal Cloud Service – Supporting Services (PFCS-SS) | [Markdown](package-information/pfcs-ss-package.md) | [JSON](data/pfcs-ss-package.json) |

---

## Maintainer Information

Every `data/<name>.json` file is rendered with `TEMPLATE.md` into `package-information/<name>.md`.

### Updating Information

To update the information, edit the JSON in the data directory, the template, or both, then regenerate the markdowns:

```bash
./godelw generate
```

Run all validation and tests:

```bash
./godelw verify --apply=false
```

Optionally, you can add the `fix-nits` label to your PR and the nitbot will render and push the updated markdowns.

### Commands

```bash
./godelw generate              # Regenerate all package-information pages
./godelw verify --apply=false  # Verify generated output and run tests
./godelw test                  # Run the Go tests
go run . --check               # Read-only public-information page check
go run . --validate            # Validate JSON
```

### Source files

- `TEMPLATE.md` is the human-readable Markdown template for the data.
- `package-information/*.md` is generated. Do not edit it directly.
- `generate.go` renders the template with the JSON data.
- `validate.go` validates the JSON against the vendored FedRAMP schemas in `schemas/`.
- `schemas/README.md` records the upstream schema revision and update instructions.
- `vendor/` contains Go dependencies so generation and validation do not download modules. Refresh it with `go mod vendor` after updating dependencies.
- `generate_test.go` and `validate_test.go` test public-information markdown synchronization and schema validation.
