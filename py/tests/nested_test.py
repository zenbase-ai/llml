from src import llml


def test_nested_dict():
    """Test nested dictionary formatting."""
    result = llml(config={"debug": True, "timeout": 30})
    expected = (
        "<config>\n"
        "  <config-debug>True</config-debug>\n"
        "  <config-timeout>30</config-timeout>\n"
        "</config>"
    )
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