package infrastructure

import (
	"back-transportes-fisi/cmd/cliente/application"
	"back-transportes-fisi/cmd/cliente/infrastructure"
	"database/sql"
)

type ClienteService struct {
	Create  *application.ClienteCreate
	GetAll  *application.ClienteGetAll
	GetById *application.ClienteGetById
	Update  *application.ClienteUpdate
	Delete  *application.ClienteDelete
}

func NewClienteService(db *sql.DB) ClienteService {
	clietneContainer := infrastructure.NewSQLiteClienteRepository(db)

	return ClienteService{
		Create:  application.NewClienteCreate(clietneContainer),
		GetAll:  application.NewClienteGetAll(clietneContainer),
		GetById: application.NewClienteGetById(clietneContainer),
		Update:  application.NewClienteUpdate(clietneContainer),
		Delete:  application.NewClienteDelete(clietneContainer),
	}
}
