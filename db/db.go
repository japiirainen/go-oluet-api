package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/japiirainen/go-oluet-api/helpers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

//Db is the database connection
type Db struct {
	conn *sql.DB
}

// Connect makes a postgres connection
func Connect() *Db {
	err := godotenv.Load(".env")
	if err != nil {
		log.Errorf("%s", err)
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
	log.Info("postgres connection succesful")

	return &Db{
		conn: conn,
	}
}

//CloseConnection closes the sql conection
func (db *Db) CloseConnection() {
	db.conn.Close()
}
