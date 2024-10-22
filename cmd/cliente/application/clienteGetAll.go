package application

import (
	"back-transportes-fisi/cmd/cliente/domain"
	"back-transportes-fisi/cmd/cliente/domain/repository"
)

type ClienteGetAll struct {
	repository repository.ClienteRepository
}

func NewClienteGetAll(repository repository.ClienteRepository) *ClienteGetAll {
	return &ClienteGetAll{repository: repository}
}

func (c *ClienteGetAll) Execute() ([]*domain.Cliente, error) {
	return c.repository.GetAll()
}
