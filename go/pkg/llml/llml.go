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
	Strict bool
}

// Sprintf converts data structures to XML-like markup using a recursive approach
// Supports various call patterns:
//   - Sprintf() -> ""
//   - Sprintf(nil) -> "nil"
//   - Sprintf([]interface{}{}) -> ""
//   - Sprintf(map[string]interface{}{}) -> ""
//   - Sprintf(map[string]interface{}{"key": "value"}) -> "<key>value</key>"
func Sprintf(data interface{}, opts ...Options) string {
	options := Options{Indent: "", Prefix: "", Strict: false}
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
	for _, key := range keys {
		value := m[key]

		// Recursively format this key-value pair
		formatted := formatKeyValue(key, value, opts)
		
		// Skip empty results (like empty arrays)
		if formatted != "" {
			if len(parts) > 0 {
				parts = append(parts, "\n")
			}
			parts = append(parts, formatted)
		}
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
		Strict: opts.Strict,
	}
	
	// In strict mode, use parent key as prefix. In non-strict mode, don't use prefix
	if opts.Strict {
		nestedOpts.Prefix = fullKey
	} else {
		nestedOpts.Prefix = ""
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
	wrapperTag := kebabKey

	if len(items) == 0 {
		return ""
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
				Strict: opts.Strict,
			}
			// In strict mode, use array item tag as prefix. In non-strict mode, don't use prefix
			if opts.Strict {
				nestedOpts.Prefix = itemTag
			} else {
				nestedOpts.Prefix = ""
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
				Strict: opts.Strict,
			}
			// In strict mode, use array item tag as prefix. In non-strict mode, don't use prefix
			if opts.Strict {
				nestedOpts.Prefix = itemTag
			} else {
				nestedOpts.Prefix = ""
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

// formatSlice handles direct slice calls with numeric tags
func formatSlice(items []any, opts Options) string {
	if len(items) == 0 {
		return ""
	}

	var parts []string
	for i, item := range items {
		itemTag := strconv.Itoa(i + 1)
		if opts.Prefix != "" {
			itemTag = opts.Prefix + "-" + itemTag
		}

		// Handle dictionary items in direct arrays
		if dict, ok := item.(map[string]any); ok {
			var content string
			if len(dict) == 0 {
				content = ""
			} else {
				nestedOpts := Options{
					Indent: opts.Indent + "  ",
					Prefix: itemTag,
					Strict: opts.Strict,
				}
				content = Sprintf(dict, nestedOpts)
			}
			
			if content == "" {
				parts = append(parts, fmt.Sprintf("%s<%s></%s>", opts.Indent, itemTag, itemTag))
			} else {
				// Force multiline format for objects in direct arrays
				parts = append(parts, fmt.Sprintf("%s<%s>\n%s\n%s</%s>",
					opts.Indent, itemTag, content, opts.Indent, itemTag))
			}
		} else if dict, ok := item.(map[string]interface{}); ok {
			anyDict := make(map[string]any)
			for k, v := range dict {
				anyDict[k] = v
			}
			var content string
			if len(anyDict) == 0 {
				content = ""
			} else {
				nestedOpts := Options{
					Indent: opts.Indent + "  ",
					Prefix: itemTag,
					Strict: opts.Strict,
				}
				content = Sprintf(anyDict, nestedOpts)
			}
			
			if content == "" {
				parts = append(parts, fmt.Sprintf("%s<%s></%s>", opts.Indent, itemTag, itemTag))
			} else {
				// Force multiline format for objects in direct arrays
				parts = append(parts, fmt.Sprintf("%s<%s>\n%s\n%s</%s>",
					opts.Indent, itemTag, content, opts.Indent, itemTag))
			}
		} else if slice, ok := item.([]any); ok {
			// Handle array items in direct arrays - skip empty arrays
			if len(slice) > 0 {
				// For non-empty arrays, format recursively
				nestedResult := formatSlice(slice, Options{
					Indent: opts.Indent + "  ",
					Prefix: "",
					Strict: opts.Strict,
				})
				if nestedResult != "" {
					parts = append(parts, fmt.Sprintf("%s<%s>\n%s\n%s</%s>",
						opts.Indent, itemTag, nestedResult, opts.Indent, itemTag))
				}
			}
			// Empty arrays are skipped implicitly
		} else if slice, ok := item.([]interface{}); ok {
			// Handle array items in direct arrays - skip empty arrays
			anySlice := make([]any, len(slice))
			copy(anySlice, slice)
			if len(anySlice) > 0 {
				// For non-empty arrays, format recursively
				nestedResult := formatSlice(anySlice, Options{
					Indent: opts.Indent + "  ",
					Prefix: "",
					Strict: opts.Strict,
				})
				if nestedResult != "" {
					parts = append(parts, fmt.Sprintf("%s<%s>\n%s\n%s</%s>",
						opts.Indent, itemTag, nestedResult, opts.Indent, itemTag))
				}
			}
			// Empty arrays are skipped implicitly
		} else {
			// Handle simple items
			formatted := Sprintf(item)
			if formatted != "" {
				parts = append(parts, fmt.Sprintf("%s<%s>%s</%s>",
					opts.Indent, itemTag, formatted, itemTag))
			}
			// Empty items are skipped implicitly
		}
	}

	// If all items were skipped, return empty string
	if len(parts) == 0 {
		return ""
	}

	return strings.Join(parts, "\n")
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
