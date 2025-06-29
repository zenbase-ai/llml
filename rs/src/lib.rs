/*!
# LLML - Lightweight Markup Language

Converts data structures to XML-like markup with specific formatting rules.

## Features

- Converts key-value pairs to XML tags: `{"key": "value"}` → `<key>value</key>`
- Formats arrays as numbered lists with wrapper tags
- Supports nested objects and kebab-case conversion
- Handles indentation, prefixes, and multiline strings

## Usage

```rust
use llml::{llml, Options};
use serde_json::json;

let data = json!({"instructions": "Follow these steps"});
let result = llml(&data, None);
// Output: "<instructions>Follow these steps</instructions>"
```
*/

use serde_json::Value;

mod formatters;
use formatters::format_value;

/// Configuration options for LLML formatting
#[derive(Debug, Clone)]
pub struct Options {
    /// Indentation string to use for nested elements
    pub indent: String,
    /// Prefix to prepend to all tags
    pub prefix: String,
}

impl Default for Options {
    fn default() -> Self {
        Options {
            indent: String::new(),
            prefix: String::new(),
        }
    }
}

/// Main LLML function - converts data structures to XML-like markup
///
/// Supports various call patterns:
/// - `llml(&Value::Null, None)` → `"null"`
/// - `llml(&json!([]), None)` → `""`
/// - `llml(&json!({}), None)` → `""`
/// - `llml(&json!({"key": "value"}), None)` → `"<key>value</key>"`
pub fn llml(data: &Value, options: Option<Options>) -> String {
    let opts = options.unwrap_or_default();
    format_value(data, &opts)
}

#[cfg(test)]
mod tests {
    use super::*;
    use serde_json::json;


    #[test]
    fn test_empty_values() {
        assert_eq!(llml(&json!({}), None), "");
        assert_eq!(llml(&json!([]), None), "");
        assert_eq!(llml(&Value::Null, None), "null");
    }

    #[test]
    fn test_simple_values() {
        let result = llml(&json!({"instructions": "Follow these steps"}), None);
        assert_eq!(result, "<instructions>Follow these steps</instructions>");

        let result = llml(&json!({"count": 42}), None);
        assert_eq!(result, "<count>42</count>");

        let result = llml(&json!({"enabled": true}), None);
        assert_eq!(result, "<enabled>true</enabled>");
    }

    #[test]
    fn test_list_formatting() {
        let result = llml(&json!({"rules": ["first", "second", "third"]}), None);
        let expected = "<rules-list>\n  <rules-1>first</rules-1>\n  <rules-2>second</rules-2>\n  <rules-3>third</rules-3>\n</rules-list>";
        assert_eq!(result, expected);
    }

    #[test]
    fn test_nested_objects() {
        let result = llml(&json!({
            "config": {
                "debug": true,
                "timeout": 30
            }
        }), None);

        assert!(result.contains("<config>"));
        assert!(result.contains("</config>"));
        assert!(result.contains("<config-debug>true</config-debug>"));
        assert!(result.contains("<config-timeout>30</config-timeout>"));
    }
}
