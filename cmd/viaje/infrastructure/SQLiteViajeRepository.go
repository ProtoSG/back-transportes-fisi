package infrastructure

import (
	"back-transportes-fisi/cmd/viaje/domain"
	"database/sql"
)

type SQLiteViajeRepository struct {
	db *sql.DB
}

func NewSQLiteViajeRepository(db *sql.DB) *SQLiteViajeRepository {
	return &SQLiteViajeRepository{db}
}

func (this *SQLiteViajeRepository) Create(viaje *domain.Viaje) error {
	stmt, err := this.db.Prepare(`
    INSERT INTO Viaje (
    id_ruta,
    fecha_salida,
    hora_embarque,
    id_terminal_embarque, 
    hora_desembarque,
    id_terminal_desembarque, 
    estado, 
    id_servicio,
    id_bus,
    precio_viaje,
    asientos_disponibles) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		viaje.IdRuta,
		viaje.FechaSalida,
		viaje.HoraEmbarque,
		viaje.IdTerminalEmbarque,
		viaje.HoraDesembarque,
		viaje.IdTerminalDesembarque,
		viaje.Estado,
		viaje.IdServicio,
		viaje.IdBus,
		viaje.PrecioViaje,
		viaje.AsientosDisponibles,
	)

	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteViajeRepository) GetAll() ([]*domain.Viaje, error) {
	rows, err := this.db.Query("SELECT * FROM Viaje")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var viajes []*domain.Viaje

	for rows.Next() {
		var viaje domain.Viaje
		err = rows.Scan(
			&viaje.IdViaje,
			&viaje.IdRuta,
			&viaje.FechaSalida,
			&viaje.HoraEmbarque,
			&viaje.IdTerminalEmbarque,
			&viaje.HoraDesembarque,
			&viaje.IdTerminalDesembarque,
			&viaje.Estado,
			&viaje.IdServicio,
			&viaje.IdBus,
			&viaje.PrecioViaje,
			&viaje.AsientosDisponibles,
		)
		if err != nil {
			return nil, err
		}
		viajes = append(viajes, &viaje)
	}

	return viajes, nil
}

func (this *SQLiteViajeRepository) GetById(id int) (*domain.Viaje, error) {
	stmt, err := this.db.Prepare("SELECT * FROM Viaje WHERE id_viaje = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var viaje domain.Viaje
	err = row.Scan(
		&viaje.IdViaje,
		&viaje.IdRuta,
		&viaje.FechaSalida,
		&viaje.HoraEmbarque,
		&viaje.IdTerminalEmbarque,
		&viaje.HoraDesembarque,
		&viaje.IdTerminalDesembarque,
		&viaje.Estado,
		&viaje.IdServicio,
		&viaje.IdBus,
		&viaje.PrecioViaje,
		&viaje.AsientosDisponibles,
	)

	if err != nil {
		return nil, err
	}
	return &viaje, nil
}

func (this *SQLiteViajeRepository) Update(viaje *domain.Viaje) error {
	stmt, err := this.db.Prepare(`
    UPDATE Viaje SET
    id_ruta = ?,
    fecha_salida = ?,
    hora_embarque = ?,
    id_terminal_embarque = ?,
    hora_desembarque = ?,
    id_terminal_desembarque = ?,
    estado = ?,
    id_servicio = ?,
    id_bus = ?,
    precio_viaje = ?,
    asientos_disponibles = ?
    WHERE id_viaje = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		viaje.IdRuta,
		viaje.FechaSalida,
		viaje.HoraEmbarque,
		viaje.IdTerminalEmbarque,
		viaje.HoraDesembarque,
		viaje.IdTerminalDesembarque,
		viaje.Estado,
		viaje.IdServicio,
		viaje.IdBus,
		viaje.PrecioViaje,
		viaje.AsientosDisponibles,
		viaje.IdViaje,
	)
	if err != nil {
		return err
	}

	return nil
}

func (this *SQLiteViajeRepository) Delete(id int) error {
	stmt, err := this.db.Prepare("DELETE FROM Viaje WHERE id_viaje = ?")
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

func (this *SQLiteViajeRepository) GetByOriginDestinationAndDate(origen, destino, fecha string) ([]*domain.ViajeOriginDestinationAndDate, error) {
	stmt, err := this.db.Prepare(`
    SELECT 
      v.id_viaje,
      v.hora_embarque,
      v.id_terminal_embarque,
      te.ubicacion as ubicacion_embarque,
      v.hora_desembarque,
      v.id_terminal_desembarque,
      td.ubicacion as ubicacion_desembarque,
      v.id_servicio,
      s.nombre as nombre_servicio,
      v.id_bus,
      b.piso as piso_bus,
      v.asientos_disponibles,
      v.precio_viaje
    FROM Viaje v
    INNER JOIN Ruta r ON v.id_ruta = r.id_ruta
    INNER JOIN Ciudad co ON r.id_ciudad_origen = co.id_ciudad 
    INNER JOIN Ciudad cd ON r.id_ciudad_destino = cd.id_ciudad
    INNER JOIN Terminal te ON v.id_terminal_embarque = te.id_terminal 
    INNER JOIN Terminal td ON v.id_terminal_desembarque = td.id_terminal 
    INNER JOIN Servicio s ON v.id_servicio = s.id_servicio 
    INNER JOIN Bus b ON v.id_bus = b.id_bus 
    WHERE 
      LOWER(co.nombre) = LOWER(?) AND 
      LOWER(cd.nombre) = LOWER(?) AND
      v.fecha_salida = ?;
    `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(origen, destino, fecha)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	viajes := []*domain.ViajeOriginDestinationAndDate{}

	for rows.Next() {
		var viaje domain.ViajeOriginDestinationAndDate
		err = rows.Scan(
			&viaje.IdViaje,
			&viaje.Embarque.Hora,
			&viaje.Embarque.IdTerminal,
			&viaje.Embarque.Ubicacion,
			&viaje.Desembarque.Hora,
			&viaje.Desembarque.IdTerminal,
			&viaje.Desembarque.Ubicacion,
			&viaje.Servicio.IdServicio,
			&viaje.Servicio.Nombre,
			&viaje.Bus.IdBus,
			&viaje.Bus.Piso,
			&viaje.AsientosDisponibles,
			&viaje.Precio,
		)
		if err != nil {
			return nil, err
		}
		viajes = append(viajes, &viaje)
	}

	return viajes, nil
}
