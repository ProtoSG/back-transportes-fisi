package router

import (
	"back-transportes-fisi/cmd/shared/infrastructure"
	"back-transportes-fisi/cmd/viaje/infrastructure/controller"

	"github.com/gorilla/mux"
)

func ViajeRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	viajeController := controller.NewViajeController(serviceContainer)

	r.HandleFunc("/viaje", viajeController.Create).Methods("POST")
	r.HandleFunc("/viaje", viajeController.GetAll).Methods("GET")
	r.HandleFunc("/viaje/search", viajeController.GetByOriginDestinationAndDate).Methods("GET")
	r.HandleFunc("/viaje/{id}", viajeController.GetById).Methods("GET")
	r.HandleFunc("/viaje/{id}", viajeController.Update).Methods("PUT")
	r.HandleFunc("/viaje/{id}", viajeController.Delete).Methods("DELETE")
}
