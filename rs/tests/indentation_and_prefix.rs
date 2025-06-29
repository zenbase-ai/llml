use zenbase_llml::{llml_with_options, Options};
use serde_json::json;

#[test]
fn should_handle_basic_indentation_with_string_indent() {
    let options = Some(Options {
        indent: "  ".to_string(),
        prefix: String::new(),
    });
    let result = llml_with_options(&json!({"message": "Hello"}), options);
    assert_eq!(result, "  <message>Hello</message>");
}

#[test]
fn should_apply_prefix_to_simple_values() {
    let options = Some(Options {
        indent: String::new(),
        prefix: "app".to_string(),
    });
    let result = llml_with_options(&json!({"config": "value"}), options);
    assert_eq!(result, "<app-config>value</app-config>");
}