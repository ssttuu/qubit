package scenes


import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"
	"cloud.google.com/go/datastore"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/scene"
	scenes_pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
	"github.com/stupschwartz/qubit/core/organization"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


type Server struct {
	DatastoreClient *datastore.Client
}

func (s *Server) List(ctx context.Context, in *scenes_pb.ListScenesRequest) (*scenes_pb.ListScenesResponse, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)

	var scenes scene.Scenes
	_, err := s.DatastoreClient.GetAll(ctx, datastore.NewQuery(scene.Kind).Ancestor(orgKey), &scenes)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	scenes_proto, err := scenes.ToProto()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert scenes to proto, %v", scenes)
	}

	return &scenes_pb.ListScenesResponse{Scenes:scenes_proto, NextPageToken:""}, nil
}

func (s *Server) Get(ctx context.Context, in *scenes_pb.GetSceneRequest) (*scenes_pb.Scene, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)

	var existingScene scene.Scene
	if err := s.DatastoreClient.Get(ctx, sceneKey, &existingScene); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingScene.ToProto()
}

func (s *Server) Create(ctx context.Context, in *scenes_pb.CreateSceneRequest) (*scenes_pb.Scene, error) {
	in.Scene.Id = fmt.Sprint(r.Int63())
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.Scene.Id, orgKey)

	newScene := scene.NewSceneFromProto(in.Scene)

	if _, err := s.DatastoreClient.Put(ctx, sceneKey, &newScene); err != nil {
		return nil, errors.Wrapf(err, "Failed to put scene, %v", newScene)
	}

	return newScene.ToProto()
}

func (s *Server) Update(ctx context.Context, in *scenes_pb.UpdateSceneRequest) (*scenes_pb.Scene, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)

	newScene := scene.NewSceneFromProto(in.Scene)

	_, err := s.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingScene scene.Scene
		if err := tx.Get(sceneKey, &existingScene); err != nil {
			return errors.Wrapf(err, "Failed to get scene in tx %v", existingScene)
		}

		existingScene.Name = newScene.Name

		_, err := tx.Put(sceneKey, &existingScene)
		if err != nil {
			return errors.Wrapf(err, "Failed to put scene in tx %v", existingScene)
		}

		newScene = existingScene

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to update scene")
	}

	return newScene.ToProto()
}

func (s *Server) Delete(ctx context.Context, in *scenes_pb.DeleteSceneRequest) (*empty.Empty, error) {
	orgKey := datastore.NameKey(organization.Kind, in.OrganizationId, nil)
	sceneKey := datastore.NameKey(scene.Kind, in.SceneId, orgKey)

	if err := s.DatastoreClient.Delete(ctx, sceneKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to deleted scene by key: %v", sceneKey)
	}

	return &empty.Empty{}, nil
}

func Register(grpcServer *grpc.Server, datastoreClient *datastore.Client) {
	scenes_pb.RegisterScenesServer(grpcServer, &Server{DatastoreClient: datastoreClient})
}