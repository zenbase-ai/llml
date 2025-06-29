import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Prefix Support", () => {
  it("should apply prefix to simple values", () => {
    const result = llml({ config: "value" }, { prefix: "app" })
    const expected = "<app-config>value</app-config>"
    expect(result).toBe(expected)
  })

  it("should apply prefix to list formatting", () => {
    const result = llml({ items: ["a", "b"] }, { prefix: "app" })
    const expected = [
      "<app-items>",
      "  <app-items-1>a</app-items-1>",
      "  <app-items-2>b</app-items-2>",
      "</app-items>",
    ].join("\n")
    expect(result).toBe(expected)
  })
})
