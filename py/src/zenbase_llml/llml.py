import typing as t

from .formatters.vibe_xml import Formatters, vibe_xml

__all__ = ["llml", "vibe_xml", "Formatters"]


# Sentinel value to detect when no arguments are provided
_SENTINEL = object()


def llml(data: t.Any = _SENTINEL, formatters: Formatters = vibe_xml) -> str:
    """
    Core LLML function - ultra-simple formatter Map approach
    Takes data and a formatter Map, applies the first matching formatter
    """
    # Handle no arguments case
    if data is _SENTINEL:
        return ""

    # Iterate through formatters in insertion order
    for predicate, format_function in formatters.items():
        if predicate(data):
            return format_function(data, llml, formatters)

    # No formatter found - fallback to string conversion
    return str(data)
