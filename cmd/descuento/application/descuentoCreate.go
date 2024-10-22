package application

import (
	"back-transportes-fisi/cmd/descuento/domain"
	"back-transportes-fisi/cmd/descuento/domain/repository"
)

type DescuentoCreate struct {
	repository repository.DescuentoRepository
}

func NewDescuentoCreate(repository repository.DescuentoRepository) *DescuentoCreate {
	return &DescuentoCreate{repository: repository}
}

func (d *DescuentoCreate) Execute(idDescuento int, codigo string, monto float32) error {
	descuento := domain.NewDescuento(idDescuento, codigo, monto)
	return d.repository.Create(descuento)
}
