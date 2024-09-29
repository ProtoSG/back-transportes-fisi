package infrastructure

import (
	"back-transportes-fisi/cmd/ciudad/application"
	"back-transportes-fisi/cmd/ciudad/infrastructure"
	"database/sql"
)

type CiudadService struct {
	Create  *application.CiudadCreate
	GetAll  *application.CiudadGetAll
	GetById *application.CiudadGetById
	Update  *application.CiudadUpdate
	Delete  *application.CiudadDelete
}

func NewCiudadService(db *sql.DB) CiudadService {
	ciudadContainer := infrastructure.NewSQLiteCiudadRepository(db)
	return CiudadService{
		Create:  application.NewCiudadCreate(ciudadContainer),
		GetAll:  application.NewCiudadGetAll(ciudadContainer),
		GetById: application.NewCiudadGetById(ciudadContainer),
		Update:  application.NewCiudadUpdate(ciudadContainer),
		Delete:  application.NewCiudadDelete(ciudadContainer),
	}
}
