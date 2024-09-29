package domain

type Bus struct {
	IdBus     int `json:"id_bus"`
	Capacidad int `json:"capacidad"`
	Piso      int `json:"piso"`
}

func NewBus(idBus int, capacidad int, piso int) *Bus {
	return &Bus{
		IdBus:     idBus,
		Capacidad: capacidad,
		Piso:      piso,
	}
}
