package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/japiirainen/go-oluet-api/graph/model"

	"github.com/japiirainen/go-oluet-api/exel"
	"github.com/japiirainen/go-oluet-api/helpers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Db is the database connection
type Db struct {
	conn *sql.DB
}

// Connect makes a postgres connection
func Connect() *Db {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%s", err)
		panic("env not found")
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		helpers.ToInt(os.Getenv("PG_PORT")),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"))

	// open database
	conn, dbErr := sql.Open("postgres", psqlconn)
	if dbErr != nil {
		log.Panic(dbErr)
	}

	if err = conn.Ping(); err != nil {
		log.Panic(err)
	}
	println("postgres connection succesful")

	return &Db{
		conn: conn,
	}
}

//CloseConnection closes the sql conection
func (db *Db) CloseConnection() {
	db.conn.Close()
}

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
