package infrastructure

import (
	"back-transportes-fisi/cmd/descuento/application"
	"back-transportes-fisi/cmd/descuento/infrastructure"
	"database/sql"
)

type DescuentoService struct {
	GetAll  *application.DescuentoGetAll
	GetById *application.DescuentoGetById
	Create  *application.DescuentoCreate
	Update  *application.DescuentoUpdate
	Delete  *application.DescuentoDelete
}

func NewDescuentoService(db *sql.DB) DescuentoService {
	descuentoContinaer := infrastructure.NewSQLiteDescuentoRepository(db)

	return DescuentoService{
		GetAll:  application.NewDescuentoGetAll(descuentoContinaer),
		GetById: application.NewDescuentoGetById(descuentoContinaer),
		Create:  application.NewDescuentoCreate(descuentoContinaer),
		Update:  application.NewDescuentoUpdate(descuentoContinaer),
		Delete:  application.NewDescuentoDelete(descuentoContinaer),
	}
}
