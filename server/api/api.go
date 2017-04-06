package api

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/api/nodes"
	"github.com/stupschwartz/qubit/server/env"
	_ "github.com/lib/pq"
	"github.com/stupschwartz/qubit/server/api/render"
	"github.com/stupschwartz/qubit/server/api/health"
)

func Handlers(environ *env.Env) *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	health.Register(apiRouter, environ)
	nodes.Register(apiRouter, environ)
	render.Register(apiRouter, environ)

	return router
}
