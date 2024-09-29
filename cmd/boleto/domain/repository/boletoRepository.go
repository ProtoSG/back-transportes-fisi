package repository

import "back-transportes-fisi/cmd/boleto/domain"

type BoletoRepository interface {
	Create(boleto *domain.Boleto) error
	GetAll() ([]*domain.Boleto, error)
	GetById(id int) (*domain.Boleto, error)
	Update(boleto *domain.Boleto) error
	Delete(id int) error
}
