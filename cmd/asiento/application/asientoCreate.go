package application

import (
	"back-transportes-fisi/cmd/asiento/domain"
	"back-transportes-fisi/cmd/asiento/domain/repository"
)

type AsientoCreate struct {
	repository repository.AsientoRepository
}

func NewAsientoCreate(repository repository.AsientoRepository) *AsientoCreate {
	return &AsientoCreate{repository: repository}
}

func (a *AsientoCreate) Execute(idAsiento int, idBus int, numeroAsiento int, piso int, precio int, disponibilidad int) error {
	asiento := domain.NewAsiento(idAsiento, idBus, numeroAsiento, piso, precio, disponibilidad)
	return a.repository.Create(asiento)
}
