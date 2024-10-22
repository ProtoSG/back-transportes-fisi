package domain

type Servicio struct {
	IdServicio  int     `json:"id_servicio"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float32 `json:"precio"`
}

func NewServicio(idServicio int, nombre string, descripcion string, precio float32) *Servicio {
	return &Servicio{
		IdServicio:  idServicio,
		Nombre:      nombre,
		Descripcion: descripcion,
		Precio:      precio,
	}
}
