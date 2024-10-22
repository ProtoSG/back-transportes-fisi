package application

import (
	"back-transportes-fisi/cmd/ruta/domain"
	"back-transportes-fisi/cmd/ruta/domain/repository"
)

type RutaCreate struct {
	repository repository.RutaRepository
}

func NewRutaCreate(repository repository.RutaRepository) *RutaCreate {
	return &RutaCreate{repository}
}

func (c *RutaCreate) Execute(idRuta, idCiudadOrigen, idCiudadDestino int, precio float32) error {
	ruta := domain.NewRuta(idRuta, idCiudadOrigen, idCiudadDestino, precio)
	return c.repository.Create(ruta)
}
