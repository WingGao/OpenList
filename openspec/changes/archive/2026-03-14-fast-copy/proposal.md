## Why

This change introduces a "Fast Copy" (快传) capability to OpenList. It allows users to quickly transfer files between different storage platforms by leveraging file metadata (MD5, SHA1, GCID) stored in a local database. This avoids unnecessary data transfers when the target platform already possesses a copy of the file (e.g., via metadata matching).

## What Changes

- **Metadata Storage**: Implement a local persistent database using [Pebble](https://github.com/cockroachdb/pebble) to store file hashes (MD5, SHA1, GCID) and associated metadata.
- **Hash Command**: Add a new CLI command to calculate file/directory hashes and store them in the metadata database.
- **Query API/Methods**: Add internal utility methods to query the metadata database by file path, size, or hash.
- **Human-Readable Output**: Provide functions to format binary hashes into human-readable hex strings for display and JSON export.

## Capabilities

### New Capabilities

- `metadata-db`: Core capability for managing the persistent metadata store (Pebble), including schema design (MD5, SHA1, GCID as keys) and CRUD operations.
- `file-hash-cli`: Command-line interface for calculating file/directory hashes, updating the metadata database, and exporting results as JSON.
- `metadata-lookup`: Internal utility methods for looking up file metadata by path/size or hash values, enabling the "Fast Copy" logic.

### Modified Capabilities

- (None)

## Impact

- **New Dependencies**: `github.com/cockroachdb/pebble` will be added to `go.mod`.
- **CLI**: A new sub-command will be added to the OpenList CLI.
- **Codebase**: New packages/logic will be added under `wing/` (core logic) and `cmd/` (CLI entry point).
- **Storage**: A new directory for the Pebble database (e.g., `data/filehash/`) will be required.
