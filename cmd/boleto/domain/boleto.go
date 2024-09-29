package domain

type Boleto struct {
	IdBoleto      int `json:"id_boleto"`
	IdPasajero    int `json:"id_pasajero"`
	IdViaje       int `json:"id_viaje"`
	IdAsiento     int `json:"id_asiento"`
	IdTransaccion int `json:"id_transaccion"`
	Fecha         int `json:"fecha"`
	SubTotal      int `json:"sub_total"`
}

func NewBoleto(idBoleto, idPasajero, idViaje, idAsiento, idTransaccion, fecha, subTotal int) *Boleto {
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
