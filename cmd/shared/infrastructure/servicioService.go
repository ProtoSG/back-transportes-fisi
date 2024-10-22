package infrastructure

import (
	"back-transportes-fisi/cmd/servicio/application"
	"back-transportes-fisi/cmd/servicio/infrastructure"
	"database/sql"
)

type ServicioService struct {
	Create  *application.ServicioCreate
	GetAll  *application.ServicioGetAll
	GetById *application.ServicioGetById
	Update  *application.ServicioUpdate
	Delete  *application.ServicioDelete
}

func NewServicioService(db *sql.DB) ServicioService {
	serviceContainer := infrastructure.NewSQLiteServicioRepository(db)

	return ServicioService{
		Create:  application.NewServicioCreate(serviceContainer),
		GetAll:  application.NewServicioGetAll(serviceContainer),
		GetById: application.NewServicioGetById(serviceContainer),
		Update:  application.NewServicioUpdate(serviceContainer),
		Delete:  application.NewServicioDelete(serviceContainer),
	}
}
