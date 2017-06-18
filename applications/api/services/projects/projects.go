package projects

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/api/lib/apiutils"
	"github.com/stupschwartz/qubit/core/project"
	projects_pb "github.com/stupschwartz/qubit/proto-gen/go/projects"
)

var projectsTable = "projects"

type Server struct {
	PostgresClient *sqlx.DB
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	projects_pb.RegisterProjectsServer(grpcServer, &Server{PostgresClient: postgresClient})
}

func (s *Server) Create(ctx context.Context, in *projects_pb.CreateProjectRequest) (*projects_pb.Project, error) {
	newObject := project.NewFromProto(in.Project)
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  projectsTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *projects_pb.DeleteProjectRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Table: projectsTable,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *projects_pb.GetProjectRequest) (*projects_pb.Project, error) {
	var obj project.Project
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: projectsTable,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *projects_pb.ListProjectsRequest) (*projects_pb.ListProjectsResponse, error) {
	var objectList project.Projects
	err := apiutils.List(&apiutils.ListConfig{
		DB:    s.PostgresClient,
		Out:   &objectList,
		Table: projectsTable,
	})
	if err != nil {
		return nil, err
	}
	return &projects_pb.ListProjectsResponse{Projects: objectList.ToProto()}, nil
}

func (s *Server) Update(ctx context.Context, in *projects_pb.UpdateProjectRequest) (*projects_pb.Project, error) {
	newObject := project.NewFromProto(in.Project)
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetId(),
		NewObject: &newObject,
		OldObject: &project.Project{},
		Table:     projectsTable,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}
