package infrastructure

type ServiceContainer struct {
	Bus         BusService
	Asiento     AsientoService
	Boleto      BoletoService
	Ciudad      CiudadService
	Cliente     ClienteService
	Descuento   DescuentoService
	Pasajero    PasajeroService
	Ruta        RutaService
	Servicio    ServicioService
	Terminal    TerminalService
	Transaccion TransaccionService
	Viaje       ViajeService
}

func NewServiceContainer() *ServiceContainer {
	container := &ServiceContainer{}

	db := NewConnection("./trans.db")

	container.Bus = NewBusService(db)
	container.Asiento = NewAsientoService(db)
	container.Boleto = NewBoletoService(db)
	container.Ciudad = NewCiudadService(db)
	container.Cliente = NewClienteService(db)
	container.Descuento = NewDescuentoService(db)
	container.Pasajero = NewPasajeroService(db)
	container.Ruta = NewRutaService(db)
	container.Servicio = NewServicioService(db)
	container.Terminal = NewTerminalService(db)
	container.Transaccion = NewTransaccionService(db)
	container.Viaje = NewViajeService(db)

	return container
}
