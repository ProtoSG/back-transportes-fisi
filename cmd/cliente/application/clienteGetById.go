package application

import (
	"back-transportes-fisi/cmd/cliente/domain"
	"back-transportes-fisi/cmd/cliente/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type ClienteGetById struct {
	repository repository.ClienteRepository
}

func NewClienteGetById(repository repository.ClienteRepository) *ClienteGetById {
	return &ClienteGetById{repository: repository}
}

func (c *ClienteGetById) Execute(id int) (*domain.Cliente, error) {
	cliente, _ := c.repository.GetById(id)

	if cliente == nil {
		return nil, domainShared.NewEntityNotFound(id, "Cliente")
	}

	return cliente, nil
}
