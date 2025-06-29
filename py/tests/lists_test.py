from src import Options, llml


def test_empty_named_list():
    """Test empty named list formatting."""
    result = llml(items=[])
    expected = ""
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
    result = llml(options=Options(indent="  "), items=["a", "b"])
    expected = (
        "  <items-list>\n"
        "    <items-1>a</items-1>\n"
        "    <items-2>b</items-2>\n"
        "  </items-list>"
    )
    assert result == expected


def test_list_with_prefix():
    """Test list formatting with prefix."""
    result = llml(options=Options(prefix="app"), items=["a", "b"])
    expected = (
        "<app-items-list>\n"
        "  <app-items-1>a</app-items-1>\n"
        "  <app-items-2>b</app-items-2>\n"
        "</app-items-list>"
    )
    assert result == expected