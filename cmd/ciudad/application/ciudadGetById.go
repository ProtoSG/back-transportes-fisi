package application

import (
	"back-transportes-fisi/cmd/ciudad/domain"
	"back-transportes-fisi/cmd/ciudad/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type CiudadGetById struct {
	repository repository.CiudadRepository
}

func NewCiudadGetById(repository repository.CiudadRepository) *CiudadGetById {
	return &CiudadGetById{repository: repository}
}

func (this *CiudadGetById) Execute(id int) (*domain.Ciudad, error) {
	ciudad, _ := this.repository.GetById(id)
	if ciudad == nil {
		return nil, domainShared.NewEntityNotFound(id, "Ciudad")
	}

	return ciudad, nil
}
