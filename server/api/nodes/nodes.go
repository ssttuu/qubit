package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/node"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"log"
	"net/http"
	"github.com/satori/go.uuid"
	"cloud.google.com/go/datastore"
	"context"
	"cloud.google.com/go/storage"
	"os"
	"github.com/stupschwartz/qubit/params"
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

// TODO: do this concurrently
// TODO: use transaction
func PutNode(n *node.Node, e *env.Env) {
	nodeKey := datastore.NameKey("Node", n.Id, nil)
	if _, err := e.DatastoreClient.Put(context.Background(), nodeKey, n); err != nil {
		log.Fatalf("Failed to put node with digest, %v", err)
	}
}

func PutParams(id string, p *params.Parameters, e *env.Env) {
	bucket := e.StorageClient.Bucket(os.Getenv("STORAGE_BUCKET"))

	paramsObj := bucket.Object("params/" + id)

	ctx := context.Background()
	w := paramsObj.NewWriter(ctx)

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		log.Fatal("Error encoding JSON")
	}

	if _, err := fmt.Fprint(w, string(jsonBytes)); err != nil {
		log.Fatal("Failed to write to Storage")
	}

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	_, err = paramsObj.Update(ctx, storage.ObjectAttrsToUpdate{
		ContentType: "application/json",
	})
	if err != nil {
		log.Fatalf("Failed to update attributes: %v", err)
	}

}

func GetAllHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	var nodes []*node.Node
	ctx := context.Background()
	_, err := env.DatastoreClient.GetAll(ctx, datastore.NewQuery("Node"), &nodes)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonData, _ := json.Marshal(&nodes)
	fmt.Fprint(w, string(jsonData))

	return nil
}

func GetHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]

	nodeKey := datastore.NameKey("Node", whereNodeId, nil)

	var existingNode node.Node
	ctx := context.Background()
	if err := env.DatastoreClient.Get(ctx, nodeKey, &existingNode); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	nodeAsJson, err := json.Marshal(&existingNode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

type RequestBody struct {
	Node   *node.Node `json:"node"`
	Params params.Parameters `json:"params"`
}

func PostHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	log.Println("Posting")

	nodeUuid := uuid.NewV4()
	log.Println("UUID: " + nodeUuid.String())

	requestBody := RequestBody{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		log.Fatal(err)
	}

	newNode := requestBody.Node
	newParams := requestBody.Params

	defer r.Body.Close()

	newNode.Id = nodeUuid.String()
	newNode.Version = 0
	PutNode(newNode, env)
	PutParams(newNode.Id, &newParams, env)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	nodeAsJson, err := json.Marshal(newNode)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(nodeAsJson))
	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

func PutHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]
	nodeKey := datastore.NameKey("Node", whereNodeId, nil)

	decoder := json.NewDecoder(r.Body)

	requestBody := RequestBody{}

	if err := decoder.Decode(&requestBody); err != nil {
		log.Fatal(err)
	}

	newNode := requestBody.Node

	ctx := context.Background()
	_, err := env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingNode node.Node
		if err := tx.Get(nodeKey, &existingNode); err != nil {
			return err
		}
		existingNode.Version += 1
		existingNode.Name = newNode.Name
		existingNode.Inputs = newNode.Inputs
		PutNode(&existingNode, env)

		_, err := tx.Put(nodeKey, &existingNode)
		return err
	})

	if err != nil {
		log.Fatalf("Failed to update node, %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	nodeAsJson, err := json.Marshal(&newNode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

func DeleteHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
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

func ConnectHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	connection := ConnectionData{}

	if err := decoder.Decode(&connection); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	_, err := env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		fromNodeKey := datastore.NameKey("Node", connection.From, nil)
		toNodeKey := datastore.NameKey("Node", connection.To, nil)

		var fromNode node.Node
		if err := tx.Get(fromNodeKey, &fromNode); err != nil {
			return err
		}

		var toNode node.Node
		if err := tx.Get(toNodeKey, &toNode); err != nil {
			return err
		}

		if !stringInSlice(connection.To, fromNode.Outputs) {
			fromNode.Outputs = append(fromNode.Outputs, connection.To)

			_, err := tx.Put(fromNodeKey, &fromNode)
			if err != nil {
				log.Fatalf("Failed to put FromNode: %v", err)
			}
		}

		if !stringInSlice(connection.From, toNode.Inputs) {
			toNode.Inputs = append(toNode.Inputs, connection.From)
			PutNode(&toNode, env)
		}

		return nil
	})

	return err
}

func DisconnectHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/nodes").Subrouter()

	s.Handle("/", handler.Handler{environ, GetAllHandler}).Methods("GET")
	s.Handle("/{id}/", handler.Handler{environ, GetHandler}).Methods("GET")

	s.Handle("/", handler.Handler{environ, PostHandler}).Methods("POST")
	s.Handle("/{id}", handler.Handler{environ, PutHandler}).Methods("PUT")
	s.Handle("/{id}", handler.Handler{environ, DeleteHandler}).Methods("DELETE")

	s.Handle("/connect/", handler.Handler{environ, ConnectHandler}).Methods("PUT")
	s.Handle("/disconnect/", handler.Handler{environ, DisconnectHandler}).Methods("PUT")
}
