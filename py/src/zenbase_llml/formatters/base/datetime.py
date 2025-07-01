import datetime

from ..types import FormatFunction, Predicate


def is_datetime(value) -> bool:
    return isinstance(value, datetime.datetime)


def format_datetime(value, llml, formatters=None) -> str:
    return str(value)


datetime_formatter: tuple[Predicate, FormatFunction] = (is_datetime, format_datetime)
