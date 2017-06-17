package scenes

import (
	"strconv"

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
	scene_id, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	var sc scene.Scene
	err = s.PostgresClient.Get(&sc, "SELECT * FROM scenes WHERE id=$1", scene_id)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not get scene with ID %v", scene_id)
	}
	return sc.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *scenes_pb.CreateSceneRequest) (*scenes_pb.Scene, error) {
	// TODO: Validation
	query := `INSERT INTO scenes (project_id, name) VALUES (:project_id, :name) RETURNING id`
	stmt, err := s.PostgresClient.PrepareNamed(query)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to prepare statement, %s", query)
	}
	var id int64
	err = stmt.Get(&id, map[string]interface{}{
		"project_id": in.Scene.ProjectId,
		"name":       in.Scene.Name,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create scene, %s", in.Scene.Name)
	}
	newScene := scene.Scene{
		Id:   strconv.FormatInt(id, 10),
		Name: in.Scene.Name,
	}
	return newScene.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *scenes_pb.UpdateSceneRequest) (*scenes_pb.Scene, error) {
	// TODO: Permissions & validation
	scene_id, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	tx, err := s.PostgresClient.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to begin transaction for scene with ID %v", scene_id)
	}
	txStmt, err := tx.Prepare(`SELECT id, name FROM scenes WHERE id=$1 FOR UPDATE`)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to select scene in tx %v", scene_id)
	}
	row := txStmt.QueryRow(scene_id)
	if row == nil {
		return nil, errors.Wrapf(err, "No scene with ID %v exists", scene_id)
	}
	var existingScene scene.Scene
	err = row.Scan(&existingScene.Id, &existingScene.Name)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load scene from row")
	}
	// TODO: Make update fields dynamic
	newScene := scene.NewSceneFromProto(in.Scene)
	if newScene.Name != existingScene.Name {
		existingScene.Name = newScene.Name
		_, err = tx.Exec("UPDATE scenes SET name=$1 WHERE id=$2", newScene.Name, scene_id)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to update scene with ID %v", scene_id)
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
	scene_id, err := strconv.ParseInt(in.GetId(), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not convert to integer %v", in.GetId())
	}
	_, err = s.PostgresClient.Queryx("DELETE FROM scenes WHERE id=$1", scene_id)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted scene by id: %v", scene_id)
	}
	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	scenes_pb.RegisterScenesServer(grpcServer, &Server{PostgresClient: postgresClient})
}
