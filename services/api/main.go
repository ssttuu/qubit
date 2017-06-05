package main

import (
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/context"
	"cloud.google.com/go/datastore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/api/option"
	"cloud.google.com/go/storage"

	"github.com/stupschwartz/qubit/services/api/organizations"
	"github.com/stupschwartz/qubit/services/api/scenes"
	"github.com/stupschwartz/qubit/services/api/operators"
	"github.com/stupschwartz/qubit/services/api/parameters"

	"github.com/stupschwartz/qubit/proto-gen/go/compute"
)

func serve(server *grpc.Server, listener net.Listener, done chan bool) {
	server.Serve(listener)
	done <- true
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

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	servingDone := make(chan bool)
	go serve(grpcServer, lis, servingDone)

	apiAddress := os.Getenv("API_SERVICE_ADDRESS")
	if apiAddress == "" {
		log.Fatal(`You need to set the environment variable "API_SERVICE_ADDRESS"`)
	}

	conn, err := grpc.Dial(apiAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Api Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		conn, err = grpc.Dial(apiAddress, grpc.WithInsecure())
	}
	defer conn.Close()

	parametersClient := parameters.NewClient(conn)

	computeAddress := os.Getenv("COMPUTE_SERVICE_ADDRESS")
	if apiAddress == "" {
		log.Fatal(`You need to set the environment variable "COMPUTE_SERVICE_ADDRESS"`)
	}

	computeConn, err := grpc.Dial(computeAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Compute Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		conn, err = grpc.Dial(computeAddress, grpc.WithInsecure())
	}
	defer conn.Close()

	computeClient := compute.NewComputeClient(computeConn)

	organizations.Register(grpcServer, datastoreClient)
	scenes.Register(grpcServer, datastoreClient)
	parameters.Register(grpcServer, datastoreClient)
	operators.Register(grpcServer, datastoreClient, storageClient, parametersClient, computeClient)
	//images.Register(grpcServer, datastoreClient, storageClient)

	<-servingDone
}