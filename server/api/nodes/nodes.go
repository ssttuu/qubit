package nodes

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/node"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"log"
	"net/http"
	"github.com/satori/go.uuid"
	"encoding/base64"
	"github.com/stupschwartz/qubit/server/api/render"
	"cloud.google.com/go/datastore"
	"strings"
)

func GetDigestFromNode(n *node.Node) string {
	digestHasher := sha256.New()
	digestHasher.Write([]byte(n.Type))
	digestHasher.Write([]byte(strings.Join(n.Inputs, ",")))

	//paramsAsJson, _ := json.Marshal(&n.Params)
	//digestHasher.Write([]byte(paramsAsJson))
	return base64.URLEncoding.EncodeToString(digestHasher.Sum(nil))
}


func GetAllHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	var nodes []*node.Node
	_, err := env.DatastoreClient.GetAll(env.Context, datastore.NewQuery("Node"), &nodes)
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
	if err := env.DatastoreClient.Get(env.Context, nodeKey, &existingNode); err != nil {
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

func PostHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	log.Println("Posting")

	nodeUuid := uuid.NewV4()
	log.Println("UUID: " + nodeUuid.String())

	decoder := json.NewDecoder(r.Body)

	newNode := node.Node{}
	log.Println("new Node")

	if err := decoder.Decode(&newNode); err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	newNode.Id = nodeUuid.String()
	newNode.Digest = GetDigestFromNode(&newNode)
	newNode.Version = 0

	log.Println("set values")

	newNodeKey := datastore.NameKey("Node", newNode.Id, nil)

	log.Println("newNodeKey")

	if _, err := env.DatastoreClient.Put(env.Context, newNodeKey, &newNode); err != nil {
		log.Fatalf("Failed to save node: %v", err)
	}

	log.Println("Put item")

	//nodeChanged(env, n, nodeDigest)
	go render.RenderNodeAndDependents(newNode.Id)

	log.Println("rendered dependents")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	nodeAsJson, err := json.Marshal(&newNode)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("json data")
	log.Println(string(nodeAsJson))

	fmt.Fprintf(w, string(nodeAsJson))

	log.Println("made it")

	return nil
}

func PutHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]
	nodeKey := datastore.NameKey("Node", whereNodeId, nil)

	decoder := json.NewDecoder(r.Body)

	newNode := node.Node{}

	if err := decoder.Decode(&newNode); err != nil {
		log.Fatal(err)
	}

	_, err := env.DatastoreClient.RunInTransaction(env.Context, func(tx *datastore.Transaction) error {
		var existingNode node.Node
		if err := tx.Get(nodeKey, &existingNode); err != nil {
			return err
		}
		existingNode.Version += 1
		existingNode.Name = newNode.Name
		//existingNode.Params = newNode.Params
		existingNode.Inputs = newNode.Inputs
		existingNode.Digest = GetDigestFromNode(&existingNode)

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
	To string `json:"to"`
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

	_, err := env.DatastoreClient.RunInTransaction(env.Context, func(tx *datastore.Transaction) error {
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
			toNode.Digest = GetDigestFromNode(&toNode)

			_, err := tx.Put(toNodeKey, &toNode)
			if err != nil {
				log.Fatalf("Failed to put ToNode: %v", err)
			}
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
