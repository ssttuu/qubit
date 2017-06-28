package renders

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
	renders_pb "github.com/stupschwartz/qubit/proto-gen/go/renders"
)

type Server struct {
	OperatorsClient operators_pb.OperatorsClient
}

func Register(grpcServer *grpc.Server, operatorsClient operators_pb.OperatorsClient) {
	renders_pb.RegisterRendersServer(grpcServer, &Server{
		OperatorsClient: operatorsClient,
	})
}

func (s *Server) DoRender(ctx context.Context, in *renders_pb.RenderRequest) (*renders_pb.RenderResponse, error) {
	pbRenderParameterRequest := operators_pb.RenderParameterRequest{}
	pbRenderParameter, err := s.OperatorsClient.GetRenderParameters(ctx, &pbRenderParameterRequest)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument")
	}
	log.Println("Render Parameters:", pbRenderParameter)
	// TODO: Use core/parameters to Unmarshal
	// TODO: Do the rendering
	return nil, nil
}
