import typing as t
from dataclasses import dataclass
from io import StringIO
from textwrap import dedent

from zenbase_llml.utils import kebab_case

SPACE = " "


@dataclass
class LLMLOptions:
    indent: str = ""
    prefix: str = ""
    strict: bool = False


def llml(
    *args: t.Any,
    options: LLMLOptions | None = None,
    io: t.IO | None = None,
    **parts: t.Any,
) -> str:
    io = io or StringIO()

    # Extract options
    indent = options.indent if options else ""
    prefix = options.prefix if options else ""
    strict = options.strict if options else False

    # Handle direct value calls: llml(), llml(None), llml([]), llml({})
    if len(args) == 0 and not parts:
        return ""
    elif len(args) == 1 and not parts:
        value = args[0]
        if isinstance(value, dict):
            if not value:  # Empty dict
                return ""
            # Non-empty dicts need keys, so return empty
            return ""
        elif isinstance(value, list):
            if not value:  # Empty list
                return ""
            # Handle direct list as numbered items
            for i, item in enumerate(value, 1):
                item_tag = f"{prefix}-{i}" if prefix else str(i)

                if isinstance(item, dict):
                    # Dict items need special formatting
                    io.write(f"{indent}<{item_tag}>\n")
                    dict_content = llml(
                        options=LLMLOptions(
                            indent=indent + "  ",
                            prefix=item_tag if strict else "",
                            strict=strict,
                        ),
                        **{str(k): v for k, v in item.items()},
                    )
                    io.write(dict_content)
                    io.write(f"\n{indent}</{item_tag}>")
                else:
                    # Simple items on one line
                    io.write(f"{indent}<{item_tag}>{str(item)}</{item_tag}>")

                if i < len(value):
                    io.write("\n")

            return io.getvalue()
        else:
            # Direct primitive value
            return str(value)

    # Handle single key-value pair formatting (recursive case)
    if len(args) == 2:
        key, value = args

        full_key = f"{prefix}-{key}" if prefix else key
        kebab_key = kebab_case(full_key)

        if isinstance(value, list):
            # Handle list formatting with wrapper tag (no -list suffix)
            wrapper_tag = kebab_key

            if not value:
                return ""

            io.write(f"{indent}<{wrapper_tag}>\n")
            inner_indent = indent + "  "

            for i, item in enumerate(value, 1):
                item_tag = f"{kebab_key}-{i}"

                if isinstance(item, dict):
                    # Dict items need special formatting
                    io.write(f"{inner_indent}<{item_tag}>\n")
                    dict_content = llml(
                        options=LLMLOptions(
                            indent=inner_indent + "  ",
                            prefix=item_tag if strict else "",
                            strict=strict,
                        ),
                        **{str(k): v for k, v in item.items()},
                    )
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
            dict_content = llml(
                options=LLMLOptions(
                    indent=indent + "  ",
                    prefix=full_key if strict else "",
                    strict=strict,
                ),
                **{str(k): v for k, v in value.items()},
            )

            # Always format dicts with newlines for proper indentation
            io.write(f"{indent}<{kebab_key}>\n{dict_content}\n{indent}</{kebab_key}>")

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
                    io.write(
                        f"{indent}<{kebab_key}>\n{formatted_content}\n{indent}</{kebab_key}>"
                    )
                else:
                    io.write(
                        f"{indent}<{kebab_key}>{dedent(value.strip())}</{kebab_key}>"
                    )
            else:
                io.write(f"{indent}<{kebab_key}>{str(value)}</{kebab_key}>")

        return io.getvalue()

    # Handle multiple key-value pairs (base case)
    for i, (part_key, part_value) in enumerate(parts.items()):
        if i > 0:
            io.write("\n")

        formatted_content = llml(
            part_key,
            part_value,
            options=LLMLOptions(indent=indent, prefix=prefix, strict=strict),
        )
        io.write(formatted_content)

    return io.getvalue()
