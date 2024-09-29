package infrastructure

import (
	"back-transportes-fisi/cmd/bus/application"
	"back-transportes-fisi/cmd/bus/infrastructure"
	"database/sql"
)

type BusService struct {
	Create  *application.BusCreate
	GetAll  *application.BusGetAll
	GetById *application.BusGetById
	Update  *application.BusUpdate
	Delete  *application.BusDelete
}

func NewBusService(db *sql.DB) BusService {
	busContainer := infrastructure.NewSQLiteBusRepository(db)
	return BusService{
		Create:  application.NewBusCreate(busContainer),
		GetAll:  application.NewBusGetAll(busContainer),
		GetById: application.NewBusGetById(busContainer),
		Update:  application.NewBusUpdate(busContainer),
		Delete:  application.NewBusDelete(busContainer),
	}
}
