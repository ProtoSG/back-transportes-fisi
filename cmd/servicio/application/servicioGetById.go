package application

import (
	"back-transportes-fisi/cmd/servicio/domain"
	"back-transportes-fisi/cmd/servicio/domain/repository"
	domainShared "back-transportes-fisi/cmd/shared/domain"
)

type ServicioGetById struct {
	repository repository.ServicioRepository
}

func NewServicioGetById(repository repository.ServicioRepository) *ServicioGetById {
	return &ServicioGetById{repository}
}

func (this *ServicioGetById) Execute(id int) (*domain.Servicio, error) {
	servicio, _ := this.repository.GetById(id)
	if servicio == nil {
		return nil, domainShared.NewEntityNotFound(id, "Servicio")
	}

	return servicio, nil
}
