package handlers

import (
	"encoding/json"
	"net/http"

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
}
