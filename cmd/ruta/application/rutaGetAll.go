package application

import (
	"back-transportes-fisi/cmd/ruta/domain"
	"back-transportes-fisi/cmd/ruta/domain/repository"
)

type RutaGetAll struct {
	repository repository.RutaRepository
}

func NewRutaGetAll(repository repository.RutaRepository) *RutaGetAll {
	return &RutaGetAll{repository}
}

func (r *RutaGetAll) Execute() ([]*domain.Ruta, error) {
	return r.repository.GetAll()
}
