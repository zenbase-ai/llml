/**
 * LLML - Lightweight Markup Language
 * Converts data structures to XML-like markup
 */

// Re-export built-in formatters for convenience
export * from "./formatters"
export { vibeXML } from "./formatters"
export { llml } from "./llml"
export type { Formatter, Formatters, Predicate } from "./types"
