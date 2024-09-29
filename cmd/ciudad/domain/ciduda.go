package domain

type Ciudad struct {
	IdCiudad int    `json:"id_ciudad"`
	Nombre   string `json:"nombre"`
}

func NewCiudad(id int, nombre string) *Ciudad {
	return &Ciudad{
		IdCiudad: id,
		Nombre:   nombre,
	}
}
