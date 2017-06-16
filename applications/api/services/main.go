package main

import (
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/storage"
	_ "github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/stupschwartz/qubit/applications/api/services/images"
	"github.com/stupschwartz/qubit/applications/api/services/operators"
	"github.com/stupschwartz/qubit/applications/api/services/organizations"
	"github.com/stupschwartz/qubit/applications/api/services/projects"
	"github.com/stupschwartz/qubit/applications/api/services/scenes"
	"github.com/stupschwartz/qubit/proto-gen/go/compute"
)

func runServer(server *grpc.Server, listener net.Listener, done chan bool) {
	server.Serve(listener)
	done <- true
}

func main() {
	projID := os.Getenv("GOOGLE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_PROJECT_ID"`)
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}
	apiAddress := os.Getenv("API_SERVICE_ADDRESS")
	if apiAddress == "" {
		log.Fatal(`You need to set the environment variable "API_SERVICE_ADDRESS"`)
	}
	computeAddress := os.Getenv("COMPUTE_SERVICE_ADDRESS")
	if computeAddress == "" {
		log.Fatal(`You need to set the environment variable "COMPUTE_SERVICE_ADDRESS"`)
	}
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatal(`You need to set the environment variable "POSTGRES_URL"`)
	}
	ctx := context.Background()
	serviceCredentials := option.WithServiceAccountFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	storageClient, err := storage.NewClient(ctx, serviceCredentials)
	for err != nil {
		log.Printf("Could not create storage client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		storageClient, err = storage.NewClient(ctx, serviceCredentials)
	}
	postgresClient, err := sqlx.Open("postgres", postgresURL)
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	servingDone := make(chan bool)
	go runServer(grpcServer, lis, servingDone)
	conn, err := grpc.Dial(apiAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Api Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		conn, err = grpc.Dial(apiAddress, grpc.WithInsecure())
	}
	defer conn.Close()
	computeConn, err := grpc.Dial(computeAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Compute Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		conn, err = grpc.Dial(computeAddress, grpc.WithInsecure())
	}
	defer conn.Close()
	computeClient := compute.NewComputeClient(computeConn)
	organizations.Register(grpcServer, postgresClient)
	projects.Register(grpcServer, postgresClient)
	scenes.Register(grpcServer, postgresClient)
	operators.Register(grpcServer, postgresClient, storageClient, computeClient)
	images.Register(grpcServer, postgresClient, storageClient)
	<-servingDone
}
