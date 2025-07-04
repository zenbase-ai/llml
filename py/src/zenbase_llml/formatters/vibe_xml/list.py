from ..types import FormatFunction, Predicate


def is_list(value) -> bool:
    return isinstance(value, list)


def format_list(value, llml, formatters=None) -> str:
    if len(value) == 0:
        return ""

    results = []
    for i, item in enumerate(value):
        item_tag = str(i + 1)

        if i > 0:
            results.append("\n")

        if isinstance(item, dict) and item is not None:
            formatted = llml(item, formatters)
            if "\n" in formatted:
                results.append(f"<{item_tag}>\n{formatted}\n</{item_tag}>")
            else:
                results.append(f"<{item_tag}>{formatted}</{item_tag}>")
        else:
            formatted = llml(item, formatters)
            results.append(f"<{item_tag}>{formatted}</{item_tag}>")

    return "".join(results)


list_formatter: tuple[Predicate, FormatFunction] = (is_list, format_list)
