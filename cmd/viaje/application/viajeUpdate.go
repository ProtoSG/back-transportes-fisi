package application

import (
	domainShared "back-transportes-fisi/cmd/shared/domain"
	"back-transportes-fisi/cmd/viaje/domain"
	"back-transportes-fisi/cmd/viaje/domain/repository"
)

type ViajeUpdate struct {
	repository repository.ViajeRepository
}

func NewViajeUpdate(repository repository.ViajeRepository) *ViajeUpdate {
	return &ViajeUpdate{repository}
}

func (this *ViajeUpdate) Execute(
	idViaje int,
	idRuta int,
	fechaSalida string,
	horaEmbarque string,
	idTerminalEmbarque int,
	horaDesembarque string,
	idTerminalDesembarque int,
	estado string,
	idServicio int,
	idBus int,
	precioViaje float32,
	asientosDisponibles int,
) error {
	if viaje, _ := this.repository.GetById(idViaje); viaje == nil {
		return domainShared.NewEntityNotFound(idViaje, "Viaje")
	}

	newViaje := domain.NewViaje(
		idViaje,
		idRuta,
		fechaSalida,
		horaEmbarque,
		idTerminalEmbarque,
		horaDesembarque,
		idTerminalDesembarque,
		estado,
		idServicio,
		idBus,
		precioViaje,
		asientosDisponibles,
	)

	return this.repository.Update(newViaje)
}
