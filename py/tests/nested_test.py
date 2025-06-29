from zenbase_llml import LLMLOptions, llml


def test_nested_dict():
    """Test nested dictionary formatting (default non-strict mode)."""
    result = llml(config={"debug": True, "timeout": 30})
    expected = (
        "<config>\n"
        "  <debug>True</debug>\n"
        "  <timeout>30</timeout>\n"
        "</config>"
    )
    assert result == expected


def test_nested_dict_with_kebab_case():
    """Test nested dictionary with kebab-case conversion (default non-strict mode)."""
    result = llml(user_config={"debug_mode": True, "maxRetries": 5})
    expected = (
        "<user-config>\n"
        "  <debug-mode>True</debug-mode>\n"
        "  <max-retries>5</max-retries>\n"
        "</user-config>"
    )
    assert result == expected


def test_list_of_dicts():
    """Test list containing dictionaries (default non-strict mode)."""
    result = llml(data=[{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}])
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
    result = llml(options=LLMLOptions(strict=True), config={"debug": True, "timeout": 30})
    expected = (
        "<config>\n"
        "  <config-debug>True</config-debug>\n"
        "  <config-timeout>30</config-timeout>\n"
        "</config>"
    )
    assert result == expected


def test_nested_dict_with_kebab_case_strict_mode():
    """Test nested dictionary with kebab-case conversion in strict mode."""
    result = llml(options=LLMLOptions(strict=True), user_config={"debug_mode": True, "maxRetries": 5})
    expected = (
        "<user-config>\n"
        "  <user-config-debug-mode>True</user-config-debug-mode>\n"
        "  <user-config-max-retries>5</user-config-max-retries>\n"
        "</user-config>"
    )
    assert result == expected


def test_list_of_dicts_strict_mode():
    """Test list containing dictionaries with strict mode enabled."""
    result = llml(options=LLMLOptions(strict=True), data=[{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}])
    expected = (
        "<data>\n"
        "  <data-1>\n"
        "    <data-1-name>Alice</data-1-name>\n"
        "    <data-1-age>30</data-1-age>\n"
        "  </data-1>\n"
        "  <data-2>\n"
        "    <data-2-name>Bob</data-2-name>\n"
        "    <data-2-age>25</data-2-age>\n"
        "  </data-2>\n"
        "</data>"
    )
    assert result == expected
