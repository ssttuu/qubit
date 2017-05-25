package image

import (
	pb "github.com/stupschwartz/qubit/server/protos/images"
)

type Frame struct {
	Width  int32
	Height int32
	Labels map[string]string
	Planes []Plane
}

func (p *Frame) ToProto() *pb.Frame {
	pb_planes := make([]*pb.Plane, len(p.Planes))

	for index, plane := range p.Planes {
		pb_planes[index] = plane.ToProto()
	}

	return &pb.Frame{Labels: p.Labels, Planes: pb_planes}
}
