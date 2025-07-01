from ..types import FormatFunction, Predicate


def is_bool(value) -> bool:
    return isinstance(value, bool)


def format_bool(value, llml, formatters=None) -> str:
    return str(value)


bool_formatter: tuple[Predicate, FormatFunction] = (is_bool, format_bool)
