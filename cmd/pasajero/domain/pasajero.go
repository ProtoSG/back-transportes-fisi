package domain

type Pasajero struct {
	IdPasajero      int    `json:"id_pasajero"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Sexo            string `json:"sexo"`
	Dni             string `json:"dni"`
}

func NewPasajero(idPasajero int, nombre, apellido, fechaNacimiento, sexo, dni string) *Pasajero {
	return &Pasajero{
		IdPasajero:      idPasajero,
		Nombre:          nombre,
		Apellido:        apellido,
		FechaNacimiento: fechaNacimiento,
		Sexo:            sexo,
		Dni:             dni,
	}
}
