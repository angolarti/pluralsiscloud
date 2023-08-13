package server

import (
	"github/angolarti/pluralcloud/cmd/api/router"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router chi.Router
}

func NewServer(router chi.Router) *Server {
	return &Server{
		Router: router,
	}
}

func (server *Server) Start() error {
	router.SetupRouter(server.Router)
	return http.ListenAndServe(":5000", server.Router)

}
