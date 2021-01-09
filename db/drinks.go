package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/japiirainen/go-oluet-api/exel"
	"github.com/japiirainen/go-oluet-api/gql/model"
	"github.com/japiirainen/go-oluet-api/helpers"
	log "github.com/sirupsen/logrus"
)

//InsertManyDrinks reads the alko file and inserts everything to postgres
func (db *Db) InsertManyDrinks() (string, error) {
	val, err := exel.ReadXlsx()
	if err != nil {
		log.Errorf("db: %s", err)
		return "err during exel read", err
	}
	OK, jerr := db.insertDrinks(&val)
	if !OK {
		log.Errorf("db: %s", jerr)
	}
	OK2, herr := db.CreatePrices(&val)
	if !OK2 {
		log.Fatal(herr)
	}
	return "OK", nil
}

func (db *Db) insertDrinks(drinks *[]exel.Drink) (OK bool, error error) {
	defer helpers.Duration(time.Now(), "insertDrinks")
	stmt, prepErr := db.conn.Prepare("INSERT INTO drink (Date," +
		" ProductID," +
		" Nimi," +
		" Valmistaja," +
		" PulloKoko," +
		" Hinta," +
		" LitraHinta," +
		" Uutuus," +
		" HinnastoJarjestysKoodi," +
		" Tyyppi," +
		" AlaTyyppi," +
		" ErityisRyhma," +
		" OlutTyyppi," +
		" ValmistusMaa," +
		" Alue," +
		" VuosiKerta," +
		" EtikettiMerkintoja," +
		" Huomautus," +
		" Rypaleet," +
		" Luonnehdinta," +
		" PakkausTyyppi," +
		" SuljentaTyyppi," +
		" AlkoholiProsentti," +
		" HapotGl," +
		" SokeriGL," +
		" Kantavierrep," +
		" Vari," +
		" Katkerot," +
		" Energia100ml," +
		" Valikoima) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30);")
	if prepErr != nil {
		return false, nil
	}
	defer stmt.Close()
	for _, v := range *drinks {
		_, err := stmt.Exec(v.Date, v.ProductID, v.Nimi, v.Valikoima, v.PulloKoko, v.Hinta, v.LitraHinta, v.Uutuus, v.HinnastoJarjestysKoodi, v.Tyyppi, v.AlaTyyppi, v.ErityisRyhma, v.OlutTyyppi, v.ValmistusMaa, v.Alue, v.VuosiKerta, v.EtikettiMerkintoja, v.Huomautus, v.Rypaleet, v.Luonnehdinta, v.PakkausTyyppi, v.SuljentaTyyppi, v.AlkoholiProsentti, v.HapotGl, v.SokeriGl, v.Kantavierrep, v.Vari, v.Katkerot, v.Energia100ml, v.Valikoima)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// GetAllDrinks finds all the drinks
func (db *Db) GetAllDrinks() ([]model.Drink, error) {
	rows, err := db.conn.Query("SELECT * FROM drink;")
	if err != nil {
		log.Errorf("db: %s", err)
	}
	defer rows.Close()
	drinks, scanErr := scanDrinks(rows)
	if scanErr != nil {
		log.Errorf("db: failed to scan juoma: %v\n", err)
	}
	return drinks, nil
}

//GetDrinkByProdID gets one juoma with specified productID
func (db *Db) GetDrinkByProdID(ProductID string) (d model.Drink, error error) {
	rows, err := db.conn.Query("SELECT * FROM drink WHERE productid = $1;", ProductID)
	defer rows.Close()
	if err != nil {
		log.Errorf("db: failed to get juoma: %v\n", err)
	}
	var drink model.Drink
	for rows.Next() {
		scanErr := rows.Scan(&drink.ID, &drink.Date, &drink.ProductID, &drink.Nimi, &drink.Valikoima, &drink.PulloKoko, &drink.Hinta, &drink.LitraHinta, &drink.Uutuus, &drink.HinnastoJarjestysKoodi, &drink.Tyyppi, &drink.AlaTyyppi, &drink.ErityisRyhma, &drink.OlutTyyppi, &drink.ValmistusMaa, &drink.Alue, &drink.VuosiKerta, &drink.EtikettiMerkintoja, &drink.Huomautus, &drink.Rypaleet, &drink.Luonnehdinta, &drink.PakkausTyyppi, &drink.SuljentaTyyppi, &drink.AlkoholiProsentti, &drink.HapotGl, &drink.SokeriGl, &drink.Kantavierrep, &drink.Vari, &drink.Katkerot, &drink.Energia100ml, &drink.Valikoima)
		if scanErr != nil {
			log.Errorf("db: failed to scan juoma: %v\n", err)
		}
	}
	return drink, nil
}

//SearchForDrink gets one or more juoma based on search results
func (db *Db) SearchForDrink(term string) (j []model.Drink, error error) {
	qstr := fmt.Sprintf("SELECT * FROM drink WHERE nimi ILIKE '%%%s%%' OR tyyppi ILIKE '%%%s%%'", term, term)
	rows, err := db.conn.Query(qstr)
	defer rows.Close()
	if err != nil {
		log.Errorf("db: failed to find juoma: %v\n", err)
	}
	drinks, scanErr := scanDrinks(rows)
	if scanErr != nil {
		log.Errorf("db: failed to scan drinks: %v\n", err)
	}
	return drinks, nil
}

//GetAllBeers gets all the beers from postgres
func (db *Db) GetAllBeers() (d []model.Drink, error error) {
	rows, err := db.conn.Query("SELECT * FROM drink WHERE tyyppi = 'oluet'")
	if err != nil {
		log.Errorf("db: %s", err)
	}
	defer rows.Close()
	drinks, scanErr := scanDrinks(rows)
	if scanErr != nil {
		log.Errorf("db: failed to scan drinks: %v\n", err)
	}
	return drinks, nil
}

//SearchForBeer gets one or more beer based on search results
func (db *Db) SearchForBeer(term string) (d []model.Drink, error error) {
	qstr := fmt.Sprintf("SELECT * FROM drink WHERE nimi ILIKE '%%%s%%' AND tyyppi = 'oluet'", term)
	rows, err := db.conn.Query(qstr)
	defer rows.Close()
	if err != nil {
		log.Errorf("db: %v\n", err)
	}
	drinks, scanErr := scanDrinks(rows)
	if scanErr != nil {
		log.Errorf("db: failed to find drinks: %v\n", err)
	}
	return drinks, nil
}

func scanDrinks(rows *sql.Rows) (ds []model.Drink, error error) {
	var drinks []model.Drink
	for rows.Next() {
		var drink model.Drink
		err := rows.Scan(&drink.ID, &drink.Date, &drink.ProductID, &drink.Nimi, &drink.Valikoima, &drink.PulloKoko, &drink.Hinta, &drink.LitraHinta, &drink.Uutuus, &drink.HinnastoJarjestysKoodi, &drink.Tyyppi, &drink.AlaTyyppi, &drink.ErityisRyhma, &drink.OlutTyyppi, &drink.ValmistusMaa, &drink.Alue, &drink.VuosiKerta, &drink.EtikettiMerkintoja, &drink.Huomautus, &drink.Rypaleet, &drink.Luonnehdinta, &drink.PakkausTyyppi, &drink.SuljentaTyyppi, &drink.AlkoholiProsentti, &drink.HapotGl, &drink.SokeriGl, &drink.Kantavierrep, &drink.Vari, &drink.Katkerot, &drink.Energia100ml, &drink.Valikoima)
		if err != nil {
			return nil, err
		}
		drinks = append(drinks, drink)
	}
	return drinks, nil
}
