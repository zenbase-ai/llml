from textwrap import dedent

from ..types import FormatFunction, Predicate


def is_str(value) -> bool:
    return isinstance(value, str) and "\n" not in value


def format_str(value, llml, formatters=None) -> str:
    return dedent(str(value).strip())


string_formatter: tuple[Predicate, FormatFunction] = (is_str, format_str)
