package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/japiirainen/go-oluet-api/db"
)

type Resolver struct {
	DB *db.Db
}
