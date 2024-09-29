package application

import (
	"back-transportes-fisi/cmd/boleto/domain"
	"back-transportes-fisi/cmd/boleto/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type BoletoGetById struct {
	repository repository.BoletoRepository
}

func NewBoletoGetById(repository repository.BoletoRepository) *BoletoGetById {
	return &BoletoGetById{repository: repository}
}

func (b *BoletoGetById) Execute(id int) (*domain.Boleto, error) {
	boleto, _ := b.repository.GetById(id)
	if boleto == nil {
		return nil, domainShared.NewEntityNotFound(id, "Boleto")
	}

	return boleto, nil
}
