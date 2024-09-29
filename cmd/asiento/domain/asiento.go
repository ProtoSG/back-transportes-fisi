package domain

type Asiento struct {
	IdAsiento      int `json:"id_asiento"`
	IdBus          int `json:"id_bus"`
	NumeroAsiento  int `json:"numero_asiento"`
	Piso           int `json:"piso"`
	Precio         int `json:"precio"`
	Disponibilidad int `json:"disponibilidad"`
}

func NewAsiento(idAsiento int, idBus int, numeroAsiento int, piso int, precio int, disponibilidad int) *Asiento {
	return &Asiento{
		IdAsiento:      idAsiento,
		IdBus:          idBus,
		NumeroAsiento:  numeroAsiento,
		Piso:           piso,
		Precio:         precio,
		Disponibilidad: disponibilidad,
	}
}
