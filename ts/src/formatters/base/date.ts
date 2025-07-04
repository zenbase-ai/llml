import type { Formatters } from "../types";

export const isDate = (v: unknown): boolean => v instanceof Date;

export const formatDate = (
	value: unknown,
	_llml: (data: unknown, formatters: Formatters) => string,
	_formatters: Formatters,
): string => (value as Date).toISOString();
