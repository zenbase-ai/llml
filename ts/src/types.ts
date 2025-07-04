import type { Formatters } from "./formatters/types"

export interface Renderer {
  render(data: unknown): string
}

export interface VibeXMLOptions {
  indent?: string
  prefix?: string
  formatters?: Formatters
}

export interface LLMLOptions {
  renderer?: Renderer
}

// Re-export formatter types for convenience
export type {
  Formatter,
  Formatters,
  Predicate,
} from "./formatters/types"
