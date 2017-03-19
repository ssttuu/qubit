package render

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"net/http"
	"github.com/stupschwartz/qubit/node"
	"log"
	"cloud.google.com/go/datastore"
)

func PostHandler(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]
	nodeKey := datastore.NameKey("Node", whereNodeId, nil)

	var existingNode node.Node
	if err := e.DatastoreClient.Get(e.Context, nodeKey, &existingNode); err != nil {
		return err
	}

	log.Printf("Rendering Task %v", existingNode)

	w.WriteHeader(http.StatusProcessing)

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/render").Subrouter()

	s.Handle("/{id}/", handler.Handler{environ, PostHandler}).Methods("POST")
}
