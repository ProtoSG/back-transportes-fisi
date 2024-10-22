package infrastructure

import (
	"back-transportes-fisi/cmd/cliente/domain"
	"database/sql"
)

type SQLiteClienteRepository struct {
	db *sql.DB
}

func NewSQLiteClienteRepository(db *sql.DB) *SQLiteClienteRepository {
	return &SQLiteClienteRepository{db: db}
}

func (this *SQLiteClienteRepository) Create(cliente *domain.Cliente) error {
	stmt, err := this.db.Prepare("INSERT INTO Cliente (nombre, apellido, fecha_nacimiento, sexo, dni, email, telefono) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cliente.Nombre, cliente.Apellido, cliente.FechaNacimiento, cliente.Sexo, cliente.Dni, cliente.Email, cliente.Telefono)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteClienteRepository) GetAll() ([]*domain.Cliente, error) {
	rows, err := this.db.Query("SELECT * FROM Cliente")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []*domain.Cliente
	for rows.Next() {
		cliente := &domain.Cliente{}
		err := rows.Scan(&cliente.IdCliente, &cliente.Nombre, &cliente.Apellido, &cliente.FechaNacimiento, &cliente.Sexo, &cliente.Dni, &cliente.Email, &cliente.Telefono)
		if err != nil {
			return nil, err
		}

		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func (this *SQLiteClienteRepository) GetById(id int) (*domain.Cliente, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Cliente WHERE id_cliente = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	cliente := &domain.Cliente{}
	err = row.Scan(&cliente.IdCliente, &cliente.Nombre, &cliente.Apellido, &cliente.FechaNacimiento, &cliente.Sexo, &cliente.Dni, &cliente.Email, &cliente.Telefono)
	if err != nil {
		return nil, err
	}

	return cliente, nil
}

func (this *SQLiteClienteRepository) Update(cliente *domain.Cliente) error {
	stmt, err := this.db.Prepare("UPDATE Cliente SET nombre = ?, apellido = ?, fecha_nacimiento = ?, sexo = ?, dni = ?, email = ?, telefono = ? WHERE id_cliente = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cliente.Nombre, cliente.Apellido, cliente.FechaNacimiento, cliente.Sexo, cliente.Dni, cliente.Email, cliente.Telefono, cliente.IdCliente)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteClienteRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Cliente WHERE id_cliente = ?")
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
