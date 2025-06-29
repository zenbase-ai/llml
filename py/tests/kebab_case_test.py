from src import llml


def test_kebab_case_conversion():
    """Test kebab-case conversion for keys."""
    result = llml(user_name="Alice", userAge=30)
    expected = "<user-name>Alice</user-name>\n<user-age>30</user-age>"
    assert result == expected


def test_string_with_spaces_in_key():
    """Test string with spaces converted to kebab-case."""
    result = llml(**{"key with spaces": "value"})
    expected = "<key-with-spaces>value</key-with-spaces>"
    assert result == expected