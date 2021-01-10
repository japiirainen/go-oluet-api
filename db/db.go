package db

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

//Db is the database connection
type Db struct {
	conn *sql.DB
}

//MigrateUp runs migrations
func MigrateUp(dbURL string) {
	m, _ := migrate.New(
		"file://db/migrations/postgres",
		dbURL)
	m.Up()
}

// Connect makes a postgres connection
func Connect(dbURL string) *Db {
	// open database
	conn, dbErr := sql.Open("postgres", dbURL)
	if dbErr != nil {
		log.Panic(dbErr)
	}

	err := conn.Ping()
	if err != nil {
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
