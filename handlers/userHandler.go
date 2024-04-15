package handlers

import (
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
