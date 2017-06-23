package compute

import (
	"encoding/json"
	"strconv"

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
	newComputation.Id = strconv.FormatInt(compKey.ID, 10)
	// TODO: Enum?
	newComputation.Status = "created"
	// TODO: Use gRPC for serialization of messages instead of JSON
	msg := map[string]string{"computation_id": newComputation.Id}
	messageData, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to Marshal JSON %v", msg)
	}
	// Ignore returned errors, which will be "already exists". If they're fatal
	// errors, subsequent calls will also fail.
	topic, _ := s.PubSubClient.CreateTopic(ctx, computation.PubsubTopicID)
	/*
		Begin Critical section:
		    - Create computation in Datastore
		    - Publish computation message in PubSub
	*/
	// TODO: Maybe operators should not be stored directly in computation
	// TODO: to avoid responding with operator data
	if _, err := s.DatastoreClient.Put(ctx, compKey, &newComputation); err != nil {
		return nil, errors.Wrapf(err, "Failed to create new computation %v", newComputation)
	}
	result := topic.Publish(ctx, &pubsub.Message{Data: messageData})
	// Wait for server confirmation
	// TODO: Retry?
	if _, err := result.Get(ctx); err != nil {
		return nil, errors.Wrapf(err, "Failed to publish message for new computation %v", newComputation)
	}
	// TODO: Update Datastore with indication that message was published? If so, use transaction for both operations
	/*
		End Critical section
	*/
	// TODO: Don't return operators
	return newComputation.ToProto(), nil
}

func (s *Server) GetComputation(ctx context.Context, in *compute_pb.GetComputationRequest) (*compute_pb.Computation, error) {
	objId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert %v to integer", in.GetId())
	}
	compKey := datastore.IDKey(computation.Kind, objId, nil)
	var comp computation.Computation
	if err := s.DatastoreClient.Get(ctx, compKey, &comp); err != nil {
		return nil, errors.Wrapf(err, "Failed to get computation ID %v", objId)
	}
	// TODO: Don't return operators
	return comp.ToProto(), nil
}
