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
use zenbase_llml::{llml, llml_with_options, LLMLOptions};
use serde_json::json;

let data = json!({"instructions": "Follow these steps"});
let result = llml(&data);
// Output: "<instructions>Follow these steps</instructions>"

// For custom formatting:
let options = LLMLOptions { indent: "  ".to_string(), prefix: String::new(), strict: false };
let result = llml_with_options(&data, Some(options));
```
*/

use serde_json::Value;

mod utils;
use utils::format_value;

/// Configuration LLMLOptions for LLML formatting
#[derive(Debug, Clone, Default)]
pub struct LLMLOptions {
    /// Indentation string to use for nested elements
    pub indent: String,
    /// Prefix to prepend to all tags
    pub prefix: String,
    /// Whether to use strict mode (include parent keys as prefixes in nested objects)
    pub strict: bool,
}

/// Main LLML function - converts data structures to XML-like markup
///
/// Supports various call patterns:
/// - `llml(&Value::Null)` → `"null"`
/// - `llml(&json!([]))` → `""`
/// - `llml(&json!({}))` → `""`
/// - `llml(&json!({"key": "value"}))` → `"<key>value</key>"`
/// - `llml_with_options(&data, LLMLOptions)` → formatted with custom LLMLOptions
pub fn llml(data: &Value) -> String {
    format_value(data, &LLMLOptions::default())
}

/// LLML function with explicit LLMLOptions - use when you need custom formatting
///
/// Examples:
/// - `llml_with_options(&data, None)` → same as `llml(&data)`
/// - `llml_with_options(&data, Some(LLMLOptions))` → formatted with custom LLMLOptions
pub fn llml_with_options(data: &Value, options: Option<LLMLOptions>) -> String {
    let opts = options.unwrap_or_default();
    format_value(data, &opts)
}

#[cfg(test)]
mod tests {
    use super::*;
    use serde_json::json;

    #[test]
    fn test_empty_values() {
        assert_eq!(llml(&json!({})), "");
        assert_eq!(llml(&json!([])), "");
        assert_eq!(llml(&Value::Null), "null");
    }

    #[test]
    fn test_simple_values() {
        let result = llml(&json!({"instructions": "Follow these steps"}));
        assert_eq!(result, "<instructions>Follow these steps</instructions>");

        let result = llml(&json!({"count": 42}));
        assert_eq!(result, "<count>42</count>");

        let result = llml(&json!({"enabled": true}));
        assert_eq!(result, "<enabled>true</enabled>");
    }

    #[test]
    fn test_list_formatting() {
        let result = llml(&json!({"rules": ["first", "second", "third"]}));
        let expected = "<rules>\n  <rules-1>first</rules-1>\n  <rules-2>second</rules-2>\n  <rules-3>third</rules-3>\n</rules>";
        assert_eq!(result, expected);
    }

    #[test]
    fn test_nested_objects() {
        let result = llml(&json!({
            "config": {
                "debug": true,
                "timeout": 30
            }
        }));

        assert!(result.contains("<config>"));
        assert!(result.contains("</config>"));
        assert!(result.contains("<debug>true</debug>"));
        assert!(result.contains("<timeout>30</timeout>"));
    }

    #[test]
    fn test_optional_second_argument() {
        let data = json!({"instructions": "Follow these steps"});

        // Test new simple syntax (no second argument)
        let result1 = llml(&data);
        assert_eq!(result1, "<instructions>Follow these steps</instructions>");

        // Test with LLMLOptions using the new function
        let options = Some(LLMLOptions {
            indent: "  ".to_string(),
            prefix: String::new(),
            strict: false,
        });
        let result2 = llml_with_options(&data, options);
        assert_eq!(result2, "  <instructions>Follow these steps</instructions>");

        // Test backward compatibility (explicit None with new function)
        let result3 = llml_with_options(&data, None);
        assert_eq!(result3, "<instructions>Follow these steps</instructions>");

        // Simple function and explicit None should be identical
        assert_eq!(result1, result3);
    }

    #[test]
    fn test_with_LLMLOptions_function() {
        let data = json!({"test": "value"});

        // Test with indentation
        let options = LLMLOptions {
            indent: "    ".to_string(),
            prefix: String::new(),
            strict: false,
        };
        let result = llml_with_options(&data, Some(options));
        assert_eq!(result, "    <test>value</test>");

        // Test with prefix
        let options = LLMLOptions {
            indent: String::new(),
            prefix: "app".to_string(),
            strict: false,
        };
        let result = llml_with_options(&data, Some(options));
        assert_eq!(result, "<app-test>value</app-test>");
    }

    #[test]
    fn test_insertion_order_preservation() {
        // Test that json! macro and our formatter preserve key insertion order
        let result = llml(&json!({
            "first": "1st",
            "second": "2nd",
            "third": "3rd"
        }));

        // Verify exact order matches insertion order
        let expected = "<first>1st</first>\n<second>2nd</second>\n<third>3rd</third>";
        assert_eq!(result, expected);
    }

    #[test]
    fn test_deterministic_output() {
        // Test that output is deterministic across multiple runs
        let data = json!({
            "alpha": "value1",
            "beta": "value2",
            "gamma": "value3",
            "delta": "value4"
        });

        let result1 = llml(&data);
        let result2 = llml(&data);
        let result3 = llml(&data);

        // All results should be identical
        assert_eq!(result1, result2);
        assert_eq!(result2, result3);

        // And should have a predictable order (alphabetical in this case)
        let expected = "<alpha>value1</alpha>\n<beta>value2</beta>\n<delta>value4</delta>\n<gamma>value3</gamma>";
        assert_eq!(result1, expected);
    }
}
