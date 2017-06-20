package geometry

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/geometry"
)

type BoundingBox2D struct {
	StartX int32
	StartY int32
	EndX   int32
	EndY   int32
}

func NewBoundingBoxFromProto(pb_bbox *pb.BoundingBox2D) *BoundingBox2D {
	return &BoundingBox2D{
		StartX: pb_bbox.StartX,
		StartY: pb_bbox.StartY,
		EndX:   pb_bbox.EndX,
		EndY:   pb_bbox.EndY,
	}
}
