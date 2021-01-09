package db

import (
	"time"

	"github.com/japiirainen/go-oluet-api/exel"
	"github.com/japiirainen/go-oluet-api/gql/model"
	"github.com/japiirainen/go-oluet-api/helpers"
	log "github.com/sirupsen/logrus"
)

//CreatePrices creates new prices for juomas
func (db *Db) CreatePrices(drinks *[]exel.Drink) (OK bool, error error) {
	defer helpers.Duration(time.Now(), "insertPrices")
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
		log.Errorf("db: %s\n", err)
	}
	defer rows.Close()
	var prices []model.Price
	for rows.Next() {
		var price model.Price
		scanErr := rows.Scan(&price.ID, &price.Date, &price.ProductID, &price.Hinta)
		if scanErr != nil {
			log.Errorf("db: %s\n", scanErr)
		}
		prices = append(prices, price)
	}
	if err = rows.Err(); err != nil {
		log.Errorf("db: %s\n", err)
	}
	return prices, nil
}

// GetPriceHistory returns the price history of one drink
func (db *Db) GetPriceHistory(productID string) ([]model.Price, error) {
	rows, err := db.conn.Query("SELECT * FROM price WHERE productid = $1", productID)

	var prices []model.Price
	for rows.Next() {
		var price model.Price
		scanErr := rows.Scan(&price.ID, &price.Date, &price.ProductID, &price.Hinta)
		if scanErr != nil {
			log.Errorf("db: %s\n", scanErr)
		}
		prices = append(prices, price)
	}
	if err = rows.Err(); err != nil {
		log.Errorf("db: %s\n", err)
	}
	return prices, nil
}
