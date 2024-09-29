package application

import (
	"back-transportes-fisi/cmd/bus/domain"
	"back-transportes-fisi/cmd/bus/domain/repository"
)

type BusGetAll struct {
	repository repository.BusRepository
}

func NewBusGetAll(repository repository.BusRepository) *BusGetAll {
	return &BusGetAll{repository: repository}
}

func (b *BusGetAll) Execute() ([]*domain.Bus, error) {
	return b.repository.GetAll()
}
