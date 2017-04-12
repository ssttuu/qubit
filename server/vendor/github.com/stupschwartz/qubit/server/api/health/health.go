package health

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/handler"
	"github.com/stupschwartz/qubit/server/env"
	"fmt"
	"net/http"
	"golang.org/x/net/context"
)

func GetAllHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, "{\"status\":\"ok\"}")

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/health").Subrouter()

	s.Handle("/", handler.Handler{environ, GetAllHandler}).Methods("GET")
}