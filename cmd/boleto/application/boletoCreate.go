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

func (b *BoletoCreate) Execute(idBoleto, idPasajero, idViaje, idAsiento, idTransaccion int, subTotal float32, fecha string) error {
	boleto := domain.NewBoleto(idBoleto, idPasajero, idViaje, idAsiento, idTransaccion, subTotal, fecha)

	return b.repository.Create(boleto)
}
