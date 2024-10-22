package router

import (
	asientoRouter "back-transportes-fisi/cmd/asiento/infrastructure/router"
	boletoRouter "back-transportes-fisi/cmd/boleto/infrastructure/router"
	busRouter "back-transportes-fisi/cmd/bus/infrastructure/router"
	ciudadRouter "back-transportes-fisi/cmd/ciudad/infrastructure/router"
	clienteRouter "back-transportes-fisi/cmd/cliente/infrastructure/router"
	descuentoRouter "back-transportes-fisi/cmd/descuento/infrastructure/router"
	pasajeroRouter "back-transportes-fisi/cmd/pasajero/infrastructure/router"
	rutaRouter "back-transportes-fisi/cmd/ruta/infrastructure/router"
	servicioRouter "back-transportes-fisi/cmd/servicio/infrastructure/router"
	"back-transportes-fisi/cmd/shared/infrastructure"
	terminalRouter "back-transportes-fisi/cmd/terminal/infrastructure/router"
	transaccionRouter "back-transportes-fisi/cmd/transaccion/infrastructure/router"
	viajeRouter "back-transportes-fisi/cmd/viaje/infrastructure/router"

	"github.com/gorilla/mux"
)

func NewRouter(serviceContainer *infrastructure.ServiceContainer) *mux.Router {
	r := mux.NewRouter()

	busRouter.BusRouter(r, serviceContainer)
	asientoRouter.AsientoRouter(r, serviceContainer)
	boletoRouter.BoletoRouter(r, serviceContainer)
	ciudadRouter.CiudadRouter(r, serviceContainer)
	clienteRouter.ClienteRouter(r, serviceContainer)
	descuentoRouter.DescuentoRouter(r, serviceContainer)
	pasajeroRouter.PasajeroRouter(r, serviceContainer)
	rutaRouter.RutaRouter(r, serviceContainer)
	servicioRouter.ServicioRouter(r, serviceContainer)
	terminalRouter.TerminalRouter(r, serviceContainer)
	transaccionRouter.TransaccionRouter(r, serviceContainer)
	viajeRouter.ViajeRouter(r, serviceContainer)

	return r
}
