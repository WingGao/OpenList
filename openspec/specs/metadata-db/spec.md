# Metadata Storage (Pebble DB)

## Purpose
This specification defines the requirements for the local persistent storage of file metadata using Pebble DB.

## Requirements

### Requirement: Metadata Storage (Pebble DB)
The system SHALL use [Pebble](https://github.com/cockroachdb/pebble) as the local persistent storage for file metadata. The data SHOULD be stored in the `data/filehash/` directory.

#### Scenario: Database Initialization
- **WHEN** the system starts the hash command or lookup logic
- **THEN** it SHOULD initialize the Pebble database in the configured directory if it does not exist

### Requirement: Key-Value Schema for File Hash Metadata
The system SHALL store file metadata using the following key-value pairs:
- **Primary Key (MD5)**: `K:MD5:{md5_binary}` -> `name={xxx}|size={xxx}|md5={md5_binary}|sha1={sha1_binary}|gcid={gcid_binary}`
- **Lookup Key (SHA1)**: `K:SHA1:{sha1_binary}` -> `{md5_binary}`
- **Lookup Key (GCID)**: `K:GCID:{gcid_binary}` -> `{md5_binary}`
- **Cache Key (Local File)**: `C:{FILE_FULLPATH}|{FILE_SIZE}` -> `{md5_binary}`

#### Scenario: Storing New File Hash Metadata
- **WHEN** a file's metadata is calculated
- **THEN** the system SHALL create the primary MD5 key and all secondary lookup/cache keys in the database

#### Scenario: Atomic Updates
- **WHEN** multiple keys for a single file are being written
- **THEN** the system SHOULD use a Pebble batch to ensure atomic updates
