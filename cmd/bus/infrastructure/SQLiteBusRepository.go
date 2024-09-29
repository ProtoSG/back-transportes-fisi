package infrastructure

import (
	"back-transportes-fisi/cmd/bus/domain"
	"database/sql"
)

type SQLiteBusRepository struct {
	db *sql.DB
}

func NewSQLiteBusRepository(db *sql.DB) *SQLiteBusRepository {
	return &SQLiteBusRepository{db}
}

func (this SQLiteBusRepository) Create(bus *domain.Bus) error {
	stmt, err := this.db.Prepare("INSERT INTO Bus (capacidad, piso) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bus.Capacidad, bus.Piso)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteBusRepository) GetAll() ([]*domain.Bus, error) {
	rows, err := this.db.Query("SELECT * FROM Bus")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buses []*domain.Bus
	for rows.Next() {
		bus := &domain.Bus{}
		err = rows.Scan(&bus.IdBus, &bus.Capacidad, &bus.Piso)
		if err != nil {
			return nil, err
		}
		buses = append(buses, bus)
	}

	return buses, nil
}

func (this SQLiteBusRepository) GetById(id int) (*domain.Bus, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Bus WHERE id_bus = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	bus := &domain.Bus{}
	err = row.Scan(&bus.IdBus, &bus.Capacidad, &bus.Piso)
	if err != nil {
		return nil, err
	}

	return bus, nil
}

func (this SQLiteBusRepository) Update(bus *domain.Bus) error {
	stmt, err := this.db.Prepare("UPDATE Bus SET capacidad = ?, piso = ? WHERE id_bus = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bus.Capacidad, bus.Piso, bus.IdBus)
	if err != nil {
		return err
	}

	return nil
}

func (this SQLiteBusRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Bus WHERE id_bus = ?")
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
