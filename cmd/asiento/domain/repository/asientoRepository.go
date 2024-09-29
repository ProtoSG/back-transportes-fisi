package repository

import "back-transportes-fisi/cmd/asiento/domain"

type AsientoRepository interface {
	GetAll() ([]*domain.Asiento, error)
	GetById(id int) (*domain.Asiento, error)
	Create(asiento *domain.Asiento) error
	Update(asiento *domain.Asiento) error
	Delete(id int) error
}
