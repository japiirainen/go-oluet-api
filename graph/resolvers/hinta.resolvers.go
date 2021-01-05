package resolvers

import (
	"context"
	"fmt"
	"log"

	"github.com/japiirainen/go-oluet-api/graph/model"
)

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
