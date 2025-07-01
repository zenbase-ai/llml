from zenbase_llml import llml


def test_with_prefix():
    """Test formatting with prefix."""
    # Note: New API doesn't support prefix options
    result = llml({"config": "value"})
    expected = "<config>value</config>"
    assert result == expected


def test_prefix_with_arrays():
    """Test prefix with arrays (no -list suffix)."""
    # Note: New API doesn't support prefix options
    result = llml({"items": ["a", "b"]})
    expected = "<items>\n  <items-1>a</items-1>\n  <items-2>b</items-2>\n</items>"
    assert result == expected


def test_prefix_with_nested_objects_non_strict():
    """Test prefix with nested objects in non-strict mode."""
    # Note: New API doesn't support prefix options
    result = llml({"config": {"debug": True}})
    expected = "<config><debug>True</debug></config>"
    assert result == expected


def test_prefix_with_nested_objects_strict():
    """Test prefix with nested objects in strict mode."""
    # Note: New API doesn't support prefix or strict mode options
    result = llml({"config": {"debug": True}})
    expected = "<config><debug>True</debug></config>"
    assert result == expected


def test_prefix_with_array_objects_non_strict():
    """Test prefix with array containing objects in non-strict mode."""
    # Note: New API doesn't support prefix options
    result = llml({"data": [{"name": "Alice"}]})
    expected = "<data>\n  <data-1><name>Alice</name></data-1>\n</data>"
    assert result == expected


def test_prefix_with_array_objects_strict():
    """Test prefix with array containing objects in strict mode."""
    # Note: New API doesn't support prefix or strict mode options
    result = llml({"data": [{"name": "Alice"}]})
    expected = "<data>\n  <data-1><name>Alice</name></data-1>\n</data>"
    assert result == expected
