## 1. Setup and Dependencies

- [x] 1.1 Add `github.com/cockroachdb/pebble` dependency to `go.mod`
- [x] 1.2 Create `wing/filehash` package directory
- [x] 1.3 Implement Pebble database initialization and configuration logic

## 2. Core Metadata Store (`wing/filehash`)

- [x] 2.1 Define metadata schema and binary serialization/deserialization logic
- [x] 2.2 Implement primary MD5 key storage and retrieval (`K:MD5:`)
- [x] 2.3 Implement lookup index keys (`K:SHA1:`, `K:GCID:`)
- [x] 2.4 Implement local path/size cache key (`C:`)
- [x] 2.5 Implement atomic batch update logic for multi-key writes

## 3. Hashing and Lookup Utilities

- [x] 3.1 Implement concurrent file hashing logic (MD5, SHA1, GCID)
- [x] 3.2 Implement path-based lookup with size validation
- [x] 3.3 Implement hash-based lookup methods
- [x] 3.4 Add human-readable hex formatting helpers for binary hashes

## 4. CLI Implementation (`cmd/hash.go`)

- [x] 4.1 Create `hash` command entry point in `cmd/`
- [x] 4.2 Implement recursive directory scanning for the `hash` command
- [x] 4.3 Add `--json` flag support for exporting metadata
- [x] 4.4 Add `--force` flag to bypass the path/size cache

## 5. Verification

- [x] 5.1 Add unit tests for Pebble KV operations
- [x] 5.2 Add unit tests for hashing and lookup logic
- [x] 5.3 Perform manual end-to-end test of the `hash` CLI command
