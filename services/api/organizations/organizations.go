package organizations

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/organization"
	organizations_pb "github.com/stupschwartz/qubit/proto-gen/go/organizations"
)

type Server struct {
	PostgresClient *sqlx.DB
}

func (s *Server) List(ctx context.Context, in *organizations_pb.ListOrganizationsRequest) (*organizations_pb.ListOrganizationsResponse, error) {
	// TODO: Permissions
	var orgs organization.Organizations
	err := s.PostgresClient.Select(&orgs, "SELECT * FROM organizations")
	if err != nil {
		return nil, errors.Wrap(err, "Could not select organizations")
	}
	return &organizations_pb.ListOrganizationsResponse{Organizations: orgs.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *organizations_pb.GetOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Permissions
	var org organization.Organization
	err := s.PostgresClient.Get(&org, "SELECT * FROM organizations WHERE id=$1", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get organization with ID %v", in.Id)
	}
	return org.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *organizations_pb.CreateOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Validation
	result, err := s.PostgresClient.NamedExec(
		`INSERT INTO organizations (name) VALUES (:name)`,
		map[string]interface{}{
			"name": in.Organization.Name,
		},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create organization, %s", in.Organization.Name)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve new ID")
	}
	newOrganization := organization.Organization{
		Id:   id,
		Name: in.Organization.Name,
	}
	return newOrganization.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *organizations_pb.UpdateOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Permissions & validation
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for organization with ID %v", in.Id)
	}
	txStmt, err := tx.Prepare(`SELECT * FROM organizations WHERE id=? FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select organization in tx %v", in.Id)
	}
	row := txStmt.QueryRow(in.Id)
	if row == nil {
		return nil, errors.Wrapf(err, "No organization with ID %v exists", in.Id)
	}
	var existingOrganization organization.Organization
	err = row.Scan(&existingOrganization)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load organization from row")
	}
	// TODO: Make update fields dynamic
	newOrganization := organization.NewOrganizationFromProto(in.Organization)
	if newOrganization.Name != existingOrganization.Name {
		existingOrganization.Name = newOrganization.Name
		_, err = tx.Exec("UPDATE organizations SET name=? WHERE id=?", newOrganization.Name, in.Id)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update organization with ID %v", in.Id)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update organization")
	}
	return existingOrganization.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *organizations_pb.DeleteOrganizationRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	_, err := s.PostgresClient.Queryx("DELETE FROM organizations WHERE id=?", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted organization by id: %v", in.Id)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	organizations_pb.RegisterOrganizationsServer(grpcServer, &Server{PostgresClient: postgresClient})
}
