package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	graphql "github.com/graph-gophers/graphql-go"
)

var (
	// Schema : GraphQL schema definition. This is an example schema
	Schema = `
		schema {
			query: Query
		}
		type Person{
			id: ID!
			firstName: String!
			lastName: String
		}
		type Query{
			person(id: ID!): Person
		}
	`
	peopleData = make(map[graphql.ID]*person)
	mainSchema *graphql.Schema
)

// Handler gets requests and returns response.
func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no query is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		log.Println("request.Body length zero.")
		return events.APIGatewayProxyResponse{}, errors.New("no query was provided in the HTTP body")
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not decode body", err)
	}

	response := mainSchema.Exec(context, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Print("Could not decode body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil

}

func init() {
	people := []*person{
		{
			ID:        "1000",
			FirstName: "Pedro",
			LastName:  "Marquez",
		},
		{
			ID:        "1001",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			ID:        "1002",
			FirstName: "Tom",
			LastName:  "Vincent",
		},
	}

	for _, p := range people {
		peopleData[p.ID] = p
	}
	mainSchema = graphql.MustParseSchema(Schema, &resolver{})
}

func main() {
	lambda.Start(Handler)
}
