use zenbase_llml::llml;
use serde_json::json;

#[test]
fn should_handle_multiline_content_with_dedent() {
    let content = "\n    Line 1\n    Line 2\n    Line 3\n    ";
    let result = llml(&json!({"description": content}));
    assert_eq!(
        result,
        "<description>\n  Line 1\n  Line 2\n  Line 3\n</description>"
    );
}