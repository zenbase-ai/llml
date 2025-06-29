import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Basic Value Formatting", () => {
  it("should handle null values", () => {
    const result = llml()
    const expected = ""
    expect(result).toBe(expected)
  })

  it("should handle basic primitives string values", () => {
    const result = llml(0)
    const expected = "0"
    expect(result).toBe(expected)
  })

  it("should handle empty list values", () => {
    const result = llml([])
    const expected = ""
    expect(result).toBe(expected)
  })

  it("should handle empty object values", () => {
    const result = llml({})
    const expected = ""
    expect(result).toBe(expected)
  })

  it("should handle zero numeric values", () => {
    const result = llml({ zero: 0 })
    const expected = "<zero>0</zero>"
    expect(result).toBe(expected)
  })

  it("should handle false boolean values", () => {
    const result = llml({ disabled: false })
    const expected = "<disabled>false</disabled>"
    expect(result).toBe(expected)
  })

  it("should handle null values", () => {
    const result = llml({ nothing: null })
    const expected = "<nothing>null</nothing>"
    expect(result).toBe(expected)
  })

  it("should handle undefined values", () => {
    const result = llml({ nothing: undefined })
    const expected = "<nothing>undefined</nothing>"
    expect(result).toBe(expected)
  })

  it("should format simple string values", () => {
    const result = llml({ instructions: "Follow these steps" })
    const expected = "<instructions>Follow these steps</instructions>"
    expect(result).toBe(expected)
  })

  it("should format integer values", () => {
    const result = llml({ count: 42 })
    const expected = "<count>42</count>"
    expect(result).toBe(expected)
  })

  it("should format float values", () => {
    const result = llml({ temperature: 98.6 })
    const expected = "<temperature>98.6</temperature>"
    expect(result).toBe(expected)
  })

  it("should format true boolean values", () => {
    const result = llml({ enabled: true })
    const expected = "<enabled>true</enabled>"
    expect(result).toBe(expected)
  })

  it("should handle multiple simple values", () => {
    const result = llml({ name: "Alice", age: 30, active: true })
    const expected = "<name>Alice</name>\n<age>30</age>\n<active>true</active>"
    expect(result).toBe(expected)
  })

  it("should handle empty arrays in objects", () => {
    const result = llml({ items: [] })
    const expected = ""
    expect(result).toBe(expected)
  })

  it("should handle empty strings", () => {
    const result = llml({ message: "" })
    const expected = "<message></message>"
    expect(result).toBe(expected)
  })
})
