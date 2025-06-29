import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Complex Scenarios", () => {
  it("should handle mixed content types", () => {
    const result = llml({
      title: "My Document",
      sections: ["intro", "body", "conclusion"],
      metadata: { author: "Alice", version: "1.0" },
    })
    const expected = [
      "<title>My Document</title>",
      "<sections>",
      "  <sections-1>intro</sections-1>",
      "  <sections-2>body</sections-2>",
      "  <sections-3>conclusion</sections-3>",
      "</sections>",
      "<metadata>",
      "  <author>Alice</author>",
      "  <version>1.0</version>",
      "</metadata>",
    ].join("\n")
    expect(result).toBe(expected)
  })

  it("should handle deeply nested structures", () => {
    const result = llml({
      level1: {
        level2: {
          items: ["a", "b"],
        },
      },
    })
    const expected = [
      "<level1>",
      "  <level2>",
      "    <items>",
      "      <items-1>a</items-1>",
      "      <items-2>b</items-2>",
      "    </items>",
      "  </level2>",
      "</level1>",
    ].join("\n")
    expect(result).toBe(expected)
  })
})
