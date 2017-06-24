package scenes

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/scene"
	scenes_pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

type Server struct {
	PostgresClient *sqlx.DB
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	scenes_pb.RegisterScenesServer(grpcServer, &Server{PostgresClient: postgresClient})
}

func (s *Server) Create(ctx context.Context, in *scenes_pb.CreateSceneRequest) (*scenes_pb.Scene, error) {
	newObject := scene.NewFromProto(in.Scene)
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     s.PostgresClient,
		Object: &newObject,
		Table:  scene.TableName,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *scenes_pb.DeleteSceneRequest) (*empty.Empty, error) {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Table: scene.TableName,
	})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) Get(ctx context.Context, in *scenes_pb.GetSceneRequest) (*scenes_pb.Scene, error) {
	var obj scene.Scene
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: scene.TableName,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *scenes_pb.ListScenesRequest) (*scenes_pb.ListScenesResponse, error) {
	var objectList scene.Scenes
	err := apiutils.List(&apiutils.ListConfig{
		DB:    s.PostgresClient,
		Out:   &objectList,
		Table: scene.TableName,
	})
	if err != nil {
		return nil, err
	}
	return &scenes_pb.ListScenesResponse{Scenes: objectList.ToProto()}, nil
}

func (s *Server) Update(ctx context.Context, in *scenes_pb.UpdateSceneRequest) (*scenes_pb.Scene, error) {
	newObject := scene.NewFromProto(in.Scene)
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        s.PostgresClient,
		Id:        in.GetId(),
		NewObject: &newObject,
		OldObject: &scene.Scene{},
		Table:     scene.TableName,
	})
	if err != nil {
		return nil, err
	}
	return newObject.ToProto(), nil
}
