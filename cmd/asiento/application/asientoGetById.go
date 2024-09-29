package application

import (
	"back-transportes-fisi/cmd/asiento/domain"
	"back-transportes-fisi/cmd/asiento/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type AsientoGetById struct {
	repository repository.AsientoRepository
}

func NewAsientoGetById(repository repository.AsientoRepository) *AsientoGetById {
	return &AsientoGetById{repository: repository}
}

func (a *AsientoGetById) Execute(id int) (*domain.Asiento, error) {
	asiento, _ := a.repository.GetById(id)
	if asiento == nil {
		return nil, domainShared.NewEntityNotFound(id, "Asiento")
	}

	return asiento, nil
}
