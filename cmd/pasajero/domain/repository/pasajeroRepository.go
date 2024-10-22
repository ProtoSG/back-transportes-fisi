package repository

import "back-transportes-fisi/cmd/pasajero/domain"

type PasajeroRepository interface {
	Create(pasajero *domain.Pasajero) error
	GetAll() ([]*domain.Pasajero, error)
	GetById(id int) (*domain.Pasajero, error)
	Update(pasajero *domain.Pasajero) error
	Delete(id int) error
}
