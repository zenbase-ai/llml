from zenbase_llml import LLMLOptions, llml


def test_direct_array_with_strings():
    """Test direct array with string values."""
    result = llml(["a", "b", "c"])
    expected = "<1>a</1>\n<2>b</2>\n<3>c</3>"
    assert result == expected


def test_direct_array_with_mixed_types():
    """Test direct array with mixed types."""
    result = llml([1, "hello", True])
    expected = "<1>1</1>\n<2>hello</2>\n<3>True</3>"
    assert result == expected


def test_direct_array_with_objects():
    """Test direct array containing dictionaries (default non-strict mode)."""
    result = llml([{"name": "Alice"}, {"name": "Bob"}])
    expected = "<1>\n  <name>Alice</name>\n</1>\n<2>\n  <name>Bob</name>\n</2>"
    assert result == expected


def test_direct_empty_array():
    """Test direct empty array."""
    result = llml([])
    expected = ""
    assert result == expected


def test_direct_array_with_indentation():
    """Test direct array with indentation options."""
    result = llml(["a", "b"], options=LLMLOptions(indent="  "))
    expected = "  <1>a</1>\n  <2>b</2>"
    assert result == expected


def test_direct_array_with_prefix():
    """Test direct array with prefix options."""
    result = llml(["a", "b"], options=LLMLOptions(prefix="item"))
    expected = "<item-1>a</item-1>\n<item-2>b</item-2>"
    assert result == expected


def test_direct_array_with_objects_strict_mode():
    """Test direct array containing dictionaries with strict mode enabled."""
    result = llml([{"name": "Alice"}, {"name": "Bob"}], options=LLMLOptions(strict=True))
    expected = "<1>\n  <1-name>Alice</1-name>\n</1>\n<2>\n  <2-name>Bob</2-name>\n</2>"
    assert result == expected
