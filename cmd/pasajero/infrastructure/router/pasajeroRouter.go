package router

import (
	"back-transportes-fisi/cmd/pasajero/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func PasajeroRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	pasajeroController := controller.NewPasajeroController(serviceContainer)

	r.HandleFunc("/pasajero", pasajeroController.Create).Methods("POST")
	r.HandleFunc("/pasajero", pasajeroController.GetAll).Methods("GET")
	r.HandleFunc("/pasajero/{id}", pasajeroController.GetById).Methods("GET")
	r.HandleFunc("/pasajero/{id}", pasajeroController.Update).Methods("PUT")
	r.HandleFunc("/pasajero/{id}", pasajeroController.Delete).Methods("DELETE")
}
