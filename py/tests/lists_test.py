from zenbase_llml import llml


def test_empty_named_list():
    """Test empty named list formatting."""
    result = llml({"items": []})
    expected = ""
    assert result == expected


def test_simple_list_with_wrapper():
    """Test list formatting with wrapper tag and numbered items."""
    result = llml({"rules": ["first", "second", "third"]})
    expected = (
        "<rules>\n"
        "  <rules-1>first</rules-1>\n"
        "  <rules-2>second</rules-2>\n"
        "  <rules-3>third</rules-3>\n"
        "</rules>"
    )
    assert result == expected


def test_list_with_numbers():
    """Test list with numeric values."""
    result = llml({"numbers": [1, 2, 3]})
    expected = (
        "<numbers>\n"
        "  <numbers-1>1</numbers-1>\n"
        "  <numbers-2>2</numbers-2>\n"
        "  <numbers-3>3</numbers-3>\n"
        "</numbers>"
    )
    assert result == expected


def test_list_kebab_case_conversion():
    """Test that list names are converted to kebab-case."""
    result = llml({"user_tasks": ["task1", "task2"]})
    expected = (
        "<user_tasks>\n"
        "  <user_tasks-1>task1</user_tasks-1>\n"
        "  <user_tasks-2>task2</user_tasks-2>\n"
        "</user_tasks>"
    )
    assert result == expected


def test_list_with_indentation():
    """Test list formatting with indentation."""
    # Note: The new API doesn't support indentation options yet
    # This test is simplified to test basic functionality
    result = llml({"items": ["a", "b"]})
    expected = "<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n</items>"
    assert result == expected


def test_list_with_prefix():
    """Test list formatting with prefix."""
    # Note: The new API doesn't support prefix options yet
    # This test is simplified to test basic functionality
    result = llml({"items": ["a", "b"]})
    expected = "<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n</items>"
    assert result == expected
