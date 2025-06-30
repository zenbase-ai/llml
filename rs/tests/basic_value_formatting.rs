use serde_json::{Value, json};
use zenbase_llml::llml;

#[test]
fn should_handle_null_values() {
    let result = llml(&Value::Null);
    assert_eq!(result, "null");
}

#[test]
fn should_handle_empty_object_values() {
    let result = llml(&json!({}));
    assert_eq!(result, "");
}

#[test]
fn should_handle_empty_list_values() {
    let result = llml(&json!([]));
    assert_eq!(result, "");
}

#[test]
fn should_handle_zero_numeric_values() {
    let result = llml(&json!({"zero": 0}));
    assert_eq!(result, "<zero>0</zero>");
}

#[test]
fn should_handle_false_boolean_values() {
    let result = llml(&json!({"disabled": false}));
    assert_eq!(result, "<disabled>false</disabled>");
}

#[test]
fn should_handle_null_values_in_object() {
    let result = llml(&json!({"nothing": null}));
    assert_eq!(result, "<nothing>null</nothing>");
}

#[test]
fn should_handle_empty_string_values() {
    let result = llml(&json!({"empty": ""}));
    assert_eq!(result, "<empty></empty>");
}

#[test]
fn should_format_simple_string_values() {
    let result = llml(&json!({"instructions": "Follow these steps"}));
    assert_eq!(result, "<instructions>Follow these steps</instructions>");
}

#[test]
fn should_format_integer_values() {
    let result = llml(&json!({"count": 42}));
    assert_eq!(result, "<count>42</count>");
}

#[test]
fn should_format_float_values() {
    let result = llml(&json!({"temperature": 98.6}));
    assert_eq!(result, "<temperature>98.6</temperature>");
}

#[test]
fn should_format_true_boolean_values() {
    let result = llml(&json!({"enabled": true}));
    assert_eq!(result, "<enabled>true</enabled>");
}

#[test]
fn should_handle_multiple_simple_values() {
    let result = llml(&json!({
        "name": "Alice",
        "age": 30,
        "active": true
    }));

    // Output order is deterministic and consistent across runs
    let expected = "<active>true</active>\n<age>30</age>\n<name>Alice</name>";
    assert_eq!(result, expected);
}
