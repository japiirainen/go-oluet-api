package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/japiirainen/go-oluet-api/db"
	"github.com/japiirainen/go-oluet-api/exel"
	log "github.com/sirupsen/logrus"
)

//GetInternal is the handler for GET /internal
func GetInternal(rw http.ResponseWriter, r *http.Request) {
	log.Info("handlers: GET /internal")
	json.NewEncoder(rw).Encode(map[string]string{"message": "you should send a post request to execute the update!"})
}

//DailyUpdate is the handler for POST /internal
func DailyUpdate(rw http.ResponseWriter, r *http.Request) {
	log.Info("handlers: POST /internal")
	var wg sync.WaitGroup
	dbURL := os.Getenv("DATABASE_URL")
	conn := db.Connect(dbURL)
	defer conn.CloseConnection()
	// download file from alko website
	dErr := exel.Download(exel.FileLocation, exel.AlkoFileURI)
	if dErr != nil {
		log.Errorf("internal: %s\n", dErr)
	}
	drinkExel, rErr := exel.ReadXlsx(exel.FileLocation)
	// read the downloaded file
	if rErr != nil {
		log.Errorf("internal: %s\n", rErr)
	}
	// delete old rows from drink table
	deleteErr := conn.DeleteDrinks()
	if deleteErr != nil {
		log.Errorf("internal: %s\n", deleteErr)
	}
	// insert new drinks
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ok, err := conn.InsertDrinks(&drinkExel)
		if !ok {
			log.Errorf("internal: %s", err)
		}
	}(&wg)
	// insert new prices
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		ok, err := conn.CreatePrices(&drinkExel)
		if !ok {
			log.Errorf("internal: %s\n", err)
		}
	}(&wg)
	wg.Wait()
	json.NewEncoder(rw).Encode(map[string]interface{}{"message": "Daily update was succesfull", "date": time.Now()})
}
