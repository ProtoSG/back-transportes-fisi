package router

import (
	"back-transportes-fisi/cmd/ciudad/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func CiudadRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	ciudadController := controller.NewCiudadController(serviceContainer)

	r.HandleFunc("/ciudad", ciudadController.Create).Methods("POST")
	r.HandleFunc("/ciudad", ciudadController.GetAll).Methods("GET")
	r.HandleFunc("/ciudad/{id}", ciudadController.GetById).Methods("GET")
	r.HandleFunc("/ciudad/{id}", ciudadController.Update).Methods("PUT")
	r.HandleFunc("/ciudad/{id}", ciudadController.Delete).Methods("DELETE")
}
