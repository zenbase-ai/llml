---
description:
globs: rs/*
alwaysApply: false
---
# Rust Coding Rules for LLML Project

These guidelines are designed to maintain the existing quality, style, and architectural patterns of the codebase.

## Project Structure & Modules

- **Public API (`src/lib.rs`):** The `src/lib.rs` file is the public face of the library. It should only contain the public API (`llml`, `llml_with_options`, `Options` struct) and its associated documentation and unit tests.
- **Core Logic (`src/formatters.rs`):** All internal formatting logic should reside in the `src/formatters.rs` module. Functions within this module should remain private to the crate (i.e., not marked with `pub`).
- **Tests:**
  - **Unit Tests:** Place unit tests within the module they are testing, using a `#[cfg(test)] mod tests { ... }` block.
  - **Integration Tests:** Place integration tests in the `tests/` directory. These tests should only call the public API from `lib.rs`.
- **Examples:** Add new usage examples to the `examples/` directory.

## Coding Style & Formatting

- **Formatting:** All code must be formatted with `rustfmt` using the default settings.
- **Naming Conventions:**
  - **Functions & Variables:** Use `snake_case` (e.g., `format_value`, `kebab_key`).
  - **Types (Structs, Enums):** Use `PascalCase` (e.g., `Options`).
  - **Constants & Statics:** Use `SCREAMING_SNAKE_CASE` (e.g., `MULTI_HYPHEN_RE`).
- **Clarity:** Prioritize clear, readable code. Use descriptive variable names.

## API Design & Patterns

- **Public Functions:** The primary public functions are `llml` (for default options) and `llml_with_options` (for custom formatting). Maintain this clear separation.
- **Configuration:** All configuration should be passed via the `Options` struct. When adding new configuration, extend this struct.
- **Optional Configuration:** Use `Option<Options>` for the `llml_with_options` function signature. Use `options.unwrap_or_default()` to handle the `None` case.
- **Immutability:** Prefer immutable variables (`let`) over mutable ones (`let mut`) unless mutability is strictly necessary (e.g., for builders or accumulators).
- **Performance:**
  - Pass complex types like `Value` and `Options` by reference (`&`) to avoid unnecessary clones.
  - For expensive initializations that can be shared (like `Regex`), use `std::sync::OnceLock`.

## Type Usage & Data Handling

- **Primary Data Type:** The library's core data input is `&serde_json::Value`. This ensures compatibility with any data that can be serialized to JSON.
- **String Handling:**
  - Use `format!` for constructing simple strings.
  - For building strings from multiple parts in a loop, collect parts into a `Vec<String>` and then use `parts.join("")` for efficiency.
  - Use `&str` as function arguments for string slices instead of `String` where possible.
- **Recursion:** The core formatting logic is recursive. New formatting rules should be integrated into the existing recursive `format_value` and `format_key_value` functions.

## Error Handling

- **No `Result` Types:** The library's public API returns a `String`. It does not return a `Result` as it's designed to format already-validated data structures. This convention should be maintained.
- **Panics:** Avoid panics in the library code. The only acceptable use is for a one-time initialization failure in `OnceLock` where a hardcoded value (like a regex pattern) is invalid, as this indicates a critical programmer error.

## Testing

- **Comprehensive Coverage:** Every new feature or bug fix must be accompanied by tests.
- **Unit Tests:** Test internal logic, such as individual formatting helpers (e.g., `to_kebab_case`), in the module's test block.
- **Integration Tests:** Test the public API from a user's perspective in `tests/integration_test.rs`. Cover a wide range of inputs, including edge cases (empty strings, empty collections, `null`, `false`, `0`).
- **Test Data:** Use the `serde_json::json!` macro to create `Value` instances for tests, as it is concise and readable.
- **Assertions:**
  - Use `assert_eq!` for exact string matches.
  - When testing the output of a `serde_json::Map` where key order is not guaranteed, use multiple `assert!(result.contains(...))` calls to verify that all expected parts are present in the output string.

## Documentation

- **Public API:** All public items (`lib.rs`) must have clear, comprehensive doc comments (`///`). Explain what the function does, its parameters, and provide a simple usage example in a `rustdoc` code block.
- **Module-Level Docs:** Use `/*! ... */` for module-level documentation that explains the purpose of the module.
- **Internal Comments:** Use comments (`//`) sparingly to explain the *why* behind complex or non-obvious code, not the *what*.
- **README:** Keep `rs/README.md` updated with any new features or API changes.
