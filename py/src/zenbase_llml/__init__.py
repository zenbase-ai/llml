"""LLML - A lightweight markup language for structured text generation."""

from zenbase_llml import formatters
from zenbase_llml.llml import llml

__version__ = "0.1.0"
__all__ = ["llml", "formatters"]

__call__ = llml
