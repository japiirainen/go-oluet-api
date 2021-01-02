package exel

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/japiirainen/go-oluet-api/helpers"
)

//FileLocation is the location of the prie file
var FileLocation = filepath.Join("exel", "data", "alkoFile.xlsx")

//AlkoFileURI is the URI that the price file gets downloaded from
const AlkoFileURI = "https://www.alko.fi/INTERSHOP/static/WFS/Alko-OnlineShop-Site/-/Alko-OnlineShop/fi_FI/Alkon%20Hinnasto%20Tekstitiedostona/alkon-hinnasto-tekstitiedostona.xlsx"

// Juoma is a type for one drink
type Juoma struct {
	Date                   time.Time
	ProductID              string
	Nimi                   string
	Valmistaja             string
	PulloKoko              string
	Hinta                  float64
	LitraHinta             float64
	Uutuus                 string
	HinnastoJarjestysKoodi string
	Tyyppi                 string
	AlaTyyppi              string
	ErityisRyhma           string
	OlutTyyppi             string
	ValmistusMaa           string
	Alue                   string
	VuosiKerta             string
	EtikettiMerkintoja     string
	Huomautus              string
	Rypaleet               string
	Luonnehdinta           string
	PakkausTyyppi          string
	SuljentaTyyppi         string
	AlkoholiProsentti      string
	HapotGl                string
	SokeriGl               int
	Kantavierrep           float64
	Vari                   string
	Katkerot               string
	Energia100ml           string
	Valikoima              string
}

//ReadXlsx returns all data from alko price file.
func ReadXlsx() ([]Juoma, error) {
	f, err := excelize.OpenFile(FileLocation)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("Alkon Hinnasto Tekstitiedostona")
	if err != nil {
		log.Fatal(err)
	}
	res := []Juoma{}

	date := strings.TrimPrefix(rows[0][0], "Alkon hinnasto ")
	parsedDate := helpers.ParseTime(date)

	for _, row := range rows[4:] {
		temp := Juoma{Date: parsedDate,
			ProductID:              row[0],
			Nimi:                   row[1],
			Valmistaja:             row[2],
			PulloKoko:              row[3],
			Hinta:                  helpers.ToFloat(row[4]),
			LitraHinta:             helpers.ToFloat(row[5]),
			Uutuus:                 row[6],
			HinnastoJarjestysKoodi: row[7],
			Tyyppi:                 row[8],
			AlaTyyppi:              row[9],
			ErityisRyhma:           row[10],
			OlutTyyppi:             row[11],
			ValmistusMaa:           row[12],
			Alue:                   row[13],
			VuosiKerta:             row[14],
			EtikettiMerkintoja:     row[15],
			Huomautus:              row[16],
			Rypaleet:               row[17],
			Luonnehdinta:           row[18],
			PakkausTyyppi:          row[19],
			SuljentaTyyppi:         row[20],
			AlkoholiProsentti:      row[21],
			HapotGl:                row[22],
			SokeriGl:               helpers.ToInt(row[23]),
			Kantavierrep:           helpers.ToFloat(row[24]),
			Vari:                   row[25],
			Katkerot:               row[26],
			Energia100ml:           row[27],
			Valikoima:              row[28]}
		res = append(res, temp)
	}
	return res, nil
}

// Download loads the file using http
func Download(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath)

	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	fmt.Println("Done downloading file: " + AlkoFileURI)
	return err
}
