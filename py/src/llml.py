import typing as t
from io import StringIO
from textwrap import dedent

from beartype import beartype

from .utils import kebab_case

SPACE = " "


@beartype
def llml(
    *args: t.Any,
    indent: str = "",
    prefix: str = "",
    io: t.IO | None = None,
    **parts: t.Any,
) -> str:
    """
    Recursively combines parts into a single string using the following rules:

    Rules:
    0. llml({}) | llml([]) -> ""
    1. llml(instructions="...") ->
        <instructions>
        ...
        </instructions>
    2. llml(["a", "b", "c"]) ->
        <item id="1">a</item>
        <item id="2">b</item>
        <item id="3">c</item>
    3. llml(rules=["a", "b", "c"]) ->
        <rule-list>
            <rule-1>a</rule-1>
            <rule-2>b</rule-2>
            <rule-3>c</rule-3>
        </rule-list>
    4. llml(data={"key": "value"}) ->
        <data>
            <key>value</key>
        </data>
    5. llml(data=[{"key with spaces": "value"}, {"key": "value"}]) ->
        <data-list>
            <data id="1">
                <data-key-with-spaces>value</data-key-with-spaces>
            </data>
        </data-list>
    """
    io = io or StringIO()

    # Handle direct value calls: llml(), llml(None), llml([]), llml({})
    if len(args) == 0 and not parts:
        return ""
    elif len(args) == 1 and not parts:
        value = args[0]
        if isinstance(value, (list, dict)):
            if not value:  # Empty list or dict
                return ""
            # Non-empty collections need a key, so return empty
            return ""
        else:
            # Direct primitive value
            return str(value)

    # Handle single key-value pair formatting (recursive case)
    if len(args) == 2:
        key, value = args
            
        full_key = f"{prefix}-{key}" if prefix else key
        kebab_key = kebab_case(full_key)

        if isinstance(value, list):
            # Handle list formatting with wrapper tag
            wrapper_tag = f"{kebab_key}-list"

            if not value:
                io.write(f"{indent}<{wrapper_tag}></{wrapper_tag}>")
                return io.getvalue()

            io.write(f"{indent}<{wrapper_tag}>\n")
            inner_indent = indent + "  "

            for i, item in enumerate(value, 1):
                item_tag = f"{kebab_key}-{i}"

                if isinstance(item, dict):
                    # Dict items need special formatting
                    io.write(f"{inner_indent}<{item_tag}>\n")
                    dict_content = llml(indent=inner_indent + "  ", prefix=item_tag, **item)
                    io.write(dict_content)
                    io.write(f"\n{inner_indent}</{item_tag}>\n")
                else:
                    # Simple items on one line
                    io.write(f"{inner_indent}<{item_tag}>")
                    # Handle primitive values inline
                    if isinstance(item, str):
                        if "\n" in item:
                            clean_text = dedent(item).strip()
                            lines = clean_text.split("\n")
                            for j, line in enumerate(lines):
                                if j > 0:
                                    io.write("\n  ")
                                else:
                                    io.write("  ")
                                io.write(line)
                        else:
                            io.write(dedent(item.strip()))
                    else:
                        io.write(str(item))
                    io.write(f"</{item_tag}>\n")

            io.write(f"{indent}</{wrapper_tag}>")

        elif isinstance(value, dict):
            # Handle dict formatting - recursively call llml
            dict_content = llml(indent=indent + "  ", prefix=full_key, **value)

            if "\n" in dict_content:
                io.write(f"{indent}<{kebab_key}>\n{dict_content}\n{indent}</{kebab_key}>")
            else:
                io.write(f"{indent}<{kebab_key}>{dict_content}</{kebab_key}>")

        else:
            # Handle primitive values (strings, numbers, booleans, None)
            if isinstance(value, str):
                if "\n" in value:
                    clean_text = dedent(value).strip()
                    lines = clean_text.split("\n")
                    formatted_lines = []
                    for line in lines:
                        formatted_lines.append("  " + line)
                    formatted_content = "\n".join(formatted_lines)
                    io.write(f"{indent}<{kebab_key}>\n{formatted_content}\n{indent}</{kebab_key}>")
                else:
                    io.write(f"{indent}<{kebab_key}>{dedent(value.strip())}</{kebab_key}>")
            else:
                io.write(f"{indent}<{kebab_key}>{str(value)}</{kebab_key}>")

        return io.getvalue()

    # Handle multiple key-value pairs (base case)
    for i, (part_key, part_value) in enumerate(parts.items()):
        if i > 0:
            io.write("\n")

        formatted_content = llml(part_key, part_value, indent=indent, prefix=prefix)
        io.write(formatted_content)

    return io.getvalue()
