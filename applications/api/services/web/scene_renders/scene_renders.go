package scene_renders

import (
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

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
	// TODO: Denormalize scene/operator data
	// TODO: Pass through to computations
	return nil, nil
}

func (s *Server) Get(ctx context.Context, in *scene_renders_pb.SceneRenderStatusRequest) (*scene_renders_pb.SceneRenderStatus, error) {
	// TODO: Pass through to computations
	return nil, nil
}
