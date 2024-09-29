package router

import (
	"back-transportes-fisi/cmd/bus/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func BusRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	busController := controller.NewHttpBusController(serviceContainer)

	r.HandleFunc("/bus", busController.Create).Methods("POST")
	r.HandleFunc("/bus", busController.GetAll).Methods("GET")
	r.HandleFunc("/bus/{id}", busController.GetById).Methods("GET")
	r.HandleFunc("/bus/{id}", busController.Update).Methods("PUT")
	r.HandleFunc("/bus/{id}", busController.Delete).Methods("DELETE")
}
