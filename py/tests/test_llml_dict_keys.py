from zenbase_llml import llml


def test_llml_handles_non_string_dict_keys():
    payload = {
        "variables": {
            1: "Foo",
            2: "Bar",
        }
    }
    # We only care that llml() *does not* crash.
    result = llml(payload)

    # Verify the result is properly formatted
    assert "<variables>" in result
    assert "</variables>" in result
    assert "Foo" in result
    assert "Bar" in result


def test_llml_handles_mixed_key_types():
    """Test that mixed string and non-string keys work correctly."""
    data = {
        "mixed": {
            42: "integer_key",
            "string_key": "string_value",
            2.5: "float_key",
            False: "bool_key",  # Using False to avoid collision with integer keys
        }
    }

    # Should not crash and should produce valid output
    result = llml(data)

    assert "<mixed>" in result
    assert "</mixed>" in result
    assert "integer_key" in result
    assert "string_value" in result
    assert "float_key" in result
    assert "bool_key" in result


def test_llml_handles_nested_non_string_keys():
    """Test nested dictionaries with non-string keys."""
    data = {"nested": {1: {2: "deeply_nested", "inner": "value"}}}

    # Should not crash
    result = llml(data)

    assert "<nested>" in result
    assert "</nested>" in result
    assert "deeply_nested" in result
    assert "value" in result
