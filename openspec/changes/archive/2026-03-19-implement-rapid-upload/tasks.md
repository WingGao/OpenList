## 1. Interface and API Setup

- [x] 1.1 Define `RapidUploader` interface in `internal/driver/driver.go`
- [x] 1.2 Add `/put_rapid` endpoint to `server/router.go`
- [x] 1.3 Implement the handler for `/put_rapid` in `server/handles/fsup.go` (or a new dedicated handle file)

## 2. Driver Implementations

- [x] 2.1 Implement `RapidUploader` for 115 driver in `drivers/115/rapid.go`
- [x] 2.2 Implement `RapidUploader` for Aliyundrive driver in `drivers/aliyundrive/rapid.go`
- [x] 2.3 Implement `RapidUploader` for Aliyundrive Open driver in `drivers/aliyundrive_open/rapid.go`
- [x] 2.4 Implement `RapidUploader` for 189 PC/TV drivers in `drivers/189pc/rapid.go` and `drivers/189_tv/rapid.go`
- [x] 2.5 Implement `RapidUploader` for 139 driver in `drivers/139/rapid.go`
- [x] 2.6 Implement `RapidUploader` for Quark drivers in `drivers/quark_open/rapid.go` and `drivers/quark_uc/rapid.go`
- [x] 2.7 Implement `RapidUploader` for PikPak driver in `drivers/pikpak/rapid.go`
- [x] 2.8 Implement `RapidUploader` for Baidu Netdisk driver in `drivers/baidu_netdisk/rapid.go`

## 3. Core Logic and Validation

- [x] 3.1 Update `internal/op/fs.go` (or equivalent) to handle the rapid upload logic, including driver detection
- [x] 3.2 Implement metadata validation (size/filename check) before calling driver's `RapidUpload`
- [x] 3.3 Ensure proper error handling for unsupported drivers (returning `errs.NotSupport`)

## 4. Verification

- [x] 4.1 Verify the `/put_rapid` API correctly identifies supported drivers
- [x] 4.2 Verify successful rapid upload with valid hash metadata
- [x] 4.3 Verify rejection/fallback when hash is not found or size mismatches
