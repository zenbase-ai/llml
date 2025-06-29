---
description:
globs: go/*
alwaysApply: false
---
# Go Coding Rules for LLML Project

## Project Structure

The project follows a standard and clean structure for a Go library.

- **`go.mod`, `go.sum`**: Reside at the root of the Go project (`go/`). All dependency management should be handled here.
- **`pkg/`**: Contains the core library code. The main package is `llml` (`go/pkg/llml/`). Any new library code should be placed within the `pkg/` directory.
- **`tests/`**: Contains the integration and unit tests. Tests are in a separate `_test` package (e.g., `llml_test`), which is a best practice for testing only the exported API of a package.
- **`README.md`**: A dedicated README for the Go implementation exists and should be updated with any new features, API changes, or usage instructions.

## Dependencies

- **Main Library**: The core library in `pkg/` has **zero external dependencies** and relies only on the Go standard library. This should be maintained to keep the library lightweight and easy to integrate.
- **Testing**: Tests use `github.com/stretchr/testify/assert` for assertions. This is the only development dependency and should be used for all new tests to maintain consistency.

## Coding Style and Conventions

- **Formatting**: All code must be formatted with `gofmt`.
- **Naming**:
  - Packages should be named in lowercase (e.g., `llml`).
  - Exported identifiers (functions, types, variables) must start with a capital letter (e.g., `LLML`, `Options`).
  - Internal (private) identifiers must start with a lowercase letter (e.g., `formatMap`, `toKebabCase`).
  - Variable names should use `camelCase` (e.g., `kebabKey`, `anyMap`).
- **Comments**:
  - All exported functions and types must have a doc comment explaining their purpose, parameters, and return values.
  - Use comments to explain complex or non-obvious logic. Avoid comments that just restate what the code does.
- **Imports**: Imports should be organized into two blocks: standard library packages first, followed by third-party packages.

## Patterns and Idioms

- **Functional Options**: The `LLML` function accepts optional configuration via `opts ...Options`. This is a clean pattern for optional parameters and should be used if new configuration options are added.
- **Recursion for Data Structures**: The core logic uses recursion to traverse nested maps and slices. This pattern should be continued for handling nested data.
- **Type Handling**:
  - The public API uses `interface{}` (or its modern alias `any`) to accept arbitrary data structures. This provides flexibility for the user.
  - Internally, the code uses type assertions (e.g., `data.(map[string]any)`) and type switches to handle different data types. This is the established pattern for working with the `any` type.
- **Deterministic Output**: To ensure consistent output, map keys are sorted before processing. This is a critical feature and must be maintained. Any processing of maps must be done in a deterministic order.
- **Helper Functions**: The logic is cleanly separated into smaller, private helper functions with specific responsibilities (e.g., `formatMap`, `formatKeyValue`, `toKebabCase`). Complex logic should continue to be broken down this way.

## Type System Usage

- **Public API**: Continue to use `any` (`interface{}`) for the main `data` parameter to maintain flexibility.
- **Internal Types**: Use concrete types (`string`, `int`, `bool`, `map[string]any`, `[]any`) internally after type assertion.
- **Numeric Types**: The implementation correctly handles a wide range of specific integer and float types. Any new numeric handling should also be comprehensive.
- **Structs**: The `Options` struct is a good example of how to group related configuration parameters.

## Error Handling

The library currently does not return errors. For unsupported types, it falls back to a default string representation using `fmt.Sprintf("%v", data)`.

- **Guideline**: This approach is acceptable for a formatting library where the goal is to always produce a string. Avoid introducing `error` return values unless a new feature can fail in a way that the caller must handle (e.g., I/O operations, invalid configuration that cannot be defaulted).

## Testing

- **Framework**: All tests must use the standard `testing` package and `github.com/stretchr/testify/assert`.
- **Test Naming**: Test functions must follow the `Test<Name>` convention.
- **Test Structure**:
  - Each test function should focus on a specific feature or edge case.
  - For testing multiple variations of a feature, prefer creating separate, descriptively named tests (e.g., `TestEmptySlice`, `TestSimpleListWithWrapper`) as is the current convention. Table-driven tests are also an acceptable alternative for very similar inputs.
- **Assertions**:
  - Use functions from `assert` (e.g., `assert.Equal`, `assert.Contains`) for readable and expressive tests.
  - Since map keys are sorted, tests for map-based inputs should be deterministic. Use `assert.Equal` with the expected, correctly ordered string output. Avoid `assert.Contains` unless specifically testing for the presence of a substring in a larger, complex output.
- **Coverage**: All new code and features must be accompanied by comprehensive tests covering functionality, edge cases (e.g., `nil`, empty strings, zero values), and different data types.
