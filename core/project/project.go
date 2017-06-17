package project

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/projects"
)

type Project struct {
	Id             string `json:"id" db:"id"`
	OrganizationId string `json:"organization_id" db:"organization_id"`
	Name           string `json:"name" db:"name"`
}

func (p *Project) ToProto() *pb.Project {
	return &pb.Project{
		Id:             p.Id,
		OrganizationId: p.OrganizationId,
		Name:           p.Name,
	}
}

func NewFromProto(pbproject *pb.Project) Project {
	return Project{
		Id:             pbproject.Id,
		OrganizationId: pbproject.OrganizationId,
		Name:           pbproject.Name,
	}
}

type Projects []Project

func (p *Projects) ToProto() []*pb.Project {
	var pbprojects []*pb.Project
	for _, project := range *p {
		project_proto := project.ToProto()
		pbprojects = append(pbprojects, project_proto)
	}
	return pbprojects
}
