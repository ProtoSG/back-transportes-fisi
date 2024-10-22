package infrastructure

import (
	"back-transportes-fisi/cmd/descuento/domain"
	"database/sql"
)

type SQLiteDescuentoRepository struct {
	db *sql.DB
}

func NewSQLiteDescuentoRepository(db *sql.DB) *SQLiteDescuentoRepository {
	return &SQLiteDescuentoRepository{db}
}

func (this *SQLiteDescuentoRepository) GetAll() ([]*domain.Descuento, error) {
	rows, err := this.db.Query("SELECT * FROM Descuento")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var descuentos []*domain.Descuento
	for rows.Next() {
		descuento := &domain.Descuento{}
		err = rows.Scan(&descuento.IdDescuento, &descuento.Codigo, &descuento.Monto)

		if err != nil {
			return nil, err
		}

		descuentos = append(descuentos, descuento)
	}

	return descuentos, nil
}

func (this *SQLiteDescuentoRepository) GetById(id int) (*domain.Descuento, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Descuento WHERE id_descuento = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	descuento := &domain.Descuento{}
	err = row.Scan(&descuento.IdDescuento, &descuento.Codigo, &descuento.Monto)
	if err != nil {
		return nil, err
	}

	return descuento, nil
}

func (this *SQLiteDescuentoRepository) Create(descuento *domain.Descuento) error {
	stmt, err := this.db.Prepare("INSERT INTO Descuento (codigo, monto) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(descuento.Codigo, descuento.Monto)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteDescuentoRepository) Update(descuento *domain.Descuento) error {
	stmt, err := this.db.Prepare("UPDATE Descuento SET codigo = ?, monto = ? WHERE id_descuento = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(descuento.Codigo, descuento.Monto, descuento.IdDescuento)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteDescuentoRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Descuento WHERE id_descuento = ?")
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
