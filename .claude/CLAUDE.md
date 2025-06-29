# LLML Project

## Overview

LLML (Lightweight Markup Language) is a multi-language project that provides elegant data-to-markup transformation. It converts Python/JavaScript data structures into beautifully formatted XML-like markup with zero configuration and maximum flexibility.

## Project Structure

```
llml/
├── py/                       # Python implementation
│   ├── src/
│   │   ├── __init__.py      # Exports llml function
│   │   └── llml.py          # Core implementation
│   ├── tests/
│   │   ├── __init__.py
│   │   └── test_llml.py     # 25 comprehensive test cases
│   ├── pyproject.toml       # Modern Python packaging
│   ├── tox.ini             # Multi-version testing (3.8-3.13)
│   ├── uv.lock             # Dependency lock file
│   ├── README.md           # Comprehensive documentation
│   └── .venv/              # Virtual environment
├── ts/                      # TypeScript implementation
│   ├── index.ts            # TypeScript version
│   ├── package.json
│   └── [other TS files]
└── [root config files]
```

## Core Functionality

### Python Implementation (`py/`)

The main function is `llml()` which transforms data into structured markup:

```python
from llml import llml

# Simple values
llml(greeting="Hello World")
# → <greeting>Hello World</greeting>

# Lists become numbered items with wrapper tags
llml(rules=["first", "second", "third"])
# → <rules-list>
#     <rules-1>first</rules-1>
#     <rules-2>second</rules-2>
#     <rules-3>third</rules-3>
#   </rules-list>

# Nested structures
llml(config={"debug": True, "timeout": 30})
# → <config>
#     <debug>True</debug>
#     <timeout>30</timeout>
#   </config>
```

### Key Features

1. **Automatic kebab-case conversion**: `user_name` → `<user-name>`
2. **List formatting**: Lists get wrapper tags with numbered items
3. **Nested structures**: Dictionaries become nested XML-like structures
4. **Proper indentation**: 2-space indentation per level
5. **Multiline content**: Automatic `dedent` and formatting
6. **Prefix support**: Optional prefix parameter for namespacing
7. **Type safety**: Uses `beartype` for runtime type checking

## Testing & Quality

### Python Testing
- **25 comprehensive test cases** covering all functionality
- **91% code coverage**
- **Multi-version testing** with tox (Python 3.8-3.13)
- **Pytest** for test execution
- **All tests passing** ✅

### Test Categories
- Basic functionality (strings, numbers, booleans)
- List formatting with wrapper tags
- Nested dictionaries and structures
- Indentation handling
- Prefix functionality
- Edge cases and complex scenarios

## Development

### Python Development
```bash
cd py/
uv run pytest tests/ -v      # Run tests
uv run tox                   # Test across Python versions
uv add --dev [package]       # Add dev dependencies
```

### Dependencies
- **Runtime**: `beartype>=0.10.0` (type checking)
- **Dev**: `pytest`, `tox`, `pytest-cov`
- **Requires**: Python 3.8+

## Architecture

The Python implementation follows a clean, functional approach:

1. **`llml.py`**: Core implementation with recursive formatting
2. **`kebab_case()`**: String transformation utility
3. **`_format_value()`**: Handles individual value formatting
4. **`_format_list_with_tag()`**: Specialized list formatting
5. **Proper indentation logic**: Consistent 2-space indentation

## Recent Updates

- ✅ Reorganized project structure (Python in `py/`, TypeScript in `ts/`)
- ✅ Updated to export `llml` function (instead of `lml`)
- ✅ Implemented new list formatting rules (no singular/plural conversion)
- ✅ Added comprehensive test coverage (25 tests)
- ✅ Set up multi-version testing with tox
- ✅ Added proper kebab-case conversion
- ✅ Fixed multiline content handling with `dedent`

## Use Cases

- **AI Prompt Engineering**: Structure complex prompts
- **Configuration Generation**: Generate clean config files
- **Document Structure**: Create structured documents
- **API Documentation**: Format API specs
- **Data Serialization**: Convert data to readable markup

The project is production-ready with comprehensive testing and clean architecture.
