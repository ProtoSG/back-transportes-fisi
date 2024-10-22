package infrastructure

import (
	"back-transportes-fisi/cmd/terminal/domain"
	"database/sql"
)

type SQLiteTerminalRepository struct {
	db *sql.DB
}

func NewSQLiteTerminalRepository(db *sql.DB) *SQLiteTerminalRepository {
	return &SQLiteTerminalRepository{db}
}

func (this *SQLiteTerminalRepository) Create(terminal *domain.Terminal) error {
	stmt, err := this.db.Prepare("INSET INTO Terminal (nombre, ubicacion) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(terminal.Nombre, terminal.Ubicacion)
	if err != nil {
		return err
	}
	return nil
}

func (this *SQLiteTerminalRepository) GetAll() ([]*domain.Terminal, error) {
	rows, err := this.db.Query("SELECT * FROM Terminal")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var terminals []*domain.Terminal
	for rows.Next() {
		terminal := &domain.Terminal{}
		err = rows.Scan(&terminal.IdTerminal, &terminal.Nombre, &terminal.Ubicacion)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}

	return terminals, nil
}

func (this *SQLiteTerminalRepository) GetById(id int) (*domain.Terminal, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Terminal WHERE id_terminal = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	terminal := &domain.Terminal{}
	err = row.Scan(&terminal.IdTerminal, &terminal.Nombre, &terminal.Ubicacion)
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (this *SQLiteTerminalRepository) Update(terminal *domain.Terminal) error {
	stmt, err := this.db.Prepare("UPDATE Terminal SET nombre = ?, ubicacion = ? WHERE id_terminal = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(terminal.Nombre, terminal.Ubicacion, terminal.IdTerminal)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteTerminalRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Terminal WHERE id_terminal = ?")
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
