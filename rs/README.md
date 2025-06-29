# LLML Rust

Rust implementation of the Lightweight Markup Language (LLML) library.

## Installation

Add this to your `Cargo.toml`:

```toml
[dependencies]
llml = "0.1.0"
```

## Usage

```rust
use llml::llml;
use serde_json::json;

fn main() {
    let data = json!({
        "instructions": "Follow these steps",
        "rules": ["first", "second", "third"],
        "config": {
            "debug": true,
            "timeout": 30
        }
    });
    
    let result = llml(&data, None);
    println!("{}", result);
}
```

## Running Tests

```bash
cargo test
```

## Examples

```bash
cargo run --example basic_usage
```