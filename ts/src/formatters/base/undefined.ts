import type { Formatters } from "../types"

export const isUndefined = (v: unknown): boolean => v === undefined

export const formatUndefined = (
  _value: unknown,
  _llml: (data: unknown, formatters: Formatters) => string,
  _formatters: Formatters,
): string => ""
