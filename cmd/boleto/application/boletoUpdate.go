package application

import (
	"back-transportes-fisi/cmd/boleto/domain"
	"back-transportes-fisi/cmd/boleto/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type BoletoUpdate struct {
	repository repository.BoletoRepository
}

func NewBoletoUpdate(repository repository.BoletoRepository) *BoletoUpdate {
	return &BoletoUpdate{repository: repository}
}

func (b *BoletoUpdate) Execute(id int, idPasajero int, idViaje int, idAsiento int, idTransaccion int, fecha int, subTotal int) error {
	if boleto, _ := b.repository.GetById(id); boleto == nil {
		return domainShared.NewEntityNotFound(id, "Boleto")
	}

	newBoleto := domain.NewBoleto(id, idPasajero, idViaje, idAsiento, idTransaccion, fecha, subTotal)
	return b.repository.Update(newBoleto)
}
