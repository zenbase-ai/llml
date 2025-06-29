from zenbase_llml import LLMLOptions, llml


def test_basic_indentation():
    """Test basic indentation."""
    result = llml(options=LLMLOptions(indent="  "), message="Hello")
    expected = "  <message>Hello</message>"
    assert result == expected
