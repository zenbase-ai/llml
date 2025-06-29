from zenbase_llml import llml


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
