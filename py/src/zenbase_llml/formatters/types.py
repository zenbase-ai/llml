import typing as t
from typing import Any, Callable

Predicate = Callable[[Any], bool]

FormatFunction = Callable[[Any, Callable, t.Optional["Formatters"]], str]

Formatters = dict[Predicate, FormatFunction]

Formatter = tuple[Predicate, FormatFunction]
