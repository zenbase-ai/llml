import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("List Formatting", () => {
  it("should handle empty arrays", () => {
    const result = llml({ items: [] })
    const expected = ""
    expect(result).toBe(expected)
  })

  it("should format simple lists with wrapper tags and numbered items", () => {
    const result = llml({ rules: ["first", "second", "third"] })
    const expected = [
      "<rules>",
      "  <rules-1>first</rules-1>",
      "  <rules-2>second</rules-2>",
      "  <rules-3>third</rules-3>",
      "</rules>",
    ].join("\n")
    expect(result).toBe(expected)
  })

  it("should format lists with numeric values", () => {
    const result = llml({ numbers: [1, 2, 3] })
    const expected = [
      "<numbers>",
      "  <numbers-1>1</numbers-1>",
      "  <numbers-2>2</numbers-2>",
      "  <numbers-3>3</numbers-3>",
      "</numbers>",
    ].join("\n")
    expect(result).toBe(expected)
  })

  it("should convert list names to kebab-case", () => {
    const result = llml({ user_tasks: ["task1", "task2"] })
    const expected = [
      "<user-tasks>",
      "  <user-tasks-1>task1</user-tasks-1>",
      "  <user-tasks-2>task2</user-tasks-2>",
      "</user-tasks>",
    ].join("\n")
    expect(result).toBe(expected)
  })
})
