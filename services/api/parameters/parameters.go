package parameters

import (
	"math/rand"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/organization"
	"github.com/stupschwartz/qubit/core/parameter"
	"github.com/stupschwartz/qubit/core/scene"
	parameters_pb "github.com/stupschwartz/qubit/proto-gen/go/parameters"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Server struct {
	DatastoreClient *datastore.Client
}

func (s *Server) List(ctx context.Context, in *parameters_pb.ListParametersRequest) (*parameters_pb.ListParametersResponse, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.SceneId, sceneKey)

	var parameters parameter.Parameters
	_, err := s.DatastoreClient.GetAll(ctx, datastore.NewQuery(parameter.Kind).Ancestor(operatorKey), &parameters)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	return &parameters_pb.ListParametersResponse{Parameters: parameters.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *parameters_pb.GetParameterRequest) (*parameters_pb.Parameter, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(parameter.Kind, in.ParameterId, operatorKey)

	var existingParameter parameter.Parameter
	if err := s.DatastoreClient.Get(ctx, parameterKey, &existingParameter); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingParameter.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *parameters_pb.CreateParameterRequest) (*parameters_pb.Parameter, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(parameter.Kind, in.Parameter.Id, operatorKey)

	newParameter := parameter.NewParameterFromProto(in.Parameter)

	if _, err := s.DatastoreClient.Put(ctx, parameterKey, &newParameter); err != nil {
		return nil, errors.Wrapf(err, "Failed to put operator %v", newParameter.Id)
	}

	return newParameter.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *parameters_pb.UpdateParameterRequest) (*parameters_pb.Parameter, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(parameter.Kind, in.ParameterId, operatorKey)

	newParameter := parameter.NewParameterFromProto(in.Parameter)

	_, err := s.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
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
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)
	operatorKey := datastore.NameKey(operator.Kind, in.OperatorId, sceneKey)
	parameterKey := datastore.NameKey(parameter.Kind, in.ParameterId, operatorKey)

	if err := s.DatastoreClient.Delete(ctx, parameterKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to delete operator by id: %v", in.OperatorId)
	}

	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, datastoreClient *datastore.Client) {
	parameters_pb.RegisterParametersServer(grpcServer, &Server{DatastoreClient: datastoreClient})
}

func NewClient(conn *grpc.ClientConn) parameters_pb.ParametersClient {
	return parameters_pb.NewParametersClient(conn)
}
