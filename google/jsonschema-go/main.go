package main  
  
import (
    "encoding/json"
    "fmt"
    "log"
    "net/url"
    "os"
    "path/filepath"

    "github.com/google/jsonschema-go/jsonschema"
)  
  
type User struct {
    Username string   `json:"username" jsonschema:"unique identifier"`
    Email    string   `json:"email"`
    Age      int      `json:"age,omitempty"`
    Tags     []string `json:"tags,omitempty"`
}

func main() {  
    // Schemaを推論  
    schema, err := jsonschema.For[User](nil)  
    if err != nil {  
        log.Fatal(err)  
    }  
      
    // map[string]anyに変換  
    jsonData, err := json.Marshal(schema)  
    if err != nil {  
        log.Fatal(err)  
    }  
      
    var schemaMap map[string]any  
    err = json.Unmarshal(jsonData, &schemaMap)  
    if err != nil {  
        log.Fatal(err)  
    }  
      
    // 結果を表示  
    prettyJSON, _ := json.MarshalIndent(schemaMap, "", "  ")  
    fmt.Printf("%s\n", prettyJSON)  

    var schema1 = jsonschema.Schema{
        Type: "string",
    }
    jsonData, err = json.Marshal(schema1)
    if err != nil {
        log.Fatal(err)
    }
    println(string(jsonData))
    // {"type":"string"}

    var schema2 = jsonschema.Schema{
        Types: []string{"string", "integer"},
    }
    jsonData, err = json.Marshal(schema2)
    if err != nil {
        log.Fatal(err)
    }
    println(string(jsonData))
    // {"type":["string","integer"]}

    // 外部参照($ref)を持つスキーマ
    // メインスキーマ: addressプロパティがローカルファイルを参照
    mainSchemaJSON := `{
        "type": "object",
        "properties": {
            "name": { "type": "string" },
            "address": { "$ref": "address.json" }
        },
        "required": ["name", "address"]
    }`

    var mainSchema jsonschema.Schema
    if err := json.Unmarshal([]byte(mainSchemaJSON), &mainSchema); err != nil {
        log.Fatal(err)
    }

    // ローカルJSONファイルからスキーマを読み込むLoader
    loader := func(uri *url.URL) (*jsonschema.Schema, error) {
        filename := filepath.Base(uri.Path)
        data, err := os.ReadFile(filename)
        if err != nil {
            return nil, fmt.Errorf("failed to read schema file %s: %w", filename, err)
        }
        var s jsonschema.Schema
        if err := json.Unmarshal(data, &s); err != nil {
            return nil, err
        }
        return &s, nil
    }

    // Loaderを使って外部参照を解決
    resolved, _ := mainSchema.Resolve(&jsonschema.ResolveOptions{
        Loader: loader,
    })

    // map[string]anyに対してバリデーションを行う
    valid := map[string]any{
        "name": "Alice",
        "address": map[string]any{
            "street": "123 Main St",
            "city":   "Tokyo",
        },
    }
    if err := resolved.Validate(valid); err != nil {
        fmt.Println("valid data failed:", err)
    }

    // バリデーション: 不正なデータ（addressにcityがない）
    invalid := map[string]any{
        "name": "Bob",
        "address": map[string]any{
            "street": "456 Oak Ave",
        },
    }
    if err := resolved.Validate(invalid); err != nil {
        fmt.Println("invalid data:", err)
    } else {
        fmt.Println("invalid data: OK (unexpected)")
    }

    type Person struct {
        Name string `json:"name" jsonschema:"person's full name"`
        Age  int    `json:"age"`
    }
    // jsonschema.ForでPersonのスキーマを生成
    // jsonschemaタグはdescriptionに設定される
    personSchema, _ := jsonschema.For[Person](nil)
    personJSON, _ := json.MarshalIndent(personSchema, "", "  ")
    fmt.Printf("Person schema:\n%s\n", personJSON)
    // Person schema:
    // {
    //     "type": "object",
    //     "properties": {
    //         "name": {
    //              "type": "string",
    //              "description": "person's full name"
    //         },
    //         "age": {
    //              "type": "integer",
    //         }
    //     },
    //     "required": [
    //         "name",
    //         "age"
    //     ],
    //     "additionalProperties": false
    // }
}