package main

import graphql "github.com/graph-gophers/graphql-go"

// Resolver : Struct with all the resolver functions
type resolver struct{}

// Person : Resolver function for the "Person" query
func (r *resolver) Person(args struct{ ID graphql.ID }) *personResolver {
	if p := peopleData[args.ID]; p != nil {
		return &personResolver{p}
	}
	return nil
}
