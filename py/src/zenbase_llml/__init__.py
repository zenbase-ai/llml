"""LLML - A lightweight markup language for structured text generation."""

from zenbase_llml.formatters import vibe_xml
from zenbase_llml.llml import llml

__version__ = "0.1.0"
__all__ = ["llml", "vibe_xml"]

__call__ = llml
