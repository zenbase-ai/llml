---
description:
globs: ts/*
alwaysApply: false
---
# TypeScript Coding Rules for LLML Project

## Overall Philosophy

The TypeScript implementation is a minimalist, functional, and robust library with zero production dependencies besides `dedent`. It prioritizes predictable output and developer experience through a simple API and strict typing, while gracefully handling varied and unpredictable input data.

## Project Structure & Conventions

### Tooling
- **Runtime and Package Management:** Use `bun` for all operations (installing dependencies, running scripts, building). The project is configured with `bun.lockb`.
- **Testing:** `vitest` is the testing framework. Use `bun test` to run the test suite.
- **Type Checking:** `typescript` is used for static analysis. Run `bun run tsc --noEmit` to check for type errors without generating JavaScript files.

### Module System
- The project uses **ES Modules (ESM)** exclusively, as defined by `"type": "module"` in `package.json`.
- Use `import`/`export` syntax. Do not use `require`/`module.exports`.

### File Organization
- **Source Code:** All core logic resides in the `src/` directory. The main entry point is `src/index.ts`.
- **Utilities:** Pure, reusable helper functions (like `kebabCase`) should be placed in `src/utils.ts`.
- **Tests:** All test files are located in the `tests/` directory. Maintain the one-to-one mapping of feature-to-test-file (e.g., `nested.ts` logic is tested in `tests/nested.test.ts`).

## Coding Style & Patterns

- **Functional Approach:** The core logic is implemented as a set of pure functions. The main `llml` function is recursive and delegates formatting tasks to helpers. Avoid classes and stateful logic.
- **Immutability:** Do not mutate input data. Functions should return new, transformed values.
- **Dependencies:** Maintain the zero-dependency principle for production code. The only exception is `dedent`, which is used for cleaning up multiline strings. Do not add new production dependencies without strong justification.
- **Readability:** Use `const` by default and `let` only when a variable must be reassigned. Code is formatted with a 2-space indent.
- **Comments:** Add JSDoc-style comments to exported functions and complex utility functions to explain their purpose, parameters, and return values, as seen in `src/utils.ts`.

## Type Usage & Safety

- **Strict Mode:** The project enforces `"strict": true` in `tsconfig.json`. All new code must adhere to strict type-checking rules.
- **Input Flexibility (`any`):** The main `llml` function accepts `any` as its primary data input. This is an intentional design choice to allow the library to serialize any valid JavaScript object structure.
- **Internal Type Guarding:** Because the input is `any`, it is **critical** to use type guards to safely handle the data internally. Use `typeof`, `Array.isArray`, `instanceof`, and `value == null` checks to narrow down the type before processing it.
- **Explicit Interfaces:** Define and export interfaces for configuration objects, as demonstrated by `LLMLOptions`.
- **Function Overloads:** Use function overloads for the main `llml` function to provide clearer type definitions and better autocompletion for consumers based on how they call it (e.g., with no arguments vs. with data).

## Error Handling

- **No-Throw Policy:** The library does not throw errors for invalid or unexpected input data (e.g., empty objects, `null` values).
- **Graceful Fallbacks:**
  - For empty objects (`{}`) or empty arrays (`[]`), return an empty string (`""`).
  - For `null` or `undefined` passed as the top-level data, return an empty string.
  - When `null` or `undefined` are values within an object, they should be stringified as `"null"` or `"undefined"`.
  - For primitive values passed directly to `llml`, return their string representation.

## Testing Approach

- **File-per-Feature:** Each distinct feature or area of concern should have its own test file in the `tests/` directory (e.g., `lists.test.ts`, `prefix.test.ts`).
- **Test Structure:** Use `describe` to group tests for a specific feature and `it` to define individual, atomic test cases. Test names should be descriptive.
- **Globals:** `vitest` is configured with `globals: true`, so you can use `describe`, `it`, and `expect` without importing them.
- **Comprehensive Coverage:** Ensure high test coverage by testing:
  - **Happy Path:** The expected output for valid inputs.
  - **Edge Cases:** `null`, `undefined`, empty strings, empty arrays, empty objects, `0`, and `false`.
  - **All Features:** Ensure every option and transformation rule (kebab-casing, lists, nesting, prefixes, indentation) is explicitly tested.
- **Assertions:** Use `expect(result).toBe(expected)` for clear, readable assertions. For multiline strings, construct the `expected` string using an array and `.join('\n')` for readability.
