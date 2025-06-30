import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Nested Structures", () => {
  it("should handle nested objects with strict:false (default)", () => {
    const result = llml({ config: { debug: true, timeout: 30 } })
    const expected = "<config>\n  <debug>true</debug>\n  <timeout>30</timeout>\n</config>"
    expect(result).toBe(expected)
  })

  it("should handle nested objects with kebab-case conversion (strict:false)", () => {
    const result = llml({ user_config: { debug_mode: true, maxRetries: 5 } })
    const expected = [
      "<user-config>",
      "  <debug-mode>true</debug-mode>",
      "  <max-retries>5</max-retries>",
      "</user-config>",
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

  describe("Strict Mode", () => {
    it("should handle nested objects with strict:true", () => {
      const result = llml({ config: { debug: true, timeout: 30 } }, { strict: true })
      const expected = "<config>\n  <config-debug>true</config-debug>\n  <config-timeout>30</config-timeout>\n</config>"
      expect(result).toBe(expected)
    })

    it("should handle nested objects with kebab-case conversion (strict:true)", () => {
      const result = llml({ user_config: { debug_mode: true, maxRetries: 5 } }, { strict: true })
      const expected = [
        "<user-config>",
        "  <user-config-debug-mode>true</user-config-debug-mode>",
        "  <user-config-max-retries>5</user-config-max-retries>",
        "</user-config>",
      ].join("\n")
      expect(result).toBe(expected)
    })

    it("should handle arrays containing objects (strict:true)", () => {
      const result = llml(
        {
          data: [
            { name: "Alice", age: 30 },
            { name: "Bob", age: 25 },
          ],
        },
        { strict: true },
      )
      const expected = [
        "<data>",
        "  <data-1>",
        "    <data-1-name>Alice</data-1-name>",
        "    <data-1-age>30</data-1-age>",
        "  </data-1>",
        "  <data-2>",
        "    <data-2-name>Bob</data-2-name>",
        "    <data-2-age>25</data-2-age>",
        "  </data-2>",
        "</data>",
      ].join("\n")
      expect(result).toBe(expected)
    })
  })
})
