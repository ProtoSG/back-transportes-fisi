package application

import (
	"back-transportes-fisi/cmd/cliente/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type ClienteDelete struct {
	repository repository.ClienteRepository
}

func NewClienteDelete(repository repository.ClienteRepository) *ClienteDelete {
	return &ClienteDelete{repository: repository}
}

func (c *ClienteDelete) Execute(id int) error {
	if cliente, _ := c.repository.GetById(id); cliente == nil {
		return domain.NewEntityNotFound(id, "Cliente")
	}
	return c.repository.Delete(id)
}
