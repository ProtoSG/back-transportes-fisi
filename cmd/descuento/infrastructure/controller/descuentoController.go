package controller

import (
	"back-transportes-fisi/cmd/descuento/domain"
	"back-transportes-fisi/cmd/shared/infrastructure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DescuentoController struct {
	ServiceContainer *infrastructure.ServiceContainer
}

func NewDescuentoController(serviceContainer *infrastructure.ServiceContainer) *DescuentoController {
	return &DescuentoController{ServiceContainer: serviceContainer}
}

func (c *DescuentoController) Create(w http.ResponseWriter, r *http.Request) {
	bus := domain.Descuento{}
	if err := json.NewDecoder(r.Body).Decode(&bus); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.ServiceContainer.Descuento.Create.Execute(bus.IdDescuento, bus.Codigo, bus.Monto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *DescuentoController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	descuentos, err := c.ServiceContainer.Descuento.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(descuentos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *DescuentoController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idDescuento, _ := strconv.Atoi(vars["id"])

	descuento, err := c.ServiceContainer.Descuento.GetById.Execute(idDescuento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(descuento); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *DescuentoController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idDescuento, _ := strconv.Atoi(vars["id"])

	var descuento domain.Descuento
	if err := json.NewDecoder(r.Body).Decode(&descuento); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.ServiceContainer.Descuento.Update.Execute(idDescuento, descuento.Codigo, descuento.Monto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *DescuentoController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idDescuento, _ := strconv.Atoi(vars["id"])

	if err := c.ServiceContainer.Descuento.Delete.Execute(idDescuento); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
