package application

import (
	"back-transportes-fisi/cmd/transaccion/domain"
	"back-transportes-fisi/cmd/transaccion/domain/repository"
)

type TransaccionCreate struct {
	repository repository.TransaccionRepository
}

func NewTransaccionCreate(repository repository.TransaccionRepository) *TransaccionCreate {
	return &TransaccionCreate{repository}
}

func (this *TransaccionCreate) Execute(idTransaccion, idCliente, idDescuento int, total float32, metodoPago, estado, fechaTransaccion string) error {
	transaccion := domain.NewTransaccion(idTransaccion, idCliente, idDescuento, total, metodoPago, estado, fechaTransaccion)
	return this.repository.Create(transaccion)
}
