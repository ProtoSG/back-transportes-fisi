package application

import (
	"back-transportes-fisi/cmd/asiento/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type AsientoDelete struct {
	repository repository.AsientoRepository
}

func NewAsientoDelete(repository repository.AsientoRepository) *AsientoDelete {
	return &AsientoDelete{repository: repository}
}

func (a *AsientoDelete) Execute(idAsiento int) error {
	if asiento, _ := a.repository.GetById(idAsiento); asiento == nil {
		return domain.NewEntityNotFound(idAsiento, "Asiento")
	}

	return a.repository.Delete(idAsiento)
}
