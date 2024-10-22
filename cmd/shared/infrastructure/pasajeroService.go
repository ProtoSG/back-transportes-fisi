package infrastructure

import (
	"back-transportes-fisi/cmd/pasajero/application"
	"back-transportes-fisi/cmd/pasajero/infrastructure"
	"database/sql"
)

type PasajeroService struct {
	Create  *application.PasajeroCreate
	GetAll  *application.PasajeroGetAll
	GetById *application.PasajeroGetById
	Update  *application.PasajeroUpdate
	Delete  *application.PasajeroDelete
}

func NewPasajeroService(db *sql.DB) PasajeroService {
	pasajeroContainer := infrastructure.NewSQLitePasajeroRepository(db)

	return PasajeroService{
		Create:  application.NewPasajeroCreate(pasajeroContainer),
		GetAll:  application.NewPasajeroGetAll(pasajeroContainer),
		GetById: application.NewPasajeroGetById(pasajeroContainer),
		Update:  application.NewPasajeroUpdate(pasajeroContainer),
		Delete:  application.NewPasajeroDelete(pasajeroContainer),
	}
}
