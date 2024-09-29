package infrastructure

type ServiceContainer struct {
	Bus     BusService
	Asiento AsientoService
	Boleto  BoletoService
	Ciudad  CiudadService
}

func NewServiceContainer() *ServiceContainer {
	container := &ServiceContainer{}

	db := NewConnection("./trans.db")

	container.Bus = NewBusService(db)
	container.Asiento = NewAsientoService(db)
	container.Boleto = NewBoletoService(db)
	container.Ciudad = NewCiudadService(db)

	return container
}
