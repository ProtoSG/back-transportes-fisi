package domain

type Terminal struct {
	IdTerminal int    `json:"id_terminal"`
	Nombre     string `json:"nombre"`
	Ubicacion  string `json:"ubicacion"`
}

func NewTerminal(idTerminal int, nombre string, ubicacion string) *Terminal {
	return &Terminal{
		IdTerminal: idTerminal,
		Nombre:     nombre,
		Ubicacion:  ubicacion,
	}
}
