package router

import (
	controllers "github/angolarti/pluralcloud/cmd/api/container/controller"

	"github.com/go-chi/chi/v5"
)

func ContainerRouter(r chi.Router) chi.Router {

	r.Post("/container", controllers.CreateContainer)
	r.Get("/containers", controllers.ListContainer)

	return r

}
