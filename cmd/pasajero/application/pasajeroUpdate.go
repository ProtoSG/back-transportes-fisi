package application

import (
	"back-transportes-fisi/cmd/pasajero/domain"
	"back-transportes-fisi/cmd/pasajero/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type PasajeroUpdate struct {
	repository repository.PasajeroRepository
}

func NewPasajeroUpdate(repository repository.PasajeroRepository) *PasajeroUpdate {
	return &PasajeroUpdate{repository}
}

func (p *PasajeroUpdate) Execute(id int, nombre, apellido, fechaNaciemiento, sexo, dni string) error {
	if pasajero, _ := p.repository.GetById(id); pasajero != nil {
		return domainShared.NewEntityNotFound(id, "Pasajero")
	}

	pasajero := domain.NewPasajero(id, nombre, apellido, fechaNaciemiento, sexo, dni)
	return p.repository.Update(pasajero)
}
