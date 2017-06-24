package computations

import (
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/computation"
	"github.com/stupschwartz/qubit/core/computation_status"
	computations_pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
)

type Server struct {
	PostgresClient *sqlx.DB
	PubSubClient   *pubsub.Client
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, pubSubClient *pubsub.Client) {
	computations_pb.RegisterComputationsServer(grpcServer, &Server{
		PostgresClient: postgresClient,
		PubSubClient:   pubSubClient,
	})
}

func (s *Server) CreateComputation(ctx context.Context, in *computations_pb.CreateComputationRequest) (*computations_pb.ComputationStatus, error) {
	// Ignore returned errors, which will be "already exists". If they're fatal
	// errors, subsequent calls will also fail.
	topic, _ := s.PubSubClient.CreateTopic(ctx, computation.PubsubTopicID)
	newComp := computation.NewFromProto(in.Computation)
	/*
		Begin Critical section:
		    - Create computation in Postgres
		    - Create computation status in Postgres
		    - Publish computation message in PubSub
	*/
	// FIXME: Use transaction for creation of computation and computation status in Postgres
	tx, err := s.PostgresClient.Beginx()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to begin transaction")
	}
	err = apiutils.Create(&apiutils.CreateConfig{
		Tx:     tx,
		Object: &newComp,
		Table:  computation.TableName,
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	newCompStatus := computation_status.ComputationStatus{
		ComputationId: newComp.Id,
		Status:        computation_status.ComputationStatusCreated,
		CreatedAt:     0, // FIXME: Epoch timestamp as int64?
	}
	err = apiutils.Create(&apiutils.CreateConfig{
		Tx:     tx,
		Object: &newCompStatus,
		Table:  computation_status.TableName,
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// TODO: Use gRPC for serialization of messages instead of JSON
	msg := map[string]string{"computation_id": newComp.Id}
	messageData, err := json.Marshal(msg)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(err, "Failed to Marshal JSON %v", msg)
	}
	result := topic.Publish(ctx, &pubsub.Message{Data: messageData})
	// Wait for server confirmation
	if _, err := result.Get(ctx); err != nil {
		// TODO: Retry?
		tx.Rollback()
		return nil, errors.Wrapf(err, "Failed to publish message for new computation %v", newComp)
	}
	// TODO: Update Postgres with indication that message was published? If so, use transaction for both operations
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to commit transaction")
	}
	/*
		End Critical section
	*/
	return newCompStatus.ToProto(), nil
}

func (s *Server) GetComputationStatus(ctx context.Context, in *computations_pb.GetComputationStatusRequest) (*computations_pb.ComputationStatus, error) {
	var obj computation_status.ComputationStatus
	err := apiutils.Get(&apiutils.GetConfig{
		DB: s.PostgresClient,
		// FIXME: This is wrong
		Id:    in.GetComputationId(),
		Out:   &obj,
		Table: computation_status.TableName,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}
