package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/japiirainen/go-oluet-api/db"
	"github.com/japiirainen/go-oluet-api/gql/generated"
)

//Resolver is the base resolver
type Resolver struct {
	DB *db.Db
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
