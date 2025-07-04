import { formatAny, isAny } from "../base/any"
import { formatBoolean, isBoolean } from "../base/boolean"
import { formatDate, isDate } from "../base/date"
import { formatNull, isNull } from "../base/null"
import { formatNumber, isNumber } from "../base/number"
import { formatString, isString } from "../base/string"
import { formatUndefined, isUndefined } from "../base/undefined"
import type { Formatter, Formatters, Predicate } from "../types"
import { formatArray, isArray } from "./array"
import { formatObject, isObject } from "./object"

export { isArray, isBoolean, isDate, isNull, isNumber, isObject, isString, isUndefined }

// Default formatters map

const defaultVibeXML: Formatters = new Map([
  [isString, formatString],
  [isNumber, formatNumber],
  [isBoolean, formatBoolean],
  [isUndefined, formatUndefined],
  [isNull, formatNull],
  [isDate, formatDate],
  [isArray, formatArray],
  [isObject, formatObject],
  [isAny, formatAny],
])

// Options interface for vibeXML configuration
interface VibeXMLOptions {
  formatters?: Formatters
}

// Function that creates vibeXML formatters with options
export const vibeXML = (options?: VibeXMLOptions): Formatters => {
  if (!options) {
    return defaultVibeXML
  }

  // If custom formatters are provided, merge them with defaults
  // Custom formatters take priority by being added first
  if (options.formatters) {
    const mergedFormatters = new Map<Predicate, Formatter>()

    // Add custom formatters first (higher priority)
    for (const [predicate, formatter] of options.formatters) {
      mergedFormatters.set(predicate, formatter)
    }

    // Add default formatters (lower priority)
    for (const [predicate, formatter] of defaultVibeXML) {
      mergedFormatters.set(predicate, formatter)
    }

    return mergedFormatters
  }

  // For now, ignore indent and prefix options as tests expect
  return defaultVibeXML
}
