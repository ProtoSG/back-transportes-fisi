package domain

type Cliente struct {
	IdCliente       int    `json:"id_cliente"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Sexo            string `json:"sexo"`
	Dni             string `json:"dni"`
	Email           string `json:"email"`
	Telefono        string `json:"telefono"`
}

func NewCliente(idCliente int, nombre string, apellido string, fechaNacimiento string, sexo string, dni string, email string, telefono string) *Cliente {
	return &Cliente{
		IdCliente:       idCliente,
		Nombre:          nombre,
		Apellido:        apellido,
		FechaNacimiento: fechaNacimiento,
		Sexo:            sexo,
		Dni:             dni,
		Email:           email,
		Telefono:        telefono,
	}
}
