package compute

import (
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/pubsub"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/computation"
	compute_pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
)

type Server struct {
	DatastoreClient *datastore.Client
	PubSubClient    *pubsub.Client
}

func Register(grpcServer *grpc.Server, datastoreClient *datastore.Client, pubSubClient *pubsub.Client) {
	compute_pb.RegisterComputeServer(grpcServer, &Server{
		DatastoreClient: datastoreClient,
		PubSubClient:    pubSubClient,
	})
}

func (s *Server) CreateComputation(ctx context.Context, in *compute_pb.CreateComputationRequest) (*compute_pb.Computation, error) {
	newComputation := computation.NewFromProto(in.Computation)
	incompleteKey := datastore.IncompleteKey(computation.Kind, nil)
	completeKeys, err := s.DatastoreClient.AllocateIDs(ctx, []*datastore.Key{incompleteKey})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to allocate IDs for new computation %v", newComputation)
	} else if len(completeKeys) != 1 {
		return nil, errors.Wrapf(err, "Allocated %v keys instead of one key for new computation %v", len(completeKeys), newComputation)
	}
	compKey := completeKeys[0]
	// TODO: Enum?
	newComputation.Status = "created"
	if _, err := s.DatastoreClient.Put(ctx, compKey, &newComputation); err != nil {
		return nil, errors.Wrapf(err, "Failed to create new computation %v", newComputation)
	}
	// TODO: Publish message to PubSub
	return newComputation.ToProto(), nil
}

func (s *Server) GetComputation(ctx context.Context, in *compute_pb.GetComputationRequest) (*compute_pb.Computation, error) {
	return &compute_pb.Computation{}, nil
}
