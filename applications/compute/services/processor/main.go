package main

import (
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/stupschwartz/qubit/applications/compute/services/processor/renders"
	render_parameters_pb "github.com/stupschwartz/qubit/proto-gen/go/render_operators"
)

func serve(server *grpc.Server, listener net.Listener, done chan bool) {
	server.Serve(listener)
	done <- true
}

func main() {
	computeProcessorAddress := os.Getenv("COMPUTE_PROCESSOR_SERVICE_ADDRESS")
	if computeProcessorAddress == "" {
		log.Fatal(`You need to set the environment variable "COMPUTE_PROCESSOR_SERVICE_ADDRESS"`)
	}
	conn, err := grpc.Dial(computeProcessorAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Compute Processor Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		conn, err = grpc.Dial(computeProcessorAddress, grpc.WithInsecure())
	}
	defer conn.Close()
	apiWebAddress := os.Getenv("API_WEB_SERVICE_ADDRESS")
	if apiWebAddress == "" {
		log.Fatal(`You need to set the environment variable "API_WEB_SERVICE_ADDRESS"`)
	}
	apiWebConn, err := grpc.Dial(apiWebAddress, grpc.WithInsecure())
	for err != nil {
		log.Printf("Could not connect to Api Web Service: %v\n", err)
		time.Sleep(100 * time.Millisecond)
		apiWebConn, err = grpc.Dial(apiWebAddress, grpc.WithInsecure())
	}
	defer apiWebConn.Close()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(`You need to set the environment variable "PORT"`)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	servingDone := make(chan bool)
	go serve(grpcServer, lis, servingDone)
	renderOperatorsClient := render_parameters_pb.NewRenderOperatorsClient(apiWebConn)
	renders.Register(grpcServer, renderOperatorsClient)
	<-servingDone
}
