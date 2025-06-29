import { describe, it, expect } from 'vitest'
import { llml } from '../src/index'

describe('Kebab Case Conversion', () => {
  it('should convert snake_case and camelCase to kebab-case', () => {
    const result = llml({ user_name: 'Alice', userAge: 30 })
    const expected = '<user-name>Alice</user-name>\n<user-age>30</user-age>'
    expect(result).toBe(expected)
  })

  it('should convert keys with spaces to kebab-case', () => {
    const result = llml({ 'key with spaces': 'value' })
    const expected = '<key-with-spaces>value</key-with-spaces>'
    expect(result).toBe(expected)
  })
})