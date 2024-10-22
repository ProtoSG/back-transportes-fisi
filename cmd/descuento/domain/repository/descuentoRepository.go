package repository

import "back-transportes-fisi/cmd/descuento/domain"

type DescuentoRepository interface {
	GetAll() ([]*domain.Descuento, error)
	GetById(id int) (*domain.Descuento, error)
	Create(descuento *domain.Descuento) error
	Update(descuento *domain.Descuento) error
	Delete(id int) error
}
