from zenbase_llml import llml


def test_no_args():
    """Test no arguments."""
    result = llml()
    expected = ""
    assert result == expected


def test_none():
    """Test None value."""
    result = llml(None)
    expected = "None"
    assert result == expected


def test_empty_list():
    """Test empty list."""
    result = llml([])
    expected = ""
    assert result == expected


def test_empty_dict():
    """Test empty dict."""
    result = llml({})
    expected = ""
    assert result == expected


def test_empty_string():
    """Test empty string value."""
    result = llml({"empty": ""})
    expected = "<empty></empty>"
    assert result == expected


def test_zero_value():
    """Test zero numeric value."""
    result = llml({"zero": 0})
    expected = "<zero>0</zero>"
    assert result == expected


def test_false_boolean():
    """Test False boolean value."""
    result = llml({"disabled": False})
    expected = "<disabled>False</disabled>"
    assert result == expected


def test_none_value():
    """Test None value handling."""
    result = llml({"nothing": None})
    expected = "<nothing>None</nothing>"
    assert result == expected


def test_simple_string_value():
    """Test simple string formatting."""
    result = llml({"instructions": "Follow these steps"})
    expected = "<instructions>Follow these steps</instructions>"
    assert result == expected


def test_integer_value():
    """Test integer formatting."""
    result = llml({"count": 42})
    expected = "<count>42</count>"
    assert result == expected


def test_float_value():
    """Test float formatting."""
    result = llml({"temperature": 98.6})
    expected = "<temperature>98.6</temperature>"
    assert result == expected


def test_boolean_value():
    """Test boolean formatting."""
    result = llml({"enabled": True})
    expected = "<enabled>True</enabled>"
    assert result == expected


def test_multiple_simple_values():
    """Test multiple simple values."""
    result = llml({"name": "Alice", "age": 30, "active": True})
    expected = "<name>Alice</name>\n<age>30</age>\n<active>True</active>"
    assert result == expected
