package infrastructure

import (
	"back-transportes-fisi/cmd/asiento/application"
	"back-transportes-fisi/cmd/asiento/infrastructure"
	"database/sql"
)

type AsientoService struct {
	Create     *application.AsientoCreate
	GetAll     *application.AsientoGetAll
	GetById    *application.AsientoGetById
	Update     *application.AsientoUpdate
	Delete     *application.AsientoDelete
	GetByIdBus *application.AsientoGetByIdBus
}

func NewAsientoService(db *sql.DB) AsientoService {
	asientoContainer := infrastructure.NewSQLiteAsientoRepository(db)
	return AsientoService{
		Create:     application.NewAsientoCreate(asientoContainer),
		GetAll:     application.NewAsientoGetAll(asientoContainer),
		GetById:    application.NewAsientoGetById(asientoContainer),
		Update:     application.NewAsientoUpdate(asientoContainer),
		Delete:     application.NewAsientoDelete(asientoContainer),
		GetByIdBus: application.NewAsientoGetByIdBus(asientoContainer),
	}
}
