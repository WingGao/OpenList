## Context

Currently, the system lacks a unified interface for rapid uploads across different storage drivers. Some drivers (115, Aliyundrive, 189, etc.) have internal rapid upload logic, but it's often tied to the standard upload flow. There is no standalone API to attempt a rapid upload given only file metadata (hashes, size).

## Goals / Non-Goals

**Goals:**
- Define a standard `RapidUploader` interface.
- Expose a `/put_rapid` endpoint in the server to allow clients to attempt rapid uploads.
- Implement the interface for key drivers that already support hash-based rapid uploads.
- Ensure the API is flexible: it can work with just a hash, or with hash + filename/size for validation.

**Non-Goals:**
- Implementing rapid upload for drivers that do not natively support it via their APIs.
- Writing unit tests (as per the requirement in `wing/_aidoc/rapid-upload.md`).
- Modifying existing standard upload flows (though they may eventually use this interface).

## Decisions

- **Interface Definition**: The `RapidUploader` interface will be placed in `internal/driver/driver.go` (or a similar central location for driver interfaces) to ensure it is accessible to both drivers and the server logic.
- **Standalone Files**: Implementation for each driver will be in a new `rapid.go` file within the driver's directory to keep the code clean and separated from the main driver logic.
- **API Endpoint**: The `/put_rapid` endpoint will be added to `server/router.go` under the `_fs` group. It will handle the parsing of hash metadata and call the corresponding driver's `RapidUpload` method.
- **Validation Logic**: If both hash and size/filename are provided, the system should perform a check before calling the driver to ensure consistency, reducing unnecessary API calls to the storage provider.

## Risks / Trade-offs

- **[Risk] Driver Support Variability** 鈫 **[Mitigation]** Clearly document which drivers support this feature and return a specific error (e.g., `errs.NotSupport`) if a driver does not implement the interface.
- **[Risk] Hash Collisions** 鈫 **[Mitigation]** Rely on the underlying storage provider's hash-based deduplication logic, which usually employs multiple hashes or size checks.
- **[Trade-off] Metadata Requirements** 鈫 Some drivers might require more than just one type of hash (e.g., MD5 + SHA1). The `filehash.FileHashMetadata` struct should be robust enough to handle these cases.
