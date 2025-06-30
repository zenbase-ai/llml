from zenbase_llml import LLMLOptions, llml


def test_with_prefix():
    """Test formatting with prefix."""
    result = llml(options=LLMLOptions(prefix="app"), config="value")
    expected = "<app-config>value</app-config>"
    assert result == expected


def test_prefix_with_arrays():
    """Test prefix with arrays (no -list suffix)."""
    result = llml(options=LLMLOptions(prefix="app"), items=["a", "b"])
    expected = (
        "<app-items>\n"
        "  <app-items-1>a</app-items-1>\n"
        "  <app-items-2>b</app-items-2>\n"
        "</app-items>"
    )
    assert result == expected


def test_prefix_with_nested_objects_non_strict():
    """Test prefix with nested objects in non-strict mode."""
    result = llml(options=LLMLOptions(prefix="app"), config={"debug": True})
    expected = "<app-config>\n  <debug>True</debug>\n</app-config>"
    assert result == expected


def test_prefix_with_nested_objects_strict():
    """Test prefix with nested objects in strict mode."""
    result = llml(
        options=LLMLOptions(prefix="app", strict=True), config={"debug": True}
    )
    expected = (
        "<app-config>\n  <app-config-debug>True</app-config-debug>\n</app-config>"
    )
    assert result == expected


def test_prefix_with_array_objects_non_strict():
    """Test prefix with array containing objects in non-strict mode."""
    result = llml(options=LLMLOptions(prefix="app"), data=[{"name": "Alice"}])
    expected = (
        "<app-data>\n"
        "  <app-data-1>\n"
        "    <name>Alice</name>\n"
        "  </app-data-1>\n"
        "</app-data>"
    )
    assert result == expected


def test_prefix_with_array_objects_strict():
    """Test prefix with array containing objects in strict mode."""
    result = llml(
        options=LLMLOptions(prefix="app", strict=True), data=[{"name": "Alice"}]
    )
    expected = (
        "<app-data>\n"
        "  <app-data-1>\n"
        "    <app-data-1-name>Alice</app-data-1-name>\n"
        "  </app-data-1>\n"
        "</app-data>"
    )
    assert result == expected
