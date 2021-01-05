package resolvers

import (
	"context"
	"fmt"

	"github.com/japiirainen/go-oluet-api/gql/model"
	log "github.com/sirupsen/logrus"
)

func (r *queryResolver) Hinnat(ctx context.Context) ([]model.Hinta, error) {
	res, err := r.DB.GetAllPrices()
	if err != nil {
		log.Error("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Hintahistoria(ctx context.Context, productID string) ([]model.Hinta, error) {
	panic(fmt.Errorf("not implemented"))
}
