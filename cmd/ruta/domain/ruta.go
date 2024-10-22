package domain

type Ruta struct {
	IdRuta          int     `json:"id_ruta"`
	IdCiudadOrigen  int     `json:"id_ciudad_origen"`
	IdCiudadDestino int     `json:"id_ciudad_destino"`
	Precio          float32 `json:"precio"`
}

func NewRuta(idRuta, idCiudadOrigen, IdCiudadDestino int, precio float32) *Ruta {
	return &Ruta{
		IdRuta:          idRuta,
		IdCiudadOrigen:  idCiudadOrigen,
		IdCiudadDestino: IdCiudadDestino,
		Precio:          precio,
	}
}
