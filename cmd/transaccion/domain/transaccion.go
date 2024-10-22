package domain

type Transaccion struct {
	IdTransaccion    int     `json:"id_transaccion"`
	IdCliente        int     `json:"id_cliente"`
	IdDescuento      int     `json:"id_descuento"`
	Total            float32 `json:"total"`
	MetodoPago       string  `json:"metodo_pago"`
	Estado           string  `json:"estado"`
	FechaTransaccion string  `json:"fecha_transaccion"`
}

func NewTransaccion(idTransaccion, idCliente, idDescuento int, total float32, metodoPago, estado, fechaTransaccion string) *Transaccion {
	return &Transaccion{
		IdTransaccion:    idTransaccion,
		IdCliente:        idCliente,
		IdDescuento:      idDescuento,
		Total:            total,
		MetodoPago:       metodoPago,
		Estado:           estado,
		FechaTransaccion: fechaTransaccion,
	}
}
