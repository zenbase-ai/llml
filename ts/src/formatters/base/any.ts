import type { Formatters } from "../types"

export const isAny = (_v: unknown): boolean => true

export const formatAny = (
  value: unknown,
  _llml: (data: unknown, formatters: Formatters) => string,
  _formatters: Formatters,
): string => String(value)
