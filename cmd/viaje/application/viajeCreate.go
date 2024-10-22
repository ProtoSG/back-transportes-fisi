package application

import (
	"back-transportes-fisi/cmd/viaje/domain"
	"back-transportes-fisi/cmd/viaje/domain/repository"
)

type ViajeCreate struct {
	repository repository.ViajeRepository
}

func NewViajeCreate(repository repository.ViajeRepository) *ViajeCreate {
	return &ViajeCreate{repository}
}

func (this *ViajeCreate) Execute(
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
	viaje := domain.NewViaje(
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
	return this.repository.Create(viaje)
}
