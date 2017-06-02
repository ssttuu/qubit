package organization

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/organizations"
	"fmt"
	"github.com/pkg/errors"
)

const Kind string = "Organization"

type Organization struct {
	Id   string `json:"id" datastore:"id"`
	Name string `json:"name" datastore:"name"`
}

func (o *Organization) ToProto() (*pb.Organization, error) {
	return &pb.Organization{Id: o.Id, Name: o.Name}, nil
}

func NewOrganizationFromProto(pborganization *pb.Organization) Organization {
	return Organization{
		Id: fmt.Sprint(pborganization.Id),
		Name: pborganization.Name,
	}
}

type Organizations []*Organization

func (o *Organizations) ToProto() ([]*pb.Organization, error) {
	var pborganizations []*pb.Organization
	for _, organization := range *o {
		organization_proto, err := organization.ToProto()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to convert organization to proto, %v", organization)
		}
		pborganizations = append(pborganizations, organization_proto)
	}

	return pborganizations, nil
}
