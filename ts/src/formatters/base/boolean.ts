import type { Formatters } from "../types";

export const isBoolean = (v: unknown): boolean => typeof v === "boolean";

export const formatBoolean = (
	value: unknown,
	_llml: (data: unknown, formatters: Formatters) => string,
	_formatters: Formatters,
): string => String(value);
