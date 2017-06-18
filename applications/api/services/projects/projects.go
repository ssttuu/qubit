package projects

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/api/lib/pgutils"
	"github.com/stupschwartz/qubit/core/project"
	projects_pb "github.com/stupschwartz/qubit/proto-gen/go/projects"
)

var projectsTable = "projects"

type Server struct {
	PostgresClient *sqlx.DB
}

func (s *Server) List(ctx context.Context, in *projects_pb.ListProjectsRequest) (*projects_pb.ListProjectsResponse, error) {
	// TODO: Permissions
	var orgs project.Projects
	err := pgutils.Select(&pgutils.SelectConfig{
		DB:    s.PostgresClient,
		Table: projectsTable,
	}, &orgs)
	if err != nil {
		return nil, err
	}
	return &projects_pb.ListProjectsResponse{Projects: orgs.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *projects_pb.GetProjectRequest) (*projects_pb.Project, error) {
	// TODO: Permissions
	var org project.Project
	err := pgutils.SelectByID(&pgutils.SelectConfig{
		DB:    s.PostgresClient,
		Table: projectsTable,
		Id:    in.GetId(),
	}, &org)
	if err != nil {
		return nil, err
	}
	return org.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *projects_pb.CreateProjectRequest) (*projects_pb.Project, error) {
	// TODO: Validation
	createConfig := pgutils.InsertConfig{
		Columns: []string{"organization_id", "name"},
		DB:      s.PostgresClient,
		Values: [][]interface{}{
			{in.Project.OrganizationId, in.Project.Name},
		},
		Table: projectsTable,
	}
	newProject := project.Project{
		Name: in.Project.Name,
	}
	err := pgutils.InsertOne(&createConfig, &newProject.Id)
	if err != nil {
		return nil, err
	}
	return newProject.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *projects_pb.UpdateProjectRequest) (*projects_pb.Project, error) {
	// TODO: Permissions & validation
	tx, err := s.PostgresClient.Beginx()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to begin transaction")
	}
	var org project.Project
	err = pgutils.SelectByID(&pgutils.SelectConfig{
		ForClause: "FOR UPDATE",
		Id:        in.GetId(),
		Table:     projectsTable,
		Tx:        tx,
	}, &org)
	if err != nil {
		return nil, err
	}
	// TODO: Make update fields dynamic
	newProject := project.NewFromProto(in.Project)
	if newProject.Name != org.Name {
		org.Name = newProject.Name
		err = pgutils.UpdateByID(&pgutils.UpdateConfig{
			Id:    org.Id,
			Table: projectsTable,
			Tx:    tx,
			Updates: map[string]interface{}{
				"name": newProject.Name,
			},
		})
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to commit transaction")
	}
	return org.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *projects_pb.DeleteProjectRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	err := pgutils.DeleteByID(&pgutils.DeleteConfig{
		DB:    s.PostgresClient,
		Table: projectsTable,
		Id:    in.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	projects_pb.RegisterProjectsServer(grpcServer, &Server{PostgresClient: postgresClient})
}
