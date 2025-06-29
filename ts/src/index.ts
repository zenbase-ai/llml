/**
 * LLML - Lightweight Markup Language
 * Converts data structures to XML-like markup
 */

import dedent from 'dedent'
import { kebabCase } from './utils'

export interface LLMLOptions {
  indent?: string
  prefix?: string
}

type LLMLValue = any

/**
 * Main LLML function - converts data to structured markup
 * Supports multiple call signatures:
 * - llml() -> ""
 * - llml(primitive) -> string representation  
 * - llml([]) -> ""
 * - llml({}) -> ""
 * - llml({key: value}) -> <key>value</key>
 */
export function llml(data?: LLMLValue, options: LLMLOptions = {}): string {
  const { indent = '', prefix = '' } = options

  // Handle no arguments
  if (arguments.length === 0) {
    return ''
  }

  // Handle null/undefined
  if (data === null || data === undefined) {
    return data === null ? 'null' : 'undefined'
  }

  // Handle primitives (non-object values)
  if (typeof data !== 'object' || data instanceof Date) {
    return String(data)
  }

  // Handle empty arrays
  if (Array.isArray(data) && data.length === 0) {
    return ''
  }

  // Handle arrays (this should be handled differently later for lists)
  if (Array.isArray(data)) {
    // For now, just return empty for arrays - we'll implement list logic later
    return ''
  }

  // Handle empty objects
  if (Object.keys(data).length === 0) {
    return ''
  }

  // Handle objects with key-value pairs
  const buffer: string[] = []
  const entries = Object.entries(data)

  for (let i = 0; i < entries.length; i++) {
    const [key, value] = entries[i]
    
    if (i > 0) {
      buffer.push('\n')
    }

    const fullKey = prefix ? `${prefix}-${key}` : key
    const kebabKey = kebabCase(fullKey)

    // Handle arrays with special list formatting
    if (Array.isArray(value)) {
      buffer.push(indent + formatList(value, indent, fullKey))
    } else {
      const formattedValue = formatValue(value, indent, fullKey)
      
      if (formattedValue.includes('\n')) {
        buffer.push(`${indent}<${kebabKey}>\n${formattedValue}\n${indent}</${kebabKey}>`)
      } else {
        buffer.push(`${indent}<${kebabKey}>${formattedValue}</${kebabKey}>`)
      }
    }
  }

  return buffer.join('')
}

/**
 * Format a single value with proper handling for different types
 */
function formatValue(value: any, currentIndent: string, currentPrefix: string = ''): string {
  if (value === null) {
    return 'null'
  }
  
  if (value === undefined) {
    return 'undefined'
  }

  if (typeof value === 'string') {
    if (value.includes('\n')) {
      // Handle multiline strings with dedent
      const clean = dedent(value).trim()
      const lines = clean.split('\n')
      return lines.map((line, i) => {
        if (i === 0) {
          return `  ${line}`
        } else {
          return `\n  ${line}`
        }
      }).join('')
    } else {
      return value.trim()
    }
  }

  if (typeof value === 'number' || typeof value === 'boolean') {
    return String(value)
  }

  // Handle arrays with list formatting
  if (Array.isArray(value)) {
    return formatList(value, currentIndent, currentPrefix)
  }

  // Handle nested objects
  if (typeof value === 'object') {
    return llml(value, { 
      indent: currentIndent + '  ', 
      prefix: currentPrefix 
    })
  }

  return String(value)
}

/**
 * Format a list with wrapper tag and numbered items
 */
function formatList(items: any[], currentIndent: string, currentPrefix: string): string {
  const innerIndent = currentIndent + '  '
  const kebabPrefix = kebabCase(currentPrefix)
  const wrapperTag = `${kebabPrefix}-list`

  // Handle empty lists
  if (items.length === 0) {
    return `<${wrapperTag}></${wrapperTag}>`
  }

  const buffer: string[] = []
  buffer.push(`<${wrapperTag}>\n`)

  for (let i = 0; i < items.length; i++) {
    const item = items[i]
    const numberedTag = `${kebabPrefix}-${i + 1}`
    
    if (typeof item === 'object' && item !== null && !Array.isArray(item)) {
      // Handle dictionary items
      buffer.push(`${innerIndent}<${numberedTag}>\n`)
      const dictContent = llml(item, { 
        indent: innerIndent + '  ', 
        prefix: numberedTag 
      })
      buffer.push(dictContent)
      buffer.push(`\n${innerIndent}</${numberedTag}>\n`)
    } else {
      // Handle simple items
      buffer.push(`${innerIndent}<${numberedTag}>`)
      const formattedItem = formatValue(item, innerIndent, numberedTag)
      buffer.push(formattedItem)
      buffer.push(`</${numberedTag}>\n`)
    }
  }

  buffer.push(`${currentIndent}</${wrapperTag}>`)
  return buffer.join('')
}

// Default export for convenience
export default llml