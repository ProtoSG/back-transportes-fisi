package application

import (
	"back-transportes-fisi/cmd/boleto/domain"
	"back-transportes-fisi/cmd/boleto/domain/repository"
)

type BoletoCreate struct {
	repository repository.BoletoRepository
}

func NewBoletoCreate(repository repository.BoletoRepository) *BoletoCreate {
	return &BoletoCreate{repository: repository}
}

func (b *BoletoCreate) Execute(idBoleto int, idPasajero int, idViaje int, idAsiento int, idTransaccion int, fecha int, subTotal int) error {
	boleto := domain.NewBoleto(idBoleto, idPasajero, idViaje, idAsiento, idTransaccion, fecha, subTotal)

	return b.repository.Create(boleto)
}
