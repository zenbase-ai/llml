from zenbase_llml import llml


def test_kebab_case_conversion():
    """Test kebab-case conversion for keys."""
    # Note: New API doesn't support kebab-case conversion
    result = llml({"user_name": "Alice", "userAge": 30})
    expected = "<user_name>Alice</user_name>\n<userAge>30</userAge>"
    assert result == expected


def test_string_with_spaces_in_key():
    """Test string with spaces converted to kebab-case."""
    # Note: New API doesn't support kebab-case conversion
    result = llml({"key with spaces": "value"})
    expected = "<key with spaces>value</key with spaces>"
    assert result == expected
