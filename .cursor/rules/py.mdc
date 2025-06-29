---
description:
globs: py/*
alwaysApply: false
---
# Python Coding Rules for LLML Project

## Project Structure and Management

- **Configuration**: All project metadata, dependencies, and tool configurations must be managed in `pyproject.toml`.
- **Dependencies**:
  - Runtime dependencies should be added to the `[project.dependencies]` section.
  - Development dependencies must be added to the `[project.optional-dependencies]` section under the `dev` group.
  - Use `uv` for managing the virtual environment and installing dependencies.
- **Source Code**: All main application logic must reside in the `py/src/` directory.
- **Tests**: All tests must be placed in the `py/tests/` directory and follow the naming convention `test_*.py`.

## Coding Style and Formatting

- **Linter/Formatter**: `ruff` is the designated tool for all linting and formatting. All code must be compliant with the rules defined in `pyproject.toml`.
- **Line Length**: The maximum line length is `88` characters.
- **Quotes**: Use double quotes (`"`) for all strings. Single quotes are not permitted.
- **Indentation**: Use 4 spaces for indentation. Tabs are not allowed.
- **Docstrings**:
  - All public modules, functions, and methods must have a docstring.
  - Use triple-double quotes (`"""Docstring goes here"""`) for docstrings.
  - Function docstrings should clearly describe the function's purpose, arguments, and return value.

## Typing and Type Safety

- **Type Hinting**: All function signatures, including arguments and return values, must have type hints from Python's `typing` module.
- **Runtime Type Checking**: The `@beartype` decorator must be applied to all functions to enforce runtime type safety.
- **Type Aliases**: Use the `t` alias for the `typing` module (e.g., `import typing as t`).

## Function and API Design

- **Core Logic**: The main logic is centered around the `llml` function, which is designed to be recursive. When adding new features, maintain this recursive pattern.
- **Immutability**: The `llml` function should be pure and not modify its inputs. It should return a new string as the result.
- **Keyword Arguments**: When calling functions, prefer keyword arguments for clarity, especially for functions with multiple parameters.

## Error Handling

- **No Explicit Error Handling**: The current implementation does not have explicit `try...except` blocks. Instead, it relies on `beartype` for type-related errors and Python's standard behavior for other issues.
- **Implicit Error Handling**: Continue this pattern of relying on runtime checks and standard error-raising behavior unless a specific need for explicit handling arises.

## Testing

- **Framework**: All tests must be written using the `pytest` framework.
- **Test Naming**: Test functions must start with `test_` (e.g., `def test_my_feature():`).
- **Assertions**: Use `assert` statements for all checks. Do not use other assertion libraries.
- **Test Coverage**:
  - Every new feature or bug fix must be accompanied by corresponding tests.
  - Aim for comprehensive test coverage, including edge cases, basic functionality, and complex scenarios.
  - Run tests using `tox` to ensure compatibility across all supported Python versions.
- **Test Structure**: Each test function should be self-contained and focus on a single piece of functionality.

## Modularity and Code Organization

- **Separation of Concerns**:
  - The core formatting logic must remain in `llml.py`.
  - Utility functions, such as `kebab_case`, should be placed in `utils.py`.
- **Imports**:
  - Use absolute imports for modules within the project (e.g., `from .utils import kebab_case`).
  - Standard library imports should be placed at the top of the file, followed by third-party imports, and then local application imports.
