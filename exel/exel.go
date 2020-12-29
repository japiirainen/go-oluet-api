package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var fileLocation = filepath.Join("exel", "data", "alkoFile.xlsx")

const alkoFileURI = "https://www.alko.fi/INTERSHOP/static/WFS/Alko-OnlineShop-Site/-/Alko-OnlineShop/fi_FI/Alkon%20Hinnasto%20Tekstitiedostona/alkon-hinnasto-tekstitiedostona.xlsx"

// Juoma is a type for one drink
type Juoma struct {
	productID              string
	nimi                   string
	valmistaja             string
	pulloKoko              string
	hinta                  float64
	litraHinta             float64
	uutuus                 string
	hinnastoJarjestysKoodi string
	tyyppi                 string
	alaTyyppi              string
	erityisRyhma           string
	olutTyyppi             string
	valmistusMaa           string
	alue                   string
	vuosiKerta             string
	etikettiMerkintoja     string
	huomautus              string
	rypaleet               string
	luonnehdinta           string
	pakkausTyyppi          string
	suljentaTyyppi         string
	alkoholiProsentti      string
	hapotGl                string
	sokeriGl               int
	kantavierrep           float64
	vari                   string
	katkerot               string
	energia100ml           string
	valikoima              string
}

//ReadXlsx returns all data from alko price file.
func ReadXlsx() ([]Juoma, error) {
	f, err := excelize.OpenFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("Alkon Hinnasto Tekstitiedostona")
	if err != nil {
		log.Fatal(err)
	}
	res := []Juoma{}
	for _, row := range rows[4:] {
		temp := Juoma{productID: row[0], nimi: row[1], valmistaja: row[2], pulloKoko: row[3], hinta: toFloat(row[4]),
			litraHinta: toFloat(row[5]), uutuus: row[6], hinnastoJarjestysKoodi: row[7], tyyppi: row[8], alaTyyppi: row[9], erityisRyhma: row[10],
			olutTyyppi: row[11], valmistusMaa: row[12], alue: row[13], vuosiKerta: row[14], etikettiMerkintoja: row[15], huomautus: row[16],
			rypaleet: row[17], luonnehdinta: row[18], pakkausTyyppi: row[19], suljentaTyyppi: row[20], alkoholiProsentti: row[21], hapotGl: row[22],
			sokeriGl: toInt(row[23]), kantavierrep: toFloat(row[24]), vari: row[25], katkerot: row[26], energia100ml: row[27], valikoima: row[28]}
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
	fmt.Println("Done downloading file: " + alkoFileURI)
	return err
}

func main() {
	err := Download(fileLocation, alkoFileURI)
	if err != nil {
		panic(err)
	}

	_, error := ReadXlsx()
	if error != nil {
		panic("error reading file")
	}
}

func toFloat(v string) float64 {
	res, _ := strconv.ParseFloat(v, 64)
	return res
}

func toInt(v string) int {
	res, _ := strconv.Atoi(v)
	return res
}
