package router

import (
	"back-transportes-fisi/cmd/boleto/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func BoletoRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	boletoController := controller.NewBoletoController(serviceContainer)

	r.HandleFunc("/boleto", boletoController.Create).Methods("POST")
	r.HandleFunc("/boleto", boletoController.GetAll).Methods("GET")
	r.HandleFunc("/boleto/{id}", boletoController.GetById).Methods("GET")
	r.HandleFunc("/boleto/{id}", boletoController.Update).Methods("PUT")
	r.HandleFunc("/boleto/{id}", boletoController.Delete).Methods("DELETE")
}
