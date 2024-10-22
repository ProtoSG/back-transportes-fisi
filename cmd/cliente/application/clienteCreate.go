package application

import (
	"back-transportes-fisi/cmd/cliente/domain"
	"back-transportes-fisi/cmd/cliente/domain/repository"
)

type ClienteCreate struct {
	repository repository.ClienteRepository
}

func NewClienteCreate(repository repository.ClienteRepository) *ClienteCreate {
	return &ClienteCreate{repository: repository}
}

func (c *ClienteCreate) Execute(idCliente int, nombre string, apellido string, fechaNacimiento string, sexo string, dni string, email string, telefone string) error {
	cliente := domain.NewCliente(idCliente, nombre, apellido, fechaNacimiento, sexo, dni, email, telefone)
	return c.repository.Create(cliente)
}
