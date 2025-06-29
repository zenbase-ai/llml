from src import llml


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