package organizations

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/api/lib/pgutils"
	"github.com/stupschwartz/qubit/core/organization"
	organizations_pb "github.com/stupschwartz/qubit/proto-gen/go/organizations"
)

var organizationsTable = "organizations"

type Server struct {
	PostgresClient *sqlx.DB
}

func (s *Server) List(ctx context.Context, in *organizations_pb.ListOrganizationsRequest) (*organizations_pb.ListOrganizationsResponse, error) {
	// TODO: Permissions
	var orgs organization.Organizations
	err := pgutils.Select(&pgutils.SelectConfig{
		DB:    s.PostgresClient,
		Table: organizationsTable,
	}, &orgs)
	if err != nil {
		return nil, err
	}
	return &organizations_pb.ListOrganizationsResponse{Organizations: orgs.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *organizations_pb.GetOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Permissions
	var org organization.Organization
	err := pgutils.SelectByID(&pgutils.SelectConfig{
		DB:    s.PostgresClient,
		Table: organizationsTable,
		Id:    in.GetId(),
	}, &org)
	if err != nil {
		return nil, err
	}
	return org.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *organizations_pb.CreateOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Validation
	createConfig := pgutils.InsertConfig{
		Columns: []string{"name"},
		DB:      s.PostgresClient,
		Values: [][]interface{}{
			{in.Organization.Name},
		},
		Table: organizationsTable,
	}
	newOrganization := organization.Organization{
		Name: in.Organization.Name,
	}
	err := pgutils.InsertOne(&createConfig, &newOrganization.Id)
	if err != nil {
		return nil, err
	}
	return newOrganization.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *organizations_pb.UpdateOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Permissions & validation
	tx, err := s.PostgresClient.Beginx()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to begin transaction")
	}
	var org organization.Organization
	err = pgutils.SelectByID(&pgutils.SelectConfig{
		ForClause: "FOR UPDATE",
		Id:        in.GetId(),
		Table:     organizationsTable,
		Tx:        tx,
	}, &org)
	if err != nil {
		return nil, err
	}
	// TODO: Make update fields dynamic
	newOrganization := organization.NewFromProto(in.Organization)
	if newOrganization.Name != org.Name {
		org.Name = newOrganization.Name
		err = pgutils.UpdateByID(&pgutils.UpdateConfig{
			Id:    org.Id,
			Table: organizationsTable,
			Tx:    tx,
			Updates: map[string]interface{}{
				"name": newOrganization.Name,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to commit transaction")
	}
	return org.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *organizations_pb.DeleteOrganizationRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	err := pgutils.DeleteByID(&pgutils.DeleteConfig{
		DB:    s.PostgresClient,
		Table: organizationsTable,
		Id:    in.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	organizations_pb.RegisterOrganizationsServer(grpcServer, &Server{PostgresClient: postgresClient})
}
