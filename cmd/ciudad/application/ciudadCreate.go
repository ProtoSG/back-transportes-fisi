package application

import (
	"back-transportes-fisi/cmd/ciudad/domain"
	"back-transportes-fisi/cmd/ciudad/domain/repository"
)

type CiudadCreate struct {
	repository repository.CiudadRepository
}

func NewCiudadCreate(repository repository.CiudadRepository) *CiudadCreate {
	return &CiudadCreate{repository: repository}
}

func (this *CiudadCreate) Execute(id int, nombre string) error {
	ciudad := domain.NewCiudad(id, nombre)
	return this.repository.Create(ciudad)
}
