import type { Formatters } from "../types"

export const isNumber = (v: unknown) => typeof v === "number"

export const formatNumber = (
  value: unknown,
  _llml: (data: unknown, formatters: Formatters) => string,
  _formatters: Formatters,
) => String(value)
