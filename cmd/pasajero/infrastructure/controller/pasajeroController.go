package controller

import (
	"back-transportes-fisi/cmd/pasajero/domain"
	"back-transportes-fisi/cmd/shared/infrastructure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PasajeroController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewPasajeroController(serviceContainer *infrastructure.ServiceContainer) *PasajeroController {
	return &PasajeroController{serviceContainer}
}

func (this *PasajeroController) Create(w http.ResponseWriter, r *http.Request) {
	pasajero := domain.Pasajero{}

	if err := json.NewDecoder(r.Body).Decode(&pasajero); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Pasajero.Create.Execute(pasajero.IdPasajero, pasajero.Nombre, pasajero.Apellido, pasajero.FechaNacimiento, pasajero.Sexo, pasajero.Dni); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (this *PasajeroController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pasajeros, err := this.serviceContainer.Pasajero.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(pasajeros); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *PasajeroController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	pasajero, err := this.serviceContainer.Pasajero.GetById.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(pasajero); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *PasajeroController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	pasajero := domain.Pasajero{}
	if err := json.NewDecoder(r.Body).Decode(&pasajero); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Pasajero.Update.Execute(id, pasajero.Nombre, pasajero.Apellido, pasajero.FechaNacimiento, pasajero.Sexo, pasajero.Dni); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *PasajeroController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := this.serviceContainer.Pasajero.Delete.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
