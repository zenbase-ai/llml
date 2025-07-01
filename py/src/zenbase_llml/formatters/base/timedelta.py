import datetime

from ..types import FormatFunction, Predicate


def is_timedelta(value) -> bool:
    return isinstance(value, datetime.timedelta)


def format_timedelta(value, llml, formatters=None) -> str:
    return str(value)


timedelta_formatter: tuple[Predicate, FormatFunction] = (is_timedelta, format_timedelta)
