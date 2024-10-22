package main

import (
	"back-transportes-fisi/cmd/shared/infrastructure"
	"back-transportes-fisi/cmd/shared/infrastructure/router"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	serviceContainer := infrastructure.NewServiceContainer()

	r := router.NewRouter(serviceContainer)

	log.Println("Servidor corriendo en puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
