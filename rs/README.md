# LLML Rust

Rust implementation of the Lightweight Markup Language (LLML) - a library that converts JSON data structures into XML-like markup with intelligent formatting.

## Features

- **Simple Key-Value Conversion**: `{"key": "value"}` → `<key>value</key>`
- **Array Formatting**: Arrays become numbered lists with wrapper tags
- **Nested Object Support**: Recursive processing with kebab-case key conversion
- **Configurable Indentation**: Customizable spacing and prefixes
- **Multiline String Handling**: Proper formatting for multiline content
- **Kebab-Case Conversion**: Automatic conversion of camelCase and snake_case keys
- **Strict Mode Control**: Choose whether nested properties include parent key prefixes

## Installation

Add this to your `Cargo.toml`:

```toml
[dependencies]
llml = "0.1.0"
```

## Quick Start

```rust
use llml::llml;
use serde_json::json;

fn main() {
    // Simple key-value pairs
    let data = json!({"instructions": "Follow these steps"});
    println!("{}", llml(&data, None));
    // Output: <instructions>Follow these steps</instructions>
    
    // Arrays become numbered lists
    let data = json!({"rules": ["first", "second", "third"]});
    println!("{}", llml(&data, None));
    // Output:
    // <rules>
    //   <rules-1>first</rules-1>
    //   <rules-2>second</rules-2>
    //   <rules-3>third</rules-3>
    // </rules>
}
```

## Advanced Usage

### Custom Formatting Options

```rust
use llml::{llml, Options};
use serde_json::json;

let data = json!({"message": "Hello World"});
let options = Some(Options {
    indent: "  ".to_string(),
    prefix: "app".to_string(),
    strict: false,
});

let result = llml(&data, options);
// Output: <app-message>Hello World</app-message>

// Example with strict mode
let data = json!({"config": {"debug": true, "timeout": 30}});
let options = Some(Options {
    indent: "".to_string(),
    prefix: "".to_string(),
    strict: true,
});
let result = llml(&data, options);
// Output: <config>
//           <config-debug>true</config-debug>
//           <config-timeout>30</config-timeout>
//         </config>

// Example with strict mode disabled (default)
let options = Some(Options {
    indent: "".to_string(),
    prefix: "".to_string(),
    strict: false,
});
let result = llml(&data, options);
// Output: <config>
//           <debug>true</debug>
//           <timeout>30</timeout>
//         </config>
```

### Handling Complex Data Structures

```rust
let complex_data = json!({
    "user_config": {
        "theme_mode": "dark",
        "max_retries": 5
    },
    "data_sources": [
        {"name": "db1", "active": true},
        {"name": "db2", "active": false}
    ]
});

let result = llml(&complex_data, None);
// Produces nested XML-like structure with kebab-case keys
```

## Key Transformations

| Input | Output |
|-------|--------|
| `user_name` | `user-name` |
| `maxRetries` | `max-retries` |
| `key with spaces` | `key-with-spaces` |

## Supported Data Types

- **Primitives**: Strings, numbers, booleans, null
- **Arrays**: Converted to numbered lists with wrapper tags
- **Objects**: Recursively processed with nested tag structure
- **Empty Values**: Handled gracefully (empty objects/arrays → empty strings)

## Examples

Run the included examples to see LLML in action:

```bash
cargo run --example basic_usage
```

This will demonstrate:
- Simple key-value formatting
- Array handling
- Nested objects
- Mixed content
- Custom indentation
- Empty value handling

## Testing

Run the comprehensive test suite:

```bash
cargo test
```

The library includes extensive tests covering:
- Edge cases (empty values, null, zero, false)
- Data type conversions
- Kebab-case transformations
- Nested structures
- Indentation and prefixes
- Multiline content

## Architecture

The library uses a recursive formatter that:
1. Processes JSON values based on their type
2. Converts keys to kebab-case format
3. Handles arrays as numbered lists with wrapper tags
4. Maintains proper indentation for nested structures
5. Preserves multiline string formatting

## Contributing

This is part of the LLML project. See the main repository for contribution guidelines.

## License

MIT