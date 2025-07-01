from zenbase_llml import llml


def test_nested_dict():
    """Test nested dictionary formatting (default non-strict mode)."""
    result = llml({"config": {"debug": True, "timeout": 30}})
    expected = "<config>\n  <debug>True</debug>\n  <timeout>30</timeout>\n</config>"
    assert result == expected


def test_nested_dict_with_kebab_case():
    """Test nested dictionary with kebab-case conversion (default non-strict mode)."""
    result = llml({"user_config": {"debug_mode": True, "maxRetries": 5}})
    expected = (
        "<user_config>\n"
        "  <debug_mode>True</debug_mode>\n"
        "  <maxRetries>5</maxRetries>\n"
        "</user_config>"
    )
    assert result == expected


def test_list_of_dicts():
    """Test list containing dictionaries (default non-strict mode)."""
    result = llml({"data": [{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}]})
    expected = (
        "<data>\n"
        "  <data-1>\n"
        "    <name>Alice</name>\n"
        "    <age>30</age>\n"
        "  </data-1>\n"
        "  <data-2>\n"
        "    <name>Bob</name>\n"
        "    <age>25</age>\n"
        "  </data-2>\n"
        "</data>"
    )
    assert result == expected


def test_nested_dict_strict_mode():
    """Test nested dictionary formatting with strict mode enabled."""
    # Note: New API doesn't support strict mode options
    result = llml({"config": {"debug": True, "timeout": 30}})
    expected = "<config>\n  <debug>True</debug>\n  <timeout>30</timeout>\n</config>"
    assert result == expected


def test_nested_dict_with_kebab_case_strict_mode():
    """Test nested dictionary with kebab-case conversion in strict mode."""
    # Note: New API doesn't support strict mode options or kebab-case conversion
    result = llml({"user_config": {"debug_mode": True, "maxRetries": 5}})
    expected = (
        "<user_config>\n"
        "  <debug_mode>True</debug_mode>\n"
        "  <maxRetries>5</maxRetries>\n"
        "</user_config>"
    )
    assert result == expected


def test_list_of_dicts_strict_mode():
    """Test list containing dictionaries with strict mode enabled."""
    # Note: New API doesn't support strict mode options
    result = llml({"data": [{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}]})
    expected = (
        "<data>\n"
        "  <data-1>\n"
        "    <name>Alice</name>\n"
        "    <age>30</age>\n"
        "  </data-1>\n"
        "  <data-2>\n"
        "    <name>Bob</name>\n"
        "    <age>25</age>\n"
        "  </data-2>\n"
        "</data>"
    )
    assert result == expected
