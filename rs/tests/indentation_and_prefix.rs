use zenbase_llml::{llml_with_options, LLMLOptions};
use serde_json::json;

#[test]
fn should_handle_basic_indentation_with_string_indent() {
    let options = Some(LLMLOptions {
        indent: "  ".to_string(),
        prefix: String::new(),
        strict: false,
    });
    let result = llml_with_options(&json!({"message": "Hello"}), options);
    assert_eq!(result, "  <message>Hello</message>");
}

#[test]
fn should_apply_prefix_to_simple_values() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: "app".to_string(),
        strict: false,
    });
    let result = llml_with_options(&json!({"config": "value"}), options);
    assert_eq!(result, "<app-config>value</app-config>");
}
