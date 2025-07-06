import { isAny } from "./base/any"
import type { Formatters } from "./types"

export const json = (
  replacer?:
    | (number | string)[]
    // biome-ignore lint/suspicious/noExplicitAny: this is literally the type signature of JSON.stringify
    | ((this: any, key: string, value: any) => any)
    | null,
  space?: string | number,
): Formatters => [
  // biome-ignore lint/suspicious/noExplicitAny: this is literally the type signature of JSON.stringify
  [isAny, (value: unknown) => JSON.stringify(value, replacer as any, space)],
]
