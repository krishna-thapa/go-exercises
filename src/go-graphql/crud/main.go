package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/graphql-go/graphql"
)

// class for one quote
type Quote struct {
	ID     int64     `json:"id"`
	Quote  string    `json:"quote"`
	Author string    `json:"author,omitempty"`
	Tags   []string  `json:"tags,omitempty"`
	Date   time.Time `json:"date"`
}

// https://www.sohamkamani.com/golang/2018-07-19-golang-omitempty/

var quotes = []Quote{
	{
		ID:     1,
		Quote:  "Hello world",
		Author: "John Will",
		Tags: []string{
			"love",
			"hate",
		},
		Date: time.Now(), //change to default date
	},
}

var quoteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Quote",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"quote": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"tags": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			/* Get (read) single quote by id
			   http://localhost:8080/quote?query={quote(id:1){quote,author,tags,date}}
			*/
			"quote": &graphql.Field{
				Type:        quoteType,
				Description: "Get quote by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// find quote
						for _, quote := range quotes {
							if int(quote.ID) == id {
								return quote, nil
							}
						}
					}
					return nil, nil
				},
			},

			/* Get (read) quote list
			   http://localhost:8080/quotet?query={list{id,quote,author,tags,date}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(quoteType),
				Description: "Get all the quotes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return quotes, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{

		// Create a new quote
		"create": &graphql.Field{
			Type:        quoteType,
			Description: "Create a new quote",
			Args: graphql.FieldConfigArgument{
				"quote": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"tags": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//koverflow.com/questions/44027826/convert-interface-to-string-in-golang
				tags := p.Args["tags"].([]interface{})
				fmt.Println("Type of tags: ",reflect.TypeOf(tags))
				stringTags := make([]string, len(tags))
				for i, v := range tags {
					stringTags[i] = fmt.Sprint(v)
				}
				//https://flaviocopes.com/go-random/
				rand.Seed(time.Now().UnixNano())
				quote := Quote{
					ID:     int64(rand.Intn(100000)), // generate random ID
					Quote:  p.Args["quote"].(string),
					Author: p.Args["author"].(string),
					Tags:   stringTags,
					Date:   time.Now(),
				}
				quotes = append(quotes, quote)
				return quote, nil
			},
		},
	},
})
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	fmt.Printf(query)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is up and running on port 8080")
	http.ListenAndServe(":8080", nil)
}
