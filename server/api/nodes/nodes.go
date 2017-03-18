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
	"github.com/gocql/gocql"
	"encoding/base64"
	"github.com/stupschwartz/qubit/server/api/render"
)

type CqlNodeRow struct {
	NodeId      gocql.UUID `json:"node_id"`
	NodeVersion int `json:"node_version"`
	NodeData    node.Node `json:"node_data"`
	NodeDigest  string `json:"node_digest"`
}

var allNodesQuery string = `
	SELECT node_id, node_version, node_data, node_digest FROM qubit.nodes;
`

var selectNodeQuery string = `
	SELECT node_id, node_version, node_data, node_digest FROM qubit.nodes WHERE node_id=?;
`

var createNodeQuery string = `
	INSERT INTO qubit.nodes (
		node_id,
		node_version,
		node_data,
		node_digest
	) VALUES (?, ?, ?, ?);
`

//func nodeChanged(env *env.Env, n node.Node, nodeDigest string) {
//	message := struct {
//		Node   node.Node `json:"node"`
//		Digest string `json:"digest"`
//	}{
//		Node: n,
//		Digest: nodeDigest,
//	}
//
//	jsonData, _ := json.Marshal(&message)
//
//	err := env.NsqProducer.Publish("node_changed", jsonData)
//	if err != nil {
//		log.Panic("Could not connect")
//	}
//
//	go render.RenderNodeAndDependents(n)
//}

func GetAllHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	var (
		nodeId gocql.UUID
		nodeVersion int
		nodeData string
		nodeDigest string
	)

	iter := env.CqlSession.Query(allNodesQuery).Iter()

	cqlNodes := []CqlNodeRow{}

	for iter.Scan(&nodeId, &nodeVersion, &nodeData, &nodeDigest) {
		n := node.Node{}
		if err := json.Unmarshal([]byte(nodeData), &n); err != nil {
			log.Fatal(err)
		}
		cqlNodes = append(cqlNodes, CqlNodeRow{NodeId: nodeId, NodeVersion: nodeVersion, NodeData: n, NodeDigest: nodeDigest})
	}

	jsonData, _ := json.Marshal(&cqlNodes)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(jsonData))

	return nil
}

func GetHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]

	var (
		nodeId gocql.UUID
		nodeVersion int
		nodeData string
		nodeDigest string
	)

	err := env.CqlSession.Query(selectNodeQuery, whereNodeId).Consistency(gocql.One).Scan(&nodeId, &nodeVersion, &nodeData, &nodeDigest)

	if err != nil {
		log.Fatal(err)
	}

	n := node.Node{}
	if err := json.Unmarshal([]byte(nodeData), &n); err != nil {
		log.Fatal(err)
	}
	cqlNodeRow := CqlNodeRow{NodeId: nodeId, NodeVersion: nodeVersion, NodeData: n, NodeDigest: nodeDigest}

	jsonData, _ := json.Marshal(&cqlNodeRow)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, string(jsonData))

	return nil
}

func PostHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	nodeUuid, err := gocql.RandomUUID()
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(r.Body)

	n := node.Node{}

	if err := decoder.Decode(&n); err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	nodeAsJson, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}

	digestHasher := sha256.New()
	digestHasher.Write([]byte(nodeAsJson))
	nodeDigest := base64.URLEncoding.EncodeToString(digestHasher.Sum(nil))

	query := env.CqlSession.Query(createNodeQuery, nodeUuid, 0, nodeAsJson, nodeDigest)

	if err := query.Exec(); err != nil {
		log.Fatal(err)
	}

	//nodeChanged(env, n, nodeDigest)
	go render.RenderNodeAndDependents(nodeUuid)


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

func PutHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	whereNodeId := vars["id"]

	selectNodeQuery := `
		SELECT node_version FROM qubit.nodes WHERE node_id=?;
	`

	var nodeVersion int

	if err := env.CqlSession.Query(selectNodeQuery, whereNodeId).Consistency(gocql.One).Scan(&nodeVersion); err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(r.Body)

	n := node.Node{}

	if err := decoder.Decode(&n); err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	nodeAsJson, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}

	digestHasher := sha256.New()
	digestHasher.Write([]byte(nodeAsJson))
	nodeDigest := base64.URLEncoding.EncodeToString(digestHasher.Sum(nil))

	query := env.CqlSession.Query(createNodeQuery, whereNodeId, nodeVersion + 1, nodeAsJson, nodeDigest)

	if err := query.Exec(); err != nil {
		log.Fatal(err)
	}

	//nodeChanged(env, n, nodeDigest)
	//go render.RenderNodeAndDependents(whereNodeId)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(nodeAsJson))

	return nil
}

func DeleteHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/nodes").Subrouter()

	s.Handle("/", handler.Handler{environ, GetAllHandler}).Methods("GET")
	s.Handle("/{id}/", handler.Handler{environ, GetHandler}).Methods("GET")

	s.Handle("/", handler.Handler{environ, PostHandler}).Methods("POST")
	s.Handle("/{id}", handler.Handler{environ, PutHandler}).Methods("PUT")
	s.Handle("/{id}", handler.Handler{environ, DeleteHandler}).Methods("DELETE")
}
