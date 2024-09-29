package application

import (
	"back-transportes-fisi/cmd/boleto/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type BoletoDelete struct {
	repository repository.BoletoRepository
}

func NewBoletoDelete(repository repository.BoletoRepository) *BoletoDelete {
	return &BoletoDelete{repository: repository}
}

func (b *BoletoDelete) Execute(id int) error {
	if boleto, _ := b.repository.GetById(id); boleto == nil {
		return domain.NewEntityNotFound(id, "Boleto")
	}

	return b.repository.Delete(id)
}
