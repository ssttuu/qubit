package organization

import (
	pb "github.com/stupschwartz/qubit/server/protos/organizations"
	"strconv"
	"fmt"
	"github.com/pkg/errors"
)

type Organization struct {
	Id   string `json:"id" datastore:"id"`
	Name string `json:"name" datastore:"name"`
}

func (s *Organization) ToProto() (*pb.Organization, error) {
	i, err := strconv.ParseInt(s.Id, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert Id from string to int64: %v", s.Id)
	}
	return &pb.Organization{Id: i}, nil
}

func NewOrganizationFromProto(pborganization *pb.Organization) Organization {
	return Organization{
		Id: fmt.Sprint(pborganization.Id),
		Name: pborganization.Name,
	}
}

type Organizations []*Organization

func (s *Organizations) ToProto() ([]*pb.Organization, error) {
	var pborganizations []*pb.Organization
	for _, organization := range *s {
		organization_proto, err := organization.ToProto()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to convert organization to proto, %v", organization)
		}
		pborganizations = append(pborganizations, organization_proto)
	}

	return pborganizations, nil
}
