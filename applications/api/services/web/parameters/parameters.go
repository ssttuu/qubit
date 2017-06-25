package parameters

import (
	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/parameter"
	parameters_pb "github.com/stupschwartz/qubit/proto-gen/go/parameters"
)

type Server struct {
	PostgresClient *sqlx.DB
	StorageClient  *storage.Client
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, storageClient *storage.Client) {
	parameters_pb.RegisterParametersServer(grpcServer, &Server{
		PostgresClient: postgresClient,
		StorageClient:  storageClient,
	})
}

func (s *Server) Create(ctx context.Context, in *parameters_pb.CreateParameterRootRequest) (*parameters_pb.ParameterRoot, error) {
	newObject, err := parameter.NewRootFromProto(in.GetParameterRoot())
	if err != nil {
		return nil, err
	}
	err = apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  parameter.TableName,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto()
}

func (s *Server) Delete(ctx context.Context, in *parameters_pb.DeleteParameterRootRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetOperatorId(),
		Table: parameter.TableName,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *parameters_pb.GetParameterRootRequest) (*parameters_pb.ParameterRoot, error) {
	var obj parameter.ParameterRoot
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetOperatorId(),
		Out:   &obj,
		Table: parameter.TableName,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto()
}

func (s *Server) Update(ctx context.Context, in *parameters_pb.UpdateParameterRootRequest) (*parameters_pb.ParameterRoot, error) {
	newObject, err := parameter.NewRootFromProto(in.GetParameterRoot())
	if err != nil {
		return nil, err
	}
	err = apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetOperatorId(),
		NewObject: &newObject,
		OldObject: &parameter.ParameterRoot{},
		Table:     parameter.TableName,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto()
}
