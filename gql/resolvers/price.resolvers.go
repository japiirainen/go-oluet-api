package resolvers

import (
	"context"
	"fmt"

	"github.com/japiirainen/go-oluet-api/gql/model"
	log "github.com/sirupsen/logrus"
)

func (r *queryResolver) Prices(ctx context.Context) ([]model.Price, error) {
	log.Info("resolvers: Prices")
	res, err := r.DB.GetAllPrices()
	if err != nil {
		log.Error("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Pricehistory(ctx context.Context, productID string) ([]model.Price, error) {
	panic(fmt.Errorf("not implemented"))
}
