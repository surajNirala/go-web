package handlers

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User Details")
}
