from ..types import FormatFunction, Predicate


def is_int(value) -> bool:
    return isinstance(value, int) and not isinstance(value, bool)


def format_int(value, llml, formatters=None) -> str:
    return str(value)


int_formatter: tuple[Predicate, FormatFunction] = (is_int, format_int)
