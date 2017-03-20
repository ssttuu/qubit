package render

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"net/http"
)

func PostHandler(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]

	go RenderNodeAndDependents(e, whereNodeId)

	w.WriteHeader(http.StatusAccepted)

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/render").Subrouter()

	s.Handle("/{id}/", handler.Handler{environ, PostHandler}).Methods("POST")
}
