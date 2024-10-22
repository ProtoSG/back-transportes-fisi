package router

import (
	"back-transportes-fisi/cmd/shared/infrastructure"
	"back-transportes-fisi/cmd/terminal/infrastructure/controller"

	"github.com/gorilla/mux"
)

func TerminalRouter(r *mux.Router, serviceContaienr *infrastructure.ServiceContainer) {
	terminalController := controller.NewTerminalController(serviceContaienr)

	r.HandleFunc("/terminal", terminalController.Create).Methods("POST")
	r.HandleFunc("/terminal", terminalController.GetAll).Methods("GET")
	r.HandleFunc("/terminal/{id}", terminalController.GetById).Methods("GET")
	r.HandleFunc("/terminal/{id}", terminalController.Update).Methods("PUT")
	r.HandleFunc("/terminal/{id}", terminalController.Delete).Methods("DELETE")
}
