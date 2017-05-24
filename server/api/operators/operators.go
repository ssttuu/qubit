package operators

import (
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/server/env"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	operators_pb "github.com/stupschwartz/qubit/server/protos/operators"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

const SceneKind string = "Scene"
const OperatorKind string = "Operator"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


// Server implements `service Health`.
type Server struct {
	env *env.Env
}

func (s *Server) List(ctx context.Context, in *operators_pb.ListOperatorsRequest) (*operators_pb.OperatorsList, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)

	var operators operator.Operators
	_, err := s.env.DatastoreClient.GetAll(ctx, datastore.NewQuery(OperatorKind).Ancestor(sceneKey), &operators)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	return operators.ToProto()
}

func (s *Server) Get(ctx context.Context, in *operators_pb.GetOperatorRequest) (*operators_pb.Operator, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)
	operatorKey := datastore.IDKey(OperatorKind, in.OperatorId, sceneKey)

	var existingOperator operator.Operator
	if err := s.env.DatastoreClient.Get(ctx, operatorKey, &existingOperator); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingOperator.ToProto()
}

func (s *Server) Create(ctx context.Context, in *operators_pb.CreateOperatorRequest) (*operators_pb.Operator, error) {
	in.Operator.Id = r.Int63()
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)
	operatorKey := datastore.IDKey(OperatorKind, in.Operator.Id, sceneKey)

	newOperator := operator.NewOperatorFromProto(in.Operator)

	if _, err := s.env.DatastoreClient.Put(ctx, operatorKey, &newOperator); err != nil {
		return nil, errors.Wrapf(err, "Failed to put operator %v", newOperator.Id)
	}

	return newOperator.ToProto()
}

func (s *Server) Update(ctx context.Context, in *operators_pb.UpdateOperatorRequest) (*operators_pb.Operator, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)
	operatorKey := datastore.IDKey(OperatorKind, in.OperatorId, sceneKey)

	newOperator := operator.NewOperatorFromProto(in.Operator)

	_, err := s.env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingOperator operator.Operator
		if err := tx.Get(operatorKey, &existingOperator); err != nil {
			return errors.Wrapf(err, "Failed to get operator in tx %v", existingOperator)
		}

		existingOperator.Name = newOperator.Name

		_, err := tx.Put(operatorKey, &existingOperator)
		if err != nil {
			return errors.Wrapf(err, "Failed to put operator in tx %v", existingOperator)
		}

		newOperator = existingOperator

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to update operator")
	}

	return newOperator.ToProto()
}

func (s *Server) Delete(ctx context.Context, in *operators_pb.DeleteOperatorRequest) (*empty.Empty, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)
	operatorKey := datastore.IDKey(OperatorKind, in.OperatorId, sceneKey)

	if err := s.env.DatastoreClient.Delete(ctx, operatorKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to delete operator by id: %v", in.OperatorId)
	}

	return &empty.Empty{}, nil
}

func newServer(e *env.Env) *Server {
	return &Server{
		env: e,
	}
}

func Register(server *grpc.Server, e *env.Env) {
	operators_pb.RegisterOperatorsServer(server, newServer(e))
}
