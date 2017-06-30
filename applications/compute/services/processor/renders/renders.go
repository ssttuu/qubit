package renders

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	render_parameters_pb "github.com/stupschwartz/qubit/proto-gen/go/render_parameters"
	renders_pb "github.com/stupschwartz/qubit/proto-gen/go/renders"
)

type Server struct {
	RenderParametersClient render_parameters_pb.RenderParametersClient
}

func Register(grpcServer *grpc.Server, renderParametersClient render_parameters_pb.RenderParametersClient) {
	renders_pb.RegisterRendersServer(grpcServer, &Server{
		RenderParametersClient: renderParametersClient,
	})
}

func (s *Server) DoRender(ctx context.Context, in *renders_pb.RenderRequest) (*renders_pb.RenderResponse, error) {
	pbRenderParameterRequest := render_parameters_pb.RenderParameterRequest{}
	pbRenderParameter, err := s.RenderParametersClient.GetRenderParameters(ctx, &pbRenderParameterRequest)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument")
	}
	log.Println("Render Parameters:", pbRenderParameter)
	// TODO: Use core/parameters to Unmarshal
	// TODO: Do the rendering
	return nil, nil
}
