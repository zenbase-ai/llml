/**
 * Utility functions for LLML
 */

/**
 * Convert text to kebab-case format
 * @param text - Input text to convert
 * @returns kebab-case formatted string
 */
export function kebabCase(text: string): string {
  // Replace spaces and underscores with hyphens
  text = text.replace(/[\s_]+/g, '-')
  // Insert hyphens before capital letters (except at the start)
  text = text.replace(/(?<!^)(?=[A-Z])/g, '-')
  // Convert to lowercase
  return text.toLowerCase()
}

