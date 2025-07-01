from ..types import FormatFunction, Predicate


def is_none(value) -> bool:
    return value is None


def format_none(value, llml, formatters=None) -> str:
    return str(value)


none_formatter: tuple[Predicate, FormatFunction] = (is_none, format_none)
