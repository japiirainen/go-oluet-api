package db

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var dbURL = os.Getenv("DATABASE_URL")

//Db is the database connection
type Db struct {
	conn *sql.DB
}

//MigrateUp runs migrations
func MigrateUp() {
	m, err := migrate.New(
		"file://db/migrations",
		dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

// Connect makes a postgres connection
func Connect() *Db {
	err := godotenv.Load(".env")
	if err != nil {
		log.Errorf("%s", err)
		panic("env not found")
	}

	// open database
	conn, dbErr := sql.Open("postgres", dbURL)
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
