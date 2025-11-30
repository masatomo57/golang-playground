package jsonschema

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Generate generates a JSON Schema from a Go struct value.
// Returns a map[string]any representing the JSON Schema in Draft 2020-12 format.
func Generate(v any) (map[string]any, error) {
	if v == nil {
		return nil, fmt.Errorf("cannot generate schema from nil value")
	}

	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct type, got %s", t.Kind())
	}

	schema := generateSchema(t)
	schema["$schema"] = "https://json-schema.org/draft/2020-12/schema"
	return schema, nil
}

// generateSchema recursively generates a JSON Schema from a reflect.Type.
func generateSchema(t reflect.Type) map[string]any {
	// Handle pointer types
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Handle basic types
	if schema := getTypeSchema(t); schema != nil {
		return schema
	}

	// Handle slices and arrays
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return map[string]any{
			"type":  "array",
			"items": generateSchema(t.Elem()),
		}
	}

	// Handle maps
	if t.Kind() == reflect.Map {
		// Only support map[string]T
		if t.Key().Kind() != reflect.String {
			return map[string]any{
				"type": "object",
			}
		}
		return map[string]any{
			"type":                 "object",
			"additionalProperties": generateSchema(t.Elem()),
		}
	}

	// Handle structs
	if t.Kind() == reflect.Struct {
		return generateStructSchema(t)
	}

	// Fallback for unknown types
	return map[string]any{
		"type": "object",
	}
}

// generateStructSchema generates a JSON Schema for a struct type.
func generateStructSchema(t reflect.Type) map[string]any {
	schema := map[string]any{
		"type":       "object",
		"properties": map[string]any{},
	}

	var required []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Parse json tag
		jsonName, omitempty, skip := parseJSONTag(field.Tag.Get("json"))
		if skip {
			continue
		}

		// Use json tag name if available, otherwise use field name
		if jsonName == "" {
			jsonName = field.Name
		}

		// Generate schema for the field type
		fieldSchema := generateSchema(field.Type)

		// Parse validation tags
		validationTag := field.Tag.Get("validate")
		if validationTag != "" {
			constraints := parseValidationTag(validationTag)
			for k, v := range constraints {
				fieldSchema[k] = v
			}
		}

		// Add to required if not omitempty and has required constraint
		if !omitempty {
			if validationTag != "" && strings.Contains(validationTag, "required") {
				required = append(required, jsonName)
			} else if !strings.Contains(validationTag, "omitempty") {
				// If no validation tag, check if it's a pointer (nullable)
				if field.Type.Kind() != reflect.Ptr {
					required = append(required, jsonName)
				}
			}
		}

		// Add property to schema
		props := schema["properties"].(map[string]any)
		props[jsonName] = fieldSchema
	}

	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

// getTypeSchema returns the JSON Schema for basic Go types.
func getTypeSchema(t reflect.Type) map[string]any {
	switch t.Kind() {
	case reflect.String:
		return map[string]any{"type": "string"}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return map[string]any{"type": "integer"}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return map[string]any{"type": "integer"}
	case reflect.Float32, reflect.Float64:
		return map[string]any{"type": "number"}
	case reflect.Bool:
		return map[string]any{"type": "boolean"}
	default:
		return nil
	}
}

// parseJSONTag parses a json tag string and returns the name, omitempty flag, and skip flag.
// Examples:
//   - `json:"name"` -> ("name", false, false)
//   - `json:"name,omitempty"` -> ("name", true, false)
//   - `json:"-"` -> ("", false, true)
func parseJSONTag(tag string) (name string, omitempty bool, skip bool) {
	if tag == "" {
		return "", false, false
	}

	if tag == "-" {
		return "", false, true
	}

	parts := strings.Split(tag, ",")
	name = parts[0]

	for _, part := range parts[1:] {
		if strings.TrimSpace(part) == "omitempty" {
			omitempty = true
		}
	}

	return name, omitempty, false
}

// parseValidationTag parses a validation tag and returns a map of JSON Schema constraints.
// Supported constraints:
//   - required: adds to required array (handled separately)
//   - minimum=N: sets minimum value for numbers
//   - maximum=N: sets maximum value for numbers
//   - minLength=N: sets minLength for strings
//   - maxLength=N: sets maxLength for strings
//   - pattern=REGEX: sets pattern for strings
func parseValidationTag(tag string) map[string]any {
	constraints := make(map[string]any)

	if tag == "" {
		return constraints
	}

	parts := strings.Split(tag, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// Handle key=value constraints
		if strings.Contains(part, "=") {
			kv := strings.SplitN(part, "=", 2)
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])

			switch key {
			case "minimum", "min":
				if num, err := strconv.ParseFloat(value, 64); err == nil {
					constraints["minimum"] = num
				}
			case "maximum", "max":
				if num, err := strconv.ParseFloat(value, 64); err == nil {
					constraints["maximum"] = num
				}
			case "minLength":
				if num, err := strconv.Atoi(value); err == nil {
					constraints["minLength"] = num
				}
			case "maxLength":
				if num, err := strconv.Atoi(value); err == nil {
					constraints["maxLength"] = num
				}
			case "pattern":
				constraints["pattern"] = value
			}
		}
		// Note: "required" is handled separately in generateStructSchema
	}

	return constraints
}
