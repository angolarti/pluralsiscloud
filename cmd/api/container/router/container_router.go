package router

import (
	controllers "github/angolarti/pluralcloud/cmd/api/container/controller/create"
	list "github/angolarti/pluralcloud/cmd/api/container/controller/list"

	"github.com/go-chi/chi/v5"
)

func ContainerRouter(r chi.Router) chi.Router {

	r.Post("/container", controllers.CreateContainer)
	r.Get("/containers", list.ListContainer)

	return r

}
