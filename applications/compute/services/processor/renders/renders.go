package renders

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	renders_pb "github.com/stupschwartz/qubit/proto-gen/go/computation_renders"
	render_operators_pb "github.com/stupschwartz/qubit/proto-gen/go/render_operators"
)

type Server struct {
	RenderOperatorsClient render_operators_pb.RenderOperatorsClient
}

func Register(grpcServer *grpc.Server, renderOperatorsClient render_operators_pb.RenderOperatorsClient) {
	renders_pb.RegisterComputationRendersServer(grpcServer, &Server{
		RenderOperatorsClient: renderOperatorsClient,
	})
}

func (s *Server) Render(ctx context.Context, in *renders_pb.ComputationRenderRequest) (*renders_pb.ComputationRenderResponse, error) {
	pbRenderOperatorRequest := render_operators_pb.RenderOperatorRequest{}
	pbRenderOperator, err := s.RenderOperatorsClient.GetRenderOperators(ctx, &pbRenderOperatorRequest)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument")
	}
	log.Println("Render Operators:", pbRenderOperator)
	// TODO: Use core/parameters to Unmarshal
	// TODO: Do the rendering
	return nil, nil
}
