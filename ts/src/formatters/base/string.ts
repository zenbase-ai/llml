import dedent from "dedent"

import type { Formatters } from "../types"

export const isString = (v: unknown) => typeof v === "string" && !v.includes("\n")

export const formatString = (
  value: unknown,
  _llml: (data: unknown, formatters: Formatters) => string,
  _formatters: Formatters,
) => dedent((value as string).trim())
