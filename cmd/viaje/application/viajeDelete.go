package application

import (
	domainShared "back-transportes-fisi/cmd/shared/domain"
	"back-transportes-fisi/cmd/viaje/domain/repository"
)

type ViajeDelete struct {
	repository repository.ViajeRepository
}

func NewViajeDelete(repository repository.ViajeRepository) *ViajeDelete {
	return &ViajeDelete{repository}
}

func (this *ViajeDelete) Execute(id int) error {
	if viaje, _ := this.repository.GetById(id); viaje == nil {
		return domainShared.NewEntityNotFound(id, "Viaje")
	}

	return this.repository.Delete(id)
}
