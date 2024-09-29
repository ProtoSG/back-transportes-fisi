package application

import (
	"back-transportes-fisi/cmd/ciudad/domain"
	"back-transportes-fisi/cmd/ciudad/domain/repository"
)

type CiudadGetAll struct {
	repository repository.CiudadRepository
}

func NewCiudadGetAll(repository repository.CiudadRepository) *CiudadGetAll {
	return &CiudadGetAll{repository: repository}
}

func (this *CiudadGetAll) Execute() ([]*domain.Ciudad, error) {
	return this.repository.GetAll()
}
