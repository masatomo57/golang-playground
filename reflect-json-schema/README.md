# JSON Schema Generator

Goの構造体からJSON Schema Draft 2020-12形式の`map[string]any`を生成するパッケージです。`reflect`パッケージを使用して、構造体の型情報とタグから自動的にJSON Schemaを生成します。

## 機能

### サポートする型

- **基本型**: `string`, `int/int8/int16/int32/int64`, `uint/uint8/uint16/uint32/uint64`, `float32/float64`, `bool`
- **配列/スライス**: `[]T` は `array`型として扱われます
- **マップ**: `map[string]T` は `object`型として扱われ、`additionalProperties`で値の型を指定します
- **構造体**: `object`型として扱われ、ネストした構造体も再帰的に処理されます
- **ポインタ**: ポインタ型は自動的に参照先の型に解決されます

### タグサポート

#### JSONタグ

- **カスタムフィールド名**: `json:"custom_name"` でJSON Schemaのプロパティ名をカスタマイズ
- **omitempty**: `json:"field,omitempty"` で必須フィールドから除外
- **除外**: `json:"-"` でフィールドをスキーマから除外

#### Validationタグ

`validate`タグを使用してJSON Schemaの制約を指定できます：

- **required**: フィールドを必須にします
- **minimum/min**: 数値の最小値を指定（例: `minimum=0`）
- **maximum/max**: 数値の最大値を指定（例: `maximum=100`）
- **minLength**: 文字列の最小長を指定（例: `minLength=1`）
- **maxLength**: 文字列の最大長を指定（例: `maxLength=100`）
- **pattern**: 正規表現パターンを指定（例: `pattern=^[a-z]+@[a-z]+\\.[a-z]+$`）

## インストール

```bash
go get github.com/masatomo57/golang-oreore-comparable/reflect-json-schema
```

## 使用方法

### 基本的な使用例

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/masatomo57/golang-oreore-comparable/reflect-json-schema"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    schema, err := jsonschema.Generate(User{})
    if err != nil {
        panic(err)
    }

    jsonBytes, _ := json.MarshalIndent(schema, "", "  ")
    fmt.Println(string(jsonBytes))
}
```

出力:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "id": {
      "type": "integer"
    },
    "name": {
      "type": "string"
    },
    "email": {
      "type": "string"
    }
  },
  "required": ["id", "name", "email"]
}
```

### Validation制約の使用例

```go
type ValidatedUser struct {
    ID       int    `json:"id" validate:"required,minimum=1,maximum=1000"`
    Name     string `json:"name" validate:"required,minLength=1,maxLength=100"`
    Email    string `json:"email" validate:"required,pattern=^[a-z]+@[a-z]+\\.[a-z]+$"`
    Age      int    `json:"age,omitempty" validate:"minimum=0,maximum=150"`
}
```

生成されるスキーマ:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "id": {
      "type": "integer",
      "minimum": 1,
      "maximum": 1000
    },
    "name": {
      "type": "string",
      "minLength": 1,
      "maxLength": 100
    },
    "email": {
      "type": "string",
      "pattern": "^[a-z]+@[a-z]+\\.[a-z]+$"
    },
    "age": {
      "type": "integer",
      "minimum": 0,
      "maximum": 150
    }
  },
  "required": ["id", "name", "email"]
}
```

### ネストした構造体の例

```go
type Address struct {
    Street string `json:"street" validate:"required"`
    City   string `json:"city" validate:"required"`
}

type User struct {
    ID      int     `json:"id" validate:"required"`
    Name    string  `json:"name" validate:"required"`
    Address Address `json:"address"`
}
```

### 配列とマップの例

```go
type User struct {
    ID       int               `json:"id"`
    Tags     []string          `json:"tags"`
    Metadata map[string]string `json:"metadata"`
}
```

生成されるスキーマ:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "properties": {
    "id": {
      "type": "integer"
    },
    "tags": {
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "metadata": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    }
  }
}
```

## API

### `Generate(v any) (map[string]any, error)`

構造体の値からJSON Schemaを生成します。

**パラメータ:**
- `v`: 構造体の値（ポインタも可）

**戻り値:**
- `map[string]any`: JSON Schema Draft 2020-12形式のスキーマ
- `error`: エラー（nil値や非構造体型の場合）

**例:**

```go
schema, err := jsonschema.Generate(User{})
// または
schema, err := jsonschema.Generate(&User{})
```

## 制限事項

- 非公開フィールド（小文字で始まるフィールド）はスキーマに含まれません
- マップは`map[string]T`形式のみサポートしています（キーがstring以外のマップは`object`型として扱われます）
- カスタム型（`type MyString string`など）は基本型として扱われます
- `validate`タグの`required`は、`omitempty`がない場合にのみ`required`配列に追加されます

## テスト

```bash
cd reflect-json-schema
go test -v
```

## ライセンス

このプロジェクトのライセンスに従います。

