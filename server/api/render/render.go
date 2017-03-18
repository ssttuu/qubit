package render

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/env"
	"github.com/stupschwartz/qubit/server/handler"
	"net/http"
	"github.com/gocql/gocql"
	"github.com/stupschwartz/qubit/node"
	"encoding/json"
	"log"
)

func PostHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]

	var (
		nodeId gocql.UUID
		nodeVersion int
		nodeData string
		nodeDigest string
	)

	selectNodeQuery := `
		SELECT node_id, node_version, node_data, node_digest FROM qubit.nodes WHERE node_id=?;
	`

	err := env.CqlSession.Query(selectNodeQuery, whereNodeId).Consistency(gocql.One).Scan(&nodeId, &nodeVersion, &nodeData, &nodeDigest)

	if err != nil {
		log.Fatal(err)
	}

	n := node.Node{}
	if err := json.Unmarshal([]byte(nodeData), &n); err != nil {
		log.Fatal(err)
	}

	message := struct {
		Node   node.Node `json:"node"`
		Digest string `json:"digest"`
	}{
		Node: n,
		Digest: nodeDigest,
	}

	jsonData, _ := json.Marshal(&message)

	err = env.NsqProducer.Publish("render", jsonData)
	if err != nil {
		log.Panic("Could not connect")
	}

	w.WriteHeader(http.StatusProcessing)

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/render").Subrouter()

	s.Handle("/{id}/", handler.Handler{environ, PostHandler}).Methods("POST")
}
