package application

import (
	"back-transportes-fisi/cmd/pasajero/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type PasajeroDelete struct {
	repository repository.PasajeroRepository
}

func NewPasajeroDelete(repository repository.PasajeroRepository) *PasajeroDelete {
	return &PasajeroDelete{repository}
}

func (p *PasajeroDelete) Execute(id int) error {
	if pasajero, _ := p.repository.GetById(id); pasajero == nil {
		return domain.NewEntityNotFound(id, "Pasajero")
	}

	return p.repository.Delete(id)
}
