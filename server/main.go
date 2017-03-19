package main

import (
	"github.com/stupschwartz/qubit/server/api"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"time"
	"github.com/stupschwartz/qubit/server/env"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"

	"golang.org/x/net/context"
	"os"
)

type Credentials struct {
    Cid string `json:"cid"`
    Csecret string `json:"csecret"`
}

func init() {

}

func main() {
	projID := os.Getenv("DATASTORE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "DATASTORE_PROJECT_ID"`)
	}

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projID, option.WithServiceAccountFile(os.Getenv("DATASTORE")))
	for err != nil {
		log.Printf("Could not create datastore client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		client, err = datastore.NewClient(ctx, projID)
	}

	environ := &env.Env{
		DatastoreClient: client,
		Context: ctx,
	}


	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", api.Handlers(environ)))
}
