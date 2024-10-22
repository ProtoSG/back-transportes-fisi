package application

import (
	"back-transportes-fisi/cmd/ruta/domain"
	"back-transportes-fisi/cmd/ruta/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type RutaGetById struct {
	repository repository.RutaRepository
}

func NewRutaGetById(repository repository.RutaRepository) *RutaGetById {
	return &RutaGetById{repository}
}

func (this *RutaGetById) Execute(id int) (*domain.Ruta, error) {
	ruta, _ := this.repository.GetById(id)

	if ruta == nil {
		return nil, domainShared.NewEntityNotFound(id, "Ruta")
	}

	return ruta, nil
}
