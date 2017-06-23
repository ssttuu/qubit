package compute

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	compute_pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
)

type Server struct {
	DatastoreClient *datastore.Client
}

func (s *Server) CreateComputation(ctx context.Context, in *compute_pb.CreateComputationRequest) (*compute_pb.ComputationStatusResponse, error) {
	return &compute_pb.ComputationStatusResponse{}, nil
}

func (s *Server) GetComputationStatus(ctx context.Context, in *compute_pb.ComputationStatusRequest) (*compute_pb.ComputationStatusResponse, error) {
	return &compute_pb.ComputationStatusResponse{}, nil
}

func Register(grpcServer *grpc.Server, datastoreClient *datastore.Client) {
	compute_pb.RegisterComputeServer(grpcServer, &Server{DatastoreClient: datastoreClient})
}
