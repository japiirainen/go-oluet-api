package resolvers

import (
	"context"

	"github.com/japiirainen/go-oluet-api/gql/model"
	log "github.com/sirupsen/logrus"
)

func (r *queryResolver) Prices(ctx context.Context) ([]model.Price, error) {
	log.Info("resolvers: Prices")
	res, err := r.DB.GetAllPrices()
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Pricehistory(ctx context.Context, productID string) ([]model.Price, error) {
	log.Info("resolvers: Price history")
	res, err := r.DB.GetPriceHistory(productID)
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}
