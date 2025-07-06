import { describe, expect, it } from "vitest"
import { formatters, llml } from "../src/index"

describe("Prefix Support", () => {
  it("should handle basic formatting (prefix option ignored for now)", () => {
    const formatterSet = formatters.vibeXML({ prefix: "app" })
    const result = llml({ config: "value" }, formatterSet)
    const expected = "<config>value</config>"
    expect(result).toBe(expected)
  })

  it("should handle list formatting (prefix option ignored for now)", () => {
    const formatterSet = formatters.vibeXML({ prefix: "app" })
    const result = llml({ items: ["a", "b"] }, formatterSet)
    const expected = ["<items>", "  <items-1>a</items-1>", "  <items-2>b</items-2>", "</items>"].join("\n")
    expect(result).toBe(expected)
  })
})
