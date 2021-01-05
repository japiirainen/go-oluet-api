package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/japiirainen/go-oluet-api/db"
)

//Resolver is the base resolver
type Resolver struct {
	DB *db.Db
}
