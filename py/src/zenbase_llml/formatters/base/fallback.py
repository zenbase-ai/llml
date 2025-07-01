from ..types import FormatFunction, Predicate


def is_fallback(value) -> bool:
    return True  # This always matches as a fallback


def format_fallback(value, llml, formatters=None) -> str:
    return str(value)


fallback_formatter: tuple[Predicate, FormatFunction] = (is_fallback, format_fallback)
