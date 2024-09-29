package application

import (
	"back-transportes-fisi/cmd/bus/domain"
	"back-transportes-fisi/cmd/bus/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type BusGetById struct {
	repository repository.BusRepository
}

func NewBusGetById(repository repository.BusRepository) *BusGetById {
	return &BusGetById{repository: repository}
}

func (b *BusGetById) Execute(id int) (*domain.Bus, error) {

	bus, err := b.repository.GetById(id)

	if err != nil {
		return nil, domainShared.NewEntityNotFound(id, "Bus")
	}

	return bus, nil
}
