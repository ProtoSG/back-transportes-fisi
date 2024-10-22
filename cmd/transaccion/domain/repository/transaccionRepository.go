package repository

import "back-transportes-fisi/cmd/transaccion/domain"

type TransaccionRepository interface {
	Create(transaccion *domain.Transaccion) error
	// GetAll() ([]*domain.Transaccion, error)
	// GetById(id int) (*domain.Transaccion, error)
	// Update(transaccion *domain.Transaccion) error
	// Delete(id int) error
}
