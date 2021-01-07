package middleware

import (
	"net/http"
	"os"
)

//CheckAuth checks if there is auth for internal routes
func CheckAuth(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Query().Get("password") != os.Getenv("INTERNAL_PASSWORD") {
		http.Error(rw, "Not Authorized", 401)
	} else {
		next(rw, r)
	}
}
