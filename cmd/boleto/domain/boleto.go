package domain

type Boleto struct {
	IdBoleto      int     `json:"id_boleto"`
	IdPasajero    int     `json:"id_pasajero"`
	IdViaje       int     `json:"id_viaje"`
	IdAsiento     int     `json:"id_asiento"`
	IdTransaccion int     `json:"id_transaccion"`
	SubTotal      float32 `json:"sub_total"`
	Fecha         string  `json:"fecha"`
}

func NewBoleto(idBoleto, idPasajero, idViaje, idAsiento, idTransaccion int, subTotal float32, fecha string) *Boleto {
	return &Boleto{
		IdBoleto:      idBoleto,
		IdPasajero:    idPasajero,
		IdViaje:       idViaje,
		IdAsiento:     idAsiento,
		IdTransaccion: idTransaccion,
		Fecha:         fecha,
		SubTotal:      subTotal,
	}
}
