import dedent from "dedent";

import type { Formatters } from "../types";

export const isObject = (v: unknown): boolean =>
	typeof v === "object" &&
	v !== null &&
	!Array.isArray(v) &&
	!(v instanceof Date);

export const formatObject = (
	value: unknown,
	llml: (data: unknown, formatters: Formatters) => string,
	formatters: Formatters,
): string => {
	// Check if object has a custom toString method (not the default Object.prototype.toString)
	if (
		value &&
		typeof value.toString === "function" &&
		value.toString !== Object.prototype.toString
	) {
		return value.toString();
	}

	const entries = Object.entries(value as Record<string, unknown>);
	if (entries.length === 0) {
		// Plain empty objects should return empty string
		// Other objects with no enumerable properties should use String(value)
		if (Object.getPrototypeOf(value) === Object.prototype) {
			return "";
		}
		return String(value);
	}
	const results: string[] = [];

	for (let i = 0; i < entries.length; i++) {
		const [key, val] = entries[i];

		if (i > 0) results.push("\n");

		// Format key-value pair with proper XML structure
		const formattedContent = formatKeyValue(key, val, llml, formatters);
		results.push(formattedContent);
	}

	return results.join("");
};

/**
 * Helper function to format key-value pairs in SwagXML style
 */
function formatKeyValue(
	key: string,
	value: unknown,
	llml: (data: unknown, formatters: Formatters) => string,
	formatters: Formatters,
): string {
	if (Array.isArray(value)) {
		if (value.length === 0) return "";

		const wrapperTag = key;
		let result = `<${wrapperTag}>\n`;

		for (let i = 0; i < value.length; i++) {
			const item = value[i];
			const itemTag = `${key}-${i + 1}`;

			if (typeof item === "object" && item !== null && !Array.isArray(item)) {
				const itemFormatted = llml(item, formatters);
				if (itemFormatted.includes("\n")) {
					result += `  <${itemTag}>\n`;
					const indentedContent = itemFormatted
						.split("\n")
						.map((line) => (line ? `    ${line}` : ""))
						.join("\n");
					result += indentedContent;
					result += `\n  </${itemTag}>\n`;
				} else {
					result += `  <${itemTag}>${itemFormatted}</${itemTag}>\n`;
				}
			} else {
				if (typeof item === "string" && item.includes("\n")) {
					const cleanText = dedent(item).trim();
					const lines = cleanText.split("\n");
					result += `  <${itemTag}>`;
					for (let j = 0; j < lines.length; j++) {
						if (j > 0) {
							result += "\n  ";
						} else {
							result += "  ";
						}
						result += lines[j];
					}
					result += `</${itemTag}>\n`;
				} else {
					const formatted = llml(item, formatters);
					result += `  <${itemTag}>${formatted}</${itemTag}>\n`;
				}
			}
		}

		result += `</${wrapperTag}>`;
		return result;
	}

	if (typeof value === "object" && value !== null) {
		const formatted = llml(value, formatters);
		if (formatted.includes("\n")) {
			const indentedContent = formatted
				.split("\n")
				.map((line) => (line ? `  ${line}` : ""))
				.join("\n");
			return `<${key}>\n${indentedContent}\n</${key}>`;
		}

		return `<${key}>${formatted}</${key}>`;
	}

	if (typeof value === "string" && value.includes("\n")) {
		const cleanText = dedent(value).trim();
		const lines = cleanText.split("\n");
		const formattedLines = lines.map((line) => `  ${line}`);
		const formattedContent = formattedLines.join("\n");
		return `<${key}>\n${formattedContent}\n</${key}>`;
	}

	const formatted = llml(value, formatters);
	return `<${key}>${formatted}</${key}>`;
}
