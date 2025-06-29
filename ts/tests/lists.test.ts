import { describe, it, expect } from 'vitest'
import { llml } from '../src/index'

describe('List Formatting', () => {
  it('should handle empty arrays', () => {
    const result = llml({ items: [] })
    const expected = '<items-list></items-list>'
    expect(result).toBe(expected)
  })

  it('should format simple lists with wrapper tags and numbered items', () => {
    const result = llml({ rules: ['first', 'second', 'third'] })
    const expected = [
      '<rules-list>',
      '  <rules-1>first</rules-1>',
      '  <rules-2>second</rules-2>',
      '  <rules-3>third</rules-3>',
      '</rules-list>'
    ].join('\n')
    expect(result).toBe(expected)
  })

  it('should format lists with numeric values', () => {
    const result = llml({ numbers: [1, 2, 3] })
    const expected = [
      '<numbers-list>',
      '  <numbers-1>1</numbers-1>',
      '  <numbers-2>2</numbers-2>',
      '  <numbers-3>3</numbers-3>',
      '</numbers-list>'
    ].join('\n')
    expect(result).toBe(expected)
  })

  it('should convert list names to kebab-case', () => {
    const result = llml({ user_tasks: ['task1', 'task2'] })
    const expected = [
      '<user-tasks-list>',
      '  <user-tasks-1>task1</user-tasks-1>',
      '  <user-tasks-2>task2</user-tasks-2>',
      '</user-tasks-list>'
    ].join('\n')
    expect(result).toBe(expected)
  })
})