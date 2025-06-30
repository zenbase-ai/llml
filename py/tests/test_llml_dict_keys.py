import pytest
from zenbase_llml import llml, LLMLOptions

def test_llml_handles_non_string_dict_keys():
    payload = {
        "variables": {
            1: "Foo",
            2: "Bar",
        }
    }
    # We only care that llml() *does not* crash.
    llml("variables", payload["variables"], options=LLMLOptions(strict=False))
