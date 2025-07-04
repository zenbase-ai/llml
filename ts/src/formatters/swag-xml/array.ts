import type { Formatters } from "../types";

export const isArray = (v: unknown): boolean => Array.isArray(v);

export const formatArray = (
	unknownValue: unknown,
	llml: (data: unknown, formatters: Formatters) => string,
	formatters: Formatters,
): string => {
	const value = unknownValue as unknown[];

	if (value.length === 0) return "";

	const results: string[] = [];
	for (let i = 0; i < value.length; i++) {
		const item = value[i];
		const itemTag = String(i + 1);

		if (i > 0) results.push("\n");

		if (typeof item === "object" && item !== null && !Array.isArray(item)) {
			const formatted = llml(item, formatters);
			if (formatted.includes("\n")) {
				results.push(`<${itemTag}>\n${formatted}\n</${itemTag}>`);
			} else {
				results.push(`<${itemTag}>${formatted}</${itemTag}>`);
			}
		} else {
			const formatted = llml(item, formatters);
			results.push(`<${itemTag}>${formatted}</${itemTag}>`);
		}
	}

	return results.join("");
};
