import re

whitespace = re.compile(r"[\s_]+")
capital_letter = re.compile(r"(?<!^)(?=[A-Z])")


def kebab_case(text: str) -> str:
    """Convert text to kebab-case format."""
    # Replace spaces and underscores with hyphens
    text = whitespace.sub("-", text)
    # Insert hyphens before capital letters (except at the start)
    text = capital_letter.sub("-", text)
    # Convert to lowercase
    return text.lower()
