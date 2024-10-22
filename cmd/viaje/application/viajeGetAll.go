package application

import (
	"back-transportes-fisi/cmd/viaje/domain"
	"back-transportes-fisi/cmd/viaje/domain/repository"
)

type ViajeGetAll struct {
	repository repository.ViajeRepository
}

func NewViajeGetAll(repository repository.ViajeRepository) *ViajeGetAll {
	return &ViajeGetAll{repository}
}

func (this *ViajeGetAll) Execute() ([]*domain.Viaje, error) {
	return this.repository.GetAll()
}
