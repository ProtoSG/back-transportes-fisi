package application

import (
	"back-transportes-fisi/cmd/asiento/domain"
	"back-transportes-fisi/cmd/asiento/domain/repository"
)

type AsientoGetAll struct {
	repository repository.AsientoRepository
}

func NewAsientoGetAll(repository repository.AsientoRepository) *AsientoGetAll {
	return &AsientoGetAll{repository: repository}
}

func (a *AsientoGetAll) Execute() ([]*domain.Asiento, error) {
	return a.repository.GetAll()
}
