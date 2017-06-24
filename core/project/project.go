package project

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/projects"
)

const TableName = "projects"

type Project struct {
	Id             string `json:"id" db:"id"`
	OrganizationId string `json:"organization_id" db:"organization_id"`
	Name           string `json:"name" db:"name"`
}

type Projects []Project

func NewFromProto(pbproject *pb.Project) Project {
	return Project{
		Id:             pbproject.GetId(),
		OrganizationId: pbproject.GetOrganizationId(),
		Name:           pbproject.GetName(),
	}
}

func (p *Project) ToProto() *pb.Project {
	return &pb.Project{
		Id:             p.Id,
		OrganizationId: p.OrganizationId,
		Name:           p.Name,
	}
}

func (p *Project) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"organization_id": p.OrganizationId,
		"name":            p.Name,
	}
}

func (p *Project) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"name": p.Name,
	}
}

func (p *Project) ValidateCreate() error {
	return nil
}

func (p *Project) ValidateUpdate(newObj interface{}) error {
	//project := newObj.(*Project)
	return nil
}

func (p *Projects) ToProto() []*pb.Project {
	var pbprojects []*pb.Project
	for _, project := range *p {
		project_proto := project.ToProto()
		pbprojects = append(pbprojects, project_proto)
	}
	return pbprojects
}
