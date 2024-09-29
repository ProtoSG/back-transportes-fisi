package application

import (
	"back-transportes-fisi/cmd/asiento/domain"
	"back-transportes-fisi/cmd/asiento/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type AsientoUpdate struct {
	repository repository.AsientoRepository
}

func NewAsientoUpdate(repository repository.AsientoRepository) *AsientoUpdate {
	return &AsientoUpdate{repository: repository}
}

func (a *AsientoUpdate) Execute(idAsiento int, idBus int, numeroAsiento int, piso int, precio int, disponibilidad int) error {
	asiento, _ := a.repository.GetById(idAsiento)
	if asiento != nil {
		return domainShared.NewEntityNotFound(idAsiento, "Asiento")
	}

	newAsiento := domain.NewAsiento(asiento.IdAsiento, asiento.IdBus, asiento.NumeroAsiento, asiento.Piso, asiento.Precio, asiento.Disponibilidad)

	return a.repository.Update(newAsiento)
}
