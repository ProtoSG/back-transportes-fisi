package application

import (
	"back-transportes-fisi/cmd/bus/domain"
	"back-transportes-fisi/cmd/bus/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type BusUpdate struct {
	repository repository.BusRepository
}

func NewBusUpdate(repository repository.BusRepository) *BusUpdate {
	return &BusUpdate{repository: repository}
}

func (b *BusUpdate) Execute(id int, capacidad int, piso int) error {
	if bus, _ := b.repository.GetById(id); bus == nil {
		return domainShared.NewEntityNotFound(id, "Bus")
	}

	newBus := domain.NewBus(id, capacidad, piso)

	return b.repository.Update(newBus)
}
