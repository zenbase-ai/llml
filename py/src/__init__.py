"""LLML - A lightweight markup language for structured text generation."""

from src.llml import Options, llml

__version__ = "0.1.0"
__all__ = ["llml", "Options"]

__call__ = llml
