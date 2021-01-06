package db

import (
	"github.com/japiirainen/go-oluet-api/exel"
	"github.com/japiirainen/go-oluet-api/gql/model"
)

//CreatePrices creates new prices for juomas
func (db *Db) CreatePrices(drinks *[]exel.Drink) (OK bool, error error) {
	stmp, stmpErr := db.conn.Prepare("INSERT INTO price (Date, ProductID, Hinta) VALUES ($1, $2, $3)")
	if stmpErr != nil {
		return false, stmpErr
	}
	defer stmp.Close()
	for _, drink := range *drinks {
		_, err := stmp.Exec(drink.Date, drink.ProductID, drink.Hinta)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

//GetAllPrices gets all the prices
func (db *Db) GetAllPrices() ([]model.Price, error) {
	rows, err := db.conn.Query("SELECT * FROM price;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var prices []model.Price
	for rows.Next() {
		var price model.Price
		scanErr := rows.Scan(&price.ID, &price.Date, &price.ProductID, &price.Hinta)
		if scanErr != nil {
			return nil, scanErr
		}
		prices = append(prices, price)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return prices, nil
}
