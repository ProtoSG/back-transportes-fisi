package infrastructure

import (
	"back-transportes-fisi/cmd/ciudad/domain"
	"database/sql"
)

type SQLiteCiudadRepository struct {
	db *sql.DB
}

func NewSQLiteCiudadRepository(db *sql.DB) *SQLiteCiudadRepository {
	return &SQLiteCiudadRepository{db: db}
}

func (this *SQLiteCiudadRepository) Create(ciudad *domain.Ciudad) error {
	stmt, err := this.db.Prepare("INSERT INTO Ciudad (nombre) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ciudad.Nombre)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteCiudadRepository) GetAll() ([]*domain.Ciudad, error) {
	rows, err := this.db.Query("SELECT * FROM Ciudad")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ciudads []*domain.Ciudad
	for rows.Next() {
		ciudad := &domain.Ciudad{}
		err := rows.Scan(&ciudad.IdCiudad, &ciudad.Nombre)
		if err != nil {
			return nil, err
		}

		ciudads = append(ciudads, ciudad)
	}

	return ciudads, nil
}

func (this *SQLiteCiudadRepository) GetById(id int) (*domain.Ciudad, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Ciudad WHERE id_ciudad = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	ciudad := &domain.Ciudad{}
	err = row.Scan(&ciudad.IdCiudad, &ciudad.Nombre)
	if err != nil {
		return nil, err
	}

	return ciudad, nil
}

func (this *SQLiteCiudadRepository) Update(ciudad *domain.Ciudad) error {
	stmt, err := this.db.Prepare("UPDATE Ciudad SET nombre = ? WHERE id_ciudad = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ciudad.Nombre, ciudad.IdCiudad)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteCiudadRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Ciudad WHERE id_ciudad = ?")
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
