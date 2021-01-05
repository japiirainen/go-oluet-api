package db

import (
	"fmt"
	"time"

	"github.com/japiirainen/go-oluet-api/exel"
	"github.com/japiirainen/go-oluet-api/gql/model"
	"github.com/japiirainen/go-oluet-api/helpers"
	log "github.com/sirupsen/logrus"
)

//InsertManyJuomas reads the alko file and inserts everything to postgres
func (db *Db) InsertManyJuomas() (string, error) {
	val, err := exel.ReadXlsx()
	if err != nil {
		log.Errorf("db: %s", err)
		return "err during exel read", err
	}
	OK, jerr := db.insertJuomas(&val)
	if !OK {
		log.Errorf("db: %s", jerr)
	}
	OK2, herr := db.CreatePrices(&val)
	if !OK2 {
		log.Fatal(herr)
	}
	return "OK", nil
}

func (db *Db) insertJuomas(juomat *[]exel.Juoma) (OK bool, error error) {
	defer helpers.Duration(time.Now(), "insertJuomas")
	stmt, prepErr := db.conn.Prepare("INSERT INTO juoma (Date," +
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
	for _, v := range *juomat {
		_, err := stmt.Exec(v.Date, v.ProductID, v.Nimi, v.Valikoima, v.PulloKoko, v.Hinta, v.LitraHinta, v.Uutuus, v.HinnastoJarjestysKoodi, v.Tyyppi, v.AlaTyyppi, v.ErityisRyhma, v.OlutTyyppi, v.ValmistusMaa, v.Alue, v.VuosiKerta, v.EtikettiMerkintoja, v.Huomautus, v.Rypaleet, v.Luonnehdinta, v.PakkausTyyppi, v.SuljentaTyyppi, v.AlkoholiProsentti, v.HapotGl, v.SokeriGl, v.Kantavierrep, v.Vari, v.Katkerot, v.Energia100ml, v.Valikoima)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// GetAllJuomas finds all the drinks
func (db *Db) GetAllJuomas() ([]model.Juoma, error) {
	rows, err := db.conn.Query("SELECT * FROM juoma;")
	if err != nil {
		log.Errorf("db: %s", err)
	}
	defer rows.Close()

	var juomat []model.Juoma
	for rows.Next() {
		var juoma model.Juoma
		err := rows.Scan(&juoma.ID, &juoma.Date, &juoma.ProductID, &juoma.Nimi, &juoma.Valikoima, &juoma.PulloKoko, &juoma.Hinta, &juoma.LitraHinta, &juoma.Uutuus, &juoma.HinnastoJarjestysKoodi, &juoma.Tyyppi, &juoma.AlaTyyppi, &juoma.ErityisRyhma, &juoma.OlutTyyppi, &juoma.ValmistusMaa, &juoma.Alue, &juoma.VuosiKerta, &juoma.EtikettiMerkintoja, &juoma.Huomautus, &juoma.Rypaleet, &juoma.Luonnehdinta, &juoma.PakkausTyyppi, &juoma.SuljentaTyyppi, &juoma.AlkoholiProsentti, &juoma.HapotGl, &juoma.SokeriGl, &juoma.Kantavierrep, &juoma.Vari, &juoma.Katkerot, &juoma.Energia100ml, &juoma.Valikoima)
		if err != nil {
			log.Errorf("db: %s", err)
		}

		juomat = append(juomat, juoma)
	}

	if err = rows.Err(); err != nil {
		log.Errorf("db: %s", err)
	}
	return juomat, nil
}

//GetJuomaByProdID gets one juoma with specified productID
func (db *Db) GetJuomaByProdID(ProductID string) (j model.Juoma, error error) {
	rows, err := db.conn.Query("SELECT * FROM juoma WHERE productid = $1;", ProductID)
	defer rows.Close()
	if err != nil {
		log.Errorf("db: failed to get juoma: %v\n", err)
	}
	var juoma model.Juoma
	for rows.Next() {
		scanErr := rows.Scan(&juoma.ID, &juoma.Date, &juoma.ProductID, &juoma.Nimi, &juoma.Valikoima, &juoma.PulloKoko, &juoma.Hinta, &juoma.LitraHinta, &juoma.Uutuus, &juoma.HinnastoJarjestysKoodi, &juoma.Tyyppi, &juoma.AlaTyyppi, &juoma.ErityisRyhma, &juoma.OlutTyyppi, &juoma.ValmistusMaa, &juoma.Alue, &juoma.VuosiKerta, &juoma.EtikettiMerkintoja, &juoma.Huomautus, &juoma.Rypaleet, &juoma.Luonnehdinta, &juoma.PakkausTyyppi, &juoma.SuljentaTyyppi, &juoma.AlkoholiProsentti, &juoma.HapotGl, &juoma.SokeriGl, &juoma.Kantavierrep, &juoma.Vari, &juoma.Katkerot, &juoma.Energia100ml, &juoma.Valikoima)
		if scanErr != nil {
			log.Errorf("db: failed to scan juoma: %v\n", err)
		}
	}
	return juoma, nil
}

//SearchForJuoma gets one or more juoma based on search results
func (db *Db) SearchForJuoma(term string) (j []model.Juoma, error error) {
	qstr := fmt.Sprintf("SELECT * FROM juoma WHERE nimi ILIKE '%%%s%%' OR tyyppi ILIKE '%%%s%%'", term, term)
	rows, err := db.conn.Query(qstr)
	defer rows.Close()
	if err != nil {
		log.Errorf("db: failed to find juoma: %v\n", err)
	}
	var juomat []model.Juoma
	for rows.Next() {
		var juoma model.Juoma
		scanErr := rows.Scan(&juoma.ID, &juoma.Date, &juoma.ProductID, &juoma.Nimi, &juoma.Valikoima, &juoma.PulloKoko, &juoma.Hinta, &juoma.LitraHinta, &juoma.Uutuus, &juoma.HinnastoJarjestysKoodi, &juoma.Tyyppi, &juoma.AlaTyyppi, &juoma.ErityisRyhma, &juoma.OlutTyyppi, &juoma.ValmistusMaa, &juoma.Alue, &juoma.VuosiKerta, &juoma.EtikettiMerkintoja, &juoma.Huomautus, &juoma.Rypaleet, &juoma.Luonnehdinta, &juoma.PakkausTyyppi, &juoma.SuljentaTyyppi, &juoma.AlkoholiProsentti, &juoma.HapotGl, &juoma.SokeriGl, &juoma.Kantavierrep, &juoma.Vari, &juoma.Katkerot, &juoma.Energia100ml, &juoma.Valikoima)
		if scanErr != nil {
			log.Errorf("db: failed to scan juoma: %v\n", err)
		}
		juomat = append(juomat, juoma)
	}
	return juomat, nil
}

//GetAllBeers gets all the beers from postgres
func (db *Db) GetAllBeers() (j []model.Juoma, error error) {
	rows, err := db.conn.Query("SELECT * FROM juoma WHERE tyyppi = 'oluet'")
	if err != nil {
		log.Errorf("db: %s", err)
	}
	defer rows.Close()
	var juomat []model.Juoma
	for rows.Next() {
		var juoma model.Juoma
		err := rows.Scan(&juoma.ID, &juoma.Date, &juoma.ProductID, &juoma.Nimi, &juoma.Valikoima, &juoma.PulloKoko, &juoma.Hinta, &juoma.LitraHinta, &juoma.Uutuus, &juoma.HinnastoJarjestysKoodi, &juoma.Tyyppi, &juoma.AlaTyyppi, &juoma.ErityisRyhma, &juoma.OlutTyyppi, &juoma.ValmistusMaa, &juoma.Alue, &juoma.VuosiKerta, &juoma.EtikettiMerkintoja, &juoma.Huomautus, &juoma.Rypaleet, &juoma.Luonnehdinta, &juoma.PakkausTyyppi, &juoma.SuljentaTyyppi, &juoma.AlkoholiProsentti, &juoma.HapotGl, &juoma.SokeriGl, &juoma.Kantavierrep, &juoma.Vari, &juoma.Katkerot, &juoma.Energia100ml, &juoma.Valikoima)
		if err != nil {
			log.Errorf("db: scan failed: %s", err)
		}
		juomat = append(juomat, juoma)
	}
	return juomat, nil
}

//SearchForBeer gets one or more beer based on search results
func (db *Db) SearchForBeer(term string) (j []model.Juoma, error error) {
	qstr := fmt.Sprintf("SELECT * FROM juoma WHERE nimi ILIKE '%%%s%%' AND tyyppi = 'oluet'", term)
	rows, err := db.conn.Query(qstr)
	defer rows.Close()
	if err != nil {
		log.Errorf("db: failed to find juoma: %v\n", err)
	}
	var juomat []model.Juoma
	for rows.Next() {
		var juoma model.Juoma
		scanErr := rows.Scan(&juoma.ID, &juoma.Date, &juoma.ProductID, &juoma.Nimi, &juoma.Valikoima, &juoma.PulloKoko, &juoma.Hinta, &juoma.LitraHinta, &juoma.Uutuus, &juoma.HinnastoJarjestysKoodi, &juoma.Tyyppi, &juoma.AlaTyyppi, &juoma.ErityisRyhma, &juoma.OlutTyyppi, &juoma.ValmistusMaa, &juoma.Alue, &juoma.VuosiKerta, &juoma.EtikettiMerkintoja, &juoma.Huomautus, &juoma.Rypaleet, &juoma.Luonnehdinta, &juoma.PakkausTyyppi, &juoma.SuljentaTyyppi, &juoma.AlkoholiProsentti, &juoma.HapotGl, &juoma.SokeriGl, &juoma.Kantavierrep, &juoma.Vari, &juoma.Katkerot, &juoma.Energia100ml, &juoma.Valikoima)
		if scanErr != nil {
			log.Errorf("db: failed to scan juoma: %v\n", err)
		}
		juomat = append(juomat, juoma)
	}
	return juomat, nil
}
