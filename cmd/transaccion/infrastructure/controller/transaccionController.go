package controller

import (
	"back-transportes-fisi/cmd/shared/infrastructure"
	"back-transportes-fisi/cmd/transaccion/domain"
	"encoding/json"
	"net/http"
)

type TransaccionController struct {
	serviceContainer *infrastructure.ServiceContainer
}

func NewTransaccionController(serviceContainer *infrastructure.ServiceContainer) *TransaccionController {
	return &TransaccionController{serviceContainer}
}

func (this *TransaccionController) Create(w http.ResponseWriter, r *http.Request) {
	transaccion := domain.Transaccion{}
	if err := json.NewDecoder(r.Body).Decode(&transaccion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := this.serviceContainer.Transaccion.Create.Execute(transaccion.IdTransaccion, transaccion.IdCliente, transaccion.IdDescuento, transaccion.Total, transaccion.MetodoPago, transaccion.Estado, transaccion.FechaTransaccion); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
