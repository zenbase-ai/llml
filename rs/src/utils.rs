use crate::Options;
use regex::Regex;
use serde_json::Value;
use std::sync::OnceLock;

static MULTI_HYPHEN_RE: OnceLock<Regex> = OnceLock::new();

/// Unified recursive formatter that handles all LLML formatting cases
pub fn format_value(data: &Value, opts: &Options) -> String {
    match data {
        // Base cases for primitive values
        Value::Null => "null".to_string(),
        Value::Bool(b) => b.to_string(),
        Value::Number(n) => n.to_string(),
        Value::String(s) => format_string_content(s),

        // Direct array/object calls without context - return empty for consistency
        Value::Array(_) if opts.prefix.is_empty() => String::new(),
        Value::Object(obj) if opts.prefix.is_empty() && obj.is_empty() => String::new(),

        // Recursive cases for collections
        Value::Array(arr) => format_with_key(arr, opts, &opts.prefix),
        Value::Object(obj) => format_object_recursive(obj, opts),
    }
}

/// Recursively format an object by processing each key-value pair
fn format_object_recursive(obj: &serde_json::Map<String, Value>, opts: &Options) -> String {
    if obj.is_empty() {
        return String::new();
    }

    let mut parts = Vec::new();
    let entries: Vec<_> = obj.iter().collect();

    for (i, (key, value)) in entries.iter().enumerate() {
        if i > 0 {
            parts.push("\n".to_string());
        }

        // Recursive call: format this key-value pair
        let formatted = format_key_value(key, value, opts);
        parts.push(formatted);
    }

    parts.join("")
}

/// Format a single key-value pair recursively (like Python's llml(key, value) call)
fn format_key_value(key: &str, value: &Value, opts: &Options) -> String {
    let full_key = if opts.prefix.is_empty() {
        key.to_string()
    } else {
        format!("{}-{}", opts.prefix, key)
    };
    let kebab_key = to_kebab_case(&full_key);

    match value {
        // Recursive case: Array becomes a list with wrapper tag
        Value::Array(arr) => {
            let wrapper_tag = format!("{kebab_key}-list");

            if arr.is_empty() {
                return String::new();
            }

            let mut parts = Vec::new();
            parts.push(format!("{}<{}>\n", opts.indent, wrapper_tag));

            let inner_indent = format!("{}  ", opts.indent);

            for (i, item) in arr.iter().enumerate() {
                let item_tag = format!("{}-{}", kebab_key, i + 1);

                match item {
                    Value::Object(_) => {
                        // Recursive call: format nested object with new context
                        parts.push(format!("{inner_indent}<{item_tag}>\n"));
                        let nested_opts = Options {
                            indent: format!("{inner_indent}  "),
                            prefix: item_tag.clone(),
                        };
                        let dict_content = format_value(item, &nested_opts);
                        parts.push(dict_content);
                        parts.push(format!("\n{inner_indent}</{item_tag}>"));
                    }
                    _ => {
                        // Recursive call: format primitive item
                        let formatted = format_value(item, &Options::default());
                        parts.push(format!(
                            "{inner_indent}<{item_tag}>{formatted}</{item_tag}>"
                        ));
                    }
                }

                if i < arr.len() - 1 {
                    parts.push("\n".to_string());
                }
            }

            parts.push(format!("\n{}</{}>", opts.indent, wrapper_tag));
            parts.join("")
        }

        // Recursive case: Object becomes nested tags
        Value::Object(_) => {
            let nested_opts = Options {
                indent: format!("{}  ", opts.indent),
                prefix: full_key,
            };
            // Recursive call: format nested object
            let formatted = format_value(value, &nested_opts);

            if formatted.contains('\n') {
                format!(
                    "{}<{}>\n{}\n{}</{}>",
                    opts.indent, kebab_key, formatted, opts.indent, kebab_key
                )
            } else {
                format!(
                    "{}<{}>{}</{}>",
                    opts.indent, kebab_key, formatted, kebab_key
                )
            }
        }

        // Recursive case: Primitive values
        _ => {
            // Recursive call: format primitive value
            let formatted = format_value(value, &Options::default());
            if formatted.contains('\n') {
                format!(
                    "{}<{}>\n{}\n{}</{}>",
                    opts.indent, kebab_key, formatted, opts.indent, kebab_key
                )
            } else {
                format!(
                    "{}<{}>{}</{}>",
                    opts.indent, kebab_key, formatted, kebab_key
                )
            }
        }
    }
}

/// Helper function to handle array formatting when called with a specific key context
fn format_with_key(arr: &[Value], opts: &Options, prefix: &str) -> String {
    // This should not normally be called in the new recursive design,
    // but kept for compatibility. Arrays should be handled via format_key_value.
    if prefix.is_empty() {
        return String::new();
    }

    // Use the recursive approach
    format_key_value(
        prefix,
        &Value::Array(arr.to_vec()),
        &Options {
            indent: opts.indent.clone(),
            prefix: String::new(), // Reset prefix since we're handling it in the key
        },
    )
}

/// Format string content with proper multiline handling
fn format_string_content(s: &str) -> String {
    let trimmed = s.trim();
    if trimmed.contains('\n') {
        trimmed
            .lines()
            .map(|line| format!("  {}", line.trim()))
            .collect::<Vec<_>>()
            .join("\n")
    } else {
        trimmed.to_string()
    }
}

/// Convert text to kebab-case format
/// Handles spaces, underscores, and camelCase conversion
fn to_kebab_case(text: &str) -> String {
    if text.is_empty() {
        return text.to_string();
    }

    let mut result = String::new();
    let chars: Vec<char> = text.chars().collect();

    for (i, &ch) in chars.iter().enumerate() {
        if ch.is_ascii_uppercase() && i > 0 && chars[i - 1].is_ascii_lowercase() {
            result.push('-');
        }

        if ch == ' ' || ch == '_' {
            result.push('-');
        } else {
            result.push(ch.to_ascii_lowercase());
        }
    }

    // Clean up multiple consecutive hyphens
    let multi_hyphen_re = MULTI_HYPHEN_RE.get_or_init(|| Regex::new(r"-+").unwrap());
    let cleaned = multi_hyphen_re.replace_all(&result, "-");

    cleaned.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;
    use serde_json::json;

    #[test]
    fn test_kebab_case_conversion() {
        assert_eq!(to_kebab_case("user_name"), "user-name");
        assert_eq!(to_kebab_case("userAge"), "user-age");
        assert_eq!(to_kebab_case("key with spaces"), "key-with-spaces");
        assert_eq!(to_kebab_case("maxRetries"), "max-retries");
    }

    #[test]
    fn test_format_string_content() {
        assert_eq!(format_string_content("simple"), "simple");
        assert_eq!(format_string_content("  trimmed  "), "trimmed");

        let multiline = "\n    Line 1\n    Line 2\n    ";
        let expected = "  Line 1\n  Line 2";
        assert_eq!(format_string_content(multiline), expected);
    }

    #[test]
    fn test_recursive_formatting() {
        let opts = Options::default();

        // Test empty cases
        assert_eq!(format_value(&json!({}), &opts), "");
        assert_eq!(format_value(&json!([]), &opts), "");

        // Test primitive values
        assert_eq!(format_value(&json!("hello"), &opts), "hello");
        assert_eq!(format_value(&json!(42), &opts), "42");
        assert_eq!(format_value(&json!(true), &opts), "true");
        assert_eq!(format_value(&json!(null), &opts), "null");
    }

    #[test]
    fn test_key_value_formatting() {
        let opts = Options::default();

        // Test simple key-value
        let result = format_key_value("test", &json!("value"), &opts);
        assert_eq!(result, "<test>value</test>");

        // Test with array
        let result = format_key_value("items", &json!(["a", "b"]), &opts);
        let expected =
            "<items-list>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n</items-list>";
        assert_eq!(result, expected);
    }
}
