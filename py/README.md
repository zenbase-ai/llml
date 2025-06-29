# 🚀 LLML - Lightweight Language Markup Language

**The most elegant way to generate structured text in Python.**

LLML transforms your data into beautifully formatted XML-like markup with zero fuss and maximum flexibility. Perfect for prompt engineering, configuration generation, and structured document creation.

## ⚡ Quick Start

```python
from zenbase_llml import llml

# Simple values
print(llml(greeting="Hello World"))
# Output: <greeting>Hello World</greeting>

# Lists become numbered items
print(llml(tasks=["Buy milk", "Walk dog", "Code LLML"]))
# Output:
# <tasks>
#   <tasks-1>Buy milk</tasks-1>
#   <tasks-2>Walk dog</tasks-2>
#   <tasks-3>Code LLML</tasks-3>
# </tasks>

# Complex nested structures
print(llml(
    title="My Project",
    features=["Fast", "Simple", "Powerful"],
    config={"debug": True, "version": "1.0"}
))
```

## 🎯 Why LLML?

- **🔥 Zero Learning Curve**: One function, infinite possibilities
- **🎨 Beautiful Output**: Automatically formatted, properly indented
- **🔧 Type Safe**: Built with beartype for runtime type checking
- **⚡ Lightning Fast**: Minimal overhead, maximum performance
- **🌟 Pythonic**: Feels natural, works everywhere
- **⚙️ Strict Mode**: Control nested property prefixes with `strict` parameter

## 🛠️ Installation

```bash
pip install zenbase-llml
```

## 📚 Advanced Usage

### Prefix Support
```python
# Add prefix to all keys
print(llml(message="Hello", prefix="app"))
# Output: <app-message>Hello</app-message>
```

### Multi-line Content
```python
instructions = """
Step 1: Install LLML
Step 2: Import lml
Step 3: Create magic
"""

llml(instructions=instructions)
# Output:
# <instructions>
# Step 1: Install LLML
# Step 2: Import lml
# Step 3: Create magic
# </instructions>
```

### Complex Nested Structures
```python
prompt_data = {
    "system": "You are a helpful assistant",
    "rules": [
        "Be concise and clear",
        "Provide examples when helpful",
        "Ask clarifying questions"
    ],
    "context": {
        "user_level": "beginner",
        "topic": "Python programming"
    }
}

print(llml(**prompt_data))

# Example with strict mode enabled
print(llml(config={"debug": True, "timeout": 30}, strict=True))
# Output: <config>
#           <config-debug>True</config-debug>
#           <config-timeout>30</config-timeout>
#         </config>

# Example with strict mode disabled (default)
print(llml(config={"debug": True, "timeout": 30}, strict=False))
# Output: <config>
#           <debug>True</debug>
#           <timeout>30</timeout>
#         </config>
```

## 🎪 Use Cases

### 🤖 AI Prompt Engineering
Perfect for structuring complex prompts:
```python
prompt = llml(
    role="Senior Python Developer",
    task="Code review the following function",
    criteria=["Performance", "Readability", "Best practices"],
    code=function_to_review
)
```

### ⚙️ Configuration Generation
Generate clean config files:
```python
config = llml(
    database={"host": "localhost", "port": 5432},
    features=["logging", "caching", "monitoring"],
    environment="production"
)
```

### 📄 Document Structure
Create structured documents:
```python
document = llml(
    title="API Documentation",
    sections=["Authentication", "Endpoints", "Examples"],
    metadata={"version": "2.1", "author": "Dev Team"}
)
```

## 🧪 Testing

Run the comprehensive test suite:

```bash
# Run tests
uv run pytest

# Run with coverage
uv run pytest --cov=src --cov-report=html

# Test across Python versions
tox
```

## 🌐 Python Compatibility

LLML supports Python 3.8+ and is tested against:
- ✅ Python 3.8
- ✅ Python 3.9
- ✅ Python 3.10
- ✅ Python 3.11
- ✅ Python 3.12
- ✅ Python 3.13

## 🏗️ Development

```bash
# Clone the repo
git clone https://github.com/yourusername/llml.git
cd llml/py

# Create virtual environment and install dependencies
uv venv
source .venv/bin/activate
uv pip install -e .[dev]

# Run linting and formatting
uv run ruff check .
uv run ruff format .

# Run tests
uv run pytest
```

## 🤝 Contributing

We love contributions! Whether it's:
- 🐛 Bug reports
- 💡 Feature requests
- 📝 Documentation improvements
- 🔧 Code contributions

Check out our [contribution guidelines](CONTRIBUTING.md) to get started.

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🌟 Star History

If LLML makes your life easier, give us a star! ⭐

---

**Made with ❤️ for the Python community**

*LLML: Because beautiful markup shouldn't be hard.*
