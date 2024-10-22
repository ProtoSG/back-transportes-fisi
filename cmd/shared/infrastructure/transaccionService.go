package infrastructure

import (
	"back-transportes-fisi/cmd/transaccion/application"
	"back-transportes-fisi/cmd/transaccion/infrastructure"
	"database/sql"
)

type TransaccionService struct {
	Create *application.TransaccionCreate
}

func NewTransaccionService(db *sql.DB) TransaccionService {
	transaccionContainer := infrastructure.NewSQLiteTransaccionRepository(db)

	return TransaccionService{
		Create: application.NewTransaccionCreate(transaccionContainer),
	}
}
