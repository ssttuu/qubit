package parameters

import (
	"github.com/stupschwartz/qubit/server/env"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	_ "github.com/stupschwartz/qubit/core/operators"
	parameters_pb "github.com/stupschwartz/qubit/server/protos/parameters"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stupschwartz/qubit/core/parameter"
)

const OrganizationKind string = "Organization"
const SceneKind string = "Scene"
const OperatorKind string = "Operator"
const ParameterKind string = "Parameter"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


// Server implements `service Health`.
type Server struct {
	env *env.Env
}

func (s *Server) List(ctx context.Context, in *parameters_pb.ListParametersRequest) (*parameters_pb.ListParametersResponse, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.SceneId, sceneKey)

	var parameters parameter.Parameters
	_, err := s.env.DatastoreClient.GetAll(ctx, datastore.NewQuery(ParameterKind).Ancestor(operatorKey), &parameters)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	return &parameters_pb.ListParametersResponse{Parameters:parameters.ToProto(), NextPageToken:""}, nil
}

func (s *Server) Get(ctx context.Context, in *parameters_pb.GetParameterRequest) (*parameters_pb.Parameter, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(ParameterKind, in.ParameterId, operatorKey)

	var existingParameter parameter.Parameter
	if err := s.env.DatastoreClient.Get(ctx, parameterKey, &existingParameter); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingParameter.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *parameters_pb.CreateParameterRequest) (*parameters_pb.Parameter, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(ParameterKind, in.Parameter.Id, operatorKey)

	newParameter := parameter.NewParameterFromProto(in.Parameter)

	if _, err := s.env.DatastoreClient.Put(ctx, parameterKey, &newParameter); err != nil {
		return nil, errors.Wrapf(err, "Failed to put operator %v", newParameter.Id)
	}

	return newParameter.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *parameters_pb.UpdateParameterRequest) (*parameters_pb.Parameter, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(ParameterKind, in.ParameterId, operatorKey)

	newParameter := parameter.NewParameterFromProto(in.Parameter)

	_, err := s.env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingParameter parameter.Parameter
		if err := tx.Get(parameterKey, &existingParameter); err != nil {
			return errors.Wrapf(err, "Failed to get operator in tx %v", existingParameter)
		}

		existingParameter.Id = newParameter.Id

		_, err := tx.Put(parameterKey, &existingParameter)
		if err != nil {
			return errors.Wrapf(err, "Failed to put operator in tx %v", existingParameter)
		}

		newParameter = &existingParameter

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to update operator")
	}

	return newParameter.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *parameters_pb.DeleteParameterRequest) (*empty.Empty, error) {
	orgKey := datastore.NameKey(OrganizationKind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(SceneKind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(OperatorKind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(ParameterKind, in.ParameterId, operatorKey)

	if err := s.env.DatastoreClient.Delete(ctx, parameterKey); err != nil {
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
	parameters_pb.RegisterParametersServer(server, newServer(e))
}
