package renders

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
	parameters_pb "github.com/stupschwartz/qubit/proto-gen/go/parameters"
	renders_pb "github.com/stupschwartz/qubit/proto-gen/go/renders"
)

type Server struct {
	OperatorsClient  operators_pb.OperatorsClient
	ParametersClient parameters_pb.ParametersClient
}

func Register(grpcServer *grpc.Server, operatorsClient operators_pb.OperatorsClient, parametersClient parameters_pb.ParametersClient) {
	renders_pb.RegisterRendersServer(grpcServer, &Server{
		OperatorsClient:  operatorsClient,
		ParametersClient: parametersClient,
	})
}

func (s *Server) DoRender(ctx context.Context, in *renders_pb.RenderRequest) (*renders_pb.RenderResponse, error) {
	// TODO
	return nil, nil
}
