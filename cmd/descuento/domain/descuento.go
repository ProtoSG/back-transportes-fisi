package domain

type Descuento struct {
	IdDescuento int     `json:"id_descuento"`
	Codigo      string  `json:"codigo"`
	Monto       float32 `json:"monto"`
}

func NewDescuento(idDescuento int, codigo string, monto float32) *Descuento {
	return &Descuento{
		IdDescuento: idDescuento,
		Codigo:      codigo,
		Monto:       monto,
	}
}
