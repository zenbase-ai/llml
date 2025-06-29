import type { LLMLOptions } from "./types"
import { formatKeyValue } from "./utils"

// Simple function overloads
export function llml(): string
export function llml(data: unknown, options?: LLMLOptions): string
export function llml(data?: unknown, options: LLMLOptions = {}): string {
  const { indent = "", prefix = "", strict = false } = options

  // Handle no arguments
  if (data == null) {
    return ""
  }

  // Handle primitives (non-object values)
  if (typeof data !== "object" || data instanceof Date) {
    return String(data)
  }

  // Handle empty arrays and objects
  if (Array.isArray(data) && data.length === 0) {
    return ""
  }
  if (typeof data === "object" && Object.keys(data).length === 0) {
    return ""
  }

  // Handle arrays - format as numbered items
  if (Array.isArray(data)) {
    const results: string[] = []

    for (let i = 0; i < data.length; i++) {
      const item = data[i]
      const itemTag = String(i + 1)

      if (i > 0) {
        results.push("\n")
      }

      if (typeof item === "object" && item !== null && !Array.isArray(item)) {
        // Handle object items with nested structure
        results.push(`${indent}<${itemTag}>\n`)
        const entries = Object.entries(item)
        const itemResults: string[] = []
        const itemIndent = `${indent}  `

        for (let j = 0; j < entries.length; j++) {
          const [subKey, subValue] = entries[j]
          if (j > 0) {
            itemResults.push("\n")
          }
          const subContent = formatKeyValue(subKey, subValue, {
            indent: itemIndent,
            prefix: itemTag,
            strict,
          })
          itemResults.push(subContent)
        }

        results.push(itemResults.join(""))
        results.push(`\n${indent}</${itemTag}>`)
      } else {
        // Handle primitive items
        results.push(`${indent}<${itemTag}>${String(item)}</${itemTag}>`)
      }
    }

    return results.join("")
  }

  // Handle objects with key-value pairs
  const entries = Object.entries(data)
  const results: string[] = []

  for (let i = 0; i < entries.length; i++) {
    const [key, value] = entries[i]

    if (i > 0) {
      results.push("\n")
    }

    const formattedContent = formatKeyValue(key, value, { indent, prefix, strict })
    results.push(formattedContent)
  }

  return results.join("")
}
