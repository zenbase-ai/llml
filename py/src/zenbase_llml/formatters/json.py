from zenbase_llml.formatters.types import Formatters

try:
    import orjson

    def format_json(value, llml, formatters=None) -> str:
        return orjson.dumps(value).decode()
except ImportError:
    import json as stdjson

    def format_json(value, llml, formatters=None) -> str:
        return stdjson.dumps(value)


# JSON formatter that matches all values
def is_json(value) -> bool:
    """Predicate that matches all values for JSON formatting."""
    return True


json: Formatters = {
    is_json: format_json,
}
