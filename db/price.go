package db

import (
	"github.com/japiirainen/go-oluet-api/exel"
	"github.com/japiirainen/go-oluet-api/gql/model"
)

//CreatePrices creates new prices for juomas
func (db *Db) CreatePrices(juomat *[]exel.Juoma) (OK bool, error error) {
	stmp, stmpErr := db.conn.Prepare("INSERT INTO Hinta (Date, ProductID, Hinta) VALUES ($1, $2, $3)")
	if stmpErr != nil {
		return false, stmpErr
	}
	defer stmp.Close()
	for _, juoma := range *juomat {
		_, err := stmp.Exec(juoma.Date, juoma.ProductID, juoma.Hinta)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

//GetAllPrices gets all the prices
func (db *Db) GetAllPrices() ([]model.Hinta, error) {
	rows, err := db.conn.Query("SELECT * FROM Hinta;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hinnat []model.Hinta
	for rows.Next() {
		var hinta model.Hinta
		scanErr := rows.Scan(&hinta.ID, &hinta.Date, &hinta.ProductID, &hinta.Hinta)
		if scanErr != nil {
			return nil, scanErr
		}
		hinnat = append(hinnat, hinta)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return hinnat, nil
}
