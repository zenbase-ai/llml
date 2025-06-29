import { describe, it, expect } from 'vitest'
import { llml } from '../src/index'

describe('Multiline Content', () => {
  it('should handle multiline content with dedent', () => {
    const content = `
    Line 1
    Line 2
    Line 3
    `
    const result = llml({ description: content })
    const expected = '<description>\n  Line 1\n  Line 2\n  Line 3\n</description>'
    expect(result).toBe(expected)
  })
})