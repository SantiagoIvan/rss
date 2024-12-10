package main

import (
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	respondeWithJSON(w, 200, struct{}{})
}
