package application

import (
	"back-transportes-fisi/cmd/asiento/domain"
	"back-transportes-fisi/cmd/asiento/domain/repository"
)

type AsientoGetByIdBus struct {
	repository repository.AsientoRepository
}

func NewAsientoGetByIdBus(repository repository.AsientoRepository) *AsientoGetByIdBus {
	return &AsientoGetByIdBus{repository: repository}
}

func (this *AsientoGetByIdBus) Execute(id int) ([]*domain.Asiento, error) {
	return this.repository.GetByIdBus(id)
}
