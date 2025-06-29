use zenbase_llml::{llml, llml_with_options, Options};
use serde_json::json;

#[test]
fn should_handle_empty_arrays() {
    let result = llml(&json!({"items": []}));
    assert_eq!(result, "");
}

#[test]
fn should_format_simple_lists_with_wrapper_tags_and_numbered_items() {
    let result = llml(&json!({"rules": ["first", "second", "third"]}));
    let expected = "<rules-list>\n  <rules-1>first</rules-1>\n  <rules-2>second</rules-2>\n  <rules-3>third</rules-3>\n</rules-list>";
    assert_eq!(result, expected);
}

#[test]
fn should_format_lists_with_numeric_values() {
    let result = llml(&json!({"numbers": [1, 2, 3]}));
    let expected = "<numbers-list>\n  <numbers-1>1</numbers-1>\n  <numbers-2>2</numbers-2>\n  <numbers-3>3</numbers-3>\n</numbers-list>";
    assert_eq!(result, expected);
}

#[test]
fn should_convert_list_names_to_kebab_case() {
    let result = llml(&json!({"user_tasks": ["task1", "task2"]}));
    let expected = "<user-tasks-list>\n  <user-tasks-1>task1</user-tasks-1>\n  <user-tasks-2>task2</user-tasks-2>\n</user-tasks-list>";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_list_with_indentation() {
    let options = Some(Options {
        indent: "  ".to_string(),
        prefix: String::new(),
    });
    let result = llml_with_options(&json!({"items": ["a", "b"]}), options);
    let expected = "  <items-list>\n    <items-1>a</items-1>\n    <items-2>b</items-2>\n  </items-list>";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_list_with_prefix() {
    let options = Some(Options {
        indent: String::new(),
        prefix: "app".to_string(),
    });
    let result = llml_with_options(&json!({"items": ["a", "b"]}), options);
    let expected = "<app-items-list>\n  <app-items-1>a</app-items-1>\n  <app-items-2>b</app-items-2>\n</app-items-list>";
    assert_eq!(result, expected);
}