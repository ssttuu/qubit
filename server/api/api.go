package api

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/api/nodes"
	"github.com/stupschwartz/qubit/server/env"
	//"github.com/stupschwartz/qubit/server/api/render"
	_ "github.com/lib/pq"
	"github.com/stupschwartz/qubit/server/api/render"
)

func Handlers(environ *env.Env) *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	nodes.Register(apiRouter, environ)
	render.Register(apiRouter, environ)

	return router
}
