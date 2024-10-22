package controller

import (
	"back-transportes-fisi/cmd/shared/infrastructure"
	"back-transportes-fisi/cmd/viaje/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ViajeController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewViajeController(serviceContainer *infrastructure.ServiceContainer) *ViajeController {
	return &ViajeController{serviceContainer}
}

func (this *ViajeController) Create(w http.ResponseWriter, r *http.Request) {
	viaje := domain.Viaje{}
	if err := json.NewDecoder(r.Body).Decode(&viaje); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Viaje.Create.Execute(
		viaje.IdViaje,
		viaje.IdRuta,
		viaje.FechaSalida,
		viaje.HoraEmbarque,
		viaje.IdTerminalEmbarque,
		viaje.HoraDesembarque,
		viaje.IdTerminalDesembarque,
		viaje.Estado,
		viaje.IdServicio,
		viaje.IdBus,
		viaje.PrecioViaje,
		viaje.AsientosDisponibles,
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (this *ViajeController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	viajes, err := this.serviceContainer.Viaje.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(viajes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *ViajeController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	viaje, err := this.serviceContainer.Viaje.GetById.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(viaje); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *ViajeController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var viaje domain.Viaje
	if err := json.NewDecoder(r.Body).Decode(&viaje); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Viaje.Update.Execute(id, viaje.IdRuta, viaje.FechaSalida, viaje.HoraEmbarque, viaje.IdTerminalEmbarque, viaje.HoraDesembarque, viaje.IdTerminalDesembarque, viaje.Estado, viaje.IdServicio, viaje.IdBus, viaje.PrecioViaje, viaje.AsientosDisponibles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *ViajeController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := this.serviceContainer.Viaje.Delete.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (this *ViajeController) GetByOriginDestinationAndDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	origen := r.URL.Query().Get("origen")
	destino := r.URL.Query().Get("destino")
	fecha := r.URL.Query().Get("fecha")

	viajes, err := this.serviceContainer.Viaje.GetByOriginDestinationAndDate.Execute(origen, destino, fecha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(viajes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
