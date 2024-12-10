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
	data, err := json.Marshal(payload) // json stringfy del objeto que te viene en el payload

	if err != nil {
		log.Printf("Error al responder con JSON %v", payload)
		writer.WriteHeader(code)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(data)
}

func respondeWithError(writer http.ResponseWriter, code int, msg string) {

	if code >= 500 {
		log.Printf("Respondiendo en Server Error con codigo %v", code)
		type errResponse struct {
			Error string `json:"error"`
			Code  int    `json:"code"`
		}

		respondeWithJSON(writer, code, errResponse{
			Error: msg,
			Code:  code,
		})

		return
	}
}
