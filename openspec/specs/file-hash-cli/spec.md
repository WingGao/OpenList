# File Hash CLI

## Purpose
This specification defines the requirements for the command-line interface to calculate and manage file hashes.

## Requirements

### Requirement: Hash Calculation Command
The CLI SHALL provide a `hash` command to calculate and store file/directory metadata.

#### Scenario: Single File Hash Calculation
- **WHEN** the `hash` command is executed with a single file path
- **THEN** the system SHALL calculate its MD5, SHA1, and GCID and store them in the metadata database

#### Scenario: Directory Recursive Hash Calculation
- **WHEN** the `hash` command is executed with a directory path
- **THEN** the system SHALL recursively calculate and store metadata for all files within that directory

### Requirement: Hash Command JSON Output
The `hash` command SHALL support a `--json <path>` flag to export calculation results to a JSON file.

#### Scenario: Exporting Results to JSON
- **WHEN** the `hash` command is executed with `--json output.json`
- **THEN** the system SHALL write a JSON file containing the calculated metadata for each file processed

### Requirement: Human-Readable Hash Formatting
The CLI output SHALL display binary hashes as human-readable hex strings.

#### Scenario: Displaying MD5 in CLI
- **WHEN** a file's MD5 is displayed in the CLI output
- **THEN** it SHALL be shown as its 32-character hexadecimal representation
