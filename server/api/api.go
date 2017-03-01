package api

import (
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/stupschwartz/qubit/server/api/nodes"
	"github.com/stupschwartz/qubit/server/env"
	"log"
)

//func RenderHandler(w http.ResponseWriter, r *http.Request) {
//	profile := operator.Operator{Name: "read1", Type: "Read", Params: map[string]interface{}{"Param1": "Stu"}, Inputs: []string{"input1", "input2"}}
//
//	js, err := json.Marshal(profile)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(js)
//}

func Handlers() *mux.Router {
	router := mux.NewRouter()

	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "qubit"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()

	environ := &env.Env{
		CqlSession: session,
	}

	createNodesTableStatement := `
		CREATE TABLE IF NOT EXISTS qubit.nodes (
			node_id uuid,
			node_version int,
			node_data text,
			node_digest text,
			PRIMARY KEY (node_id, node_version)
		) WITH CLUSTERING ORDER BY (node_version ASC);
	`

	if err := environ.CqlSession.Query(createNodesTableStatement).Exec(); err != nil {
		log.Fatal(err)
	}

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	nodes.Register(apiRouter, environ)

	//apiRouter.HandleFunc("/render", NodeHandler).Methods("POST")
	//apiRouter.HandleFunc("/render", RenderHandler).Methods("POST")

	return router
}
