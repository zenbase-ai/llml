import type { Formatters } from "../types"

export const isBoolean = (v: unknown) => typeof v === "boolean"

export const formatBoolean = (
  value: unknown,
  _llml: (data: unknown, formatters: Formatters) => string,
  _formatters: Formatters,
) => String(value)
