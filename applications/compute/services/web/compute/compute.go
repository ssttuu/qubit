package compute

import (
	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/stupschwartz/qubit/core/geometry"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/parameter"
	compute_pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
)

type Server struct {
}

func (s *Server) RenderImage(ctx context.Context, in *compute_pb.RenderImageRequest) (*compute_pb.RenderImageResponse, error) {
	op, err := operator.GetOperation(in.Operator.Type)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Operation for rendering, %v", in.Operator.Type)
	}
	var p parameter.Parameter
	err = json.Unmarshal(in.Operator.ParameterRoot, &p)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to unmarshal ParameterRoot for rendering, %v", in.Operator.Type)
	}
	renderContext := operator.RenderImageContext{
		ParameterRoot: &p,
		BoundingBox:   geometry.NewBoundingBoxFromProto(in.BoundingBox),
		Time:          in.Time,
	}
	imagePlane, err := op.Process(&renderContext)
	return &compute_pb.RenderImageResponse{ImagePlane: imagePlane.ToProto()}, nil
}

func Register(grpcServer *grpc.Server) {
	compute_pb.RegisterComputeServer(grpcServer, &Server{})
}
