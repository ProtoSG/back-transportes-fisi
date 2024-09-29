package repository

import "back-transportes-fisi/cmd/ciudad/domain"

type CiudadRepository interface {
	Create(ciudad *domain.Ciudad) error
	GetAll() ([]*domain.Ciudad, error)
	GetById(id int) (*domain.Ciudad, error)
	Update(*domain.Ciudad) error
	Delete(id int) error
}
