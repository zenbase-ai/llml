from src import Options, llml


def test_with_prefix():
    """Test formatting with prefix."""
    result = llml(options=Options(prefix="app"), config="value")
    expected = "<app-config>value</app-config>"
    assert result == expected