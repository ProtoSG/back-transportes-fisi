package application

import (
	"back-transportes-fisi/cmd/cliente/domain"
	"back-transportes-fisi/cmd/cliente/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type ClienteUpdate struct {
	repository repository.ClienteRepository
}

func NewClienteUpdate(repository repository.ClienteRepository) *ClienteUpdate {
	return &ClienteUpdate{repository}
}

func (c *ClienteUpdate) Execute(idCliente int, nombre string, apellido string, fechaNacimiento string, sexo string, dni string, email string, telefone string) error {
	if cliente, _ := c.repository.GetById(idCliente); cliente != nil {
		return domainShared.NewEntityNotFound(idCliente, "Cliente")
	}

	newCliente := domain.NewCliente(idCliente, nombre, apellido, fechaNacimiento, sexo, dni, email, telefone)

	return c.repository.Update(newCliente)
}
