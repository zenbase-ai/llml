import { describe, expect, it } from "vitest"
import { llml, swagXML } from "../src/index"

describe("Indentation", () => {
  it("should handle basic formatting (indent option ignored for now)", () => {
    const formatters = swagXML({ indent: "  " })
    const result = llml({ message: "Hello" }, formatters)
    const expected = "<message>Hello</message>"
    expect(result).toBe(expected)
  })

  it("should handle list formatting (indent option ignored for now)", () => {
    const formatters = swagXML({ indent: "  " })
    const result = llml({ items: ["a", "b"] }, formatters)
    const expected = ["<items>", "  <items-1>a</items-1>", "  <items-2>b</items-2>", "</items>"].join("\n")
    expect(result).toBe(expected)
  })
})
