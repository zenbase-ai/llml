import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Indentation", () => {
  it("should handle basic indentation with string indent", () => {
    const result = llml({ message: "Hello" }, { indent: "  " })
    const expected = "  <message>Hello</message>"
    expect(result).toBe(expected)
  })

  it("should handle list formatting with indentation", () => {
    const result = llml({ items: ["a", "b"] }, { indent: "  " })
    const expected = ["  <items-list>", "    <items-1>a</items-1>", "    <items-2>b</items-2>", "  </items-list>"].join(
      "\n",
    )
    expect(result).toBe(expected)
  })
})
