package infrastructure

import (
	"back-transportes-fisi/cmd/ruta/domain"
	"database/sql"
)

type SQLiteRutaRepository struct {
	db *sql.DB
}

func NewSQLiteRutaRepository(db *sql.DB) *SQLiteRutaRepository {
	return &SQLiteRutaRepository{db}
}

func (this *SQLiteRutaRepository) Create(ruta *domain.Ruta) error {
	stmt, err := this.db.Prepare("INSERT INTO Ruta (id_ciudad_origen, id_ciudad_destino, precio) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ruta.IdCiudadOrigen, ruta.IdCiudadDestino, ruta.Precio)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteRutaRepository) GetAll() ([]*domain.Ruta, error) {
	rows, err := this.db.Query("SELECT * FROM Ruta")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rutas []*domain.Ruta
	for rows.Next() {
		var ruta domain.Ruta
		err = rows.Scan(&ruta.IdRuta, &ruta.IdCiudadOrigen, &ruta.IdCiudadDestino, &ruta.Precio)
		if err != nil {
			return nil, err
		}
		rutas = append(rutas, &ruta)
	}

	return rutas, nil
}

func (this *SQLiteRutaRepository) GetById(id int) (*domain.Ruta, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Ruta WHERE id_ruta = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ruta domain.Ruta
	err = stmt.QueryRow(id).Scan(&ruta.IdRuta, &ruta.IdCiudadOrigen, &ruta.IdCiudadDestino, &ruta.Precio)
	if err != nil {
		return nil, err
	}

	return &ruta, nil
}

func (this *SQLiteRutaRepository) Update(ruta *domain.Ruta) error {
	stmt, err := this.db.Prepare("UPDATE Ruta SET id_ciudad_origen = ?, id_ciudad_destino = ?, precio = ? WHERE id_ruta = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ruta.IdCiudadOrigen, ruta.IdCiudadDestino, ruta.Precio, ruta.IdRuta)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteRutaRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Ruta WHERE id_ruta = ?")
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
