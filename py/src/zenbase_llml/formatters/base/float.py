from ..types import FormatFunction, Predicate


def is_float(value) -> bool:
    return isinstance(value, float)


def format_float(value, llml, formatters=None) -> str:
    return str(value)


float_formatter: tuple[Predicate, FormatFunction] = (is_float, format_float)
