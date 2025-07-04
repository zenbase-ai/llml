# Go Formatter System Design Analysis

This document explores three idiomatic Go approaches for implementing a formatter system similar to the TypeScript implementation, which uses an array of predicate-formatter pairs.

## Current State

The existing `llml.go` uses a large switch statement approach for type detection and formatting. While functional, it's not easily extensible and requires modifying the core function to add new formatters.

## Approach 1: Interface-Based Formatter Registry

### Overview
Uses Go interfaces to define contracts for predicates and formatters, with a registry to manage them.

### Code Example

```go
package llml

import (
    "fmt"
    "strings"
)

// Predicate interface for type checking
type Predicate interface {
    Test(value interface{}) bool
}

// Formatter interface for value formatting
type Formatter interface {
    Format(value interface{}, llml func(interface{}, *Registry) string, registry *Registry) string
}

// FormatterRule combines a predicate with its formatter
type FormatterRule struct {
    Predicate Predicate
    Formatter Formatter
}

// Registry manages formatter rules
type Registry struct {
    rules []FormatterRule
}

// NewRegistry creates a new formatter registry
func NewRegistry() *Registry {
    return &Registry{
        rules: make([]FormatterRule, 0),
    }
}

// Register adds a new formatter rule
func (r *Registry) Register(predicate Predicate, formatter Formatter) {
    r.rules = append(r.rules, FormatterRule{
        Predicate: predicate,
        Formatter: formatter,
    })
}

// Format finds and applies the first matching formatter
func (r *Registry) Format(value interface{}, llmlFunc func(interface{}, *Registry) string) string {
    for _, rule := range r.rules {
        if rule.Predicate.Test(value) {
            return rule.Formatter.Format(value, llmlFunc, r)
        }
    }
    return fmt.Sprintf("%v", value) // fallback
}

// Concrete implementations
type StringPredicate struct{}
func (p StringPredicate) Test(value interface{}) bool {
    _, ok := value.(string)
    return ok
}

type StringFormatter struct{}
func (f StringFormatter) Format(value interface{}, llml func(interface{}, *Registry) string, registry *Registry) string {
    s := value.(string)
    return strings.TrimSpace(s)
}

type MapPredicate struct{}
func (p MapPredicate) Test(value interface{}) bool {
    _, ok := value.(map[string]interface{})
    return ok
}

type MapFormatter struct{}
func (f MapFormatter) Format(value interface{}, llml func(interface{}, *Registry) string, registry *Registry) string {
    m := value.(map[string]interface{})
    if len(m) == 0 {
        return ""
    }
    
    var parts []string
    for key, val := range m {
        formatted := llml(val, registry)
        if formatted != "" {
            parts = append(parts, fmt.Sprintf("<%s>%s</%s>", key, formatted, key))
        }
    }
    return strings.Join(parts, "\n")
}

// Usage example
func SprintfWithRegistry(data interface{}, registry *Registry) string {
    if registry == nil {
        registry = DefaultRegistry()
    }
    return registry.Format(data, SprintfWithRegistry)
}

func DefaultRegistry() *Registry {
    registry := NewRegistry()
    registry.Register(StringPredicate{}, StringFormatter{})
    registry.Register(MapPredicate{}, MapFormatter{})
    // Add more default formatters...
    return registry
}
```

### Pros
- Clear separation of concerns
- Easy to test individual components
- Extensible through interface implementation
- Type-safe with compile-time checks
- Familiar OOP patterns

### Cons
- More verbose than other approaches
- Requires more boilerplate code
- Interface overhead (minimal performance impact)

## Approach 2: Function Type-Based Registry

### Overview
Uses function types for predicates and formatters, leveraging Go's first-class functions.

### Code Example

```go
package llml

import (
    "fmt"
    "strings"
)

// PredicateFunc is a function that tests if a value matches a type
type PredicateFunc func(value interface{}) bool

// FormatterFunc is a function that formats a value
type FormatterFunc func(value interface{}, llml func(interface{}, *FormatterRegistry) string, registry *FormatterRegistry) string

// FormatterPair represents a predicate-formatter pair
type FormatterPair struct {
    Predicate PredicateFunc
    Formatter FormatterFunc
}

// FormatterRegistry manages formatter pairs
type FormatterRegistry struct {
    formatters []FormatterPair
}

// NewFormatterRegistry creates a new formatter registry
func NewFormatterRegistry() *FormatterRegistry {
    return &FormatterRegistry{
        formatters: make([]FormatterPair, 0),
    }
}

// Register adds a new formatter pair
func (r *FormatterRegistry) Register(predicate PredicateFunc, formatter FormatterFunc) {
    r.formatters = append(r.formatters, FormatterPair{
        Predicate: predicate,
        Formatter: formatter,
    })
}

// Format applies the first matching formatter
func (r *FormatterRegistry) Format(value interface{}, llmlFunc func(interface{}, *FormatterRegistry) string) string {
    for _, pair := range r.formatters {
        if pair.Predicate(value) {
            return pair.Formatter(value, llmlFunc, r)
        }
    }
    return fmt.Sprintf("%v", value) // fallback
}

// Predicate functions
var (
    IsString = func(value interface{}) bool {
        _, ok := value.(string)
        return ok
    }
    
    IsMap = func(value interface{}) bool {
        _, ok := value.(map[string]interface{})
        return ok
    }
    
    IsSlice = func(value interface{}) bool {
        _, ok := value.([]interface{})
        return ok
    }
    
    IsNumber = func(value interface{}) bool {
        switch value.(type) {
        case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
            return true
        default:
            return false
        }
    }
)

// Formatter functions
var (
    FormatString = func(value interface{}, llml func(interface{}, *FormatterRegistry) string, registry *FormatterRegistry) string {
        s := value.(string)
        return strings.TrimSpace(s)
    }
    
    FormatMap = func(value interface{}, llml func(interface{}, *FormatterRegistry) string, registry *FormatterRegistry) string {
        m := value.(map[string]interface{})
        if len(m) == 0 {
            return ""
        }
        
        var parts []string
        for key, val := range m {
            formatted := llml(val, registry)
            if formatted != "" {
                parts = append(parts, fmt.Sprintf("<%s>%s</%s>", key, formatted, key))
            }
        }
        return strings.Join(parts, "\n")
    }
    
    FormatSlice = func(value interface{}, llml func(interface{}, *FormatterRegistry) string, registry *FormatterRegistry) string {
        slice := value.([]interface{})
        if len(slice) == 0 {
            return ""
        }
        
        var parts []string
        for i, item := range slice {
            formatted := llml(item, registry)
            if formatted != "" {
                parts = append(parts, fmt.Sprintf("<%d>%s</%d>", i+1, formatted, i+1))
            }
        }
        return strings.Join(parts, "\n")
    }
    
    FormatNumber = func(value interface{}, llml func(interface{}, *FormatterRegistry) string, registry *FormatterRegistry) string {
        return fmt.Sprintf("%v", value)
    }
)

// Usage example
func SprintfWithFunctions(data interface{}, registry *FormatterRegistry) string {
    if registry == nil {
        registry = DefaultFormatterRegistry()
    }
    return registry.Format(data, SprintfWithFunctions)
}

func DefaultFormatterRegistry() *FormatterRegistry {
    registry := NewFormatterRegistry()
    registry.Register(IsString, FormatString)
    registry.Register(IsNumber, FormatNumber)
    registry.Register(IsSlice, FormatSlice)
    registry.Register(IsMap, FormatMap) // Map last for fallback
    return registry
}

// Custom formatter example
func RegisterCustomFormatter(registry *FormatterRegistry) {
    // Custom predicate for email strings
    isEmail := func(value interface{}) bool {
        if s, ok := value.(string); ok {
            return strings.Contains(s, "@")
        }
        return false
    }
    
    // Custom formatter for emails
    formatEmail := func(value interface{}, llml func(interface{}, *FormatterRegistry) string, registry *FormatterRegistry) string {
        email := value.(string)
        return fmt.Sprintf("<email>%s</email>", strings.TrimSpace(email))
    }
    
    registry.Register(isEmail, formatEmail)
}
```

### Pros
- Concise and readable
- Leverages Go's functional programming capabilities
- Easy to create inline formatters
- Minimal boilerplate
- Closures allow for stateful formatters

### Cons
- Less type safety than interfaces
- Function signatures can be complex
- Harder to document than interfaces

## Approach 3: Generics-Based Type-Safe Registry

### Overview
Uses Go 1.18+ generics for type-safe formatter registration and execution.

### Code Example

```go
package llml

import (
    "fmt"
    "reflect"
    "strings"
)

// TypeConstraint defines what types can be formatted
type TypeConstraint interface {
    ~string | ~int | ~float64 | ~bool | 
    map[string]interface{} | []interface{} | interface{}
}

// TypedPredicate is a generic predicate function
type TypedPredicate[T TypeConstraint] func(value interface{}) (T, bool)

// TypedFormatter is a generic formatter function
type TypedFormatter[T TypeConstraint] func(value T, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string

// GenericFormatterRule represents a type-safe formatter rule
type GenericFormatterRule struct {
    predicate func(interface{}) (interface{}, bool)
    formatter func(interface{}, func(interface{}, *GenericRegistry) string, *GenericRegistry) string
    typeName  string
}

// GenericRegistry manages type-safe formatters
type GenericRegistry struct {
    rules []GenericFormatterRule
}

// NewGenericRegistry creates a new generic registry
func NewGenericRegistry() *GenericRegistry {
    return &GenericRegistry{
        rules: make([]GenericFormatterRule, 0),
    }
}

// Register adds a type-safe formatter
func Register[T TypeConstraint](r *GenericRegistry, predicate TypedPredicate[T], formatter TypedFormatter[T]) {
    rule := GenericFormatterRule{
        predicate: func(value interface{}) (interface{}, bool) {
            return predicate(value)
        },
        formatter: func(value interface{}, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string {
            typedValue, ok := predicate(value)
            if !ok {
                return ""
            }
            return formatter(typedValue, llml, registry)
        },
        typeName: fmt.Sprintf("%T", *new(T)),
    }
    r.rules = append(r.rules, rule)
}

// Format applies the first matching formatter
func (r *GenericRegistry) Format(value interface{}, llmlFunc func(interface{}, *GenericRegistry) string) string {
    for _, rule := range r.rules {
        if _, ok := rule.predicate(value); ok {
            return rule.formatter(value, llmlFunc, r)
        }
    }
    return fmt.Sprintf("%v", value) // fallback
}

// Type-safe predicates
var (
    StringPredicate = func(value interface{}) (string, bool) {
        s, ok := value.(string)
        return s, ok
    }
    
    MapPredicate = func(value interface{}) (map[string]interface{}, bool) {
        m, ok := value.(map[string]interface{})
        return m, ok
    }
    
    SlicePredicate = func(value interface{}) ([]interface{}, bool) {
        s, ok := value.([]interface{})
        return s, ok
    }
    
    IntPredicate = func(value interface{}) (int, bool) {
        i, ok := value.(int)
        return i, ok
    }
    
    Float64Predicate = func(value interface{}) (float64, bool) {
        f, ok := value.(float64)
        return f, ok
    }
)

// Type-safe formatters
var (
    StringFormatter = func(value string, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string {
        return strings.TrimSpace(value)
    }
    
    MapFormatter = func(value map[string]interface{}, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string {
        if len(value) == 0 {
            return ""
        }
        
        var parts []string
        for key, val := range value {
            formatted := llml(val, registry)
            if formatted != "" {
                parts = append(parts, fmt.Sprintf("<%s>%s</%s>", key, formatted, key))
            }
        }
        return strings.Join(parts, "\n")
    }
    
    SliceFormatter = func(value []interface{}, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string {
        if len(value) == 0 {
            return ""
        }
        
        var parts []string
        for i, item := range value {
            formatted := llml(item, registry)
            if formatted != "" {
                parts = append(parts, fmt.Sprintf("<%d>%s</%d>", i+1, formatted, i+1))
            }
        }
        return strings.Join(parts, "\n")
    }
    
    IntFormatter = func(value int, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string {
        return fmt.Sprintf("%d", value)
    }
    
    Float64Formatter = func(value float64, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string {
        return fmt.Sprintf("%g", value)
    }
)

// Usage example
func SprintfWithGenerics(data interface{}, registry *GenericRegistry) string {
    if registry == nil {
        registry = DefaultGenericRegistry()
    }
    return registry.Format(data, SprintfWithGenerics)
}

func DefaultGenericRegistry() *GenericRegistry {
    registry := NewGenericRegistry()
    Register(registry, StringPredicate, StringFormatter)
    Register(registry, IntPredicate, IntFormatter)
    Register(registry, Float64Predicate, Float64Formatter)
    Register(registry, SlicePredicate, SliceFormatter)
    Register(registry, MapPredicate, MapFormatter)
    return registry
}

// Advanced: Reflection-based type registration
func RegisterByType[T TypeConstraint](r *GenericRegistry, formatter TypedFormatter[T]) {
    var zero T
    targetType := reflect.TypeOf(zero)
    
    predicate := func(value interface{}) (T, bool) {
        if reflect.TypeOf(value) == targetType {
            return value.(T), true
        }
        return zero, false
    }
    
    Register(r, predicate, formatter)
}

// Custom type example
type EmailAddress string

func (e EmailAddress) String() string {
    return string(e)
}

// Register custom type formatter
func RegisterEmailFormatter(registry *GenericRegistry) {
    emailPredicate := func(value interface{}) (EmailAddress, bool) {
        if s, ok := value.(string); ok && strings.Contains(s, "@") {
            return EmailAddress(s), true
        }
        return "", false
    }
    
    emailFormatter := func(value EmailAddress, llml func(interface{}, *GenericRegistry) string, registry *GenericRegistry) string {
        return fmt.Sprintf("<email>%s</email>", strings.TrimSpace(string(value)))
    }
    
    Register(registry, emailPredicate, emailFormatter)
}
```

### Pros
- Type safety at compile time
- Generic constraints prevent runtime errors
- Modern Go practices
- Excellent IDE support
- Reflection capabilities for advanced use cases

### Cons
- Requires Go 1.18+
- More complex syntax
- Longer compile times
- May be overkill for simple use cases

## Comparison Matrix

| Feature | Interface-Based | Function-Based | Generics-Based |
|---------|----------------|----------------|----------------|
| Type Safety | High | Medium | Highest |
| Performance | Good | Best | Good |
| Extensibility | High | High | Highest |
| Simplicity | Medium | High | Low |
| Boilerplate | High | Low | Medium |
| Go Version | 1.0+ | 1.0+ | 1.18+ |
| Memory Usage | Medium | Low | Medium |
| IDE Support | Excellent | Good | Excellent |

## Integration with Current API

All three approaches can maintain backward compatibility with the existing `Sprintf` function:

```go
// Backward compatible wrapper
func Sprintf(data interface{}, opts ...Options) string {
    // Use default registry based on chosen approach
    registry := DefaultRegistry() // or DefaultFormatterRegistry() or DefaultGenericRegistry()
    
    formatted := registry.Format(data, formatWithRegistry)
    
    // Apply existing options (indent, prefix, etc.)
    if len(opts) > 0 {
        return applyOptions(formatted, opts[0])
    }
    
    return formatted
}
```

## Recommendation

For the LLML project, I recommend **Approach 2: Function Type-Based Registry** for the following reasons:

1. **Simplicity**: Matches the TypeScript implementation's simplicity
2. **Flexibility**: Easy to add custom formatters inline
3. **Performance**: Minimal overhead, direct function calls
4. **Maintainability**: Clear, readable code
5. **Compatibility**: Works with all Go versions

The function-based approach strikes the best balance between power and simplicity, making it ideal for a library that needs to be both extensible and easy to use.