import { describe, it, expect } from 'vitest'
import { llml } from '../src/index'

describe('Complex Scenarios', () => {
  it('should handle mixed content types', () => {
    const result = llml({
      title: 'My Document',
      sections: ['intro', 'body', 'conclusion'],
      metadata: { author: 'Alice', version: '1.0' }
    })
    const expected = [
      '<title>My Document</title>',
      '<sections-list>',
      '  <sections-1>intro</sections-1>',
      '  <sections-2>body</sections-2>',
      '  <sections-3>conclusion</sections-3>',
      '</sections-list>',
      '<metadata>',
      '  <metadata-author>Alice</metadata-author>',
      '  <metadata-version>1.0</metadata-version>',
      '</metadata>'
    ].join('\n')
    expect(result).toBe(expected)
  })

  it('should handle deeply nested structures', () => {
    const result = llml({
      level1: {
        level2: {
          items: ['a', 'b']
        }
      }
    })
    const expected = [
      '<level1>',
      '  <level1-level2>',
      '    <level1-level2-items-list>',
      '      <level1-level2-items-1>a</level1-level2-items-1>',
      '      <level1-level2-items-2>b</level1-level2-items-2>',
      '    </level1-level2-items-list>',
      '  </level1-level2>',
      '</level1>'
    ].join('\n')
    expect(result).toBe(expected)
  })
})