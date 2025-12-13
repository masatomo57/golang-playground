package main  
  
import (  
    "encoding/json"  
    "fmt"  
    "log"  
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
}