package main

import (
	"net/http"
)

func handleError(w http.ResponseWriter, r *http.Request) {
	respondeWithError(w, 501, "error test log")
}
