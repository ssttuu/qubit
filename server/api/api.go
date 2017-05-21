package api

import (
	"net"
	"google.golang.org/grpc/grpclog"
	"os"
)

func Handlers(environ *env.Env)  {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background()

	traceClient, err := trace.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"))
	if err != nil {
		log.Fatalf("Could not create trace client: %v\n", err)
	}
	p, err := trace.NewLimitedSampler(1.0, 10)
	if err != nil {
		log.Fatalf("Could not create limited sampler: %v\n", err)
	}
	traceClient.SetSamplingPolicy(p)

	log.Println("MAIN")

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(serverInterceptor(traceClient)))
	pb.RegisterComputeServer(grpcServer, newServer(traceClient))
	grpcServer.Serve(lis)


}
