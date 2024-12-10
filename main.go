package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// server que responde en formato json
func main() {
	godotenv.Load() //con esto me carga las variables de entorno
	var port string = os.Getenv("port")

	if port == "" {
		log.Fatalln("F de fatal")
	}

	// Enrutador principal
	router := chi.NewRouter() // router principal

	// V1 router
	v1router := chi.NewRouter() // para cuando tenes diferentes versiones
	v1router.Get("/health", handlerHealth)

	// configuracion del middleware
	router.Use(cors.Default().Handler)
	// setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	// Mas info en https://pkg.go.dev/github.com/rs/cors#readme-allow-with-credentials-security-protection

	//server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	router.Mount("/v1", v1router) // para todas als request que vayan a /v1, enrutalas segun la configuracion del v1router.

	log.Printf("Server listening on port: %v", port)
	err := server.ListenAndServe() // deberia correr por siempre

	if err != nil {
		log.Fatalln("F listener")
	}
}
