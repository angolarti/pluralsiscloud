package main

import (
	"github/angolarti/pluralcloud/cmd/api/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	server := server.NewServer(r)
	server.Start()
}
