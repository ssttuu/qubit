package main

import (
	"log"
	_ "github.com/lib/pq"
	"time"
	"github.com/stupschwartz/qubit/server/env"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/trace"
	"google.golang.org/api/option"

	"golang.org/x/net/context"
	"os"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/server/api/health"
	"github.com/stupschwartz/qubit/server/api/organizations"
	"github.com/stupschwartz/qubit/server/api/scenes"
	"github.com/stupschwartz/qubit/server/api/operators"
	"github.com/stupschwartz/qubit/server/api/parameters"
	"net"
)

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func init() {

}

func main() {
	projID := os.Getenv("GOOGLE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_PROJECT_ID"`)
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

	traceClient, err := trace.NewClient(ctx, projID)
	if err != nil {
		log.Fatalf("Could not create trace client: %v\n", err)
	}
	p, err := trace.NewLimitedSampler(1.0, 10)
	if err != nil {
		log.Fatalf("Could not create limited sampler: %v\n", err)
	}
	traceClient.SetSamplingPolicy(p)

	environ := &env.Env{
		DatastoreClient: datastoreClient,
		StorageClient: storageClient,
		TraceClient: traceClient,
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	health.Register(grpcServer, environ)
	organizations.Register(grpcServer, environ)
	scenes.Register(grpcServer, environ)
	operators.Register(grpcServer, environ)
	parameters.Register(grpcServer, environ)

	grpcServer.Serve(lis)
}
