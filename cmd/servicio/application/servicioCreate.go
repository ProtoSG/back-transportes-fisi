package application

import (
	"back-transportes-fisi/cmd/servicio/domain"
	"back-transportes-fisi/cmd/servicio/domain/repository"
)

type ServicioCreate struct {
	repository repository.ServicioRepository
}

func NewServicioCreate(repository repository.ServicioRepository) *ServicioCreate {
	return &ServicioCreate{repository}
}

func (this *ServicioCreate) Execute(idServicio int, nombre, descripcion string, precio float32) error {
	servicio := domain.NewServicio(idServicio, nombre, descripcion, precio)
	return this.repository.Create(servicio)
}
