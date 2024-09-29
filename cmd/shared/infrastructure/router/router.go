package router

import (
	asientoRouter "back-transportes-fisi/cmd/asiento/infrastructure/router"
	boletoRouter "back-transportes-fisi/cmd/boleto/infrastructure/router"
	busRouter "back-transportes-fisi/cmd/bus/infrastructure/router"
	"back-transportes-fisi/cmd/shared/infrastructure"

	"github.com/gorilla/mux"
)

func NewRouter(serviceContainer *infrastructure.ServiceContainer) *mux.Router {
	r := mux.NewRouter()

	busRouter.BusRouter(r, serviceContainer)
	asientoRouter.AsientoRouter(r, serviceContainer)
	boletoRouter.BoletoRouter(r, serviceContainer)

	return r
}
