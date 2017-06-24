package computations

import (
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/applications/lib/pgutils"
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
	newComp := computation.NewFromProto(in.Computation)
	// TODO: Validation
	topic, err := s.PubSubClient.CreateTopic(ctx, computation.PubSubTopicID)
	if err != nil {
		// 409 ALREADY_EXISTS is an inevitable and harmless error
		// https://cloud.google.com/pubsub/docs/reference/error-codes
		if statusErr, ok := status.FromError(err); !ok || statusErr.Code() != 409 {
			log.Println(err)
			return nil, status.Errorf(codes.Internal, "Internal error")
		}
	}
	tx, err := s.PostgresClient.Beginx()
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}
	newCompStatus := computation_status.ComputationStatus{
		Status:    computation_status.ComputationStatusCreated,
		CreatedAt: time.Now(),
	}
	// If anon func returns an error, rollback the transaction
	err = func() error {
		err = apiutils.Create(&apiutils.CreateConfig{
			Tx:     tx,
			Object: &newComp,
			Table:  computation.TableName,
		})
		if err != nil {
			return err
		}
		newCompStatus.ComputationId = newComp.Id
		err = apiutils.Create(&apiutils.CreateConfig{
			Tx:     tx,
			Object: &newCompStatus,
			Table:  computation_status.TableName,
		})
		if err != nil {
			return err
		}
		pbCompStatus := newCompStatus.ToProto()
		pbCompStatusData, err := proto.Marshal(pbCompStatus)
		if err != nil {
			return errors.Wrapf(err, "Failed to marshal computation status %v", pbCompStatus)
		}
		result := topic.Publish(ctx, &pubsub.Message{Data: pbCompStatusData})
		// Wait for server confirmation
		if _, err := result.Get(ctx); err != nil {
			// TODO: Retry?
			return errors.Wrapf(err, "Failed to publish message for new computation %v", newComp)
		}
		return nil
	}()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}
	// Committing after publishing means that a message may be published referencing
	// a missing record. However, the client will get a failure properly, so this
	// is an acceptable failure mode.
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}
	return newCompStatus.ToProto(), nil
}

func (s *Server) GetComputationStatus(ctx context.Context, in *computations_pb.GetComputationStatusRequest) (*computations_pb.ComputationStatus, error) {
	var compStatuses computation_status.ComputationStatuses
	err := pgutils.Select(&pgutils.SelectConfig{
		Args:        []interface{}{in.GetComputationId()},
		DB:          s.PostgresClient,
		Limit:       1,
		Out:         &compStatuses,
		Table:       computation_status.TableName,
		WhereClause: "computation_id=$1",
	})
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	} else if len(compStatuses) < 1 {
		return nil, status.Errorf(codes.NotFound, "Not found")
	}
	return compStatuses[0].ToProto(), nil
}
