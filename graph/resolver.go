package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/japiirainen/go-oluet-api/db"

	"github.com/japiirainen/go-oluet-api/graph/model"
)

type Resolver struct{}

var psql = db.Connect()

func (r *mutationResolver) NewJuomas(ctx context.Context) (string, error) {
	res, err := psql.InsertManyJuomas()
	if err != nil {
		log.Fatal(err)
		return "error", err
	}
	return res, nil
}

func (r *queryResolver) Juoma(ctx context.Context, id string) (*model.Juoma, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Juomat(ctx context.Context) ([]*model.Juoma, error) {
	res, err := psql.GetAllJuomas()
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
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
