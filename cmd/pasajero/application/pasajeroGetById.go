package application

import (
	"back-transportes-fisi/cmd/pasajero/domain"
	"back-transportes-fisi/cmd/pasajero/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type PasajeroGetById struct {
	repository repository.PasajeroRepository
}

func NewPasajeroGetById(repository repository.PasajeroRepository) *PasajeroGetById {
	return &PasajeroGetById{repository}
}

func (p *PasajeroGetById) Execute(id int) (*domain.Pasajero, error) {
	pasajero, _ := p.repository.GetById(id)
	if pasajero == nil {
		return nil, domainShared.NewEntityNotFound(id, "Pasajero")
	}

	return pasajero, nil
}
