from textwrap import dedent


def format_key_value(key: str, value, llml, formatters):
    """Helper function to format key-value pairs in SwagXML style"""
    if isinstance(value, list):
        if len(value) == 0:
            return ""

        wrapper_tag = key
        result = f"<{wrapper_tag}>\n"

        for i, item in enumerate(value):
            item_tag = f"{key}-{i + 1}"

            if isinstance(item, dict) and item is not None:
                item_formatted = llml(item, formatters)
                if "\n" in item_formatted:
                    result += f"  <{item_tag}>\n"
                    indented_content = "\n".join(
                        f"    {line}" if line else ""
                        for line in item_formatted.split("\n")
                    )
                    result += indented_content
                    result += f"\n  </{item_tag}>\n"
                else:
                    result += f"  <{item_tag}>{item_formatted}</{item_tag}>\n"
            else:
                if isinstance(item, str) and "\n" in item:
                    clean_text = dedent(item).strip()
                    lines = clean_text.split("\n")
                    result += f"  <{item_tag}>"
                    for j, line in enumerate(lines):
                        if j > 0:
                            result += "\n  "
                        else:
                            result += "  "
                        result += line
                    result += f"</{item_tag}>\n"
                else:
                    formatted = llml(item, formatters)
                    result += f"  <{item_tag}>{formatted}</{item_tag}>\n"

        result += f"</{wrapper_tag}>"
        return result
    elif isinstance(value, dict) and value is not None:
        formatted = llml(value, formatters)
        if "\n" in formatted:
            indented_content = "\n".join(
                f"  {line}" if line else "" for line in formatted.split("\n")
            )
            return f"<{key}>\n{indented_content}\n</{key}>"
        else:
            return f"<{key}>{formatted}</{key}>"
    else:
        if isinstance(value, str) and "\n" in value:
            clean_text = dedent(value).strip()
            lines = clean_text.split("\n")
            formatted_lines = [f"  {line}" for line in lines]
            formatted_content = "\n".join(formatted_lines)
            return f"<{key}>\n{formatted_content}\n</{key}>"
        else:
            formatted = llml(value, formatters)
            return f"<{key}>{formatted}</{key}>"
