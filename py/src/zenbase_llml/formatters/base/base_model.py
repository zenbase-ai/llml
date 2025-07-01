import typing as t

try:
    from pydantic import BaseModel
except ImportError:
    BaseModel = type()


def is_base_model(value: t.Any) -> bool:
    return isinstance(value, BaseModel)


def format_base_model(
    value: BaseModel, llml: t.Callable, formatters: t.Dict[t.Callable, t.Callable]
) -> str:
    return llml(value.model_dump(), formatters)
