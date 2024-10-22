package controller

import (
	"back-transportes-fisi/cmd/shared/infrastructure"
	"back-transportes-fisi/cmd/terminal/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TerminalController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewTerminalController(serviceContainer *infrastructure.ServiceContainer) *TerminalController {
	return &TerminalController{serviceContainer}
}

func (this *TerminalController) Create(w http.ResponseWriter, r *http.Request) {
	terminal := domain.Terminal{}
	if err := json.NewDecoder(r.Body).Decode(&terminal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Terminal.Create.Execute(terminal.IdTerminal, terminal.Nombre, terminal.Ubicacion); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (this *TerminalController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	terminal, err := this.serviceContainer.Terminal.GetAll.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(terminal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *TerminalController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	terminal, err := this.serviceContainer.Terminal.GetById.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(terminal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (this *TerminalController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	terminal := domain.Terminal{}
	if err := json.NewDecoder(r.Body).Decode(&terminal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Terminal.Update.Execute(id, terminal.Nombre, terminal.Ubicacion); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (this *TerminalController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Terminal.Delete.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
