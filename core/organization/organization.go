package organization

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/organizations"
)

type Organization struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (o *Organization) ToProto() *pb.Organization {
	return &pb.Organization{Id: o.Id, Name: o.Name}
}

func NewOrganizationFromProto(pborganization *pb.Organization) Organization {
	return Organization{
		Id:   pborganization.Id,
		Name: pborganization.Name,
	}
}

type Organizations []*Organization

func (o *Organizations) ToProto() []*pb.Organization {
	var pborganizations []*pb.Organization
	for _, organization := range *o {
		pborganizations = append(pborganizations, organization.ToProto())
	}
	return pborganizations
}
