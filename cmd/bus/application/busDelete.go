package application

import (
	"back-transportes-fisi/cmd/bus/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type BusDelete struct {
	repository repository.BusRepository
}

func NewBusDelete(repository repository.BusRepository) *BusDelete {
	return &BusDelete{
		repository: repository,
	}
}

func (b *BusDelete) Execute(id int) error {
	if bus, _ := b.repository.GetById(id); bus == nil {
		return domain.NewEntityNotFound(id, "Bus")
	}

	return b.repository.Delete(id)
}
