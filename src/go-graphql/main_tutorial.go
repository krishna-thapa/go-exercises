/*
Source: https://tutorialedge.net/golang/go-graphql-beginners-tutorial/
We’ll be creating a GraphQL server that returns a series of in-memory tutorials as well as their Author, and any comments made on those particular tutorials.
*/
package main

import (
	"fmt"
	"encoding/json"
	"log"
	"github.com/graphql-go/graphql"
)

//define some struct’s that will represent a Tutorial, an Author, and a Comment:

type Tutorial struct {
	ID       int
	Title    string
	Author   Author
	Comments []Comment
}

type Author struct {
	Name      string
	Tutorials []int
}

type Comment struct {
	Body string
}

//  populate() function which will return an array of type Tutorial
func populate() []Tutorial {
	author := &Author{Name: "John Will", Tutorials: []int{1, 2}}
	tutorial := Tutorial{
		ID:     1,
		Title:  "Go GraphQL tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "First comment"},
		},
	}

	tutorial2 := Tutorial{
		ID:     2,
		Title:  "Go GraphQL tutorial - part2",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "First comment - part 2"},
		},
	}

	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	tutorials = append(tutorials, tutorial2)

	return tutorials
}

/*
Create a new object in GraphQL using graphql.NewObject(), define 3 different Types using GraphQL’s strict typing, these will match up with the 3 structs we’ve already defined.
*/

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)

func main() {

	tutorials := populate()

	// Schema
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type:        tutorialType,
			Description: "Get Tutorial by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// take in the ID argument
				id, ok := p.Args["id"].(int)
				if ok {
					// Parse our tutorial array for the matching id
					for _, tutorial := range tutorials {
						if int(tutorial.ID) == id {
							// return our tutorial
							return tutorial, nil
						}
					}
				}
				return nil, nil
			},
		},

		"list": &graphql.Field{
			Type:        graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return tutorials, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	/*
	So within our query we have a special root object. Within this we then say that we want the list field on that object. On the list returned by list, we want to see the id, title, comments and the author.*/
	// query := `
	// 	{
	// 		list {
	// 			id
	// 			title
	// 			comments{
	// 				body
	// 			}
	// 			author {
	// 				Name
	// 				Tutorials
	// 			}
	// 		}
	// 	}
	// `

	// query against our tutorial schema:
	query1 := `
    {
        tutorial(id:2) {
            title
            author {
                Name
                Tutorials
			}
			comments {
				body
			}
        }
    }
`

	params := graphql.Params{Schema: schema, RequestString: query1}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("faliked to execute graphql operations, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)

}
