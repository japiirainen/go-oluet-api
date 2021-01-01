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

func (db *Db) InsertManyJuomas() (string, error) {
	val, err := exel.ReadXlsx()
	if err != nil {
		log.Fatal(err)
		return "err during exel read", err
	}
	stmt, err := db.conn.Prepare("INSERT INTO Juoma(Date, ProductID, Nimi, Valmistaja, PulloKoko, Hinta, LitraHinta, Uutuus, HinnastoJarjestysKoodi, Tyyppi, AlaTyyppi, ErityisRyhma, OlutTyyppi, ValmistusMaa, Alue, VuosiKerta, EtikettiMerkintoja, Huomautus, Rypaleet, Luonnehdinta, PakkausTyyppi, SuljentaTyyppi, AlkoholiProsentti, HapotGl, SokeriGL, Kantavierrep, Vari, Katkerot, Energia100ml, Valikoima) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
		return "err during prepare", err
	}
	defer stmt.Close()

	for _, v := range val {
		fmt.Printf("%v+", v)
		_, err := stmt.Exec(v.Date, v.ProductID, v.Nimi, v.Valikoima, v.PulloKoko, v.Hinta, v.LitraHinta, v.Uutuus, v.HinnastoJarjestysKoodi, v.Tyyppi, v.AlaTyyppi, v.ErityisRyhma, v.OlutTyyppi, v.ValmistusMaa, v.Alue, v.VuosiKerta, v.EtikettiMerkintoja, v.Huomautus, v.Rypaleet, v.Luonnehdinta, v.PakkausTyyppi, v.SuljentaTyyppi, v.AlkoholiProsentti, v.HapotGl, v.SokeriGl, v.Kantavierrep, v.Vari, v.Katkerot, v.Energia100ml, v.Valikoima)
		if err != nil {
			log.Fatal(err)
			return "err during exec", err
		}
	}

	return "OK", nil
}

func (db *Db) GetAllJuomas() ([]*model.Juoma, error) {
	rows, err := db.conn.Query("SELECT * FROM Juoma;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var juomat []*model.Juoma
	for rows.Next() {
		var juoma *model.Juoma
		err := rows.Scan(&juoma.ID, &juoma.ProductID, &juoma.Nimi, &juoma.Valikoima, &juoma.PulloKoko, &juoma.Hinta, &juoma.LitraHinta, juoma.Uutuus, &juoma.HinnastoJarjestysKoodi, &juoma.Tyyppi, &juoma.AlaTyyppi, &juoma.ErityisRyhma, &juoma.OlutTyyppi, &juoma.ValmistusMaa, &juoma.Alue, &juoma.VuosiKerta, &juoma.EtikettiMerkintoja, &juoma.Huomautus, &juoma.Rypaleet, &juoma.Luonnehdinta, &juoma.PakkausTyyppi, &juoma.SuljentaTyyppi, &juoma.AlkoholiProsentti, &juoma.HapotGl, &juoma.SokeriGl, &juoma.Kantavierrep, &juoma.Vari, &juoma.Katkerot, &juoma.Energia100ml, &juoma.Valikoima)
		if err != nil {
			log.Fatal(err)
		}

		juomat = append(juomat, juoma)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return juomat, nil
}
