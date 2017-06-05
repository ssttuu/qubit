package main


import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/stupschwartz/qubit/services/compute/compute"
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

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	servingDone := make(chan bool)
	go serve(grpcServer, lis, servingDone)

	compute.Register(grpcServer)

	<-servingDone
}
