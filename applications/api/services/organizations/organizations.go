package organizations

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/api/lib/apiutils"
	"github.com/stupschwartz/qubit/core/organization"
	organizations_pb "github.com/stupschwartz/qubit/proto-gen/go/organizations"
)

var organizationsTable = "organizations"

type Server struct {
	PostgresClient *sqlx.DB
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	organizations_pb.RegisterOrganizationsServer(grpcServer, &Server{PostgresClient: postgresClient})
}

func (s *Server) Create(ctx context.Context, in *organizations_pb.CreateOrganizationRequest) (*organizations_pb.Organization, error) {
	newObject := organization.NewFromProto(in.Organization)
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  organizationsTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *organizations_pb.DeleteOrganizationRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Table: organizationsTable,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *organizations_pb.GetOrganizationRequest) (*organizations_pb.Organization, error) {
	var obj organization.Organization
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: organizationsTable,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *organizations_pb.ListOrganizationsRequest) (*organizations_pb.ListOrganizationsResponse, error) {
	var objectList organization.Organizations
	err := apiutils.List(&apiutils.ListConfig{
		DB:    s.PostgresClient,
		Out:   &objectList,
		Table: organizationsTable,
	})
	if err != nil {
		return nil, err
	}
	return &organizations_pb.ListOrganizationsResponse{Organizations: objectList.ToProto()}, nil
}

func (s *Server) Update(ctx context.Context, in *organizations_pb.UpdateOrganizationRequest) (*organizations_pb.Organization, error) {
	newObject := organization.NewFromProto(in.Organization)
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetId(),
		NewObject: &newObject,
		OldObject: &organization.Organization{},
		Table:     organizationsTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}
