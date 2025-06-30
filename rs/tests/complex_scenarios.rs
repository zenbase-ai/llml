use serde_json::json;
use zenbase_llml::llml;

#[test]
fn should_handle_mixed_content_types() {
    let result = llml(&json!({
        "title": "My Document",
        "sections": ["intro", "body", "conclusion"],
        "metadata": {
            "author": "Alice",
            "version": "1.0"
        }
    }));

    // Output order is deterministic and consistent across runs (strict: false default)
    let expected = "<metadata>\n  <author>Alice</author>\n  <version>1.0</version>\n</metadata>\n<sections>\n  <sections-1>intro</sections-1>\n  <sections-2>body</sections-2>\n  <sections-3>conclusion</sections-3>\n</sections>\n<title>My Document</title>";
    assert_eq!(result, expected);
}

#[test]
fn should_handle_deeply_nested_structures() {
    let result = llml(&json!({
        "level1": {
            "level2": {
                "items": ["a", "b"]
            }
        }
    }));

    // Verify exact deeply nested structure with preserved order (strict: false default)
    let expected = "<level1>\n  <level2>\n    <items>\n      <items-1>a</items-1>\n      <items-2>b</items-2>\n    </items>\n  </level2>\n</level1>";
    assert_eq!(result, expected);
}
