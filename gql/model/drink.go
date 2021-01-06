package model

import (
	"time"
)

// Drink is the type of one Drink
type Drink struct {
	ID                     int       `json:"id"`
	Date                   time.Time `json:"date"`
	ProductID              string    `json:"productId"`
	Nimi                   string    `json:"nimi"`
	Valmistaja             string    `json:"valmistaja"`
	PulloKoko              string    `json:"pulloKoko"`
	Hinta                  float64   `json:"hinta"`
	LitraHinta             float64   `json:"litraHinta"`
	Uutuus                 string    `json:"uutuus"`
	HinnastoJarjestysKoodi string    `json:"hinnastoJarjestysKoodi"`
	Tyyppi                 string    `json:"tyyppi"`
	AlaTyyppi              string    `json:"alaTyyppi"`
	ErityisRyhma           string    `json:"erityisRyhma"`
	OlutTyyppi             string    `json:"olutTyyppi"`
	ValmistusMaa           string    `json:"valmistusMaa"`
	Alue                   string    `json:"alue"`
	VuosiKerta             string    `json:"vuosiKerta"`
	EtikettiMerkintoja     string    `json:"etikettiMerkintoja"`
	Huomautus              string    `json:"huomautus"`
	Rypaleet               string    `json:"rypaleet"`
	Luonnehdinta           string    `json:"luonnehdinta"`
	PakkausTyyppi          string    `json:"pakkausTyyppi"`
	SuljentaTyyppi         string    `json:"suljentaTyyppi"`
	AlkoholiProsentti      string    `json:"alkoholiProsentti"`
	HapotGl                string    `json:"hapotGL"`
	SokeriGl               int       `json:"sokeriGL"`
	Kantavierrep           float64   `json:"kantavierrep"`
	Vari                   string    `json:"vari"`
	Katkerot               string    `json:"katkerot"`
	Energia100ml           string    `json:"energia100ML"`
	Valikoima              string    `json:"valikoima"`
}
