## Why

Enable rapid upload functionality across supported storage drivers to optimize data transfer by reusing existing files based on their hash metadata, significantly reducing bandwidth and time for redundant uploads.

## What Changes

- **RapidUploader Interface**: Define a new interface `RapidUploader` to be implemented by drivers supporting this feature.
- **Driver Implementations**: Implement `RapidUpload` for all drivers that support hash-based rapid uploads (e.g., 115, Aliyundrive, Quark).
- **Server API**: Add a new `/put_rapid` endpoint to the file system group in the server router to expose the rapid upload capability.
- **File Metadata Usage**: Utilize existing `filehash.FileHashMetadata` for hash-based lookups and validation.

## Capabilities

### New Capabilities
- `rapid-upload`: Defines the interface and core requirements for drivers and the server to support rapid file uploads using file hashes.

### Modified Capabilities
- (None)

## Impact

- **Drivers**: Multiple drivers will need to implement the `RapidUploader` interface in a new `rapid.go` file.
- **Server**: `server/router.go` will be updated to include the `/put_rapid` endpoint.
- **Storage/FS**: New logic for handling rapid upload requests, including hash-based validation and comparison.
