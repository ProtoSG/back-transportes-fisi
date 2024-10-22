package repository

import "back-transportes-fisi/cmd/terminal/domain"

type TerminalRepository interface {
	Create(terminal *domain.Terminal) error
	GetAll() ([]*domain.Terminal, error)
	GetById(id int) (*domain.Terminal, error)
	Update(terminal *domain.Terminal) error
	Delete(id int) error
}
