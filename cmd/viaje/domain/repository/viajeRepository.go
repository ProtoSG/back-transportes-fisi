package repository

import "back-transportes-fisi/cmd/viaje/domain"

type ViajeRepository interface {
	Create(viaje *domain.Viaje) error
	GetAll() ([]*domain.Viaje, error)
	GetById(id int) (*domain.Viaje, error)
	Update(viaje *domain.Viaje) error
	Delete(id int) error
	GetByOriginDestinationAndDate(origen, destino, fecha string) ([]*domain.ViajeOriginDestinationAndDate, error)
}
