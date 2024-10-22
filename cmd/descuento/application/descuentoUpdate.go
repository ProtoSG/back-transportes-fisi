package application

import (
	"back-transportes-fisi/cmd/descuento/domain"
	"back-transportes-fisi/cmd/descuento/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type DescuentoUpdate struct {
	repository repository.DescuentoRepository
}

func NewDescuentoUpdate(repository repository.DescuentoRepository) *DescuentoUpdate {
	return &DescuentoUpdate{repository: repository}
}

func (d *DescuentoUpdate) Execute(idDescuento int, codigo string, monto float32) error {
	if descuento, _ := d.repository.GetById(idDescuento); descuento == nil {
		return domainShared.NewEntityNotFound(idDescuento, "Descuento")
	}

	descuento := domain.NewDescuento(idDescuento, codigo, monto)
	return d.repository.Update(descuento)
}
