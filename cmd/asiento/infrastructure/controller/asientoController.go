package controller

import (
	"back-transportes-fisi/cmd/asiento/domain"
	"back-transportes-fisi/cmd/shared/infrastructure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AsientoController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewAsientoController(serviceContainer *infrastructure.ServiceContainer) *AsientoController {
	return &AsientoController{serviceContainer: serviceContainer}
}

func (this *AsientoController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var asiento domain.Asiento
	if err := json.NewDecoder(r.Body).Decode(&asiento); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Asiento.Create.Execute(asiento.IdAsiento, asiento.IdBus, asiento.NumeroAsiento, asiento.Piso, asiento.Precio, asiento.Disponibilidad); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (this *AsientoController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	asientos, err := this.serviceContainer.Asiento.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(asientos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *AsientoController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	asiento, err := this.serviceContainer.Asiento.GetById.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(asiento); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *AsientoController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var asiento domain.Asiento
	if err := json.NewDecoder(r.Body).Decode(&asiento); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Asiento.Update.Execute(id, asiento.IdBus, asiento.NumeroAsiento, asiento.Piso, asiento.Precio, asiento.Disponibilidad); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *AsientoController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := this.serviceContainer.Asiento.Delete.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (this *AsientoController) GetByIdBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	asientos, err := this.serviceContainer.Asiento.GetByIdBus.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(asientos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
