package main

import graphql "github.com/graph-gophers/graphql-go"

type person struct {
	ID        graphql.ID
	FirstName string
	LastName  string
}
