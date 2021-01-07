package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HomeHandler is the handler for GET /
func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Info("handlers: GET /")
	json.NewEncoder(rw).Encode(map[string]bool{"ok": true})
}
