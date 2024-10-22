package application

import (
	domainShared "back-transportes-fisi/cmd/shared/domain"
	"back-transportes-fisi/cmd/viaje/domain"
	"back-transportes-fisi/cmd/viaje/domain/repository"
)

type ViajeGetById struct {
	repository repository.ViajeRepository
}

func NewViajeGetById(repository repository.ViajeRepository) *ViajeGetById {
	return &ViajeGetById{repository}
}

func (this *ViajeGetById) Execute(id int) (*domain.Viaje, error) {
	viaje, _ := this.repository.GetById(id)
	if viaje == nil {
		return nil, domainShared.NewEntityNotFound(id, "Viaje")
	}

	return viaje, nil
}
