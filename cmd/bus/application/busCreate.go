package application

import (
	"back-transportes-fisi/cmd/bus/domain"
	"back-transportes-fisi/cmd/bus/domain/repository"
)

type BusCreate struct {
	repository repository.BusRepository
}

func NewBusCreate(repository repository.BusRepository) *BusCreate {
	return &BusCreate{
		repository: repository,
	}
}

func (b *BusCreate) Execute(idBus int, capacidad int, piso int) error {
	bus := domain.NewBus(idBus, capacidad, piso)
	return b.repository.Create(bus)
}
