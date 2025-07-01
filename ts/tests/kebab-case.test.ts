import { describe, expect, it } from "vitest"
import { llml } from "../src/index"

describe("Key Preservation", () => {
  it("should preserve snake_case and camelCase keys as-is", () => {
    const result = llml({ user_name: "Alice", userAge: 30 })
    const expected = "<user_name>Alice</user_name>\n<userAge>30</userAge>"
    expect(result).toBe(expected)
  })

  it("should preserve keys with spaces as-is", () => {
    const result = llml({ "key with spaces": "value" })
    const expected = "<key with spaces>value</key with spaces>"
    expect(result).toBe(expected)
  })

  describe("Key Preservation Details", () => {
    it("should preserve camelCase keys", () => {
      const result = llml({ userName: "Alice", firstName: "Bob" })
      const expected = "<userName>Alice</userName>\n<firstName>Bob</firstName>"
      expect(result).toBe(expected)
    })

    it("should preserve multiple word camel case", () => {
      const result = llml({ getUserName: "function", setUserAge: "method" })
      const expected = "<getUserName>function</getUserName>\n<setUserAge>method</setUserAge>"
      expect(result).toBe(expected)
    })

    it("should preserve acronyms as-is", () => {
      const result = llml({ XMLHttpRequest: "api", HTMLElement: "dom" })
      const expected = "<XMLHttpRequest>api</XMLHttpRequest>\n<HTMLElement>dom</HTMLElement>"
      expect(result).toBe(expected)
    })

    it("should preserve mixed cases with acronyms", () => {
      const result = llml({ XMLParser: "tool", HTTPSConnection: "secure" })
      const expected = "<XMLParser>tool</XMLParser>\n<HTTPSConnection>secure</HTTPSConnection>"
      expect(result).toBe(expected)
    })

    it("should preserve numbers in camelCase", () => {
      const result = llml({ user2Name: "test", config3Value: "data" })
      const expected = "<user2Name>test</user2Name>\n<config3Value>data</config3Value>"
      expect(result).toBe(expected)
    })

    it("should preserve single letter prefixes", () => {
      const result = llml({ iPhone: "device", iPad: "tablet" })
      const expected = "<iPhone>device</iPhone>\n<iPad>tablet</iPad>"
      expect(result).toBe(expected)
    })

    it("should preserve kebab-case keys", () => {
      const result = llml({ "user-name": "Alice", "first-name": "Bob" })
      const expected = "<user-name>Alice</user-name>\n<first-name>Bob</first-name>"
      expect(result).toBe(expected)
    })

    it("should preserve short uppercase sequences", () => {
      const result = llml({ A: "single", AB: "double", ABC: "triple" })
      const expected = "<A>single</A>\n<AB>double</AB>\n<ABC>triple</ABC>"
      expect(result).toBe(expected)
    })

    it("should preserve mixed patterns as-is", () => {
      const result = llml({
        camelCase: "test1",
        snake_case: "test2",
        "kebab-case": "test3",
        PascalCase: "test4",
        UPPER_SNAKE: "test5",
      })
      const expected = [
        "<camelCase>test1</camelCase>",
        "<snake_case>test2</snake_case>",
        "<kebab-case>test3</kebab-case>",
        "<PascalCase>test4</PascalCase>",
        "<UPPER_SNAKE>test5</UPPER_SNAKE>",
      ].join("\n")
      expect(result).toBe(expected)
    })
  })

  describe("Nested Key Preservation", () => {
    it("should preserve keys in nested objects", () => {
      const result = llml({
        userConfig: {
          debugMode: true,
          maxRetries: 5,
          XMLParser: "enabled",
        },
      })
      const expected = [
        "<userConfig>",
        "  <debugMode>true</debugMode>",
        "  <maxRetries>5</maxRetries>",
        "  <XMLParser>enabled</XMLParser>",
        "</userConfig>",
      ].join("\n")
      expect(result).toBe(expected)
    })

    it("should preserve keys in arrays", () => {
      const result = llml({
        userTasks: ["task1", "task2"],
        XMLElements: ["element1", "element2"],
      })
      const expected = [
        "<userTasks>",
        "  <userTasks-1>task1</userTasks-1>",
        "  <userTasks-2>task2</userTasks-2>",
        "</userTasks>",
        "<XMLElements>",
        "  <XMLElements-1>element1</XMLElements-1>",
        "  <XMLElements-2>element2</XMLElements-2>",
        "</XMLElements>",
      ].join("\n")
      expect(result).toBe(expected)
    })
  })
})
