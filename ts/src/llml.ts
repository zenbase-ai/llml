import { swagXML } from "./formatters";
import type { Formatters } from "./types";

/**
 * Core LLML function - simple formatter-based API
 */
export const llml = (data: unknown, formatters?: Formatters): string => {
	// Handle no arguments or undefined data
	if (data === undefined) {
		return "";
	}

	// Use default SwagXML formatters if none provided
	const activeFormatters = formatters || swagXML();

	// Iterate through formatters in insertion order
	for (const [predicate, formatFunction] of activeFormatters) {
		if (predicate(data)) {
			return formatFunction(data, llml, activeFormatters);
		}
	}

	// No formatter found - fallback to string conversion
	return String(data);
};
