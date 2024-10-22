package repository

import "back-transportes-fisi/cmd/cliente/domain"

type ClienteRepository interface {
	GetAll() ([]*domain.Cliente, error)
	GetById(id int) (*domain.Cliente, error)
	Create(cliente *domain.Cliente) error
	Update(cliente *domain.Cliente) error
	Delete(id int) error
}
