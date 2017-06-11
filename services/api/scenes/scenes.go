package scenes

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/scene"
	scenes_pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

type Server struct {
	PostgresClient *sqlx.DB
}

func (s *Server) List(ctx context.Context, in *scenes_pb.ListScenesRequest) (*scenes_pb.ListScenesResponse, error) {
	// TODO: Permissions
	var sceneList scene.Scenes
	err := s.PostgresClient.Select(&sceneList, "SELECT * FROM scenes")
	if err != nil {
		return nil, errors.Wrap(err, "Could not select scenes")
	}
	return &scenes_pb.ListScenesResponse{Scenes: sceneList.ToProto(), NextPageToken: ""}, nil
}

func (s *Server) Get(ctx context.Context, in *scenes_pb.GetSceneRequest) (*scenes_pb.Scene, error) {
	// TODO: Permissions
	var sc scene.Scene
	err := s.PostgresClient.Get(&sc, "SELECT * FROM scenes WHERE id=$1", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get scene with ID %v", in.Id)
	}
	return sc.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *scenes_pb.CreateSceneRequest) (*scenes_pb.Scene, error) {
	// TODO: Validation
	result, err := s.PostgresClient.NamedExec(
		`INSERT INTO scenes (project_id, name) VALUES (:project_id, :name)`,
		map[string]interface{}{
			"project_id": in.Scene.ProjectId,
			"name":       in.Scene.Name,
		},
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create scene, %s", in.Scene.Name)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve new ID")
	}
	newScene := scene.Scene{
		Id:   string(id),
		Name: in.Scene.Name,
	}
	return newScene.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *scenes_pb.UpdateSceneRequest) (*scenes_pb.Scene, error) {
	// TODO: Permissions & validation
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for scene with ID %v", in.Id)
	}
	txStmt, err := tx.Prepare(`SELECT * FROM scenes WHERE id=? FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select scene in tx %v", in.Id)
	}
	row := txStmt.QueryRow(in.Id)
	if row == nil {
		return nil, errors.Wrapf(err, "No scene with ID %v exists", in.Id)
	}
	var existingScene scene.Scene
	err = row.Scan(&existingScene)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load scene from row")
	}
	// TODO: Make update fields dynamic
	newScene := scene.NewSceneFromProto(in.Scene)
	if newScene.Name != existingScene.Name {
		existingScene.Name = newScene.Name
		_, err = tx.Exec("UPDATE scenes SET name=? WHERE id=", newScene.Name, in.Id)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update scene with ID %v", in.Id)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update scene")
	}
	return existingScene.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *scenes_pb.DeleteSceneRequest) (*empty.Empty, error) {
	// TODO: Permissions
	// TODO: Delete dependent entities with service calls
	_, err := s.PostgresClient.Queryx("DELETE FROM scenes WHERE id=?", in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted scene by id: %v", in.Id)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	scenes_pb.RegisterScenesServer(grpcServer, &Server{PostgresClient: postgresClient})
}
