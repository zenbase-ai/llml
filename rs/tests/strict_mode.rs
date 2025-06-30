use serde_json::json;
use zenbase_llml::{LLMLOptions, llml, llml_with_options};

#[test]
fn should_use_non_strict_mode_by_default() {
    let result = llml(&json!({
        "config": {
            "debug": true,
            "timeout": 30
        }
    }));

    // Default behavior should not include parent key prefixes
    let expected = "<config>\n  <debug>true</debug>\n  <timeout>30</timeout>\n</config>";
    assert_eq!(result, expected);
}

#[test]
fn should_support_strict_mode_for_nested_objects() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: String::new(),
        strict: true,
    });
    let result = llml_with_options(
        &json!({
            "config": {
                "debug": true,
                "timeout": 30
            }
        }),
        options,
    );

    // Strict mode should include parent key prefixes
    let expected = "<config>\n  <config-debug>true</config-debug>\n  <config-timeout>30</config-timeout>\n</config>";
    assert_eq!(result, expected);
}

#[test]
fn should_support_strict_mode_with_kebab_case_conversion() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: String::new(),
        strict: true,
    });
    let result = llml_with_options(
        &json!({
            "user_config": {
                "debug_mode": true,
                "maxRetries": 5
            }
        }),
        options,
    );

    // Strict mode with kebab-case conversion
    let expected = "<user-config>\n  <user-config-debug-mode>true</user-config-debug-mode>\n  <user-config-max-retries>5</user-config-max-retries>\n</user-config>";
    assert_eq!(result, expected);
}

#[test]
fn should_support_strict_mode_for_arrays_containing_objects() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: String::new(),
        strict: true,
    });
    let result = llml_with_options(
        &json!({
            "data": [
                {"name": "Alice", "age": 30},
                {"name": "Bob", "age": 25}
            ]
        }),
        options,
    );

    // Strict mode should include array item prefixes for object properties
    let expected = "<data>\n  <data-1>\n    <data-1-age>30</data-1-age>\n    <data-1-name>Alice</data-1-name>\n  </data-1>\n  <data-2>\n    <data-2-age>25</data-2-age>\n    <data-2-name>Bob</data-2-name>\n  </data-2>\n</data>";
    assert_eq!(result, expected);
}

#[test]
fn should_support_non_strict_mode_for_arrays_containing_objects() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: String::new(),
        strict: false,
    });
    let result = llml_with_options(
        &json!({
            "data": [
                {"name": "Alice", "age": 30},
                {"name": "Bob", "age": 25}
            ]
        }),
        options,
    );

    // Non-strict mode should not include array item prefixes for object properties
    let expected = "<data>\n  <data-1>\n    <age>30</age>\n    <name>Alice</name>\n  </data-1>\n  <data-2>\n    <age>25</age>\n    <name>Bob</name>\n  </data-2>\n</data>";
    assert_eq!(result, expected);
}

#[test]
fn should_support_strict_mode_with_deeply_nested_structures() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: String::new(),
        strict: true,
    });
    let result = llml_with_options(
        &json!({
            "level1": {
                "level2": {
                    "items": ["a", "b"]
                }
            }
        }),
        options,
    );

    // Strict mode with deep nesting should include all parent key prefixes
    let expected = "<level1>\n  <level1-level2>\n    <level1-level2-items>\n      <level1-level2-items-1>a</level1-level2-items-1>\n      <level1-level2-items-2>b</level1-level2-items-2>\n    </level1-level2-items>\n  </level1-level2>\n</level1>";
    assert_eq!(result, expected);
}

#[test]
fn should_support_non_strict_mode_with_deeply_nested_structures() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: String::new(),
        strict: false,
    });
    let result = llml_with_options(
        &json!({
            "level1": {
                "level2": {
                    "items": ["a", "b"]
                }
            }
        }),
        options,
    );

    // Non-strict mode with deep nesting should not include parent key prefixes
    let expected = "<level1>\n  <level2>\n    <items>\n      <items-1>a</items-1>\n      <items-2>b</items-2>\n    </items>\n  </level2>\n</level1>";
    assert_eq!(result, expected);
}

#[test]
fn should_support_strict_mode_with_mixed_content() {
    let options = Some(LLMLOptions {
        indent: String::new(),
        prefix: String::new(),
        strict: true,
    });
    let result = llml_with_options(
        &json!({
            "title": "My Document",
            "sections": ["intro", "body", "conclusion"],
            "metadata": {
                "author": "Alice",
                "version": "1.0"
            }
        }),
        options,
    );

    // Strict mode with mixed content should include parent key prefixes for nested objects
    let expected = "<metadata>\n  <metadata-author>Alice</metadata-author>\n  <metadata-version>1.0</metadata-version>\n</metadata>\n<sections>\n  <sections-1>intro</sections-1>\n  <sections-2>body</sections-2>\n  <sections-3>conclusion</sections-3>\n</sections>\n<title>My Document</title>";
    assert_eq!(result, expected);
}
