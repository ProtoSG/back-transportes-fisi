package application

import (
	"back-transportes-fisi/cmd/ciudad/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type CiudadDelete struct {
	repository repository.CiudadRepository
}

func NewCiudadDelete(repository repository.CiudadRepository) *CiudadDelete {
	return &CiudadDelete{repository: repository}
}

func (this *CiudadDelete) Execute(id int) error {
	if ciudad, _ := this.repository.GetById(id); ciudad == nil {
		return domain.NewEntityNotFound(id, "Ciudad")
	}

	return this.repository.Delete(id)
}
