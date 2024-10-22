package infrastructure

import (
	"back-transportes-fisi/cmd/pasajero/domain"
	"database/sql"
)

type SQLitePasajeroRepository struct {
	db *sql.DB
}

func NewSQLitePasajeroRepository(db *sql.DB) *SQLitePasajeroRepository {
	return &SQLitePasajeroRepository{db}
}

func (this *SQLitePasajeroRepository) Create(pasajero *domain.Pasajero) error {
	stmt, err := this.db.Prepare("INSERT INTO Pasajero (nombre, apellido, fecha_nacimiento, sexo, dni) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pasajero.Nombre, pasajero.Apellido, pasajero.FechaNacimiento, pasajero.Sexo, pasajero.Dni)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLitePasajeroRepository) GetAll() ([]*domain.Pasajero, error) {
	rows, err := this.db.Query("SELECT * FROM Pasajero")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pasajeros []*domain.Pasajero
	for rows.Next() {
		pasajero := &domain.Pasajero{}
		err = rows.Scan(&pasajero.IdPasajero, &pasajero.Nombre, &pasajero.Apellido, &pasajero.FechaNacimiento, &pasajero.Sexo, &pasajero.Dni)
		if err != nil {
			return nil, err
		}
		pasajeros = append(pasajeros, pasajero)
	}

	return pasajeros, nil
}

func (this *SQLitePasajeroRepository) GetById(id int) (*domain.Pasajero, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Pasajero WHERE id_pasajero = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	pasajero := &domain.Pasajero{}
	err = row.Scan(&pasajero.IdPasajero, &pasajero.Nombre, &pasajero.Apellido, &pasajero.FechaNacimiento, &pasajero.Sexo, &pasajero.Dni)
	if err != nil {
		return nil, err
	}

	return pasajero, nil
}

func (this *SQLitePasajeroRepository) Update(pasajero *domain.Pasajero) error {
	stmt, err := this.db.Prepare("UPDATE Pasajero SET nombre = ?, apellido = ?, fecha_nacimiento = ?, sexo = ?, dni = ? WHERE id_pasajero = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pasajero.Nombre, pasajero.Apellido, pasajero.FechaNacimiento, pasajero.Sexo, pasajero.Dni, pasajero.IdPasajero)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLitePasajeroRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Pasajero WHERE id_pasajero = ?")
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
