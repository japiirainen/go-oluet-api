package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/japiirainen/go-oluet-api/graph/generated"
	"github.com/japiirainen/go-oluet-api/graph/model"
)

func (r *mutationResolver) UusiJuoma(ctx context.Context, input model.UusiJuoma) (*model.Juoma, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Juoma(ctx context.Context, id string) (*model.Juoma, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Juomat(ctx context.Context) ([]*model.Juoma, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hinta(ctx context.Context, id string) (*model.Hinta, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hinnat(ctx context.Context) ([]*model.Hinta, error) {
	var hinnat []*model.Hinta
	dummyHinta := model.Hinta{
		ID:        1,
		Date:      time.Now(),
		ProductID: "1337",
		Hinta:     10.5,
	}
	hinnat = append(hinnat, &dummyHinta)
	return hinnat, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
