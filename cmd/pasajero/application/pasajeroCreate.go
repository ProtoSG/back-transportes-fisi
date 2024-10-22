package application

import (
	"back-transportes-fisi/cmd/pasajero/domain"
	"back-transportes-fisi/cmd/pasajero/domain/repository"
)

type PasajeroCreate struct {
	repository repository.PasajeroRepository
}

func NewPasajeroCreate(repository repository.PasajeroRepository) *PasajeroCreate {
	return &PasajeroCreate{repository: repository}
}

func (p *PasajeroCreate) Execute(idPasajero int, nombre, apellido, fechaNacimiento, sexo, dni string) error {
	pasajero := domain.NewPasajero(idPasajero, nombre, apellido, fechaNacimiento, sexo, dni)
	return p.repository.Create(pasajero)
}
