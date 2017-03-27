package render

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"net/http"
	"github.com/stupschwartz/qubit/node"
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"github.com/stupschwartz/qubit/operator"
	"github.com/stupschwartz/qubit/params"
	"github.com/stupschwartz/qubit/image"
	"encoding/json"
	"image/png"
	"os"
)

func Render(n *node.Node, p params.Parameters, e *env.Env) image.Plane {
	op := operator.GetOperation(n.Type)

	return op([]image.Plane{}, p, 1920, 1080)
}

func PostHandler(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	nodeUuid := vars["id"]

	nodeKey := datastore.NameKey("Node", nodeUuid, nil)

	var existingNode node.Node
	if err := e.DatastoreClient.Get(context.Background(), nodeKey, &existingNode); err != nil {
		log.Fatalf("Failed to get node to be rendered, %v", err)
	}

	var p params.Parameters
	bucket := e.StorageClient.Bucket(os.Getenv("STORAGE_BUCKET"))
	paramsObj := bucket.Object("params/" + existingNode.Id)

	ctx := context.Background()
	reader, err := paramsObj.NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to create Reader, %v", err)
	}

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&p); err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	imagePlane := Render(&existingNode, p, e)

	w.Header().Set("Content-Type", "image/png")

	if err := png.Encode(w, imagePlane.ToNRGBA()); err != nil {
		log.Fatal(err)
	}

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/render").Subrouter()

	s.Handle("/{id}/", handler.Handler{environ, PostHandler}).Methods("POST")
}
