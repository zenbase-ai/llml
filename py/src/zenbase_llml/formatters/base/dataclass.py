import typing as t
from dataclasses import asdict, is_dataclass

__all__ = ["is_dataclass", "format_dataclass"]


def format_dataclass(
    value: t.Any,
    llml: t.Callable,
    formatters: t.Dict[t.Callable, t.Callable],
) -> str:
    return llml(asdict(value), formatters)
