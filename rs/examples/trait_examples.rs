use serde_json::json;
use zenbase_llml::{Prompt, llml};

// Example from the documentation
#[derive(Debug)]
struct User {
    id: u32,
    name: String,
    email: String,
}

impl Prompt for User {
    fn to_prompt(&self) -> String {
        format!(
            "<user id=\"{}\">\n  <name>{}</name>\n  <email>{}</email>\n</user>",
            self.id, self.name, self.email
        )
    }
}

#[derive(Debug)]
struct TaskStatus {
    task: String,
    completed: bool,
    priority: Priority,
}

#[derive(Debug)]
enum Priority {
    Low,
    Medium,
    High,
    Critical,
}

impl Prompt for Priority {
    fn to_prompt(&self) -> String {
        match self {
            Priority::Low => "🟢 low",
            Priority::Medium => "🟡 medium",
            Priority::High => "🟠 high",
            Priority::Critical => "🔴 critical",
        }
        .to_string()
    }
}

impl Prompt for TaskStatus {
    fn to_prompt(&self) -> String {
        let status_icon = if self.completed { "✅" } else { "⏳" };
        let priority = llml(&self.priority);

        format!(
            "<task status=\"{}\">\n  <description>{}</description>\n  <priority>{}</priority>\n</task>",
            status_icon, self.task, priority
        )
    }
}

fn main() {
    // Test basic user example
    let user = User {
        id: 123,
        name: "Alice".to_string(),
        email: "alice@example.com".to_string(),
    };

    println!("User example:");
    println!("{}", llml(&user));
    println!();

    // Test task status example
    let task = TaskStatus {
        task: "Implement trait system".to_string(),
        completed: true,
        priority: Priority::High,
    };

    println!("Task status example:");
    println!("{}", llml(&task));
    println!();

    // Test basic types
    println!("Basic types:");
    println!("String: {}", llml(&"Hello World"));
    println!("Number: {}", llml(&42));
    println!("Boolean: {}", llml(&true));
    println!();

    // Test JSON values
    let json_data = json!({
        "instructions": "Follow these steps",
        "rules": ["first", "second", "third"]
    });

    println!("JSON example:");
    println!("{}", llml(&json_data));
}
