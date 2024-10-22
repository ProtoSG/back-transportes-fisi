package infrastructure

import (
	"back-transportes-fisi/cmd/servicio/domain"
	"database/sql"
)

type SQLiteServicioRepository struct {
	db *sql.DB
}

func NewSQLiteServicioRepository(db *sql.DB) *SQLiteServicioRepository {
	return &SQLiteServicioRepository{db: db}
}

func (this *SQLiteServicioRepository) Create(servicio *domain.Servicio) error {
	stmt, err := this.db.Prepare("INSERT INTO Servicio (nombre, descripcion, precio) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(servicio.Nombre, servicio.Descripcion, servicio.Precio)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteServicioRepository) GetAll() ([]*domain.Servicio, error) {
	rows, err := this.db.Query("SELECT * FROM Servicio")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	servicios := make([]*domain.Servicio, 0)
	for rows.Next() {
		servicio := &domain.Servicio{}
		err = rows.Scan(&servicio.IdServicio, &servicio.Nombre, &servicio.Descripcion, &servicio.Precio)
		if err != nil {
			return nil, err
		}
		servicios = append(servicios, servicio)
	}

	return servicios, nil
}

func (this *SQLiteServicioRepository) GetById(id int) (*domain.Servicio, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Servicio WHERE id_servicio = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	servicio := &domain.Servicio{}
	err = row.Scan(&servicio.IdServicio, &servicio.Nombre, &servicio.Descripcion, &servicio.Precio)
	if err != nil {
		return nil, err
	}
	return servicio, nil
}

func (this *SQLiteServicioRepository) Update(servicio *domain.Servicio) error {
	stmt, err := this.db.Prepare("UPDATE Servicio SET nombre = ?, descripcion = ?, precio = ? WHERE id_sevicio = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(servicio.Nombre, servicio.Descripcion, servicio.Precio, servicio.IdServicio)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteServicioRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Servicio WHERE id_servicio = ?")
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
