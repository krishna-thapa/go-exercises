// create a really simple GraphQL server that features a really simple resolver

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/graphql-go/graphql"
)

func main() {
	/* Schema
	When we make queries against our GraphQL API, we essentially define what fields on objects we want returned to us, so we have to define these fields within our Schema.
	*/
	fields := graphql.Fields {
		"hello": &graphql.Field {
			Type: graphql.String,
			// resolver function that is triggered whenever this particular field is requested
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world!", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields} 
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query: define a query that requests the field hello.
	query := `
	{
		hello
	}
	`
	// create a params struct which contains a reference to our defined Schema as well as our RequestString request.
	params := graphql.Params{Schema: schema, RequestString: query}

	//execute the request and the results of the request are populated into r and do some error handling 
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}

	//Marshal the response into JSON and print it out to our console.
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)  // {“data”:{“hello”:”world”}}
}


/*
GraphQL is a way to query data. A flexible API protocol that lets you query just the data that you need and get that in return. Any storage can be behind this data, GraphQL is not responsible for storage.
*/