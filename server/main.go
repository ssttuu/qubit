package main

import (
	"github.com/stupschwartz/qubit/server/api"
	"log"
	"net/http"
	"github.com/gocql/gocql"
	"github.com/bitly/go-nsq"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
	"github.com/stupschwartz/qubit/server/env"
)

func main() {
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "qubit"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()

	for err != nil {
		time.Sleep(100 * time.Millisecond)
		session, err = cluster.CreateSession()
	}

	config := nsq.NewConfig()
	nsqProducer, err := nsq.NewProducer("nsqd:4150", config)

	for err != nil {
		time.Sleep(100 * time.Millisecond)
		nsqProducer, err = nsq.NewProducer("nsqd:4150", config)
	}


	postgresdb, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres dbname=qubit sslmode=disable")

	log.Println(postgresdb)
	log.Println(err)

	err = postgresdb.Ping()

	failedTimes := 0
	for err != nil {
		failedTimes++
		log.Println("Fail! Sad!", failedTimes)
		time.Sleep(100 * time.Millisecond)
		err = postgresdb.Ping()
	}

	defer postgresdb.Close()

	environ := &env.Env{
		CqlSession: session,
		NsqProducer: nsqProducer,
		Postgres: postgresdb,
	}

	// Cassandra
	createNodesTableStatement := `
		CREATE TABLE IF NOT EXISTS qubit.node (
			node_id uuid,
			node_version int,
			data text,
			digest text,
			labels map<text, text>,
			inputs list<uuid>,
			outputs list<uuid>,
			PRIMARY KEY (node_id, node_version)
		) WITH CLUSTERING ORDER BY (node_version ASC);
	`

	// Postgres
	createTaskTableStatement := `
		CREATE TABLE IF NOT EXISTS "task" (
			task_id BIGSERIAL PRIMARY KEY,
			task_graph_id INT,
			node_id UUID,
			node_version INT,
			frame INT,
			bbox BOX,
			start_time TIMESTAMP,
			end_time TIMESTAMP,
			completed BOOLEAN
		);
	`

	// Postgres
	createTaskDependencyTableStatement := `
		CREATE TABLE IF NOT EXISTS task_dependency (
			task_dependency_id BIGSERIAL PRIMARY KEY,
			task_id BIGINT REFERENCES task (task_id),
			depends_on BIGINT REFERENCES task (task_id)
		);
	`


	if err := environ.CqlSession.Query(createNodesTableStatement).Exec(); err != nil {
		log.Fatal(err)
	}

	if _, err := environ.Postgres.Exec(createTaskTableStatement); err != nil {
		log.Fatal(err)
	}

	if _, err := environ.Postgres.Exec(createTaskDependencyTableStatement); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", api.Handlers(environ)))
}
