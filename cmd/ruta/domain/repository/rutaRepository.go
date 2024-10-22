package repository

import "back-transportes-fisi/cmd/ruta/domain"

type RutaRepository interface {
	Create(ruta *domain.Ruta) error
	GetAll() ([]*domain.Ruta, error)
	GetById(id int) (*domain.Ruta, error)
	Update(ruta *domain.Ruta) error
	Delete(id int) error
}
