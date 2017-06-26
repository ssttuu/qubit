package renders

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	renders_pb "github.com/stupschwartz/qubit/proto-gen/go/renders"
)

type Server struct{}

func Register(grpcServer *grpc.Server) {
	renders_pb.RegisterRendersServer(grpcServer, &Server{})
}

func (s *Server) DoRender(ctx context.Context, in *renders_pb.RenderRequest) (*renders_pb.RenderResponse, error) {
	// TODO
	return nil, nil
}
