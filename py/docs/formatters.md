# LLML Formatters Guide

Learn how to use and extend LLML's powerful formatter system to customize data serialization.

## Overview

LLML uses a **formatter system** to convert different data types into XML-like markup. Each formatter consists of:

1. **Predicate function**: Tests if a value matches a specific type
2. **Format function**: Converts the value to a string representation

When you call `llml(data)`, it iterates through the formatters in order until it finds one whose predicate returns `True`, then uses that formatter's format function.

```python
from zenbase_llml import llml

# Uses default VibeXML formatters
result = llml({
    "name": "Alice",
    "age": 30,
    "active": True,
    "tasks": ["code", "test", "deploy"]
})
```

## How Formatters Work

### The Formatter Structure

Each formatter is a tuple of `(predicate, format_function)`:

```python
def is_string(value) -> bool:
    return isinstance(value, str) and "\n" not in value

def format_string(value, llml, formatters=None) -> str:
    return value.strip()

string_formatter = (is_string, format_string)
```

### Default VibeXML Formatters

The `vibe_xml()` function creates a formatter map with built-in formatters:

```python
from zenbase_llml import llml
from zenbase_llml.formatters import vibe_xml

# Get the default formatter map
formatters = vibe_xml()
print(type(formatters))  # <class 'dict'>

# Use custom formatters
result = llml(data, formatters)
```

The default formatters handle:
- **Strings**: Simple text (multiline strings handled specially)
- **Numbers**: `int`, `float`
- **Booleans**: `True`/`False`
- **None**: Converts to `"None"`
- **Dates**: `datetime.datetime`, `datetime.timedelta`
- **Collections**: `list`, `dict`
- **Fallback**: Any other type via `str()`

## Extending Formatters

### Method 1: Modify Existing Formatter Map

Replace or add formatters to the default map:

```python
from datetime import datetime
from zenbase_llml import llml
from zenbase_llml.formatters.vibe_xml import vibe_xml, is_datetime

def format_datetime_custom(value, _llml, _formatters=None):
    return value.strftime("%Y-%m-%d %H:%M:%S")

# Get default formatters and override datetime formatting
formatters = {
    is_datetime: format_datetime_custom
    **vibe_xml,
}

# Use with custom datetime formatting
data = {"created": datetime(2023, 12, 25, 10, 30)}
result = llml(data, formatters)
# Output: <created>2023-12-25 10:30:00</created>
```

### Method 2: Create New Formatter Map

Put your custom formatters first (higher priority):

```python
class User:
    def __init__(self, name, email):
        self.name = name
        self.email = email

def is_user(value):
    return isinstance(value, User)

def format_user(value, llml, formatters=None):
    return f"{value.name} ({value.email})"

# Insertion order is important! This takes precedence so User doesn't hit the fallback __str__
formatters = {
    is_user: format_user,
    **vibe_xml,  # Include all default formatters
}

data = {"admin": User("Alice", "alice@example.com")}
result = llml(data, formatters)
# Output: <admin>Alice (alice@example.com)</admin>
```

## Custom Formatter Examples

### Advanced Date Formatting

```python
from zenbase_llml.formatters.vibe_xml import vibe_xml, is_date, is_datetime

from datetime import datetime, date

def format_date_iso(value, _llml, _formatters=None):
    return value.isoformat()

def format_datetime_readable(value, _llml, _formatters=None):
    return value.strftime("%B %d, %Y at %I:%M %p")

formatters = {
    is_date: format_date_iso,
    is_datetime: format_datetime_readable,
    **vibe_xml,
}

data = {
    "birth_date": date(1990, 5, 15),
    "last_login": datetime(2023, 12, 25, 14, 30)
}

result = llml(data, formatters)
# Output:
# <birth_date>1990-05-15</birth_date>
# <last_login>December 25, 2023 at 02:30 PM</last_login>
```

### URL and Email Formatting

```python
import re

def is_email(value):
    if not isinstance(value, str):
        return False
    return re.match(r'^[^@]+@[^@]+\.[^@]+$', value) is not None

def format_email(value, _llml, _formatters=None):
    return f"mailto:{value}"

def is_url(value):
    if not isinstance(value, str):
        return False
    return value.startswith(('http://', 'https://'))

def format_url(value, _llml, _formatters=None):
    return f"[{value}]"

# These take precedence over the builtin is_str / format_str
formatters = {
    is_email: format_email,
    is_url: format_url,
    **vibe_xml,
}

data = {
    "contact": "alice@example.com",
    "website": "https://example.com",
    "name": "Alice"  # Regular string, uses default formatter
}

result = llml(data, formatters)
# Output:
# <contact>mailto:alice@example.com</contact>
# <website>[https://example.com]</website>
# <name>Alice</name>
```

### Number Formatting with Units

```python
class Temperature:
    def __init__(self, celsius):
        self.celsius = celsius

    @property
    def fahrenheit(self):
        return (self.celsius * 9/5) + 32

class Distance:
    def __init__(self, meters):
        self.meters = meters

    @property
    def kilometers(self):
        return self.meters / 1000

def format_temperature(value, _llml, _formatters=None):
    return f"{value.celsius}°C ({value.fahrenheit:.1f}°F)"

def format_distance(value, _llml, _formatters=None):
    if value.meters >= 1000:
        return f"{value.kilometers:.2f} km"
    return f"{value.meters} m"

formatters = {
    (lambda value: isinstance(value, Temperature), format_temperature),
    (lambda value: isinstance(value, Distance), format_distance),
    **vibe_xml,
}

data = {
    "temperature": Temperature(25),
    "distance": Distance(1500)
}

result = llml(data, formatters)
# Output:
# <temperature>25°C (77.0°F)</temperature>
# <distance>1.50 km</distance>
```

## Advanced Usage

### Recursive Formatting

Format functions receive the `llml` function for recursive processing:

```python
class Project:
    def __init__(self, name, tasks):
        self.name = name
        self.tasks = tasks

def is_project(value):
    return isinstance(value, Project)

def format_project(value, llml, formatters):
    # Use the llml function recursively for the tasks
    tasks_formatted = llml({"tasks": value.tasks}, formatters)
    return f"Project: {value.name}\n{tasks_formatted}"

formatters = {
    is_project: format_project,
    **vibe_xml(),
}

project = Project("LLML", ["implement", "test", "document"])
result = llml({"current": project}, formatters)
# Output:
# <current>Project: LLML
# <tasks>
#   <tasks-1>implement</tasks-1>
#   <tasks-2>test</tasks-2>
#   <tasks-3>document</tasks-3>
# </tasks></current>
```

### Conditional Formatting

Create formatters that behave differently based on context:

```python
def is_sensitive_string(value):
    if not isinstance(value, str):
        return False
    sensitive_words = ["password", "token", "secret", "key"]
    return any(word in value.lower() for word in sensitive_words)

def format_sensitive(value, llml, formatters=None):
    return "[REDACTED]"

def is_large_number(value):
    return isinstance(value, (int, float)) and abs(value) >= 1000000

def format_large_number(value, llml, formatters=None):
    if value >= 1000000:
        return f"{value/1000000:.1f}M"
    elif value >= 1000:
        return f"{value/1000:.1f}K"
    return str(value)

formatters = {
    is_sensitive_string: format_sensitive,
    is_large_number: format_large_number,
    **vibe_xml(),
}

data = {
    "api_key": "secret_12345",
    "user_count": 1500000,
    "name": "Alice"
}

result = llml(data, formatters)
# Output:
# <api_key>[REDACTED]</api_key>
# <user_count>1.5M</user_count>
# <name>Alice</name>
```

## Order Matters

Formatters are processed in order - the first matching predicate wins:

```python
def is_positive_number(value):
    return isinstance(value, (int, float)) and value > 0

def format_positive(value, llml, formatters=None):
    return f"+{value}"

def is_number(value):
    return isinstance(value, (int, float))

def format_number(value, llml, formatters=None):
    return str(value)

# Correct order: more specific first
formatters = {
    is_positive_number: format_positive,  # More specific
    is_number: format_number,             # Less specific
}

result = llml({"score": 85}, formatters)
# Output: <score>+85</score>

# Wrong order: general formatter catches everything
wrong_formatters = {
    is_number: format_number,             # Too general, catches everything
    is_positive_number: format_positive,  # Never reached!
}

result = llml({"score": 85}, wrong_formatters)
# Output: <score>85</score> (positive formatter never used)
```

## Best Practices

### 1. Specific Before General
Always put more specific formatters before general ones.

### 2. Use Type Guards
Write clear, efficient predicate functions:

```python
def is_non_empty_string(value):
    return isinstance(value, str) and len(value) > 0

def is_email_address(value):
    return (isinstance(value, str) and
            "@" in value and
            "." in value.split("@")[-1])
```

### 3. Handle Edge Cases
Make your formatters robust:

```python
def format_safe_url(value, llml, formatters=None):
    if not value:
        return "[empty URL]"

    try:
        # Validate URL format
        if not value.startswith(('http://', 'https://')):
            value = f"https://{value}"
        return f"[{value}]"
    except Exception:
        return f"[invalid URL: {value}]"
```

### 4. Document Your Formatters
Make it clear what each formatter does:

```python
def is_currency_amount(value):
    """Detects numeric values that should be formatted as currency."""
    return isinstance(value, (int, float)) and hasattr(value, '_currency_type')

def format_currency(value, llml, formatters=None):
    """Formats currency amounts with appropriate symbol and precision."""
    currency_type = getattr(value, '_currency_type', 'USD')
    symbols = {'USD': '$', 'EUR': '€', 'GBP': '£'}
    symbol = symbols.get(currency_type, '$')
    return f"{symbol}{value:.2f}"
```

## Complete Example: Custom Data Types

Here's a complete example showing how to handle custom data types:

```python
from datetime import datetime
from zenbase_llml import llml
from zenbase_llml.formatters import vibe_xml

class BlogPost:
    def __init__(self, title, author, published, tags):
        self.title = title
        self.author = author
        self.published = published
        self.tags = tags

class Author:
    def __init__(self, name, email):
        self.name = name
        self.email = email

# Custom formatters
def is_blog_post(value):
    return isinstance(value, BlogPost)

def format_blog_post(value, llml, formatters):
    # Format the blog post with nested data
    post_data = {
        "title": value.title,
        "author": value.author,
        "published": value.published,
        "tags": value.tags
    }
    return llml(post_data, formatters)

def is_author(value):
    return isinstance(value, Author)

def format_author(value, llml, formatters=None):
    return f"{value.name} <{value.email}>"

def is_datetime(value):
    return isinstance(value, datetime)

def format_blog_datetime(value, llml, formatters=None):
    return value.strftime("%B %d, %Y")

# Create formatter map
formatters = {
    is_blog_post: format_blog_post,
    is_author: format_author,
    is_datetime: format_blog_datetime,
    **vibe_xml(),
}

# Use the formatters
post = BlogPost(
    title="LLML Formatters Guide",
    author=Author("Alice", "alice@example.com"),
    published=datetime(2023, 12, 25),
    tags=["llml", "python", "formatters"]
)

result = llml({"post": post}, formatters)
print(result)

# Output:
# <post>
#   <title>LLML Formatters Guide</title>
#   <author>Alice <alice@example.com></author>
#   <published>December 25, 2023</published>
#   <tags>
#     <tags-1>llml</tags-1>
#     <tags-2>python</tags-2>
#     <tags-3>formatters</tags-3>
#   </tags>
# </post>
```

## Next Steps

- Explore the built-in formatters in `zenbase_llml.formatters.vibe_xml`
- Create your own formatter library for common data types
- Consider performance implications for large datasets
- Share useful formatters with the community!

The formatter system makes LLML incredibly flexible while maintaining the clean, readable output format. Happy formatting! 🎉
