from src import Options, llml


def test_basic_indentation():
    """Test basic indentation."""
    result = llml(options=Options(indent="  "), message="Hello")
    expected = "  <message>Hello</message>"
    assert result == expected