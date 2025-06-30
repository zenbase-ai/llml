use serde_json::json;
use zenbase_llml::{LLMLOptions, llml, llml_with_options};

// Note: The current Rust implementation does not support direct array formatting
// Arrays passed as root elements return empty string for consistency
// This matches the current implementation behavior

#[test]
fn should_handle_direct_arrays_return_empty() {
    let result = llml(&json!(["a", "b", "c"]));
    let expected = "";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_arrays_with_different_types_return_empty() {
    let result = llml(&json!([1, "hello", true]));
    let expected = "";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_arrays_with_objects_return_empty() {
    let result = llml(&json!([{"name": "Alice"}, {"name": "Bob"}]));
    let expected = "";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_empty_arrays() {
    let result = llml(&json!([]));
    let expected = "";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_arrays_with_indentation_LLMLOptions_return_empty() {
    let options = Some(LLMLOptions {
        indent: "  ".to_string(),
        prefix: String::new(),
        strict: false,
    });
    let result = llml_with_options(&json!(["a", "b"]), options);
    let expected = "";
    assert_eq!(result, expected);
}
