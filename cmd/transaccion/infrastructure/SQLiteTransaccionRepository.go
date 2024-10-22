package infrastructure

import (
	"back-transportes-fisi/cmd/transaccion/domain"
	"database/sql"
)

type SQLiteTransaccionRepository struct {
	db *sql.DB
}

func NewSQLiteTransaccionRepository(db *sql.DB) *SQLiteTransaccionRepository {
	return &SQLiteTransaccionRepository{db}
}

func (this *SQLiteTransaccionRepository) Create(transaccion *domain.Transaccion) error {
	stmt, err := this.db.Prepare("INSERT INTO Transaccion (id_cliente, id_descuento, total, metodo_pago, estado, fecha_transaccion) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(transaccion.IdCliente, transaccion.IdDescuento, transaccion.Total, transaccion.MetodoPago, transaccion.Estado, transaccion.FechaTransaccion)
	if err != nil {
		return err
	}

	return nil
}
