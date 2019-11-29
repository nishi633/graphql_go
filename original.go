package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"

	"./query"
)

func main() {
	// ***スキーマ***
	// graphql.ObjectConfigは？？ https://github.com/graphql-go/graphql/blob/2e2b648ecbe42b217c0d7c2bf45fa607bf3f1201/definition.go#L371-L377
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: query.SetUserQuery()}
	// graphql.SchemaConfig https://github.com/graphql-go/graphql/blob/2b0b7340d2285b861482047fbf5de0488e021bea/schema.go#L7-L14
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	// graphql.NewSchemaは新しくスキーマを生成するもの https://github.com/graphql-go/graphql/blob/2b0b7340d2285b861482047fbf5de0488e021bea/schema.go#L50
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// ***リクエスト***
	query := `
		{
      User(ID: 3) {Name, Email}, Product
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
