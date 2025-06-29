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

    // Verify exact nested structure with preserved order (strict: false default)
    let expected = "<config>\n  <debug>true</debug>\n  <timeout>30</timeout>\n</config>";
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

    // Verify exact output with kebab-case conversion and preserved order (strict: false default)
    let expected = "<user-config>\n  <debug-mode>true</debug-mode>\n  <max-retries>5</max-retries>\n</user-config>";
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

    // Output order is deterministic and consistent across runs (strict: false default)
    let expected = "<data>\n  <data-1>\n    <age>30</age>\n    <name>Alice</name>\n  </data-1>\n  <data-2>\n    <age>25</age>\n    <name>Bob</name>\n  </data-2>\n</data>";
    assert_eq!(result, expected);
}