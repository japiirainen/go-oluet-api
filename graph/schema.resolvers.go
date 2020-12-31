package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/japiirainen/go-oluet-api/db"
	"github.com/japiirainen/go-oluet-api/exel"
	"github.com/japiirainen/go-oluet-api/graph/generated"
	"github.com/japiirainen/go-oluet-api/graph/model"
)

func (r *mutationResolver) DailyJuomas(ctx context.Context) (string, error) {
	val, err := exel.ReadXlsx()
	if err != nil {
		log.Fatal(err)
	}
	//! bug here
	// stmt, err := db.Db.Prepare("INSERT INTO Juoma(Date, ProductID, Nimi, Valmistaja, PulloKoko, Hinta, LitraHinta, Uutuus, HinnastoJarjestysKoodi, Tyyppi, AlaTyyppi, ErityisRyhma, OlutTyyppi, ValmistusMaa, Alue, VuosiKerta, EtikettiMerkintoja, Huomautus, Rypaleet, Luonnehdinta, PakkausTyyppi, SuljentaTyyppi, AlkoholiProsentti, HapotGl, SokeriGL, Kantavierrep, Vari, Katkerot, Energia100ml, Valikoima) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()

	for _, v := range val {
		fmt.Printf("%v+", v)
		return "lol", nil
		// _, err := stmt.Exec(v.Date, v.ProductID, v.Nimi, v.Valikoima, v.PulloKoko, v.Hinta, v.LitraHinta, v.Uutuus, v.HinnastoJarjestysKoodi, v.Tyyppi, v.AlaTyyppi, v.ErityisRyhma, v.OlutTyyppi, v.ValmistusMaa, v.Alue, v.VuosiKerta, v.EtikettiMerkintoja, v.Huomautus, v.Rypaleet, v.Luonnehdinta, v.PakkausTyyppi, v.SuljentaTyyppi, v.AlkoholiProsentti, v.HapotGl, v.SokeriGl, v.Kantavierrep, v.Vari, v.Katkerot, v.Energia100ml, v.Valikoima)
		if err != nil {
			log.Fatal(err)
		}
	}

	return "success", nil
}

func (r *queryResolver) Juoma(ctx context.Context, id string) (*model.Juoma, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Juomat(ctx context.Context) ([]*model.Juoma, error) {
	rows, err := db.Db.Query("SELECT * FROM Juoma WHERE")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var juomat []*model.Juoma
	for rows.Next() {
		var juoma *model.Juoma
		err := rows.Scan(&juoma.ID, &juoma.ProductID, &juoma.Nimi, &juoma.Valikoima, &juoma.PulloKoko, &juoma.Hinta, &juoma.LitraHinta, juoma.Uutuus, &juoma.HinnastoJarjestysKoodi, &juoma.Tyyppi, &juoma.AlaTyyppi, &juoma.ErityisRyhma, &juoma.OlutTyyppi, &juoma.ValmistusMaa, &juoma.Alue, &juoma.VuosiKerta, &juoma.EtikettiMerkintoja, &juoma.Huomautus, &juoma.Rypaleet, &juoma.Luonnehdinta, &juoma.PakkausTyyppi, &juoma.SuljentaTyyppi, &juoma.AlkoholiProsentti, &juoma.HapotGl, &juoma.SokeriGl, &juoma.Kantavierrep, &juoma.Vari, &juoma.Katkerot, &juoma.Energia100ml, &juoma.Valikoima)
		if err != nil {
			log.Fatal(err)
		}

		juomat = append(juomat, juoma)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print("Row inserted!")
	return juomat, nil
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
