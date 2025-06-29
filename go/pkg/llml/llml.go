package llml

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Options holds configuration for LLML formatting
type Options struct {
	Indent string
	Prefix string
}

// Sprintf converts data structures to XML-like markup using a recursive approach
// Supports various call patterns:
//   - Sprintf() -> ""
//   - Sprintf(nil) -> "nil"
//   - Sprintf([]interface{}{}) -> ""
//   - Sprintf(map[string]interface{}{}) -> ""
//   - Sprintf(map[string]interface{}{"key": "value"}) -> "<key>value</key>"
func Sprintf(data interface{}, opts ...Options) string {
	options := Options{Indent: "", Prefix: ""}
	if len(opts) > 0 {
		options = opts[0]
	}

	// Handle nil
	if data == nil {
		return "nil"
	}

	// Handle maps (the main case)
	if m, ok := data.(map[string]any); ok {
		return formatMap(m, options)
	}
	if m, ok := data.(map[string]interface{}); ok {
		// Convert to map[string]any
		anyMap := make(map[string]any)
		for k, v := range m {
			anyMap[k] = v
		}
		return formatMap(anyMap, options)
	}

	// Handle slices
	if s, ok := data.([]any); ok {
		return formatSlice(s, options)
	}
	if s, ok := data.([]interface{}); ok {
		// Convert to []any
		anySlice := make([]any, len(s))
		copy(anySlice, s)
		return formatSlice(anySlice, options)
	}

	// Handle primitive types
	switch v := data.(type) {
	case string:
		return formatString(v, options.Indent)
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	default:
		return fmt.Sprintf("%v", data)
	}
}

// formatMap handles the core recursive case: formatting key-value pairs
func formatMap(m map[string]any, opts Options) string {
	if len(m) == 0 {
		return ""
	}

	// Get sorted keys for consistent output
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for i, key := range keys {
		value := m[key]

		if i > 0 {
			parts = append(parts, "\n")
		}

		// Recursively format this key-value pair
		formatted := formatKeyValue(key, value, opts)
		parts = append(parts, formatted)
	}

	return strings.Join(parts, "")
}

// formatKeyValue handles a single key-value pair (the recursive unit)
func formatKeyValue(key string, value any, opts Options) string {
	fullKey := key
	if opts.Prefix != "" {
		fullKey = opts.Prefix + "-" + key
	}
	kebabKey := toKebabCase(fullKey)

	// Handle lists with wrapper tags
	if slice, ok := value.([]any); ok {
		return formatList(slice, key, opts)
	}
	if slice, ok := value.([]interface{}); ok {
		anySlice := make([]any, len(slice))
		copy(anySlice, slice)
		return formatList(anySlice, key, opts)
	}

	// Handle nested maps
	if nested, ok := value.(map[string]any); ok {
		return formatNestedMap(nested, kebabKey, fullKey, opts)
	}
	if nested, ok := value.(map[string]interface{}); ok {
		anyMap := make(map[string]any)
		for k, v := range nested {
			anyMap[k] = v
		}
		return formatNestedMap(anyMap, kebabKey, fullKey, opts)
	}

	// Handle primitive values
	formatted := Sprintf(value)
	if strings.Contains(formatted, "\n") {
		return fmt.Sprintf("%s<%s>\n%s\n%s</%s>",
			opts.Indent, kebabKey, formatted, opts.Indent, kebabKey)
	}
	return fmt.Sprintf("%s<%s>%s</%s>",
		opts.Indent, kebabKey, formatted, kebabKey)
}

// formatNestedMap handles nested map formatting
func formatNestedMap(nested map[string]any, kebabKey, fullKey string, opts Options) string {
	nestedOpts := Options{
		Indent: opts.Indent + "  ",
		Prefix: fullKey,
	}
	content := Sprintf(nested, nestedOpts)

	if strings.Contains(content, "\n") {
		return fmt.Sprintf("%s<%s>\n%s\n%s</%s>",
			opts.Indent, kebabKey, content, opts.Indent, kebabKey)
	}
	return fmt.Sprintf("%s<%s>%s</%s>",
		opts.Indent, kebabKey, content, kebabKey)
}

// formatList handles list formatting with wrapper tags
func formatList(items []any, key string, opts Options) string {
	fullKey := key
	if opts.Prefix != "" {
		fullKey = opts.Prefix + "-" + key
	}
	kebabKey := toKebabCase(fullKey)
	wrapperTag := kebabKey + "-list"

	if len(items) == 0 {
		return fmt.Sprintf("%s<%s></%s>", opts.Indent, wrapperTag, wrapperTag)
	}

	var parts []string
	parts = append(parts, fmt.Sprintf("%s<%s>\n", opts.Indent, wrapperTag))

	innerIndent := opts.Indent + "  "
	for i, item := range items {
		itemTag := fmt.Sprintf("%s-%d", kebabKey, i+1)

		// Handle dictionary items
		if dict, ok := item.(map[string]any); ok {
			parts = append(parts, fmt.Sprintf("%s<%s>\n", innerIndent, itemTag))
			nestedOpts := Options{
				Indent: innerIndent + "  ",
				Prefix: itemTag,
			}
			content := Sprintf(dict, nestedOpts)
			parts = append(parts, content)
			parts = append(parts, fmt.Sprintf("\n%s</%s>\n", innerIndent, itemTag))
		} else if dict, ok := item.(map[string]interface{}); ok {
			anyDict := make(map[string]any)
			for k, v := range dict {
				anyDict[k] = v
			}
			parts = append(parts, fmt.Sprintf("%s<%s>\n", innerIndent, itemTag))
			nestedOpts := Options{
				Indent: innerIndent + "  ",
				Prefix: itemTag,
			}
			content := Sprintf(anyDict, nestedOpts)
			parts = append(parts, content)
			parts = append(parts, fmt.Sprintf("\n%s</%s>\n", innerIndent, itemTag))
		} else {
			// Handle simple items
			formatted := Sprintf(item)
			parts = append(parts, fmt.Sprintf("%s<%s>%s</%s>\n",
				innerIndent, itemTag, formatted, itemTag))
		}
	}

	parts = append(parts, fmt.Sprintf("%s</%s>", opts.Indent, wrapperTag))
	return strings.Join(parts, "")
}

// formatSlice handles direct slice calls (returns empty as per spec)
func formatSlice(_items []any, _opts Options) string {
	// Direct slice calls without context should return empty
	return ""
}

// formatString handles string formatting with multiline support
func formatString(s string, _indent string) string {
	s = strings.TrimSpace(s)
	if strings.Contains(s, "\n") {
		lines := strings.Split(s, "\n")
		var formatted []string
		for _, line := range lines {
			formatted = append(formatted, "  "+strings.TrimSpace(line))
		}
		return strings.Join(formatted, "\n")
	}
	return s
}

// toKebabCase converts text to kebab-case format
// Handles camelCase, PascalCase, snake_case, spaces, and acronyms correctly
func toKebabCase(text string) string {
	if text == "" {
		return text
	}

	// Replace spaces and underscores with hyphens
	result := regexp.MustCompile(`[\s_]+`).ReplaceAllString(text, "-")

	// Handle sequences of uppercase letters followed by lowercase (acronyms)
	// e.g., "XMLHttpRequest" -> "XML-Http-Request"
	result = regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`).ReplaceAllString(result, "$1-$2")

	// Handle lowercase followed by uppercase  
	// e.g., "camelCase" -> "camel-Case"
	result = regexp.MustCompile(`([a-z\d])([A-Z])`).ReplaceAllString(result, "$1-$2")

	// Convert to lowercase
	return strings.ToLower(result)
}

// LLML is a backwards compatibility alias for Sprintf
// Deprecated: Use Sprintf instead
func LLML(data interface{}, opts ...Options) string {
	return Sprintf(data, opts...)
}
