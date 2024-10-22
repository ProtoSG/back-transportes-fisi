package controller

import (
	"back-transportes-fisi/cmd/cliente/domain"
	"back-transportes-fisi/cmd/shared/infrastructure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ClienteController struct {
	serviceController *infrastructure.ServiceContainer
}

func NewClienteController(serviceController *infrastructure.ServiceContainer) *ClienteController {
	return &ClienteController{serviceController: serviceController}
}

func (this *ClienteController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cliente domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceController.Cliente.Create.Execute(cliente.IdCliente, cliente.Nombre, cliente.Apellido, cliente.FechaNacimiento, cliente.Sexo, cliente.Dni, cliente.Email, cliente.Telefono); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (this *ClienteController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clientes, err := this.serviceController.Cliente.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clientes)
}

func (this *ClienteController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	cliente, err := this.serviceController.Cliente.GetById.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cliente)
}

func (this *ClienteController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var cliente domain.Cliente
	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceController.Cliente.Update.Execute(id, cliente.Nombre, cliente.Apellido, cliente.FechaNacimiento, cliente.Sexo, cliente.Dni, cliente.Email, cliente.Telefono); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *ClienteController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := this.serviceController.Cliente.Delete.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
