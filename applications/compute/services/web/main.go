package main

import (
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/stupschwartz/qubit/applications/compute/services/web/computations"
)

func serve(server *grpc.Server, listener net.Listener, done chan bool) {
	server.Serve(listener)
	done <- true
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}
	projID := os.Getenv("GOOGLE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`You need to set the environment variable "GOOGLE_PROJECT_ID"`)
	}
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
	ctx := context.Background()
	serviceCredentials := option.WithServiceAccountFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	pubSubClient, err := pubsub.NewClient(ctx, projID, serviceCredentials)
	for err != nil {
		log.Printf("Could not create pubsub client: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		pubSubClient, err = pubsub.NewClient(ctx, projID, serviceCredentials)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	servingDone := make(chan bool)
	go serve(grpcServer, lis, servingDone)
	computations.Register(grpcServer, postgresClient, pubSubClient)
	<-servingDone
}
