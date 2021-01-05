package resolvers

import (
	"context"
	"log"

	"github.com/japiirainen/go-oluet-api/graph/model"
)

func (r *mutationResolver) NewJuomas(ctx context.Context) (string, error) {
	res, err := r.DB.InsertManyJuomas()
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (r *queryResolver) Juoma(ctx context.Context, productID string) (*model.Juoma, error) {
	res, err := r.DB.GetJuomaByProdID(productID)
	if err != nil {
		log.Fatal(err)
	}
	return &res, nil
}

func (r *queryResolver) JuomaSearch(ctx context.Context, term string) ([]model.Juoma, error) {
	res, err := r.DB.SearchForJuoma(term)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (r *queryResolver) Juomat(ctx context.Context) ([]model.Juoma, error) {
	res, err := r.DB.GetAllJuomas()
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
