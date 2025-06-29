import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Kebab Case Conversion", () => {
  it("should convert snake_case and camelCase to kebab-case", () => {
    const result = llml({ user_name: "Alice", userAge: 30 })
    const expected = "<user-name>Alice</user-name>\n<user-age>30</user-age>"
    expect(result).toBe(expected)
  })

  it("should convert keys with spaces to kebab-case", () => {
    const result = llml({ "key with spaces": "value" })
    const expected = "<key-with-spaces>value</key-with-spaces>"
    expect(result).toBe(expected)
  })

  describe("Advanced camelCase Conversion", () => {
    it("should handle basic camelCase", () => {
      const result = llml({ userName: "Alice", firstName: "Bob" })
      const expected = "<user-name>Alice</user-name>\n<first-name>Bob</first-name>"
      expect(result).toBe(expected)
    })

    it("should handle multiple word camel case", () => {
      const result = llml({ getUserName: "function", setUserAge: "method" })
      const expected = "<get-user-name>function</get-user-name>\n<set-user-age>method</set-user-age>"
      expect(result).toBe(expected)
    })

    it("should handle acronyms correctly", () => {
      const result = llml({ XMLHttpRequest: "api", HTMLElement: "dom" })
      const expected = "<xml-http-request>api</xml-http-request>\n<html-element>dom</html-element>"
      expect(result).toBe(expected)
    })

    it("should handle mixed cases with acronyms", () => {
      const result = llml({ XMLParser: "tool", HTTPSConnection: "secure" })
      const expected = "<xml-parser>tool</xml-parser>\n<https-connection>secure</https-connection>"
      expect(result).toBe(expected)
    })

    it("should handle numbers in camelCase", () => {
      const result = llml({ user2Name: "test", config3Value: "data" })
      const expected = "<user2-name>test</user2-name>\n<config3-value>data</config3-value>"
      expect(result).toBe(expected)
    })

    it("should handle single letter prefixes", () => {
      const result = llml({ iPhone: "device", iPad: "tablet" })
      const expected = "<i-phone>device</i-phone>\n<i-pad>tablet</i-pad>"
      expect(result).toBe(expected)
    })

    it("should preserve already kebab-case keys", () => {
      const result = llml({ "user-name": "Alice", "first-name": "Bob" })
      const expected = "<user-name>Alice</user-name>\n<first-name>Bob</first-name>"
      expect(result).toBe(expected)
    })

    it("should handle short uppercase sequences", () => {
      const result = llml({ A: "single", AB: "double", ABC: "triple" })
      const expected = "<a>single</a>\n<ab>double</ab>\n<abc>triple</abc>"
      expect(result).toBe(expected)
    })

    it("should handle mixed patterns", () => {
      const result = llml({
        camelCase: "test1",
        snake_case: "test2",
        "kebab-case": "test3",
        PascalCase: "test4",
        UPPER_SNAKE: "test5",
      })
      const expected = [
        "<camel-case>test1</camel-case>",
        "<snake-case>test2</snake-case>",
        "<kebab-case>test3</kebab-case>",
        "<pascal-case>test4</pascal-case>",
        "<upper-snake>test5</upper-snake>",
      ].join("\n")
      expect(result).toBe(expected)
    })
  })

  describe("Nested camelCase Keys", () => {
    it("should convert camelCase in nested objects", () => {
      const result = llml({
        userConfig: {
          debugMode: true,
          maxRetries: 5,
          XMLParser: "enabled",
        },
      })
      const expected = [
        "<user-config>",
        "  <debug-mode>true</debug-mode>",
        "  <max-retries>5</max-retries>",
        "  <xml-parser>enabled</xml-parser>",
        "</user-config>",
      ].join("\n")
      expect(result).toBe(expected)
    })

    it("should convert camelCase in array keys", () => {
      const result = llml({
        userTasks: ["task1", "task2"],
        XMLElements: ["element1", "element2"],
      })
      const expected = [
        "<user-tasks>",
        "  <user-tasks-1>task1</user-tasks-1>",
        "  <user-tasks-2>task2</user-tasks-2>",
        "</user-tasks>",
        "<xml-elements>",
        "  <xml-elements-1>element1</xml-elements-1>",
        "  <xml-elements-2>element2</xml-elements-2>",
        "</xml-elements>",
      ].join("\n")
      expect(result).toBe(expected)
    })
  })
})
