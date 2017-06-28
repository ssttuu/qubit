package operators

import (
	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/operator"
	_ "github.com/stupschwartz/qubit/core/operators"
	computations_pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

type Server struct {
	PostgresClient     *sqlx.DB
	StorageClient      *storage.Client
	ComputationsClient computations_pb.ComputationsClient
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, storageClient *storage.Client, computationsClient computations_pb.ComputationsClient) {
	operators_pb.RegisterOperatorsServer(grpcServer, &Server{
		ComputationsClient: computationsClient,
		PostgresClient:     postgresClient,
		StorageClient:      storageClient,
	})
}

func (s *Server) Create(ctx context.Context, in *operators_pb.CreateOperatorRequest) (*operators_pb.Operator, error) {
	newObject := operator.NewFromProto(in.Operator)
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  operator.TableName,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *operators_pb.DeleteOperatorRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Table: operator.TableName,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *operators_pb.GetOperatorRequest) (*operators_pb.Operator, error) {
	var obj operator.Operator
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: operator.TableName,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *operators_pb.ListOperatorsRequest) (*operators_pb.ListOperatorsResponse, error) {
	var objectList operator.Operators
	err := apiutils.List(&apiutils.ListConfig{
		// Don't load parameters for list because it's a large column
		Columns: []string{"id", "scene_id", "context", "type", "name"},
		DB:      s.PostgresClient,
		Out:     &objectList,
		Table:   operator.TableName,
	})
	if err != nil {
		return nil, err
	}
	return &operators_pb.ListOperatorsResponse{Operators: objectList.ToProto()}, nil
}

func (s *Server) Render(ctx context.Context, in *operators_pb.RenderOperatorRequest) (*operators_pb.RenderOperatorResponse, error) {
	// TODO: Forward request to compute-web
	return nil, nil
}

func (s *Server) GetRenderParameters(ctx context.Context, in *operators_pb.RenderParameterRequest) (*operators_pb.RenderParameter, error) {
	// TODO: Use operator key to get parameter inputs and configuration
	return nil, nil
}

func (s *Server) Update(ctx context.Context, in *operators_pb.UpdateOperatorRequest) (*operators_pb.Operator, error) {
	newObject := operator.NewFromProto(in.Operator)
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetId(),
		NewObject: &newObject,
		OldObject: &operator.Operator{},
		Table:     operator.TableName,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}
