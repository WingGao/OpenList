## ADDED Requirements

### Requirement: RapidUploader Interface
The system MUST provide a `RapidUploader` interface that drivers can implement to support hash-based rapid uploads.

#### Scenario: Interface availability
- **WHEN** a driver implements the `RapidUploader` interface
- **THEN** the system can detect and call its `RapidUpload` method when a rapid upload request is received

### Requirement: Rapid Upload API Endpoint
The system MUST expose a `/put_rapid` endpoint in the server's file system router to allow clients to initiate a rapid upload.

#### Scenario: Successful rapid upload via API
- **WHEN** a client sends a valid hash and destination directory to `/put_rapid`
- **THEN** the system attempts a rapid upload using the corresponding driver and returns the created object on success

### Requirement: Hash-based Object Lookup
The system SHALL use provided file hashes to attempt to locate and reuse existing data on the storage provider's side.

#### Scenario: Rapid upload success with correct hash
- **WHEN** a client provides a hash that exists on the storage provider
- **THEN** the driver's `RapidUpload` method succeeds and creates a new file entry without re-uploading the data

### Requirement: Metadata Validation
The system SHOULD allow providing filename and size alongside the hash to validate the rapid upload request before execution.

#### Scenario: Size mismatch validation
- **WHEN** a client provides a hash and a size that does not match the metadata of the file found by the hash
- **THEN** the system SHOULD return an error or fallback to prevent incorrect metadata association

### Requirement: Fallback for Unsupported Drivers
The system MUST gracefully handle cases where a driver does not support rapid upload.

#### Scenario: Attempting rapid upload on unsupported driver
- **WHEN** a client attempts a rapid upload to a driver that does not implement `RapidUploader`
- **THEN** the system returns a `NotSupport` error or similar indicator
