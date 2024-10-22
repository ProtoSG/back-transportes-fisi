package domain

type Viaje struct {
	IdViaje               int     `json:"id_viaje"`
	IdRuta                int     `json:"id_ruta"`
	FechaSalida           string  `json:"fecha_salida"`
	HoraEmbarque          string  `json:"hora_embarque"`
	IdTerminalEmbarque    int     `json:"id_terminal_embarque"`
	HoraDesembarque       string  `json:"hora_desembarque"`
	IdTerminalDesembarque int     `json:"id_terminal_desembarque"`
	Estado                string  `json:"estado"`
	IdServicio            int     `json:"id_servicio"`
	IdBus                 int     `json:"id_bus"`
	PrecioViaje           float32 `json:"precio_viaje"`
	AsientosDisponibles   int     `json:"asientos_disponibles"`
}

func NewViaje(
	idViaje int,
	idRuta int,
	fechaSalida string,
	horaEmbarque string,
	idTerminalEmbarque int,
	horaDesembarque string,
	idTerminalDesembarque int,
	estado string,
	idServicio int,
	idBus int,
	precioViaje float32,
	asientosDisponibles int,
) *Viaje {
	return &Viaje{
		IdViaje:               idViaje,
		IdRuta:                idRuta,
		FechaSalida:           fechaSalida,
		HoraEmbarque:          horaEmbarque,
		IdTerminalEmbarque:    idTerminalEmbarque,
		HoraDesembarque:       horaDesembarque,
		IdTerminalDesembarque: idTerminalDesembarque,
		Estado:                estado,
		IdServicio:            idServicio,
		IdBus:                 idBus,
		PrecioViaje:           precioViaje,
		AsientosDisponibles:   asientosDisponibles,
	}
}
