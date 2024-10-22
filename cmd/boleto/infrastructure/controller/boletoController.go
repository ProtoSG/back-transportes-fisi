package controller

import (
	"back-transportes-fisi/cmd/boleto/domain"
	"back-transportes-fisi/cmd/shared/infrastructure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BoletoController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewBoletoController(serviceContainer *infrastructure.ServiceContainer) *BoletoController {
	return &BoletoController{serviceContainer: serviceContainer}
}

func (this *BoletoController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var boleto domain.Boleto
	if err := json.NewDecoder(r.Body).Decode(&boleto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Boleto.Create.Execute(boleto.IdBoleto, boleto.IdPasajero, boleto.IdViaje, boleto.IdAsiento, boleto.IdTransaccion, boleto.SubTotal, boleto.Fecha); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (this *BoletoController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	boletos, err := this.serviceContainer.Boleto.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(boletos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *BoletoController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	boleto, err := this.serviceContainer.Boleto.GetById.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(boleto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *BoletoController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var boleto domain.Boleto
	if err := json.NewDecoder(r.Body).Decode(&boleto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Boleto.Update.Execute(id, boleto.IdPasajero, boleto.IdViaje, boleto.IdAsiento, boleto.IdTransaccion, boleto.SubTotal, boleto.Fecha); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *BoletoController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := this.serviceContainer.Boleto.Delete.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
