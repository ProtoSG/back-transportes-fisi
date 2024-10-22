package infrastructure

import (
	"back-transportes-fisi/cmd/ruta/application"
	"back-transportes-fisi/cmd/ruta/infrastructure"
	"database/sql"
)

type RutaService struct {
	Create  *application.RutaCreate
	GetAll  *application.RutaGetAll
	GetById *application.RutaGetById
	Update  *application.RutaUpdate
	Delete  *application.RutaDelete
}

func NewRutaService(db *sql.DB) RutaService {
	rutaContainer := infrastructure.NewSQLiteRutaRepository(db)

	return RutaService{
		Create:  application.NewRutaCreate(rutaContainer),
		GetAll:  application.NewRutaGetAll(rutaContainer),
		GetById: application.NewRutaGetById(rutaContainer),
		Update:  application.NewRutaUpdate(rutaContainer),
		Delete:  application.NewRutaDelete(rutaContainer),
	}
}
