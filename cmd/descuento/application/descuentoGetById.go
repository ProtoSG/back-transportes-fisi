package application

import (
	"back-transportes-fisi/cmd/descuento/domain"
	"back-transportes-fisi/cmd/descuento/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type DescuentoGetById struct {
	repository repository.DescuentoRepository
}

func NewDescuentoGetById(repository repository.DescuentoRepository) *DescuentoGetById {
	return &DescuentoGetById{repository: repository}
}

func (d *DescuentoGetById) Execute(id int) (*domain.Descuento, error) {
	descuento, _ := d.repository.GetById(id)
	if descuento == nil {
		return nil, domainShared.NewEntityNotFound(id, "Descuneto")
	}

	return descuento, nil
}
