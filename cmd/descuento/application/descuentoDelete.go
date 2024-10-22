package application

import (
	"back-transportes-fisi/cmd/descuento/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type DescuentoDelete struct {
	repository repository.DescuentoRepository
}

func NewDescuentoDelete(repository repository.DescuentoRepository) *DescuentoDelete {
	return &DescuentoDelete{repository: repository}
}

func (d *DescuentoDelete) Execute(id int) error {
	if descuento, _ := d.repository.GetById(id); descuento == nil {
		return domain.NewEntityNotFound(id, "Descuento")
	}

	return d.repository.Delete(id)
}
