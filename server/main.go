package main

import (
	"github.com/stupschwartz/qubit/server/api"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"time"
	"github.com/stupschwartz/qubit/server/env"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
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

	serviceCredentials := option.WithServiceAccountFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	datastoreClient, err := datastore.NewClient(ctx, projID, serviceCredentials)
	for err != nil {
		log.Printf("Could not create datastore client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		datastoreClient, err = datastore.NewClient(ctx, projID, serviceCredentials)
	}

	storageClient, err := storage.NewClient(ctx, serviceCredentials)
	for err != nil {
		log.Printf("Could not create storage client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		storageClient, err = storage.NewClient(ctx, serviceCredentials)
	}

	environ := &env.Env{
		DatastoreClient: datastoreClient,
		StorageClient: storageClient,
	}


	log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", api.Handlers(environ)))
}
