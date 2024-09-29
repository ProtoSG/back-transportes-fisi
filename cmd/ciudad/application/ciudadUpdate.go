package application

import (
	"back-transportes-fisi/cmd/ciudad/domain"
	"back-transportes-fisi/cmd/ciudad/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type CiudadUpdate struct {
	repository repository.CiudadRepository
}

func NewCiudadUpdate(repository repository.CiudadRepository) *CiudadUpdate {
	return &CiudadUpdate{repository: repository}
}

func (this *CiudadUpdate) Execute(id int, nombre string) error {
	if ciudad, _ := this.repository.GetById(id); ciudad == nil {
		return domainShared.NewEntityNotFound(id, "Ciudad")
	}

	newCiudad := domain.NewCiudad(id, nombre)

	return this.repository.Update(newCiudad)
}
