package application

import (
	"back-transportes-fisi/cmd/servicio/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type ServicioDelete struct {
	repository repository.ServicioRepository
}

func NewServicioDelete(repository repository.ServicioRepository) *ServicioDelete {
	return &ServicioDelete{repository}
}

func (this *ServicioDelete) Execute(id int) error {
	if servicio, _ := this.repository.GetById(id); servicio == nil {
		return domain.NewEntityNotFound(id, "Servicio")
	}
	return this.repository.Delete(id)
}
