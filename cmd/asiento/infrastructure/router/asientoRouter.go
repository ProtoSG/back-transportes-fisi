package router

import (
	"back-transportes-fisi/cmd/asiento/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func AsientoRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	asientoController := controller.NewAsientoController(serviceContainer)

	r.HandleFunc("/asiento", asientoController.Create).Methods("POST")
	r.HandleFunc("/asiento", asientoController.GetAll).Methods("GET")
	r.HandleFunc("/asiento/{id}", asientoController.GetById).Methods("GET")
	r.HandleFunc("/asiento/{id}", asientoController.Update).Methods("PUT")
	r.HandleFunc("/asiento/{id}", asientoController.Delete).Methods("DELETE")
	r.HandleFunc("/asiento/bus/{id}", asientoController.GetByIdBus).Methods("GET")
}
