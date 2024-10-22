package router

import (
	"back-transportes-fisi/cmd/ruta/infrastructure/controller"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func RutaRouter(r *mux.Router, serviceContainer *infrastructure.ServiceContainer) {
	rutaController := controller.NewRutaController(serviceContainer)

	r.HandleFunc("/ruta", rutaController.Create).Methods("POST")
	r.HandleFunc("/ruta", rutaController.GetAll).Methods("GET")
	r.HandleFunc("/ruta/{id}", rutaController.GetById).Methods("GET")
	r.HandleFunc("/ruta/{id}", rutaController.Update).Methods("PUT")
	r.HandleFunc("/ruta/{id}", rutaController.Delete).Methods("DELETE")
}
