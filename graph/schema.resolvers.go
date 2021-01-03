package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/japiirainen/go-oluet-api/graph/generated"
	"github.com/japiirainen/go-oluet-api/graph/model"
)

func (r *mutationResolver) NewJuomas(ctx context.Context) (string, error) {
	res, err := r.DB.InsertManyJuomas()
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (r *queryResolver) Juoma(ctx context.Context, id string) (*model.Juoma, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Juomat(ctx context.Context) ([]model.Juoma, error) {
	res, err := r.DB.GetAllJuomas()
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (r *queryResolver) Hinta(ctx context.Context, id string) (*model.Hinta, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hinnat(ctx context.Context) ([]model.Hinta, error) {
	res, err := r.DB.GetAllPrices()
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
