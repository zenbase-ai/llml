from ..types import FormatFunction, Predicate
from .helpers import format_key_value


def is_dict(value) -> bool:
    return isinstance(value, dict)


def format_dict(value, llml, formatters=None) -> str:
    if len(value) == 0:
        return ""

    entries = list(value.items())
    results = []

    for i, (key, val) in enumerate(entries):
        if i > 0:
            results.append("\n")

        # Format key-value pair with proper XML structure
        formatted_content = format_key_value(key, val, llml, formatters)
        results.append(formatted_content)

    return "".join(results)


dict_formatter: tuple[Predicate, FormatFunction] = (is_dict, format_dict)
