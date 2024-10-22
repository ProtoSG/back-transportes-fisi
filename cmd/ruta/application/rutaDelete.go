package application

import (
	"back-transportes-fisi/cmd/ruta/domain/repository"
	"back-transportes-fisi/cmd/shared/domain"
)

type RutaDelete struct {
	repository repository.RutaRepository
}

func NewRutaDelete(repository repository.RutaRepository) *RutaDelete {
	return &RutaDelete{repository}
}

func (this *RutaDelete) Execute(id int) error {
	if ruta, _ := this.repository.GetById(id); ruta == nil {
		return domain.NewEntityNotFound(id, "Ruta")
	}
	return this.repository.Delete(id)
}
