package repository

import "back-transportes-fisi/cmd/servicio/domain"

type ServicioRepository interface {
	Create(servicio *domain.Servicio) error
	GetAll() ([]*domain.Servicio, error)
	GetById(id int) (*domain.Servicio, error)
	Update(servicio *domain.Servicio) error
	Delete(id int) error
}
