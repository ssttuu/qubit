package organizations

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"
	"cloud.google.com/go/datastore"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/organization"
	organizations_pb "github.com/stupschwartz/qubit/proto-gen/go/organizations"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Server struct {
	DatastoreClient *datastore.Client
}

func (s *Server) List(ctx context.Context, in *organizations_pb.ListOrganizationsRequest) (*organizations_pb.ListOrganizationsResponse, error) {
	var organizations organization.Organizations
	_, err := s.DatastoreClient.GetAll(ctx, datastore.NewQuery(organization.Kind), &organizations)
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
	organizationKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)

	var existingOrganization organization.Organization
	if err := s.DatastoreClient.Get(ctx, organizationKey, &existingOrganization); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingOrganization.ToProto()
}

func (s *Server) Create(ctx context.Context, in *organizations_pb.CreateOrganizationRequest) (*organizations_pb.Organization, error) {
	in.Organization.Id = fmt.Sprint(r.Int63())
	organizationKey := datastore.NameKey(organization.Kind, in.Organization.Id, nil)

	newOrganization := organization.NewOrganizationFromProto(in.Organization)

	if _, err := s.DatastoreClient.Put(ctx, organizationKey, &newOrganization); err != nil {
		return nil, errors.Wrapf(err, "Failed to put organization, %v", newOrganization)
	}

	return newOrganization.ToProto()
}

func (s *Server) Update(ctx context.Context, in *organizations_pb.UpdateOrganizationRequest) (*organizations_pb.Organization, error) {
	organizationKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)

	newOrganization := organization.NewOrganizationFromProto(in.Organization)

	_, err := s.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
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
	organizationKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)

	if err := s.DatastoreClient.Delete(ctx, organizationKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted organization by id: %v", in.OrganizationId)
	}

	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, datastoreClient *datastore.Client) {
	organizations_pb.RegisterOrganizationsServer(grpcServer, &Server{DatastoreClient: datastoreClient})
}