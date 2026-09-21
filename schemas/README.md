# Vendored FedRAMP schemas

Unmodified snapshots from [FedRAMP/schemas](https://github.com/FedRAMP/schemas), branch `main`, retrieved 2026-09-21.

These files are embedded by `validate.go`; validation does not fetch schemas at runtime.

| File | SHA-256 |
|---|---|
| [fedramp-certification-package-overview-schema-2026-06-24.json](https://raw.githubusercontent.com/FedRAMP/schemas/main/fedramp-certification-package-overview-schema-2026-06-24.json) | `d92fa33f04629a65f85f2f070418310603bd9d3668ee66da31f9c0566b11be5a` |
| [fedramp-common-definitions-schema-2026-06-24.json](https://raw.githubusercontent.com/FedRAMP/schemas/main/fedramp-common-definitions-schema-2026-06-24.json) | `ed810d60584580fb86cda3f846d94504a14e09538122c27c5b5231d41151d6de` |

To update, replace both JSON files with upstream copies, update the retrieval date and checksums above, and run `./godelw verify --apply=false`.
