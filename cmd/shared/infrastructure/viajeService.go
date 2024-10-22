package infrastructure

import (
	"back-transportes-fisi/cmd/viaje/application"
	"back-transportes-fisi/cmd/viaje/infrastructure"
	"database/sql"
)

type ViajeService struct {
	Create                        *application.ViajeCreate
	GetAll                        *application.ViajeGetAll
	GetById                       *application.ViajeGetById
	Update                        *application.ViajeUpdate
	Delete                        *application.ViajeDelete
	GetByOriginDestinationAndDate *application.ViajeGetByOriginDestinationAndDate
}

func NewViajeService(db *sql.DB) ViajeService {
	viajeContainer := infrastructure.NewSQLiteViajeRepository(db)
	return ViajeService{
		Create:                        application.NewViajeCreate(viajeContainer),
		GetAll:                        application.NewViajeGetAll(viajeContainer),
		GetById:                       application.NewViajeGetById(viajeContainer),
		Update:                        application.NewViajeUpdate(viajeContainer),
		Delete:                        application.NewViajeDelete(viajeContainer),
		GetByOriginDestinationAndDate: application.NewViajeGetByOriginDestinationAndDate(viajeContainer),
	}
}
