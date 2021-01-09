package resolvers

import (
	"context"

	"github.com/japiirainen/go-oluet-api/gql/model"
	log "github.com/sirupsen/logrus"
)

func (r *queryResolver) Drink(ctx context.Context, productID string) (*model.Drink, error) {
	log.Info("resolvers: Drink")
	res, err := r.DB.GetDrinkByProdID(productID)
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return &res, nil
}

func (r *queryResolver) Drinksearch(ctx context.Context, term string) ([]model.Drink, error) {
	log.Info("resolvers: DrinkSearch")
	res, err := r.DB.SearchForDrink(term)
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Drinks(ctx context.Context) ([]model.Drink, error) {
	log.Info("resolvers: Drinks")
	res, err := r.DB.GetAllDrinks()
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Beers(ctx context.Context) ([]model.Drink, error) {
	log.Info("resolvers: Beers")
	res, err := r.DB.GetAllBeers()
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}

func (r *queryResolver) Beersearch(ctx context.Context, term string) ([]model.Drink, error) {
	log.Info("resolvers: BeerSearch")
	res, err := r.DB.SearchForBeer(term)
	if err != nil {
		log.Errorf("resolvers: %s", err)
	}
	return res, nil
}
