package infrastructure

import (
	"back-transportes-fisi/cmd/boleto/domain"
	"database/sql"
)

type SQLiteBoletoRepository struct {
	db *sql.DB
}

func NewSQLiteBoletoRepository(db *sql.DB) *SQLiteBoletoRepository {
	return &SQLiteBoletoRepository{db: db}
}

func (this *SQLiteBoletoRepository) Create(boleto *domain.Boleto) error {
	stmt, err := this.db.Prepare("INSERT INTO boleto (id_boleto, id_pasajero, id_viaje, id_asiento, id_transaccion, fecha, sub_total) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(boleto.IdBoleto, boleto.IdPasajero, boleto.IdViaje, boleto.IdAsiento, boleto.IdTransaccion, boleto.Fecha, boleto.SubTotal)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteBoletoRepository) GetAll() ([]*domain.Boleto, error) {
	rows, err := this.db.Query("SELECT * FROM boleto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boletos []*domain.Boleto

	for rows.Next() {
		boleto := &domain.Boleto{}
		rows.Scan(&boleto.IdBoleto, &boleto.IdPasajero, &boleto.IdViaje, &boleto.IdAsiento, &boleto.IdTransaccion, &boleto.Fecha, &boleto.SubTotal)
		boletos = append(boletos, boleto)
	}

	return boletos, nil
}

func (this *SQLiteBoletoRepository) GetById(id int) (*domain.Boleto, error) {
	rows, err := this.db.Query("SELECT * FROM boleto WHERE id_boleto = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		boleto := &domain.Boleto{}
		rows.Scan(&boleto.IdBoleto, &boleto.IdPasajero, &boleto.IdViaje, &boleto.IdAsiento, &boleto.IdTransaccion, &boleto.Fecha, &boleto.SubTotal)
		return boleto, nil
	}

	return nil, nil
}

func (this *SQLiteBoletoRepository) Update(boleto *domain.Boleto) error {
	stmt, err := this.db.Prepare("UPDATE boleto SET id_pasajero = ?, id_viaje = ?, id_asiento = ?, id_transaccion = ?, fecha = ?, sub_total = ? WHERE id_boleto = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(boleto.IdPasajero, boleto.IdViaje, boleto.IdAsiento, boleto.IdTransaccion, boleto.Fecha, boleto.SubTotal, boleto.IdBoleto)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteBoletoRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM boleto WHERE id_boleto = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
