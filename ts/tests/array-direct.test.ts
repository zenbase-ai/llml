import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Direct Array Formatting", () => {
  it("should format direct arrays with numbered tags", () => {
    const result = llml(["a", "b", "c"])
    const expected = "<1>a</1>\n<2>b</2>\n<3>c</3>"
    expect(result).toBe(expected)
  })

  it("should format arrays with different types", () => {
    const result = llml([1, "hello", true])
    const expected = "<1>1</1>\n<2>hello</2>\n<3>true</3>"
    expect(result).toBe(expected)
  })

  it("should handle arrays with objects", () => {
    const result = llml([{ name: "Alice" }, { name: "Bob" }])
    const expected = ["<1>", "  <1-name>Alice</1-name>", "</1>", "<2>", "  <2-name>Bob</2-name>", "</2>"].join("\n")
    expect(result).toBe(expected)
  })

  it("should handle empty arrays", () => {
    const result = llml([])
    const expected = ""
    expect(result).toBe(expected)
  })

  it("should handle arrays with indentation options", () => {
    const result = llml(["a", "b"], { indent: "  " })
    const expected = "  <1>a</1>\n  <2>b</2>"
    expect(result).toBe(expected)
  })
})
