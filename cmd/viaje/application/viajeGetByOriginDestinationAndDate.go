package application

import (
	"back-transportes-fisi/cmd/viaje/domain"
	"back-transportes-fisi/cmd/viaje/domain/repository"
)

type ViajeGetByOriginDestinationAndDate struct {
	repository repository.ViajeRepository
}

func NewViajeGetByOriginDestinationAndDate(repository repository.ViajeRepository) *ViajeGetByOriginDestinationAndDate {
	return &ViajeGetByOriginDestinationAndDate{repository}
}

func (this *ViajeGetByOriginDestinationAndDate) Execute(origen, destino, fecha string) ([]*domain.ViajeOriginDestinationAndDate, error) {
	return this.repository.GetByOriginDestinationAndDate(origen, destino, fecha)
}
