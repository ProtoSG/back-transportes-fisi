package controller

import (
	"back-transportes-fisi/cmd/servicio/domain"
	"back-transportes-fisi/cmd/shared/infrastructure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ServicioController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewServicioController(serviceContainer *infrastructure.ServiceContainer) *ServicioController {
	return &ServicioController{serviceContainer: serviceContainer}
}

func (this *ServicioController) Create(w http.ResponseWriter, r *http.Request) {
	servicio := domain.Servicio{}
	if err := json.NewDecoder(r.Body).Decode(&servicio); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Servicio.Create.Execute(servicio.IdServicio, servicio.Nombre, servicio.Descripcion, servicio.Precio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *ServicioController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	servicios, err := this.serviceContainer.Servicio.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(servicios); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *ServicioController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idServicio, _ := strconv.Atoi(vars["id"])

	servicio, err := this.serviceContainer.Servicio.GetById.Execute(idServicio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(servicio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *ServicioController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idServicio, _ := strconv.Atoi(vars["id"])

	servicio := domain.Servicio{}
	if err := json.NewDecoder(r.Body).Decode(&servicio); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Servicio.Update.Execute(idServicio, servicio.Nombre, servicio.Descripcion, servicio.Precio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *ServicioController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idServicio, _ := strconv.Atoi(vars["id"])

	if err := this.serviceContainer.Servicio.Delete.Execute(idServicio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
