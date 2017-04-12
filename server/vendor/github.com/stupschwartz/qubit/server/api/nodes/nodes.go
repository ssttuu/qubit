package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/core/node"
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"net/http"
	"github.com/satori/go.uuid"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"cloud.google.com/go/storage"
	"os"
	"github.com/pkg/errors"
	"cloud.google.com/go/trace"
)

// TODO: in memory caching of Nodes (or even just the digests)
// TODO: use the input digest to lookup the content digest and use that.
// TODO: add params back to digest value
//func GetDigestFromNode(n *node.Node, e *env.Env) string {
//	digestHasher := sha256.New()
//	digestHasher.Write([]byte(n.Type))
//
//	var inputDigests []string
//	for _, inputId := range n.Inputs {
//		log.Printf("inputId: %v", inputId)
//		inputNodeKey := datastore.NameKey("Node", inputId, nil)
//		var inputNode node.Node
//		if err := e.DatastoreClient.Get(e.Context, inputNodeKey, &inputNode); err != nil {
//			log.Fatalf("Failed to get input node, %v", err)
//		}
//
//		inputDigests = append(inputDigests, inputNode.Digest)
//	}
//
//	digestHasher.Write([]byte(strings.Join(inputDigests, ",")))
//
//	//paramsAsJson, _ := json.Marshal(&n.Params)
//	//digestHasher.Write([]byte(paramsAsJson))
//	return base64.URLEncoding.EncodeToString(digestHasher.Sum(nil))
//}

func PutNode(ctx context.Context, sceneKey *datastore.Key, n *node.Node, e *env.Env) error {
	span := trace.FromContext(ctx).NewChild("PutNode")
	defer span.Finish()

	nodeKey := datastore.NameKey("Node", n.Id, sceneKey)
	if _, err := e.DatastoreClient.Put(ctx, nodeKey, n); err != nil {
		return errors.Wrapf(err, "Failed to put node with digest %v", n.Id)
	}

	return nil
}

func PutParams(ctx context.Context, id string, p *params.Parameters, e *env.Env) error {
	span := trace.FromContext(ctx).NewChild("PutParams")
	defer span.Finish()

	bucket := e.StorageClient.Bucket(os.Getenv("STORAGE_BUCKET"))

	paramsObj := bucket.Object("params/" + id)

	w := paramsObj.NewWriter(ctx)

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return errors.Wrap(err, "Error encoding JSON")
	}

	if _, err := fmt.Fprint(w, string(jsonBytes)); err != nil {
		return errors.Wrap(err, "Failed to write to Storage")
	}

	if err := w.Close(); err != nil {
		return errors.Wrap(err, "Failed to close params file")
	}

	_, err = paramsObj.Update(ctx, storage.ObjectAttrsToUpdate{
		ContentType: "application/json",
	})
	if err != nil {
		return errors.Wrap(err, "Failed to update attributes")
	}

	return nil
}

func GetAllHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("GetAllHandler")
	defer span.Finish()

	vars := mux.Vars(r)

	sceneId := vars["scene_id"]
	sceneKey := datastore.NameKey("Scene", sceneId, nil)

	var nodes []*node.Node
	_, err := env.DatastoreClient.GetAll(ctx, datastore.NewQuery("Node").Ancestor(sceneKey), &nodes)
	if err != nil {
		return errors.Wrap(err, "Could not get all")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonData, _ := json.Marshal(&nodes)
	fmt.Fprint(w, string(jsonData))

	return nil
}

func GetHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("GetHandler")
	defer span.Finish()

	vars := mux.Vars(r)

	sceneId := vars["scene_id"]
	nodeId := vars["node_id"]

	sceneKey := datastore.NameKey("Scene", sceneId, nil)
	nodeKey := datastore.NameKey("Node", nodeId, sceneKey)

	var existingNode node.Node
	if err := env.DatastoreClient.Get(ctx, nodeKey, &existingNode); err != nil {
		return errors.Wrap(err, "Could not get datastore entity")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	nodeAsJson, err := json.Marshal(&existingNode)
	if err != nil {
		return errors.Wrap(err, "failed to marshal json")
	}

	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

type RequestBody struct {
	Node   *node.Node `json:"node"`
	Params params.Parameters `json:"params"`
}

func PostHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("PostHandler")
	defer span.Finish()

	vars := mux.Vars(r)

	sceneId := vars["scene_id"]
	nodeUuid := uuid.NewV4()

	sceneKey := datastore.NameKey("Scene", sceneId, nil)

	requestBody := RequestBody{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		return errors.Wrap(err, "Failed to decode json request body")
	}

	newNode := requestBody.Node
	newParams := requestBody.Params

	defer r.Body.Close()

	newNode.Id = nodeUuid.String()
	newNode.Version = 0
	PutNode(ctx, sceneKey, newNode, env)
	PutParams(ctx, newNode.Id, &newParams, env)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	nodeAsJson, err := json.Marshal(newNode)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal json")
	}

	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

func PutHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("PutHandler")
	defer span.Finish()

	vars := mux.Vars(r)
	sceneId := vars["scene_id"]
	nodeId := vars["node_id"]

	sceneKey := datastore.NameKey("Scene", sceneId, nil)
	nodeKey := datastore.NameKey("Node", nodeId, sceneKey)

	decoder := json.NewDecoder(r.Body)

	requestBody := RequestBody{}

	if err := decoder.Decode(&requestBody); err != nil {
		return errors.Wrap(err, "Failed to decode request body")
	}

	newNode := requestBody.Node

	_, err := env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingNode node.Node
		if err := tx.Get(nodeKey, &existingNode); err != nil {
			return errors.Wrap(err, "Failed to get node in tx")
		}
		existingNode.Version += 1
		existingNode.Name = newNode.Name
		existingNode.Inputs = newNode.Inputs
		PutNode(ctx, sceneKey, &existingNode, env)

		_, err := tx.Put(nodeKey, &existingNode)
		if err != nil {
			return errors.Wrap(err, "Failed to put node in tx")
		}
		return nil
	})

	if err != nil {
		return errors.Wrap(err, "Failed to update node")
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	nodeAsJson, err := json.Marshal(&newNode)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal json")
	}

	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

func DeleteAllHandler(ctx context.Context, e *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("DeleteAllHandler")
	defer span.Finish()

	vars := mux.Vars(r)
	sceneId := vars["scene_id"]

	sceneKey := datastore.NameKey("Scene", sceneId, nil)

	var nodes interface{}
	nodeIds, err := e.DatastoreClient.GetAll(ctx, datastore.NewQuery("Node").Ancestor(sceneKey).KeysOnly(), nodes)
	if err != nil {
		return errors.Wrap(err, "Failed to get all keys only")
	}

	err = e.DatastoreClient.DeleteMulti(ctx, nodeIds)
	if err != nil {
		return errors.Wrap(err, "Failed to delete multi")
	}

	return nil
}

func DeleteHandler(ctx context.Context, e *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("DeleteHandler")
	defer span.Finish()

	return nil
}

type ConnectionData struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ConnectHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("ConnectHandler")
	defer span.Finish()

	decoder := json.NewDecoder(r.Body)

	connection := ConnectionData{}

	if err := decoder.Decode(&connection); err != nil {
		return errors.Wrap(err, "Failed to decode connection data")
	}

	vars := mux.Vars(r)
	sceneId := vars["scene_id"]

	sceneKey := datastore.NameKey("Scene", sceneId, nil)

	_, err := env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		fromNodeKey := datastore.NameKey("Node", connection.From, sceneKey)
		toNodeKey := datastore.NameKey("Node", connection.To, sceneKey)

		var fromNode node.Node
		if err := tx.Get(fromNodeKey, &fromNode); err != nil {
			return errors.Wrap(err, "Failed to get fromNode in tx")
		}

		var toNode node.Node
		if err := tx.Get(toNodeKey, &toNode); err != nil {
			return errors.Wrap(err, "Failed to get toNode in tx")
		}

		if !stringInSlice(connection.To, fromNode.Outputs) {
			fromNode.Outputs = append(fromNode.Outputs, connection.To)

			_, err := tx.Put(fromNodeKey, &fromNode)
			if err != nil {
				return errors.Wrap(err, "Failed to put FromNode")
			}
		}

		if !stringInSlice(connection.From, toNode.Inputs) {
			toNode.Inputs = append(toNode.Inputs, connection.From)

			if err := PutNode(ctx, sceneKey, &toNode, env); err != nil {
				return errors.Wrap(err, "Failed to put toNode")
			}
		}

		return nil
	})

	return err
}

func DisconnectHandler(ctx context.Context, env *env.Env, w http.ResponseWriter, r *http.Request) error {
	span := trace.FromContext(ctx).NewChild("DisconnectHandler")
	defer span.Finish()

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/scenes/{scene_id}/nodes").Subrouter()

	s.Handle("/", handler.Handler{environ, GetAllHandler}).Methods("GET")
	s.Handle("/{node_id}/", handler.Handler{environ, GetHandler}).Methods("GET")

	s.Handle("/", handler.Handler{environ, PostHandler}).Methods("POST")
	s.Handle("/{node_id}/", handler.Handler{environ, PutHandler}).Methods("PUT")
	s.Handle("/DELETE/", handler.Handler{environ, DeleteAllHandler}).Methods("DELETE")
	s.Handle("/{node_id}/", handler.Handler{environ, DeleteHandler}).Methods("DELETE")

	s.Handle("/connect/", handler.Handler{environ, ConnectHandler}).Methods("PUT")
	s.Handle("/disconnect/", handler.Handler{environ, DisconnectHandler}).Methods("PUT")
}
