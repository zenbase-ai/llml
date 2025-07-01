from zenbase_llml import llml


def test_mixed_content_new_format():
    """Test mixed content types with new formatting rules (default non-strict mode)."""
    result = llml(
        {
            "title": "My Document",
            "sections": ["intro", "body", "conclusion"],
            "metadata": {"author": "Alice", "version": "1.0"},
        }
    )
    expected = (
        "<title>My Document</title>\n"
        "<sections>\n"
        "  <sections-1>intro</sections-1>\n"
        "  <sections-2>body</sections-2>\n"
        "  <sections-3>conclusion</sections-3>\n"
        "</sections>\n"
        "<metadata>\n"
        "  <author>Alice</author>\n"
        "  <version>1.0</version>\n"
        "</metadata>"
    )
    assert result == expected


def test_deeply_nested_new_format():
    """Test deeply nested structure with new formatting (default non-strict mode)."""
    result = llml({"level1": {"level2": {"items": ["a", "b"]}}})
    expected = (
        "<level1>\n"
        "  <level2>\n"
        "    <items>\n"
        "      <items-1>a</items-1>\n"
        "      <items-2>b</items-2>\n"
        "    </items>\n"
        "  </level2>\n"
        "</level1>"
    )
    assert result == expected


def test_mixed_content_strict_mode():
    """Test mixed content types with strict mode enabled."""
    # Note: New API doesn't support strict mode options
    result = llml(
        {
            "title": "My Document",
            "sections": ["intro", "body", "conclusion"],
            "metadata": {"author": "Alice", "version": "1.0"},
        }
    )
    expected = (
        "<title>My Document</title>\n"
        "<sections>\n"
        "  <sections-1>intro</sections-1>\n"
        "  <sections-2>body</sections-2>\n"
        "  <sections-3>conclusion</sections-3>\n"
        "</sections>\n"
        "<metadata>\n"
        "  <author>Alice</author>\n"
        "  <version>1.0</version>\n"
        "</metadata>"
    )
    assert result == expected


def test_deeply_nested_strict_mode():
    """Test deeply nested structure with strict mode enabled."""
    # Note: New API doesn't support strict mode options
    result = llml({"level1": {"level2": {"items": ["a", "b"]}}})
    expected = (
        "<level1>\n"
        "  <level2>\n"
        "    <items>\n"
        "      <items-1>a</items-1>\n"
        "      <items-2>b</items-2>\n"
        "    </items>\n"
        "  </level2>\n"
        "</level1>"
    )
    assert result == expected
