package router

import (
	"back-transportes-fisi/cmd/servicio/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func ServicioRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	servicioController := controller.NewServicioController(serviceContainer)

	r.HandleFunc("/servicio", servicioController.GetAll).Methods("GET")
	r.HandleFunc("/servicio/{id}", servicioController.GetById).Methods("GET")
	r.HandleFunc("/servicio", servicioController.Create).Methods("POST")
	r.HandleFunc("/servicio/{id}", servicioController.Update).Methods("PUT")
	r.HandleFunc("/servicio/{id}", servicioController.Delete).Methods("DELETE")
}
