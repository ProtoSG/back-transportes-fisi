package repository

import "back-transportes-fisi/cmd/bus/domain"

type BusRepository interface {
	GetAll() ([]*domain.Bus, error)
	GetById(id int) (*domain.Bus, error)
	Create(bus *domain.Bus) error
	Update(bus *domain.Bus) error
	Delete(id int) error
}
