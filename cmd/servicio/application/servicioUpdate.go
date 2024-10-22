package application

import (
	"back-transportes-fisi/cmd/servicio/domain"
	"back-transportes-fisi/cmd/servicio/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type ServicioUpdate struct {
	repository repository.ServicioRepository
}

func NewServicioUpdate(repository repository.ServicioRepository) *ServicioUpdate {
	return &ServicioUpdate{repository}
}

func (this *ServicioUpdate) Execute(idServicio int, name string, description string, precio float32) error {
	if servicio, _ := this.repository.GetById(idServicio); servicio == nil {
		return domainShared.NewEntityNotFound(idServicio, "Servicio")
	}

	servicio := domain.NewServicio(idServicio, name, description, precio)
	return this.repository.Update(servicio)
}
