package application

import (
	"back-transportes-fisi/cmd/boleto/domain"
	"back-transportes-fisi/cmd/boleto/domain/repository"
)

type BoletoGetAll struct {
	repository repository.BoletoRepository
}

func NewBoletoGetAll(repository repository.BoletoRepository) *BoletoGetAll {
	return &BoletoGetAll{repository: repository}
}

func (b *BoletoGetAll) Execute() ([]*domain.Boleto, error) {
	return b.repository.GetAll()
}
