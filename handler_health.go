package main

import (
	"net/http"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	respondeWithJSON(w, 200, struct{}{})
}
