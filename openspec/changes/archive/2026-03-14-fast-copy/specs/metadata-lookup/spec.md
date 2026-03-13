## ADDED Requirements

### Requirement: Lookup by Path and Size
The system SHALL provide a lookup method that finds file metadata by combining its absolute path and size.

#### Scenario: Cached Hash Lookup
- **WHEN** the system performs a lookup with a matching absolute path and file size in the cache (`C:`)
- **THEN** it SHALL return the corresponding file metadata

#### Scenario: Stale Cache Handling
- **WHEN** a file's path matches but its size has changed
- **THEN** it SHALL NOT return the cached metadata and SHALL re-calculate the hash

### Requirement: Lookup by Hash Values
The system SHALL provide lookup methods that find file metadata by its MD5, SHA1, or GCID binary values.

#### Scenario: Lookup by SHA1
- **WHEN** the system is queried with a SHA1 binary value
- **THEN** it SHALL find the corresponding MD5 from the SHA1 lookup key (`K:SHA1:`) and then retrieve the full metadata from the primary MD5 key (`K:MD5:`)

#### Scenario: Batch Lookup Support
- **WHEN** queried with multiple file paths or hash values
- **THEN** the system SHALL perform lookup operations in batch to optimize performance
