package scenes

import (
	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/scene"
	"github.com/stupschwartz/qubit/core/scene_event"
	scenes_pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

type Server struct {
	PostgresClient *sqlx.DB
	StorageClient  *storage.Client
}

// Utilities

// API

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, storageClient *storage.Client) {
	scenes_pb.RegisterScenesServer(grpcServer, &Server{
		PostgresClient: postgresClient,
		StorageClient:  storageClient,
	})
}

// Create adds a new CreateScene event to the SceneEvent table and updates the current snapshot.
func (s *Server) Create(ctx context.Context, in *scenes_pb.CreateSceneRequest) (*scenes_pb.Scene, error) {
	newScene := scene.NewFromProto(in.Scene)
	_, err := scene_event.Append(
		s.PostgresClient,
		newScene.Id, // TODO: create unique id
		scene_event.SceneEventChange{
			Action: scene.SceneCreate,
			Changes: map[string]interface{}{
				"name": newScene.Name,
			},
		},
		scene_event.SceneEventChange{
			Action: scene.SceneUnCreate,
		},
	)
	if err != nil {
		return nil, err
	}
	// TODO: set id after creation
	err = scene.CreateScene(s.PostgresClient, in.GetScene())
	if err != nil {
		return nil, err
	}
	return newScene.ToProto(), nil
}

// Delete adds a DeleteScene event to the SceneEvent table.
func (s *Server) Delete(ctx context.Context, in *scenes_pb.DeleteSceneRequest) (*empty.Empty, error) {
	_, err := scene_event.Append(
		s.PostgresClient,
		in.GetId(),
		scene_event.SceneEventChange{
			Action: scene.SceneDelete,
		},
		scene_event.SceneEventChange{
			Action: scene.SceneUnDelete,
		},
	)
	if err != nil {
		return nil, err
	}
	err = scene.DeleteScene(s.PostgresClient, in.GetId())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// Get selects the scene by the given id from the Scenes table.
func (s *Server) Get(ctx context.Context, in *scenes_pb.GetSceneRequest) (*scenes_pb.Scene, error) {
	s, err := scene.GetScene(s.PostgresClient, in.GetId())
	if err != nil {
		return nil, err
	}
	return s.ToProto(), nil
}

// List selects all scenes from the Scenes table.
func (s *Server) List(ctx context.Context, in *scenes_pb.ListScenesRequest) (*scenes_pb.ListScenesResponse, error) {
	objectList, err := scene.ListScene(s.PostgresClient)
	if err != nil {
		return nil, err
	}
	return &scenes_pb.ListScenesResponse{Scenes: objectList.ToProto()}, nil
}

// Rename creates a new rename event on the scene_event table and updates the latest scene object with the new name.
func (s *Server) Rename(ctx context.Context, in *scenes_pb.RenameSceneRequest) (*scenes_pb.Scene, error) {
	sceneId := in.GetId()
	previousScene, err := scene.GetScene(s.PostgresClient, sceneId)
	_, err = scene_event.Append(
		s.PostgresClient,
		sceneId,
		scene_event.SceneEventChange{
			Action: scene.SceneRename,
			Changes: map[string]interface{}{
				"name": in.GetName(),
			},
		},
		scene_event.SceneEventChange{
			Action: scene.SceneUnRename,
			Changes: map[string]interface{}{
				"name": previousScene.Name,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	// TODO: apply event logic in Scene object
	newScene := previousScene
	newScene.Name = in.GetName()
	err = scene.UpdateScene(s.PostgresClient, &newScene)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
