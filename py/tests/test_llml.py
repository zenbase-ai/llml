from src import llml


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
    result = llml(empty="")
    expected = "<empty></empty>"
    assert result == expected


def test_zero_value():
    """Test zero numeric value."""
    result = llml(zero=0)
    expected = "<zero>0</zero>"
    assert result == expected


def test_false_boolean():
    """Test False boolean value."""
    result = llml(disabled=False)
    expected = "<disabled>False</disabled>"
    assert result == expected


def test_none_value():
    """Test None value handling."""
    result = llml(nothing=None)
    expected = "<nothing>None</nothing>"
    assert result == expected


def test_simple_string_value():
    """Test simple string formatting."""
    result = llml(instructions="Follow these steps")
    expected = "<instructions>Follow these steps</instructions>"
    assert result == expected


def test_integer_value():
    """Test integer formatting."""
    result = llml(count=42)
    expected = "<count>42</count>"
    assert result == expected


def test_float_value():
    """Test float formatting."""
    result = llml(temperature=98.6)
    expected = "<temperature>98.6</temperature>"
    assert result == expected


def test_boolean_value():
    """Test boolean formatting."""
    result = llml(enabled=True)
    expected = "<enabled>True</enabled>"
    assert result == expected


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


def test_multiple_simple_values():
    """Test multiple simple values."""
    result = llml(name="Alice", age=30, active=True)
    expected = "<name>Alice</name>\n<age>30</age>\n<active>True</active>"
    assert result == expected


def test_basic_indentation():
    """Test basic indentation."""
    result = llml(indent="  ", message="Hello")
    expected = "  <message>Hello</message>"
    assert result == expected


def test_with_prefix():
    """Test formatting with prefix."""
    result = llml(prefix="app", config="value")
    expected = "<app-config>value</app-config>"
    assert result == expected


def test_multiline_content():
    """Test multiline content formatting with dedent."""
    content = """
    Line 1
    Line 2
    Line 3
    """
    result = llml(description=content)
    expected = "<description>\n  Line 1\n  Line 2\n  Line 3\n</description>"
    assert result == expected


def test_empty_list():
    """Test empty list formatting."""
    result = llml(items=[])
    expected = "<items-list></items-list>"
    assert result == expected


def test_simple_list_with_wrapper():
    """Test list formatting with wrapper tag and numbered items."""
    result = llml(rules=["first", "second", "third"])
    expected = (
        "<rules-list>\n"
        "  <rules-1>first</rules-1>\n"
        "  <rules-2>second</rules-2>\n"
        "  <rules-3>third</rules-3>\n"
        "</rules-list>"
    )
    assert result == expected


def test_list_with_numbers():
    """Test list with numeric values."""
    result = llml(numbers=[1, 2, 3])
    expected = (
        "<numbers-list>\n"
        "  <numbers-1>1</numbers-1>\n"
        "  <numbers-2>2</numbers-2>\n"
        "  <numbers-3>3</numbers-3>\n"
        "</numbers-list>"
    )
    assert result == expected


def test_list_kebab_case_conversion():
    """Test that list names are converted to kebab-case."""
    result = llml(user_tasks=["task1", "task2"])
    expected = (
        "<user-tasks-list>\n"
        "  <user-tasks-1>task1</user-tasks-1>\n"
        "  <user-tasks-2>task2</user-tasks-2>\n"
        "</user-tasks-list>"
    )
    assert result == expected


def test_list_with_indentation():
    """Test list formatting with indentation."""
    result = llml(indent="  ", items=["a", "b"])
    expected = (
        "  <items-list>\n"
        "    <items-1>a</items-1>\n"
        "    <items-2>b</items-2>\n"
        "  </items-list>"
    )
    assert result == expected


def test_list_with_prefix():
    """Test list formatting with prefix."""
    result = llml(prefix="app", items=["a", "b"])
    expected = (
        "<app-items-list>\n"
        "  <app-items-1>a</app-items-1>\n"
        "  <app-items-2>b</app-items-2>\n"
        "</app-items-list>"
    )
    assert result == expected


def test_nested_dict():
    """Test nested dictionary formatting."""
    result = llml(config={"debug": True, "timeout": 30})
    expected = "<config>\n  <config-debug>True</config-debug>\n  <config-timeout>30</config-timeout>\n</config>"
    assert result == expected


def test_nested_dict_with_kebab_case():
    """Test nested dictionary with kebab-case conversion."""
    result = llml(user_config={"debug_mode": True, "maxRetries": 5})
    expected = (
        "<user-config>\n"
        "  <user-config-debug-mode>True</user-config-debug-mode>\n"
        "  <user-config-max-retries>5</user-config-max-retries>\n"
        "</user-config>"
    )
    assert result == expected


def test_list_of_dicts():
    """Test list containing dictionaries."""
    result = llml(data=[{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}])
    expected = (
        "<data-list>\n"
        "  <data-1>\n"
        "    <data-1-name>Alice</data-1-name>\n"
        "    <data-1-age>30</data-1-age>\n"
        "  </data-1>\n"
        "  <data-2>\n"
        "    <data-2-name>Bob</data-2-name>\n"
        "    <data-2-age>25</data-2-age>\n"
        "  </data-2>\n"
        "</data-list>"
    )
    assert result == expected


def test_mixed_content_new_format():
    """Test mixed content types with new formatting rules."""
    result = llml(
        title="My Document",
        sections=["intro", "body", "conclusion"],
        metadata={"author": "Alice", "version": "1.0"},
    )
    expected = (
        "<title>My Document</title>\n"
        "<sections-list>\n"
        "  <sections-1>intro</sections-1>\n"
        "  <sections-2>body</sections-2>\n"
        "  <sections-3>conclusion</sections-3>\n"
        "</sections-list>\n"
        "<metadata>\n"
        "  <metadata-author>Alice</metadata-author>\n"
        "  <metadata-version>1.0</metadata-version>\n"
        "</metadata>"
    )
    assert result == expected


def test_deeply_nested_new_format():
    """Test deeply nested structure with new formatting."""
    result = llml(level1={"level2": {"items": ["a", "b"]}})
    expected = (
        "<level1>\n"
        "  <level1-level2>\n"
        "    <level1-level2-items-list>\n"
        "      <level1-level2-items-1>a</level1-level2-items-1>\n"
        "      <level1-level2-items-2>b</level1-level2-items-2>\n"
        "    </level1-level2-items-list>\n"
        "  </level1-level2>\n"
        "</level1>"
    )
    assert result == expected
    assert result == expected
