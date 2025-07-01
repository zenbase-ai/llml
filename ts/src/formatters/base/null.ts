import type { Formatters } from "../types"

export const isNull = (v: unknown) => v === null

export const formatNull = (
  _value: unknown,
  _llml: (data: unknown, formatters: Formatters) => string,
  _formatters: Formatters,
) => "null"
