package resolvers

import (
	"context"

	"github.com/japiirainen/go-oluet-api/graph/model"
	log "github.com/sirupsen/logrus"
)

func (r *mutationResolver) Newjuomas(ctx context.Context) (string, error) {
	log.Info("resolvers: NewJuoma")
	res, err := r.DB.InsertManyJuomas()
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Juoma(ctx context.Context, productID string) (*model.Juoma, error) {
	log.Info("resolvers: Juoma")
	res, err := r.DB.GetJuomaByProdID(productID)
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return &res, nil
}

func (r *queryResolver) Juomasearch(ctx context.Context, term string) ([]model.Juoma, error) {
	log.Info("resolvers: JuomaSearch")
	res, err := r.DB.SearchForJuoma(term)
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Juomat(ctx context.Context) ([]model.Juoma, error) {
	log.Info("resolvers: Juomat")
	res, err := r.DB.GetAllJuomas()
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Oluet(ctx context.Context) ([]model.Juoma, error) {
	log.Info("resolvers: Oluet")
	res, err := r.DB.GetAllBeers()
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Olutsearch(ctx context.Context, term string) ([]model.Juoma, error) {
	log.Info("resolvers: OluetSearch")
	res, err := r.DB.SearchForBeer(term)
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}
