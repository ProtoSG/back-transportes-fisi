package application

import (
	"back-transportes-fisi/cmd/servicio/domain"
	"back-transportes-fisi/cmd/servicio/domain/repository"
)

type ServicioGetAll struct {
	repository repository.ServicioRepository
}

func NewServicioGetAll(repository repository.ServicioRepository) *ServicioGetAll {
	return &ServicioGetAll{repository}
}

func (this *ServicioGetAll) Execute() ([]*domain.Servicio, error) {
	return this.repository.GetAll()
}
