package infrastructure

import (
	"back-transportes-fisi/cmd/boleto/application"
	"back-transportes-fisi/cmd/boleto/infrastructure"
	"database/sql"
)

type BoletoService struct {
	Create  *application.BoletoCreate
	GetAll  *application.BoletoGetAll
	GetById *application.BoletoGetById
	Update  *application.BoletoUpdate
	Delete  *application.BoletoDelete
}

func NewBoletoService(db *sql.DB) BoletoService {
	boletoContainer := infrastructure.NewSQLiteBoletoRepository(db)
	return BoletoService{
		Create:  application.NewBoletoCreate(boletoContainer),
		GetAll:  application.NewBoletoGetAll(boletoContainer),
		GetById: application.NewBoletoGetById(boletoContainer),
		Update:  application.NewBoletoUpdate(boletoContainer),
		Delete:  application.NewBoletoDelete(boletoContainer),
	}
}
