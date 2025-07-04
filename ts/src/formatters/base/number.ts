import type { Formatters } from "../types";

export const isNumber = (v: unknown): boolean => typeof v === "number";

export const formatNumber = (
	value: unknown,
	_llml: (data: unknown, formatters: Formatters) => string,
	_formatters: Formatters,
): string => String(value);
