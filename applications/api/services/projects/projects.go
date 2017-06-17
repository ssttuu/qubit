package projects

import (
	"strconv"

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
	projectId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	var sc project.Project
	err = s.PostgresClient.Get(&sc, "SELECT * FROM projects WHERE id=$1", projectId)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get project with ID %v", projectId)
	}
	return sc.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *projects_pb.CreateProjectRequest) (*projects_pb.Project, error) {
	// TODO: Validation
	query := `INSERT INTO projects (organization_id, name) VALUES (:organization_id, :name) RETURNING id`
	stmt, err := s.PostgresClient.PrepareNamed(query)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to prepare statement, %s", query)
	}
	var id int64
	err = stmt.Get(&id, map[string]interface{}{
		"organization_id": in.Project.OrganizationId,
		"name":            in.Project.Name,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create project, %s", in.Project.Name)
	}
	newProject := project.Project{
		Id:   strconv.FormatInt(id, 10),
		Name: in.Project.Name,
	}
	return newProject.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *projects_pb.UpdateProjectRequest) (*projects_pb.Project, error) {
	// TODO: Permissions & validation
	projectId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for project with ID %v", projectId)
	}
	txStmt, err := tx.Prepare(`SELECT id, name FROM projects WHERE id=$1 FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select project in tx %v", projectId)
	}
	row := txStmt.QueryRow(projectId)
	if row == nil {
		return nil, errors.Wrapf(err, "No project with ID %v exists", projectId)
	}
	var existingProject project.Project
	err = row.Scan(&existingProject.Id, &existingProject.Name)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load project from row")
	}
	// TODO: Make update fields dynamic
	newProject := project.NewFromProto(in.Project)
	if newProject.Name != existingProject.Name {
		existingProject.Name = newProject.Name
		_, err = tx.Exec("UPDATE projects SET name=$1 WHERE id=$2", newProject.Name, projectId)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update project with ID %v", projectId)
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
	projectId, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	_, err = s.PostgresClient.Queryx("DELETE FROM projects WHERE id=$1", projectId)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted project by id: %v", projectId)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	projects_pb.RegisterProjectsServer(grpcServer, &Server{PostgresClient: postgresClient})
}
