package application

import (
	"back-transportes-fisi/cmd/pasajero/domain"
	"back-transportes-fisi/cmd/pasajero/domain/repository"
)

type PasajeroGetAll struct {
	repository repository.PasajeroRepository
}

func NewPasajeroGetAll(repository repository.PasajeroRepository) *PasajeroGetAll {
	return &PasajeroGetAll{repository}
}

func (p *PasajeroGetAll) Execute() ([]*domain.Pasajero, error) {
	return p.repository.GetAll()
}
