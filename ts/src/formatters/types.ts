/**
 * Predicate function that checks if a value matches a specific type
 */
export type Predicate = (value: unknown) => boolean

/**
 * Format function that transforms a value to a string
 */
export type Formatter = (
  value: unknown,
  llml: (data: unknown, formatters: Formatters) => string,
  formatters: Formatters,
) => string

/**
 * Formatters as an iterable of [predicate, formatter] pairs
 */
export type Formatters = Iterable<[Predicate, Formatter]>
