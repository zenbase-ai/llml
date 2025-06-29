use llml::{llml, Options};
use serde_json::json;

fn main() {
    println!("LLML Rust Examples\n");

    // Basic key-value pairs
    let simple_data = json!({
        "instructions": "Follow these steps",
        "count": 42,
        "enabled": true
    });
    println!("Simple data:");
    println!("{}\n", llml(&simple_data, None));

    // Lists
    let list_data = json!({
        "rules": ["first", "second", "third"],
        "numbers": [1, 2, 3]
    });
    println!("Lists:");
    println!("{}\n", llml(&list_data, None));

    // Nested objects
    let nested_data = json!({
        "config": {
            "debug": true,
            "timeout": 30
        },
        "user_settings": {
            "theme_mode": "dark",
            "maxRetries": 5
        }
    });
    println!("Nested objects:");
    println!("{}\n", llml(&nested_data, None));

    // Mixed content
    let mixed_data = json!({
        "title": "My Document",
        "sections": ["intro", "body", "conclusion"],
        "metadata": {
            "author": "Alice",
            "version": "1.0"
        }
    });
    println!("Mixed content:");
    println!("{}\n", llml(&mixed_data, None));

    // With indentation
    println!("With indentation:");
    let options = Some(Options {
        indent: "  ".to_string(),
        prefix: String::new(),
    });
    println!("{}\n", llml(&simple_data, options));

    // Empty values
    let empty_data = json!({
        "empty_string": "",
        "empty_array": [],
        "zero": 0,
        "false_value": false,
        "null_value": null
    });
    println!("Empty/falsy values:");
    println!("{}", llml(&empty_data, None));
}