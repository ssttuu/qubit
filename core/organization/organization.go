package organization

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/organizations"
)

type Organization struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Organizations []Organization

func NewFromProto(pborganization *pb.Organization) Organization {
	return Organization{
		Id:   pborganization.GetId(),
		Name: pborganization.GetName(),
	}
}

func (o *Organization) ToProto() *pb.Organization {
	return &pb.Organization{
		Id:   o.Id,
		Name: o.Name,
	}
}

func (o *Organization) ValidateCreate() error {
	return nil
}

func (o *Organization) ValidateUpdate(newOrg interface{}) error {
	//org := newOrg.(*Organization)
	return nil
}

func (o *Organizations) ToProto() []*pb.Organization {
	var pborganizations []*pb.Organization
	for _, organization := range *o {
		pborganizations = append(pborganizations, organization.ToProto())
	}
	return pborganizations
}
