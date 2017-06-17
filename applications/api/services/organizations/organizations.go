package organizations

import (
	"strconv"

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
	orgId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	var org organization.Organization
	err = s.PostgresClient.Get(&org, "SELECT * FROM organizations WHERE id=$1", int64(orgId))
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get organization with ID %v", orgId)
	}
	return org.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *organizations_pb.CreateOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Validation
	query := `INSERT INTO organizations (name) VALUES (:name) RETURNING id`
	stmt, err := s.PostgresClient.PrepareNamed(query)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to prepare statement, %s", query)
	}
	var id int64
	err = stmt.Get(&id, map[string]interface{}{
		"name": in.Organization.Name,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create organization, %s", in.Organization.Name)
	}
	newOrganization := organization.Organization{
		Id:   strconv.FormatInt(id, 10),
		Name: in.Organization.Name,
	}
	return newOrganization.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *organizations_pb.UpdateOrganizationRequest) (*organizations_pb.Organization, error) {
	// TODO: Permissions & validation
	orgId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %s", in.GetId())
	}
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for organization with ID %v", orgId)
	}
	txStmt, err := tx.Prepare(`SELECT id, name FROM organizations WHERE id=$1 FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select organization in tx %v", orgId)
	}
	row := txStmt.QueryRow(orgId)
	if row == nil {
		return nil, errors.Wrapf(err, "No organization with ID %v exists", orgId)
	}
	var existingOrganization organization.Organization
	err = row.Scan(&existingOrganization.Id, &existingOrganization.Name)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load organization from row")
	}
	// TODO: Make update fields dynamic
	newOrganization := organization.NewFromProto(in.Organization)
	if newOrganization.Name != existingOrganization.Name {
		existingOrganization.Name = newOrganization.Name
		_, err = tx.Exec("UPDATE organizations SET name=$1 WHERE id=$2", newOrganization.Name, orgId)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update organization with ID %v", orgId)
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
	orgId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %s", in.GetId())
	}
	_, err = s.PostgresClient.Queryx("DELETE FROM organizations WHERE id=$1", orgId)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted organization by id: %v", orgId)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	organizations_pb.RegisterOrganizationsServer(grpcServer, &Server{PostgresClient: postgresClient})
}
