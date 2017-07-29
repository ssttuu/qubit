package scene_renders

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/scene"
	computations_pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
	scene_renders_pb "github.com/stupschwartz/qubit/proto-gen/go/scene_renders"
)

type Server struct {
	ComputationsClient computations_pb.ComputationsClient
	PostgresClient     *sqlx.DB
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, computationsClient computations_pb.ComputationsClient) {
	scene_renders_pb.RegisterSceneRendersServer(grpcServer, &Server{
		ComputationsClient: computationsClient,
		PostgresClient:     postgresClient,
	})
}

func (s *Server) Create(ctx context.Context, in *scene_renders_pb.SceneRenderRequest) (*scene_renders_pb.SceneRenderStatus, error) {
	// TODO: Pass through to computations

	// TODO: check for / create SceneSnapshot
	var obj scene.Scene
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.Render.GetSceneId(),
		Out:   &obj,
		Table: scene.ScenesTableName,
	})
	if err != nil {
		return nil, err
	}

	// TODO: return long running task data to client
	s.ComputationsClient.CreateComputation(ctx, &computations_pb.CreateComputationRequest{
		&computations_pb.Computation{
			Scene:       obj.ToProto(),
			OperatorId:  in.Render.GetOperatorId(),
			BoundingBox: in.Render.GetBoundingBox(),
		},
	})

	return nil, nil
}

func (s *Server) Get(ctx context.Context, in *scene_renders_pb.SceneRenderStatusRequest) (*scene_renders_pb.SceneRenderStatus, error) {
	// TODO: Pass through to computations
	return nil, nil
}
