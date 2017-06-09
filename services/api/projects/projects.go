package projects

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/project"
	projects_pb "github.com/stupschwartz/qubit/proto-gen/go/projects"
)

type Server struct {
	PostgresClient *sqlx.DB
}

func (s *Server) List(ctx context.Context, in *projects_pb.ListProjectsRequest) (*projects_pb.ListProjectsResponse, error) {
	// TODO: Permissions
	var projectList project.Projects
	err := s.PostgresClient.Select(&projectList, "SELECT * FROM projects")
	if err != nil {
		return nil, errors.Wrap(err, "Could not select projects")
	}
	return &projects_pb.ListProjectsResponse{Projects: projectList.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *projects_pb.GetProjectRequest) (*projects_pb.Project, error) {
	// TODO: Permissions
	var sc project.Project
	err := s.PostgresClient.Get(&sc, "SELECT * FROM projects WHERE id=$1", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get project with ID %v", in.Id)
	}
	return sc.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *projects_pb.CreateProjectRequest) (*projects_pb.Project, error) {
	// TODO: Validation
	result, err := s.PostgresClient.NamedExec(
		`INSERT INTO projects (organization_id, name) VALUES (:organization_id, :name)`,
		map[string]interface{}{
			"organization_id": in.Project.OrganizationId,
			"name":            in.Project.Name,
		},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create project, %s", in.Project.Name)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve new ID")
	}
	newProject := project.Project{
		Id:   id,
		Name: in.Project.Name,
	}
	return newProject.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *projects_pb.UpdateProjectRequest) (*projects_pb.Project, error) {
	// TODO: Permissions & validation
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for project with ID %v", in.Id)
	}
	txStmt, err := tx.Prepare(`SELECT * FROM projects WHERE id=? FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select project in tx %v", in.Id)
	}
	row := txStmt.QueryRow(in.Id)
	if row == nil {
		return nil, errors.Wrapf(err, "No project with ID %v exists", in.Id)
	}
	var existingProject project.Project
	err = row.Scan(&existingProject)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load project from row")
	}
	// TODO: Make update fields dynamic
	newProject := project.NewProjectFromProto(in.Project)
	if newProject.Name != existingProject.Name {
		existingProject.Name = newProject.Name
		_, err = tx.Exec("UPDATE projects SET name=? WHERE id=?", newProject.Name, in.Id)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update project with ID %v", in.Id)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update project")
	}
	return existingProject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *projects_pb.DeleteProjectRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	_, err := s.PostgresClient.Queryx("DELETE FROM projects WHERE id=?", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted project by id: %v", in.Id)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	projects_pb.RegisterProjectsServer(grpcServer, &Server{PostgresClient: postgresClient})
}
