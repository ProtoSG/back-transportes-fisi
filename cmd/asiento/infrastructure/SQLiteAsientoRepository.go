package infrastructure

import (
	"back-transportes-fisi/cmd/asiento/domain"
	"database/sql"
)

type SQLiteAsientoRepository struct {
	db *sql.DB
}

func NewSQLiteAsientoRepository(db *sql.DB) *SQLiteAsientoRepository {
	return &SQLiteAsientoRepository{db: db}
}

func (this *SQLiteAsientoRepository) GetAll() ([]*domain.Asiento, error) {
	rows, err := this.db.Query("SELECT * FROM asiento")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var asientos []*domain.Asiento

	for rows.Next() {
		asiento := &domain.Asiento{}
		rows.Scan(&asiento.IdAsiento, &asiento.IdBus, &asiento.NumeroAsiento, &asiento.Piso, &asiento.Precio, &asiento.Disponibilidad)
		asientos = append(asientos, asiento)
	}

	return asientos, nil
}

func (this *SQLiteAsientoRepository) GetById(id int) (*domain.Asiento, error) {
	rows, err := this.db.Query("SELECT * FROM asiento WHERE id_asiento = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		asiento := &domain.Asiento{}
		rows.Scan(&asiento.IdAsiento, &asiento.IdBus, &asiento.NumeroAsiento, &asiento.Piso, &asiento.Precio, &asiento.Disponibilidad)
		return asiento, nil
	}

	return nil, nil
}

func (this *SQLiteAsientoRepository) Create(asiento *domain.Asiento) error {
	stmt, err := this.db.Prepare("INSERT INTO asiento (id_bus, numero_asiento, piso, precio, disponibilidad) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(asiento.IdBus, asiento.NumeroAsiento, asiento.Piso, asiento.Precio, asiento.Disponibilidad)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteAsientoRepository) Update(asiento *domain.Asiento) error {
	stmt, err := this.db.Prepare("UPDATE asiento SET id_bus = ?, numero_asiento = ?, piso = ?, precio = ?, disponibilidad = ? WHERE id_asiento = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(asiento.IdBus, asiento.NumeroAsiento, asiento.Piso, asiento.Precio, asiento.Disponibilidad, asiento.IdAsiento)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteAsientoRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM asiento WHERE id_asiento = ?")
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

func (this *SQLiteAsientoRepository) GetByIdBus(id int) ([]*domain.Asiento, error) {
	stmt, err := this.db.Prepare(`
    SELECT 
      a.id_asiento,
      a.id_bus,
      a.numero_asiento,
      a.piso,
      a.precio,
      a.disponibilidad
    FROM Asiento a
    INNER JOIN Bus b ON a.id_bus = b.id_bus
    WHERE b.id_bus = ?;
  `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var asientos []*domain.Asiento

	for rows.Next() {
		asiento := &domain.Asiento{}
		rows.Scan(&asiento.IdAsiento, &asiento.IdBus, &asiento.NumeroAsiento, &asiento.Piso, &asiento.Precio, &asiento.Disponibilidad)
		asientos = append(asientos, asiento)
	}

	return asientos, nil
}
