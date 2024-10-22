package domain

type Terminal struct {
	IdTerminal int    `json:"id_embarque"`
	Hora       string `json:"hora"`
	Ubicacion  string `json:"ubicacion"`
}

type Servicio struct {
	IdServicio int    `json:"id_servicio"`
	Nombre     string `json:"nombre"`
}

type Bus struct {
	IdBus int `json:"id_bus"`
	Piso  int `json:"piso"`
}

type ViajeOriginDestinationAndDate struct {
	IdViaje             int      `json:"id_viaje"`
	Embarque            Terminal `json:"embarque"`
	Desembarque         Terminal `json:"desembarque"`
	Servicio            Servicio `json:"servicio"`
	Bus                 Bus      `json:"bus"`
	AsientosDisponibles int      `json:"asientos_disponibles"`
	Precio              float32  `json:"precio"`
}

// func NewViajeOriginDestinationAndDate(
// 	idViaje int,
// 	embarque Terminal,
// 	desembarque Terminal,
// 	servicio servicio,
// 	bus bus,
// 	asientosDisponibles int,
// 	precio float32,
// ) *ViajeOriginDestinationAndDate {
// 	return &ViajeOriginDestinationAndDate{
// 		IdViaje:             idViaje,
// 		Embarque:            embarque,
// 		Desembarque:         desembarque,
// 		Servicio:            servicio,
// 		Bus:                 bus,
// 		AsientosDisponibles: asientosDisponibles,
// 		Precio:              precio,
// 	}
// }
