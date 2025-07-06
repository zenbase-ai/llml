import json as stdjson
from dataclasses import dataclass
from datetime import datetime, timezone

import pytest

from zenbase_llml import llml
from zenbase_llml.formatters.json import json

try:
    import orjson

    HAS_ORJSON = True
except ImportError:
    HAS_ORJSON = False


class TestJSONFormatter:
    """Test JSON formatter functionality."""

    def test_basic_json_formatting(self):
        """Test basic JSON formatting with simple data."""
        result = llml({"user": "alice", "count": 42}, json)
        expected = (
            '{"user": "alice", "count": 42}'
            if not HAS_ORJSON
            else '{"user":"alice","count":42}'
        )
        # Parse and compare to handle different JSON formatting
        assert stdjson.loads(result) == stdjson.loads(expected)

    def test_nested_objects_json(self):
        """Test JSON formatting with nested objects."""
        data = {
            "user": {
                "name": "Alice",
                "settings": {"theme": "dark", "notifications": True},
            },
            "version": 2.0,
        }
        result = llml(data, json)
        # Parse and compare to ensure correct structure
        parsed = stdjson.loads(result)
        assert parsed["user"]["name"] == "Alice"
        assert parsed["user"]["settings"]["theme"] == "dark"
        assert parsed["user"]["settings"]["notifications"] is True
        assert parsed["version"] == 2.0

    def test_array_json_formatting(self):
        """Test JSON formatting with arrays."""
        data = {
            "items": ["apple", "banana", "cherry"],
            "numbers": [1, 2, 3, 4, 5],
            "mixed": [1, "two", True, None],
        }
        result = llml(data, json)
        parsed = stdjson.loads(result)
        assert parsed["items"] == ["apple", "banana", "cherry"]
        assert parsed["numbers"] == [1, 2, 3, 4, 5]
        assert parsed["mixed"] == [1, "two", True, None]

    def test_direct_array_json(self):
        """Test JSON formatting with direct array input."""
        data = ["first", "second", "third"]
        result = llml(data, json)
        assert stdjson.loads(result) == ["first", "second", "third"]

    def test_empty_structures_json(self):
        """Test JSON formatting with empty structures."""
        # Empty dict
        result = llml({}, json)
        assert result == "{}"

        # Empty list
        result = llml([], json)
        assert result == "[]"

        # Dict with empty values
        data = {"empty_list": [], "empty_dict": {}}
        result = llml(data, json)
        parsed = stdjson.loads(result)
        assert parsed["empty_list"] == []
        assert parsed["empty_dict"] == {}

    def test_primitive_values_json(self):
        """Test JSON formatting with primitive values."""
        # String
        result = llml("hello world", json)
        assert result == '"hello world"'

        # Number
        result = llml(42, json)
        assert result == "42"

        # Float
        result = llml(3.14159, json)
        assert stdjson.loads(result) == 3.14159

        # Boolean
        result = llml(True, json)
        assert result == "true"

        result = llml(False, json)
        assert result == "false"

        # None
        result = llml(None, json)
        assert result == "null"

    def test_special_characters_json(self):
        """Test JSON formatting with special characters."""
        data = {
            "quotes": 'He said "Hello"',
            "newline": "Line 1\nLine 2",
            "tab": "Column1\tColumn2",
            "unicode": "Hello 世界 🌍",
            "backslash": "path\\to\\file",
        }
        result = llml(data, json)
        parsed = stdjson.loads(result)
        assert parsed["quotes"] == 'He said "Hello"'
        assert parsed["newline"] == "Line 1\nLine 2"
        assert parsed["tab"] == "Column1\tColumn2"
        assert parsed["unicode"] == "Hello 世界 🌍"
        assert parsed["backslash"] == "path\\to\\file"

    def test_numeric_edge_cases_json(self):
        """Test JSON formatting with numeric edge cases."""
        data = {
            "zero": 0,
            "negative": -42,
            "float_zero": 0.0,
            "large_int": 9007199254740991,  # Max safe integer in JS
            "negative_float": -3.14159,
            "scientific": 1.23e-4,
        }
        result = llml(data, json)
        parsed = stdjson.loads(result)
        assert parsed["zero"] == 0
        assert parsed["negative"] == -42
        assert parsed["float_zero"] == 0.0
        assert parsed["large_int"] == 9007199254740991
        assert parsed["negative_float"] == -3.14159
        assert abs(parsed["scientific"] - 1.23e-4) < 1e-10

    def test_custom_objects_json(self):
        """Test JSON formatting with custom objects that have __dict__."""

        @dataclass
        class User:
            name: str
            email: str
            active: bool = True

        user = User("Alice", "alice@example.com")
        result = llml(user.__dict__, json)
        parsed = stdjson.loads(result)
        assert parsed["name"] == "Alice"
        assert parsed["email"] == "alice@example.com"
        assert parsed["active"] is True

    def test_complex_nested_structure_json(self):
        """Test JSON formatting with complex nested structures."""
        data = {
            "company": {
                "name": "TechCorp",
                "employees": [
                    {"id": 1, "name": "Alice", "roles": ["developer", "lead"]},
                    {"id": 2, "name": "Bob", "roles": ["developer"]},
                    {"id": 3, "name": "Charlie", "roles": ["manager", "architect"]},
                ],
                "departments": {
                    "engineering": {
                        "head": "Alice",
                        "budget": 1000000,
                        "projects": ["AI", "Backend", "Mobile"],
                    },
                    "sales": {
                        "head": "David",
                        "budget": 500000,
                        "regions": ["NA", "EU", "APAC"],
                    },
                },
                "founded": 2020,
                "public": False,
            }
        }
        result = llml(data, json)
        parsed = stdjson.loads(result)

        # Verify structure
        assert parsed["company"]["name"] == "TechCorp"
        assert len(parsed["company"]["employees"]) == 3
        assert parsed["company"]["employees"][0]["name"] == "Alice"
        assert "developer" in parsed["company"]["employees"][0]["roles"]
        assert parsed["company"]["departments"]["engineering"]["budget"] == 1000000
        assert parsed["company"]["founded"] == 2020
        assert parsed["company"]["public"] is False

    def test_json_formatter_preserves_data_integrity(self):
        """Test that JSON formatter preserves data integrity through round-trip."""
        original_data = {
            "string": "test",
            "int": 42,
            "float": 3.14,
            "bool": True,
            "null": None,
            "array": [1, 2, 3],
            "object": {"nested": "value"},
        }

        # Format to JSON
        json_str = llml(original_data, json)

        # Parse back
        parsed_data = stdjson.loads(json_str)

        # Verify all data is preserved
        assert parsed_data == original_data

    def test_json_formatter_with_unicode(self):
        """Test JSON formatter with various Unicode characters."""
        data = {
            "english": "Hello World",
            "chinese": "你好世界",
            "arabic": "مرحبا بالعالم",
            "emoji": "🌍🚀💻",
            "mixed": "Hello 世界 🌍",
        }
        result = llml(data, json)
        parsed = stdjson.loads(result)

        # Verify all Unicode is preserved
        assert parsed["english"] == "Hello World"
        assert parsed["chinese"] == "你好世界"
        assert parsed["arabic"] == "مرحبا بالعالم"
        assert parsed["emoji"] == "🌍🚀💻"
        assert parsed["mixed"] == "Hello 世界 🌍"

    def test_json_formatter_direct_call(self):
        """Test calling the JSON formatter function directly."""
        from zenbase_llml.formatters.json import format_json

        # Test with simple data
        result = format_json({"test": "value"}, llml, json)
        assert stdjson.loads(result) == {"test": "value"}

        # Test with None formatters argument
        result = format_json([1, 2, 3], llml, None)
        assert stdjson.loads(result) == [1, 2, 3]

    @pytest.mark.skipif(not HAS_ORJSON, reason="orjson not installed")
    def test_orjson_specific_behavior(self):
        """Test behavior specific to orjson when available."""
        # orjson produces more compact output
        result = llml({"a": 1, "b": 2}, json)
        assert result == '{"a":1,"b":2}'  # No spaces after colons

    @pytest.mark.skipif(HAS_ORJSON, reason="Testing standard json fallback")
    def test_standard_json_behavior(self):
        """Test behavior when using standard json library."""
        # Standard json includes spaces
        result = llml({"a": 1, "b": 2}, json)
        # Just verify it's valid JSON
        assert stdjson.loads(result) == {"a": 1, "b": 2}
