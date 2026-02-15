package main

// スキーマ内の参照を使用する例
// homeとofficeのプロパティでaddressの定義を参照している
var schemaWithRef = `{
    "type": "object",
    "definitions": {
        "address": {
            "type": "object",
            "properties": {
                "street": { "type": "string" },
                "city": { "type": "string" }
            }
        }
    }
	"properties": {
        "home": { "$ref": "#/definitions/address" },
		"office": { "$ref": "#/definitions/address" }
    }
}`