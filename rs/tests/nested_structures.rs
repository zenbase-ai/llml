use zenbase_llml::llml;
use serde_json::json;

#[test]
fn should_handle_nested_objects() {
    let result = llml(&json!({
        "config": {
            "debug": true,
            "timeout": 30
        }
    }));

    // Verify exact nested structure with preserved order
    let expected = "<config>\n  <config-debug>true</config-debug>\n  <config-timeout>30</config-timeout>\n</config>";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_nested_objects_with_kebab_case_conversion() {
    let result = llml(&json!({
        "user_config": {
            "debug_mode": true,
            "maxRetries": 5
        }
    }));

    // Verify exact output with kebab-case conversion and preserved order
    let expected = "<user-config>\n  <user-config-debug-mode>true</user-config-debug-mode>\n  <user-config-max-retries>5</user-config-max-retries>\n</user-config>";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_arrays_containing_objects() {
    let result = llml(&json!({
        "data": [
            {"name": "Alice", "age": 30},
            {"name": "Bob", "age": 25}
        ]
    }));

    // Output order is deterministic and consistent across runs
    let expected = "<data-list>\n  <data-1>\n    <data-1-age>30</data-1-age>\n    <data-1-name>Alice</data-1-name>\n  </data-1>\n  <data-2>\n    <data-2-age>25</data-2-age>\n    <data-2-name>Bob</data-2-name>\n  </data-2>\n</data-list>";
    assert_eq!(result, expected);
}