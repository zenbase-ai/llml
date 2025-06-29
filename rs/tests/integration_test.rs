use llml::{llml, Options};
use serde_json::{json, Value};

#[test]
fn test_empty_values() {
    // Empty object
    let result = llml(&json!({}), None);
    assert_eq!(result, "");

    // Empty array
    let result = llml(&json!([]), None);
    assert_eq!(result, "");

    // Null value
    let result = llml(&Value::Null, None);
    assert_eq!(result, "null");
}

#[test]
fn test_empty_string_value() {
    let result = llml(&json!({"empty": ""}), None);
    assert_eq!(result, "<empty></empty>");
}

#[test]
fn test_zero_value() {
    let result = llml(&json!({"zero": 0}), None);
    assert_eq!(result, "<zero>0</zero>");
}

#[test]
fn test_false_boolean() {
    let result = llml(&json!({"disabled": false}), None);
    assert_eq!(result, "<disabled>false</disabled>");
}

#[test]
fn test_null_value() {
    let result = llml(&json!({"nothing": null}), None);
    assert_eq!(result, "<nothing>null</nothing>");
}

#[test]
fn test_simple_string_value() {
    let result = llml(&json!({"instructions": "Follow these steps"}), None);
    assert_eq!(result, "<instructions>Follow these steps</instructions>");
}

#[test]
fn test_integer_value() {
    let result = llml(&json!({"count": 42}), None);
    assert_eq!(result, "<count>42</count>");
}

#[test]
fn test_float_value() {
    let result = llml(&json!({"temperature": 98.6}), None);
    assert_eq!(result, "<temperature>98.6</temperature>");
}

#[test]
fn test_boolean_value() {
    let result = llml(&json!({"enabled": true}), None);
    assert_eq!(result, "<enabled>true</enabled>");
}

#[test]
fn test_kebab_case_conversion() {
    let result = llml(&json!({
        "user_name": "Alice",
        "userAge": 30
    }), None);
    
    // Since JSON object iteration order is not guaranteed, check both parts
    assert!(result.contains("<user-name>Alice</user-name>"));
    assert!(result.contains("<user-age>30</user-age>"));
}

#[test]
fn test_string_with_spaces_in_key() {
    let result = llml(&json!({"key with spaces": "value"}), None);
    assert_eq!(result, "<key-with-spaces>value</key-with-spaces>");
}

#[test]
fn test_multiple_simple_values() {
    let result = llml(&json!({
        "name": "Alice",
        "age": 30,
        "active": true
    }), None);
    
    // Check that all parts are present
    assert!(result.contains("<name>Alice</name>"));
    assert!(result.contains("<age>30</age>"));
    assert!(result.contains("<active>true</active>"));
    // Check for newline separators
    assert!(result.contains("\n"));
}

#[test]
fn test_basic_indentation() {
    let options = Some(Options {
        indent: "  ".to_string(),
        prefix: String::new(),
    });
    let result = llml(&json!({"message": "Hello"}), options);
    assert_eq!(result, "  <message>Hello</message>");
}

#[test]
fn test_with_prefix() {
    let options = Some(Options {
        indent: String::new(),
        prefix: "app".to_string(),
    });
    let result = llml(&json!({"config": "value"}), options);
    assert_eq!(result, "<app-config>value</app-config>");
}

#[test]
fn test_multiline_content() {
    let content = "\n    Line 1\n    Line 2\n    Line 3\n    ";
    let result = llml(&json!({"description": content}), None);
    assert_eq!(result, "<description>\n  Line 1\n  Line 2\n  Line 3\n</description>");
}

#[test]
fn test_empty_list() {
    let result = llml(&json!({"items": []}), None);
    assert_eq!(result, "<items-list></items-list>");
}

#[test]
fn test_simple_list_with_wrapper() {
    let result = llml(&json!({"rules": ["first", "second", "third"]}), None);
    let expected = "<rules-list>\n  <rules-1>first</rules-1>\n  <rules-2>second</rules-2>\n  <rules-3>third</rules-3>\n</rules-list>";
    assert_eq!(result, expected);
}

#[test]
fn test_list_with_numbers() {
    let result = llml(&json!({"numbers": [1, 2, 3]}), None);
    let expected = "<numbers-list>\n  <numbers-1>1</numbers-1>\n  <numbers-2>2</numbers-2>\n  <numbers-3>3</numbers-3>\n</numbers-list>";
    assert_eq!(result, expected);
}

#[test]
fn test_list_kebab_case_conversion() {
    let result = llml(&json!({"user_tasks": ["task1", "task2"]}), None);
    let expected = "<user-tasks-list>\n  <user-tasks-1>task1</user-tasks-1>\n  <user-tasks-2>task2</user-tasks-2>\n</user-tasks-list>";
    assert_eq!(result, expected);
}

#[test]
fn test_list_with_indentation() {
    let options = Some(Options {
        indent: "  ".to_string(),
        prefix: String::new(),
    });
    let result = llml(&json!({"items": ["a", "b"]}), options);
    let expected = "  <items-list>\n    <items-1>a</items-1>\n    <items-2>b</items-2>\n  </items-list>";
    assert_eq!(result, expected);
}

#[test]
fn test_list_with_prefix() {
    let options = Some(Options {
        indent: String::new(),
        prefix: "app".to_string(),
    });
    let result = llml(&json!({"items": ["a", "b"]}), options);
    let expected = "<app-items-list>\n  <app-items-1>a</app-items-1>\n  <app-items-2>b</app-items-2>\n</app-items-list>";
    assert_eq!(result, expected);
}

#[test]
fn test_nested_dict() {
    let result = llml(&json!({
        "config": {
            "debug": true,
            "timeout": 30
        }
    }), None);
    
    // Check for the main structure
    assert!(result.contains("<config>"));
    assert!(result.contains("</config>"));
    assert!(result.contains("<config-debug>true</config-debug>"));
    assert!(result.contains("<config-timeout>30</config-timeout>"));
}

#[test]
fn test_nested_dict_with_kebab_case() {
    let result = llml(&json!({
        "user_config": {
            "debug_mode": true,
            "maxRetries": 5
        }
    }), None);
    
    assert!(result.contains("<user-config>"));
    assert!(result.contains("</user-config>"));
    assert!(result.contains("<user-config-debug-mode>true</user-config-debug-mode>"));
    assert!(result.contains("<user-config-max-retries>5</user-config-max-retries>"));
}

#[test]
fn test_list_of_dicts() {
    let result = llml(&json!({
        "data": [
            {"name": "Alice", "age": 30},
            {"name": "Bob", "age": 25}
        ]
    }), None);
    
    // Check for the list structure
    assert!(result.contains("<data-list>"));
    assert!(result.contains("</data-list>"));
    assert!(result.contains("<data-1>"));
    assert!(result.contains("</data-1>"));
    assert!(result.contains("<data-2>"));
    assert!(result.contains("</data-2>"));
    
    // Check for nested data
    assert!(result.contains("<data-1-name>Alice</data-1-name>"));
    assert!(result.contains("<data-1-age>30</data-1-age>"));
    assert!(result.contains("<data-2-name>Bob</data-2-name>"));
    assert!(result.contains("<data-2-age>25</data-2-age>"));
}

#[test]
fn test_mixed_content() {
    let result = llml(&json!({
        "title": "My Document",
        "sections": ["intro", "body", "conclusion"],
        "metadata": {
            "author": "Alice",
            "version": "1.0"
        }
    }), None);
    
    // Check for all components
    assert!(result.contains("<title>My Document</title>"));
    assert!(result.contains("<sections-list>"));
    assert!(result.contains("<sections-1>intro</sections-1>"));
    assert!(result.contains("<sections-2>body</sections-2>"));
    assert!(result.contains("<sections-3>conclusion</sections-3>"));
    assert!(result.contains("</sections-list>"));
    assert!(result.contains("<metadata>"));
    assert!(result.contains("<metadata-author>Alice</metadata-author>"));
    assert!(result.contains("<metadata-version>1.0</metadata-version>"));
    assert!(result.contains("</metadata>"));
}

#[test]
fn test_deeply_nested() {
    let result = llml(&json!({
        "level1": {
            "level2": {
                "items": ["a", "b"]
            }
        }
    }), None);
    
    assert!(result.contains("<level1>"));
    assert!(result.contains("</level1>"));
    assert!(result.contains("<level1-level2>"));
    assert!(result.contains("</level1-level2>"));
    assert!(result.contains("<level1-level2-items-list>"));
    assert!(result.contains("<level1-level2-items-1>a</level1-level2-items-1>"));
    assert!(result.contains("<level1-level2-items-2>b</level1-level2-items-2>"));
    assert!(result.contains("</level1-level2-items-list>"));
}