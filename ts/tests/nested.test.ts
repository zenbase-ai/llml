import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Nested Structures", () => {
  it("should handle nested objects with strict:false (default)", () => {
    const result = llml({ config: { debug: true, timeout: 30 } })
    const expected = "<config>\n  <debug>true</debug>\n  <timeout>30</timeout>\n</config>"
    expect(result).toBe(expected)
  })

  it("should preserve keys in nested objects", () => {
    const result = llml({ user_config: { debug_mode: true, maxRetries: 5 } })
    const expected = [
      "<user_config>",
      "  <debug_mode>true</debug_mode>",
      "  <maxRetries>5</maxRetries>",
      "</user_config>",
    ].join("\n")
    expect(result).toBe(expected)
  })

  it("should handle arrays containing objects (strict:false)", () => {
    const result = llml({
      data: [
        { name: "Alice", age: 30 },
        { name: "Bob", age: 25 },
      ],
    })
    const expected = [
      "<data>",
      "  <data-1>",
      "    <name>Alice</name>",
      "    <age>30</age>",
      "  </data-1>",
      "  <data-2>",
      "    <name>Bob</name>",
      "    <age>25</age>",
      "  </data-2>",
      "</data>",
    ].join("\n")
    expect(result).toBe(expected)
  })
})
