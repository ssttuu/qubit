package main

import (
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/stupschwartz/qubit/applications/api/services/web/image_sequences"
	"github.com/stupschwartz/qubit/applications/api/services/web/images"
	"github.com/stupschwartz/qubit/applications/api/services/web/organizations"
	"github.com/stupschwartz/qubit/applications/api/services/web/projects"
	"github.com/stupschwartz/qubit/applications/api/services/web/scenes"
	"github.com/stupschwartz/qubit/proto-gen/go/computations"
)

func runServer(server *grpc.Server, listener net.Listener, done chan bool) {
	server.Serve(listener)
	done <- true
}

func main() {
	ctx := context.Background()
	serviceCredentials := option.WithServiceAccountFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	storageClient, err := storage.NewClient(ctx, serviceCredentials)
	for err != nil {
		log.Printf("Could not create storage client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		storageClient, err = storage.NewClient(ctx, serviceCredentials)
	}
	apiAddress := os.Getenv("API_WEB_SERVICE_ADDRESS")
	if apiAddress == "" {
		log.Fatal(`You need to set the environment variable "API_WEB_SERVICE_ADDRESS"`)
	}
	conn, err := grpc.Dial(apiAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Api Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		conn, err = grpc.Dial(apiAddress, grpc.WithInsecure())
	}
	defer conn.Close()
	computeAddress := os.Getenv("COMPUTE_WEB_SERVICE_ADDRESS")
	if computeAddress == "" {
		log.Fatal(`You need to set the environment variable "COMPUTE_WEB_SERVICE_ADDRESS"`)
	}
	computeConn, err := grpc.Dial(computeAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Compute Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		computeConn, err = grpc.Dial(computeAddress, grpc.WithInsecure())
	}
	defer computeConn.Close()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	servingDone := make(chan bool)
	grpcServer := grpc.NewServer()
	go runServer(grpcServer, lis, servingDone)
	computationsClient := computations.NewComputationsClient(computeConn)
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL == "" {
		log.Fatal(`You need to set the environment variable "POSTGRES_URL"`)
	}
	parsedURL, err := pq.ParseURL(postgresURL)
	if err != nil {
		log.Fatal(err)
	}
	postgresClient, err := sqlx.Open("postgres", parsedURL)
	if err != nil {
		log.Fatal(err)
	}
	organizations.Register(grpcServer, postgresClient)
	projects.Register(grpcServer, postgresClient)
	scenes.Register(grpcServer, postgresClient, storageClient, computationsClient)
	images.Register(grpcServer, postgresClient)
	image_sequences.Register(grpcServer, postgresClient)
	<-servingDone
}
