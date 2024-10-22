package application

import (
	"back-transportes-fisi/cmd/ruta/domain"
	"back-transportes-fisi/cmd/ruta/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type RutaUpdate struct {
	repository repository.RutaRepository
}

func NewRutaUpdate(repository repository.RutaRepository) *RutaUpdate {
	return &RutaUpdate{repository}
}

func (this *RutaUpdate) Execute(idRuta, idCiudadOrigen, idCiudadDestino int, precio float32) error {
	if ruta, _ := this.repository.GetById(idRuta); ruta == nil {
		return domainShared.NewEntityNotFound(idRuta, "Ruta")
	}

	ruta := domain.NewRuta(idRuta, idCiudadOrigen, idCiudadDestino, precio)
	return this.repository.Update(ruta)
}
