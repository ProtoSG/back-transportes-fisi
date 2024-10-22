package router

import (
	"back-transportes-fisi/cmd/descuento/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func DescuentoRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	descuentoController := controller.NewDescuentoController(serviceContainer)

	r.HandleFunc("/descuento", descuentoController.Create).Methods("POST")
	r.HandleFunc("/descuento", descuentoController.GetAll).Methods("GET")
	r.HandleFunc("/descuento/{id}", descuentoController.GetById).Methods("GET")
	r.HandleFunc("/descuento/{id}", descuentoController.Update).Methods("PUT")
	r.HandleFunc("/descuento/{id}", descuentoController.Delete).Methods("DELETE")
}
