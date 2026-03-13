## Context

OpenList currently lacks a centralized local metadata store for file hashes (MD5, SHA1, GCID). Implementing a "Fast Copy" feature requires a persistent, high-performance local database to avoid redundant hash calculations and enable metadata-based cloud transfers (e.g., 115, Aliyundrive).

## Goals / Non-Goals

**Goals:**
- Implement a persistent metadata store using Pebble DB.
- Provide a CLI command `hash` for calculating and indexing file metadata.
- Provide internal lookup utilities for metadata retrieval by path/size or hash.
- Support recursive directory scanning for metadata indexing.
- Ensure thread-safe access to the metadata database.

**Non-Goals:**
- Automating the "Fast Copy" process within specific drivers (this change focuses on the *infrastructure* and *CLI* for metadata).
- Implementing a GUI for metadata management.
- Synchronization of metadata across multiple devices.

## Decisions

### 1. Database Choice: Pebble
- **Rationale**: Pebble (by CockroachDB) is a Go-native, high-performance key-value store inspired by LevelDB/RocksDB. It offers excellent write performance (critical for batch hashing) and clean Go integration without CGO.
- **Alternatives**: 
  - *BoltDB/Bbolt*: Simpler but slower for high-volume writes.
  - *LevelDB (Go port)*: Less active development compared to Pebble.

### 2. Schema Design (Key-Value)
- **Primary Metadata**: `K:MD5:{md5_binary}` -> `name={xxx}|size={xxx}|md5={md5_binary}|sha1={sha1_binary}|gcid={gcid_binary}`. Using MD5 as the primary key allows easy retrieval of full metadata once an MD5 is identified.
- **Lookup Indexes**:
  - `K:SHA1:{sha1_binary}` -> `md5_binary`
  - `K:GCID:{gcid_binary}` -> `md5_binary`
  - `C:{FILE_FULLPATH}|{FILE_SIZE}` -> `md5_binary` (Cache to avoid re-hashing unchanged local files).
- **Rationle**: This "pointer" system minimizes data duplication while allowing lookups by any supported hash or path.

### 3. Package Structure
- **Core Logic (`wing/filehash`)**: Contains the Pebble wrapper, schema definitions, and hashing logic. This allows reuse across CLI and future server-side features.
- **CLI (`cmd/hash.go`)**: Wrapper around `wing/filehash` to provide the user interface.

### 4. Hash Calculation Strategy
- **Concurrent Processing**: Use a worker pool to calculate hashes for multiple files in parallel during recursive directory scans.
- **Batch Writes**: Use Pebble batches for every N files or per directory to optimize disk I/O.

## Risks / Trade-offs

- **[Risk] Database Corruption on Crash** 鈫 **Mitigation**: Use Pebble's write-ahead log (WAL) and ensure `Sync()` is called during critical batch completions.
- **[Risk] Large Database Size** 鈫 **Mitigation**: Store hashes in binary format rather than hex strings to save space.
- **[Risk] Stale Cache** 鈫 **Mitigation**: The cache key includes both `FILE_FULLPATH` and `FILE_SIZE`. While not perfect (doesn't detect content changes with same size), it's a standard trade-off for performance. A `--force` flag will be added to the `hash` command to bypass cache.
