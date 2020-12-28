package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var fileLocation = filepath.Join("exel", "data", "alkon-hinnasto-tekstitiedostona.xlsx")

// Juoma is a type for one drink
type Juoma struct {
	productID              string
	nimi                   string
	valmistaja             string
	pulloKoko              string
	hinta                  string
	litraHinta             string
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
	sokeriGl               string
	kantavierrep           string
	vari                   string
	katkerot               string
	energia100ml           string
	valikoima              string
	ean                    string
}

func read() (Juoma, error) {
	f, err := excelize.OpenFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("Alkon Hinnasto Tekstitiedostona")
	if err != nil {
		log.Fatal(err)
	}
	res := Juoma{}
	for _, row := range rows[4:] {
		res = Juoma{productID: row[0], nimi: row[1], valmistaja: row[2], pulloKoko: row[3], hinta: row[4],
			litraHinta: row[5], uutuus: row[6], hinnastoJarjestysKoodi: row[7], tyyppi: row[8], alaTyyppi: row[9], erityisRyhma: row[10],
			olutTyyppi: row[11], valmistusMaa: row[12], alue: row[13], vuosiKerta: row[14], etikettiMerkintoja: row[15], huomautus: row[16],
			rypaleet: row[17], luonnehdinta: row[18], pakkausTyyppi: row[19], suljentaTyyppi: row[20], alkoholiProsentti: row[21], hapotGl: row[22],
			sokeriGl: row[23], kantavierrep: row[24], vari: row[25], katkerot: row[26], energia100ml: row[27], valikoima: row[28], ean: row[29]}
		fmt.Printf("%v", res)
	}
	return res, nil
}

func main() {
	some, e := read()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("%+v", some)
}
