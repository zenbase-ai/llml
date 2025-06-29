use regex::Regex;
use serde_json::Value;
use std::sync::OnceLock;
use crate::Options;

static MULTI_HYPHEN_RE: OnceLock<Regex> = OnceLock::new();

pub fn format_value(data: &Value, opts: &Options) -> String {
    match data {
        Value::Null => "null".to_string(),
        Value::Bool(b) => b.to_string(),
        Value::Number(n) => n.to_string(),
        Value::String(s) => format_string(s, &opts.indent),
        Value::Array(_arr) => {
            // Direct array call without context - return empty
            if opts.prefix.is_empty() {
                String::new()
            } else {
                // This case should be handled by format_list when called from format_object
                String::new()
            }
        }
        Value::Object(obj) => format_object(obj, opts),
    }
}

pub fn format_object(obj: &serde_json::Map<String, Value>, opts: &Options) -> String {
    if obj.is_empty() {
        return String::new();
    }

    let mut parts = Vec::new();
    let entries: Vec<_> = obj.iter().collect();

    for (i, (key, value)) in entries.iter().enumerate() {
        if i > 0 {
            parts.push("\n".to_string());
        }

        let full_key = if opts.prefix.is_empty() {
            key.to_string()
        } else {
            format!("{}-{}", opts.prefix, key)
        };
        let kebab_key = to_kebab_case(&full_key);

        match value {
            Value::Array(arr) => {
                let formatted = format_list(arr, opts, &full_key);
                parts.push(format!("{}{}", opts.indent, formatted));
            }
            Value::Object(_) => {
                let nested_opts = Options {
                    indent: format!("{}  ", opts.indent),
                    prefix: full_key.clone(),
                };
                let formatted = format_value(value, &nested_opts);

                if formatted.contains('\n') {
                    parts.push(format!(
                        "{}<{}>\n{}\n{}</{}>",
                        opts.indent, kebab_key, formatted, opts.indent, kebab_key
                    ));
                } else {
                    parts.push(format!(
                        "{}<{}>{}</{}>",
                        opts.indent, kebab_key, formatted, kebab_key
                    ));
                }
            }
            _ => {
                let formatted = format_value(value, &Options::default());
                if formatted.contains('\n') {
                    parts.push(format!(
                        "{}<{}>\n{}\n{}</{}>",
                        opts.indent, kebab_key, formatted, opts.indent, kebab_key
                    ));
                } else {
                    parts.push(format!(
                        "{}<{}>{}</{}>",
                        opts.indent, kebab_key, formatted, kebab_key
                    ));
                }
            }
        }
    }

    parts.join("")
}

pub fn format_list(arr: &[Value], opts: &Options, prefix: &str) -> String {
    let kebab_prefix = to_kebab_case(prefix);
    let wrapper_tag = format!("{}-list", kebab_prefix);

    if arr.is_empty() {
        return format!("<{}></{}>", wrapper_tag, wrapper_tag);
    }

    let mut parts = Vec::new();
    parts.push(format!("<{}>\n", wrapper_tag));

    let inner_indent = format!("{}  ", opts.indent);

    for (i, item) in arr.iter().enumerate() {
        let item_tag = format!("{}-{}", kebab_prefix, i + 1);

        match item {
            Value::Object(_) => {
                parts.push(format!("{}<{}>\n", inner_indent, item_tag));
                let nested_opts = Options {
                    indent: format!("{}  ", inner_indent),
                    prefix: item_tag.clone(),
                };
                let dict_content = format_value(item, &nested_opts);
                parts.push(dict_content);
                parts.push(format!("\n{}</{}>", inner_indent, item_tag));
                if i < arr.len() - 1 {
                    parts.push("\n".to_string());
                }
            }
            _ => {
                let formatted = format_value(item, &Options::default());
                parts.push(format!(
                    "{}<{}>{}</{}>",
                    inner_indent, item_tag, formatted, item_tag
                ));
                if i < arr.len() - 1 {
                    parts.push("\n".to_string());
                }
            }
        }
    }

    parts.push(format!("\n{}</{}>", opts.indent, wrapper_tag));
    parts.join("")
}

pub fn format_string(s: &str, _indent: &str) -> String {
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
        if ch.is_ascii_uppercase() && i > 0 && chars[i-1].is_ascii_lowercase() {
            result.push('-');
        }
        
        if ch == ' ' || ch == '_' {
            result.push('-');
        } else {
            result.push(ch.to_ascii_lowercase());
        }
    }
    
    // Clean up multiple consecutive hyphens
    let multi_hyphen_re = MULTI_HYPHEN_RE.get_or_init(|| {
        Regex::new(r"-+").unwrap()
    });
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
    fn test_format_string() {
        assert_eq!(format_string("simple", ""), "simple");
        assert_eq!(format_string("  trimmed  ", ""), "trimmed");
        
        let multiline = "\n    Line 1\n    Line 2\n    ";
        let expected = "  Line 1\n  Line 2";
        assert_eq!(format_string(multiline, ""), expected);
    }

    #[test]
    fn test_format_list_empty() {
        let opts = Options::default();
        let result = format_list(&[], &opts, "items");
        assert_eq!(result, "<items-list></items-list>");
    }

    #[test]
    fn test_format_list_simple() {
        let opts = Options::default();
        let arr = vec![json!("first"), json!("second")];
        let result = format_list(&arr, &opts, "rules");
        let expected = "<rules-list>\n  <rules-1>first</rules-1>\n  <rules-2>second</rules-2>\n</rules-list>";
        assert_eq!(result, expected);
    }
}