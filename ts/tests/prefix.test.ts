import { describe, expect, it } from "vitest"
import { llml, swagXML } from "../src/index"

describe("Prefix Support", () => {
  it("should handle basic formatting (prefix option ignored for now)", () => {
    const formatters = swagXML({ prefix: "app" })
    const result = llml({ config: "value" }, formatters)
    const expected = "<config>value</config>"
    expect(result).toBe(expected)
  })

  it("should handle list formatting (prefix option ignored for now)", () => {
    const formatters = swagXML({ prefix: "app" })
    const result = llml({ items: ["a", "b"] }, formatters)
    const expected = ["<items>", "  <items-1>a</items-1>", "  <items-2>b</items-2>", "</items>"].join("\n")
    expect(result).toBe(expected)
  })
})
