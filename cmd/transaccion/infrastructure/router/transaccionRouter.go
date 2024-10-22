package router

import (
	"back-transportes-fisi/cmd/shared/infrastructure"
	"back-transportes-fisi/cmd/transaccion/infrastructure/controller"

	"github.com/gorilla/mux"
)

func TransaccionRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	transaccionController := controller.NewTransaccionController(serviceContainer)

	r.HandleFunc("/transaccion", transaccionController.Create).Methods("POST")
}
