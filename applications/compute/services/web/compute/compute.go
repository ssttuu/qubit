package compute

import (
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/api/lib/apiutils"
	"github.com/stupschwartz/qubit/core/computation"
	compute_pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
)

var computationsTable = "computations"

type Server struct {
	PostgresClient *sqlx.DB
	PubSubClient   *pubsub.Client
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, pubSubClient *pubsub.Client) {
	compute_pb.RegisterComputeServer(grpcServer, &Server{
		PostgresClient: postgresClient,
		PubSubClient:   pubSubClient,
	})
}

func (s *Server) CreateComputation(ctx context.Context, in *compute_pb.CreateComputationRequest) (*compute_pb.Computation, error) {
	// Ignore returned errors, which will be "already exists". If they're fatal
	// errors, subsequent calls will also fail.
	topic, _ := s.PubSubClient.CreateTopic(ctx, computation.PubsubTopicID)
	newObject := computation.NewFromProto(in.Computation)
	// TODO: Enum?
	newObject.Status = "created"
	/*
		Begin Critical section:
		    - Create computation in Postgres
		    - Publish computation message in PubSub
	*/
	// TODO: Maybe operators should not be stored directly in computation
	// TODO: to avoid responding with operator data
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  computationsTable,
	})
	if err != nil {
		return nil, err
	}
	// TODO: Use gRPC for serialization of messages instead of JSON
	msg := map[string]string{"computation_id": newObject.Id}
	messageData, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to Marshal JSON %v", msg)
	}
	result := topic.Publish(ctx, &pubsub.Message{Data: messageData})
	// Wait for server confirmation
	if _, err := result.Get(ctx); err != nil {
		// TODO: Retry?
		return nil, errors.Wrapf(err, "Failed to publish message for new computation %v", newObject)
	}
	// TODO: Update Postgres with indication that message was published? If so, use transaction for both operations
	/*
		End Critical section
	*/
	// TODO: Don't return operators
	return newObject.ToProto(), nil
}

func (s *Server) GetComputation(ctx context.Context, in *compute_pb.GetComputationRequest) (*compute_pb.Computation, error) {
	var obj computation.Computation
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: computationsTable,
	})
	if err != nil {
		return nil, err
	}
	// TODO: Don't return operators
	return obj.ToProto(), nil
}
