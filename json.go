package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
interface {} es como una struct comodin, que acepta cualquier formato.
*/
func respondeWithJSON(writer http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error al responder con JSON %v", payload)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}
