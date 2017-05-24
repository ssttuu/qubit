package organizations

import (
	"github.com/stupschwartz/qubit/server/env"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/organization"
	organizations_pb "github.com/stupschwartz/qubit/server/protos/organizations"
	"github.com/golang/protobuf/ptypes/empty"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	"cloud.google.com/go/trace"
)

const OrganizationKind string = "Organization"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


// Server implements `service Health`.
type Server struct {
	env *env.Env
}

func (s *Server) List(ctx context.Context, in *organizations_pb.ListOrganizationsRequest) (*organizations_pb.ListOrganizationsResponse, error) {
	span := trace.FromContext(ctx).NewChild("organizations.List")
	defer span.Finish()

	var organizations organization.Organizations
	_, err := s.env.DatastoreClient.GetAll(ctx, datastore.NewQuery(OrganizationKind), &organizations)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	organizations_proto, err := organizations.ToProto()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert organizations to proto, %v", organizations)
	}

	return &organizations_pb.ListOrganizationsResponse{Organizations:organizations_proto, NextPageToken:""}, nil
}

func (s *Server) Get(ctx context.Context, in *organizations_pb.GetOrganizationRequest) (*organizations_pb.Organization, error) {
	organizationKey := datastore.IDKey(OrganizationKind, in.OrganizationId, nil)

	var existingOrganization organization.Organization
	if err := s.env.DatastoreClient.Get(ctx, organizationKey, &existingOrganization); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingOrganization.ToProto()
}

func (s *Server) Create(ctx context.Context, in *organizations_pb.CreateOrganizationRequest) (*organizations_pb.Organization, error) {
	in.Organization.Id = r.Int63()
	organizationKey := datastore.IDKey(OrganizationKind, in.Organization.Id, nil)

	newOrganization := organization.NewOrganizationFromProto(in.Organization)

	if _, err := s.env.DatastoreClient.Put(ctx, organizationKey, &newOrganization); err != nil {
		return nil, errors.Wrapf(err, "Failed to put organization, %v", newOrganization)
	}

	return newOrganization.ToProto()
}

func (s *Server) Update(ctx context.Context, in *organizations_pb.UpdateOrganizationRequest) (*organizations_pb.Organization, error) {
	organizationKey := datastore.IDKey(OrganizationKind, in.OrganizationId, nil)

	newOrganization := organization.NewOrganizationFromProto(in.Organization)

	_, err := s.env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingOrganization organization.Organization
		if err := tx.Get(organizationKey, &existingOrganization); err != nil {
			return errors.Wrapf(err, "Failed to get organization in tx %v", existingOrganization)
		}

		existingOrganization.Name = newOrganization.Name

		_, err := tx.Put(organizationKey, &existingOrganization)
		if err != nil {
			return errors.Wrapf(err, "Failed to put organization in tx %v", existingOrganization)
		}

		newOrganization = existingOrganization

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to update organization")
	}

	return newOrganization.ToProto()
}

func (s *Server) Delete(ctx context.Context, in *organizations_pb.DeleteOrganizationRequest) (*empty.Empty, error) {
	organizationKey := datastore.IDKey(OrganizationKind, in.OrganizationId, nil)

	if err := s.env.DatastoreClient.Delete(ctx, organizationKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted organization by id: %v", in.OrganizationId)
	}

	return &empty.Empty{}, nil
}

func newServer(e *env.Env) *Server {
	return &Server{
		env: e,
	}
}

func Register(server *grpc.Server, e *env.Env) {
	organizations_pb.RegisterOrganizationsServer(server, newServer(e))
}
