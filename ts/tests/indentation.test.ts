import { describe, expect, it } from "vitest"
import { formatters, llml } from "../src/index"

describe("Indentation", () => {
  it("should handle basic formatting (indent option ignored for now)", () => {
    const formatterSet = formatters.vibeXML({ indent: "  " })
    const result = llml({ message: "Hello" }, formatterSet)
    const expected = "<message>Hello</message>"
    expect(result).toBe(expected)
  })

  it("should handle list formatting (indent option ignored for now)", () => {
    const formatterSet = formatters.vibeXML({ indent: "  " })
    const result = llml({ items: ["a", "b"] }, formatterSet)
    const expected = ["<items>", "  <items-1>a</items-1>", "  <items-2>b</items-2>", "</items>"].join("\n")
    expect(result).toBe(expected)
  })
})
