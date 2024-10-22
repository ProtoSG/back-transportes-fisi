package router

import (
	"back-transportes-fisi/cmd/cliente/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func ClienteRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	clienteController := controller.NewClienteController(serviceContainer)

	r.HandleFunc("/cliente", clienteController.Create).Methods("POST")
	r.HandleFunc("/cliente", clienteController.GetAll).Methods("GET")
	r.HandleFunc("/cliente/{id}", clienteController.GetById).Methods("GET")
	r.HandleFunc("/cliente/{id}", clienteController.Update).Methods("PUT")
	r.HandleFunc("/cliente/{id}", clienteController.Delete).Methods("DELETE")
}
