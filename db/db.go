package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/japiirainen/go-oluet-api/helpers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Db is the database connection
var Db *sql.DB

// Init makes a postgres connection
func Init() {
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
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	println("postgres connection succesful")
	Db = db
}
