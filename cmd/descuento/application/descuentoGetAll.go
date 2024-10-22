package application

import (
	"back-transportes-fisi/cmd/descuento/domain"
	"back-transportes-fisi/cmd/descuento/domain/repository"
)

type DescuentoGetAll struct {
	repository repository.DescuentoRepository
}

func NewDescuentoGetAll(repository repository.DescuentoRepository) *DescuentoGetAll {
	return &DescuentoGetAll{repository: repository}
}

func (d *DescuentoGetAll) Execute() ([]*domain.Descuento, error) {
	return d.repository.GetAll()
}
