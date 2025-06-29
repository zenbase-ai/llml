/**
 * Utility functions for LLML
 */

import dedent from "dedent"

import type { LLMLOptions } from "./types"

/**
 * Convert text to kebab-case format
 * @param text - Input text to convert
 * @returns kebab-case formatted string
 */
export const kebabCase = (text: string): string => {
  // Ensure we have a string
  const str = String(text)

  // Replace spaces and underscores with hyphens
  let result = str.replace(/[\s_]+/g, "-")

  // Handle sequences of uppercase letters followed by lowercase (acronyms)
  // e.g., "XMLHttpRequest" -> "XML-Http-Request"
  result = result.replace(/([A-Z]+)([A-Z][a-z])/g, "$1-$2")

  // Handle lowercase followed by uppercase
  // e.g., "camelCase" -> "camel-Case"
  result = result.replace(/([a-z\d])([A-Z])/g, "$1-$2")

  // Convert to lowercase
  return result.toLowerCase()
}

/**
 * Helper function to format a single key-value pair
 */
export const formatKeyValue = (key: string, value: unknown, options: LLMLOptions = {}): string => {
  const { indent = "", prefix = "", strict = false } = options

  const fullKey = prefix ? `${prefix}-${key}` : key
  const kebabKey = kebabCase(fullKey)
  let result = ""

  if (Array.isArray(value)) {
    // Handle list formatting with wrapper tag
    const wrapperTag = kebabKey

    if (value.length === 0) {
      return ""
    }

    result += `${indent}<${wrapperTag}>\n`
    const innerIndent = `${indent}  `

    for (let i = 0; i < value.length; i++) {
      const item = value[i]
      const itemTag = `${kebabKey}-${i + 1}`

      if (typeof item === "object" && item !== null && !Array.isArray(item)) {
        // Dict items need special formatting
        result += `${innerIndent}<${itemTag}>\n`
        const entries = Object.entries(item)
        const itemResults: string[] = []
        const itemIndent = `${innerIndent}  `

        for (let j = 0; j < entries.length; j++) {
          const [subKey, subValue] = entries[j]
          if (j > 0) {
            itemResults.push("\n")
          }
          const subContent = formatKeyValue(subKey, subValue, {
            indent: itemIndent,
            prefix: strict ? itemTag : "",
            strict,
          })
          itemResults.push(subContent)
        }

        const dictContent = itemResults.join("")
        result += dictContent
        result += `\n${innerIndent}</${itemTag}>\n`
      } else {
        // Simple items on one line
        result += `${innerIndent}<${itemTag}>`
        // Handle primitive values inline
        if (typeof item === "string") {
          if (item.includes("\n")) {
            const cleanText = dedent(item).trim()
            const lines = cleanText.split("\n")
            for (let j = 0; j < lines.length; j++) {
              if (j > 0) {
                result += "\n  "
              } else {
                result += "  "
              }
              result += lines[j]
            }
          } else {
            result += dedent(item.trim())
          }
        } else {
          result += String(item)
        }
        result += `</${itemTag}>\n`
      }
    }

    result += `${indent}</${wrapperTag}>`
  } else if (typeof value === "object" && value !== null) {
    // Handle dict formatting - process each key-value pair
    const entries = Object.entries(value)
    const dictResults: string[] = []
    const newIndent = `${indent}  `

    for (let i = 0; i < entries.length; i++) {
      const [subKey, subValue] = entries[i]
      if (i > 0) {
        dictResults.push("\n")
      }
      const subContent = formatKeyValue(subKey, subValue, {
        indent: newIndent,
        prefix: strict ? fullKey : "",
        strict,
      })
      dictResults.push(subContent)
    }

    const dictContent = dictResults.join("")

    if (dictContent.includes("\n")) {
      result += `${indent}<${kebabKey}>\n${dictContent}\n${indent}</${kebabKey}>`
    } else {
      result += `${indent}<${kebabKey}>${dictContent}</${kebabKey}>`
    }
  } else {
    // Handle primitive values (strings, numbers, booleans, null)
    if (typeof value === "string") {
      if (value.includes("\n")) {
        const cleanText = dedent(value).trim()
        const lines = cleanText.split("\n")
        const formattedLines = lines.map(line => `  ${line}`)
        const formattedContent = formattedLines.join("\n")
        result += `${indent}<${kebabKey}>\n${formattedContent}\n${indent}</${kebabKey}>`
      } else {
        result += `${indent}<${kebabKey}>${dedent(value.trim())}</${kebabKey}>`
      }
    } else {
      result += `${indent}<${kebabKey}>${String(value)}</${kebabKey}>`
    }
  }

  return result
}
