package main

import graphql "github.com/graph-gophers/graphql-go"

type personResolver struct {
	p *person
}

func (r *personResolver) ID() graphql.ID {
	return r.p.ID
}

func (r *personResolver) FirstName() string {
	return r.p.FirstName
}

func (r *personResolver) LastName() *string {
	return &r.p.LastName
}
