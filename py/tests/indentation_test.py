from zenbase_llml import llml


def test_basic_indentation():
    """Test basic indentation."""
    # Note: New API doesn't support indentation options
    result = llml({"message": "Hello"})
    expected = "<message>Hello</message>"
    assert result == expected
