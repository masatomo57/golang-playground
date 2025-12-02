package jsonschema

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Generate はGoの構造体からJSON Schemaを生成する。
// map[string]any形式のJSON Schemaを返す。
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

	return generateSchema(t), nil
}

// generateSchema はreflect.TypeからJSON Schemaを再帰的に生成する。
func generateSchema(t reflect.Type) map[string]any {
	// ポインタ型の処理
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 基本型の処理
	if schema := getTypeSchema(t); schema != nil {
		return schema
	}

	// スライスと配列の処理
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return map[string]any{
			"type":  "array",
			"items": generateSchema(t.Elem()),
		}
	}

	// マップの処理
	if t.Kind() == reflect.Map {
		// map[string]T のみサポート
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

	// 構造体の処理
	if t.Kind() == reflect.Struct {
		return generateStructSchema(t)
	}

	// 未知の型のフォールバック
	return map[string]any{
		"type": "object",
	}
}

// generateStructSchema は構造体型のJSON Schemaを生成する。
func generateStructSchema(t reflect.Type) map[string]any {
	schema := map[string]any{
		"type":       "object",
		"properties": map[string]any{},
	}

	var required []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// 非公開フィールドをスキップ
		if !field.IsExported() {
			continue
		}

		// jsonタグを解析
		jsonName, omitempty, skip := parseJSONTag(field.Tag.Get("json"))
		if skip {
			continue
		}

		// jsonタグがあればその名前を使用、なければフィールド名を使用
		if jsonName == "" {
			jsonName = field.Name
		}

		// フィールドの型からスキーマを生成
		fieldSchema := generateSchema(field.Type)

		// validateタグを解析
		validationTag := field.Tag.Get("validate")
		if validationTag != "" {
			constraints := parseValidationTag(validationTag)
			for k, v := range constraints {
				fieldSchema[k] = v
			}
		}

		// omitemptyでなく、required制約がある場合はrequiredに追加
		if !omitempty {
			if validationTag != "" && strings.Contains(validationTag, "required") {
				required = append(required, jsonName)
			} else if !strings.Contains(validationTag, "omitempty") {
				// validateタグがない場合、ポインタ型かどうかをチェック
				if field.Type.Kind() != reflect.Ptr {
					required = append(required, jsonName)
				}
			}
		}

		// プロパティをスキーマに追加
		props := schema["properties"].(map[string]any)
		props[jsonName] = fieldSchema
	}

	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

// getTypeSchema は基本的なGo型のJSON Schemaを返す。
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

// parseJSONTag はjsonタグ文字列を解析し、名前、omitemptyフラグ、スキップフラグを返す。
// 例:
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

// parseValidationTag はvalidateタグを解析し、JSON Schema制約のマップを返す。
// サポートする制約:
//   - required: required配列に追加（別途処理）
//   - minimum=N: 数値の最小値を設定
//   - maximum=N: 数値の最大値を設定
//   - minLength=N: 文字列の最小長を設定
//   - maxLength=N: 文字列の最大長を設定
//   - pattern=REGEX: 文字列のパターンを設定
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

		// key=value形式の制約を処理
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
		// 注: "required"はgenerateStructSchemaで別途処理
	}

	return constraints
}
