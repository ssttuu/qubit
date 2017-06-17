package image

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/images"
)

type Image struct {
	Id              string            `json:"id" db:"id"`
	ImageSequenceId string            `json:"image_sequence_id" db:"image_sequence_id"`
	Name            string            `json:"name" db:"name"`
	Width           int32             `json:"width" db:"width"`
	Height          int32             `json:"height" db:"height"`
	Labels          map[string]string `json:"labels" db:"labels"`
	Planes          []Plane           `json:"planes" db:"planes"`
}

func (i *Image) ToProto() *pb.Image {
	pb_planes := make([]*pb.Plane, len(i.Planes))
	for index, plane := range i.Planes {
		pb_planes[index] = plane.ToProto()
	}
	return &pb.Image{
		Id:              i.Id,
		ImageSequenceId: i.ImageSequenceId,
		Name:            i.Name,
		Width:           i.Width,
		Height:          i.Height,
		Labels:          i.Labels,
		Planes:          pb_planes,
	}
}

func NewFromProto(pbimage *pb.Image) Image {
	planes := make([]Plane, len(pbimage.Planes))
	for index, plane := range pbimage.Planes {
		planes[index] = *NewPlaneFromProto(plane)
	}
	return Image{
		Id:              pbimage.Id,
		ImageSequenceId: pbimage.ImageSequenceId,
		Name:            pbimage.Name,
		Width:           pbimage.Width,
		Height:          pbimage.Height,
		Labels:          pbimage.Labels,
		Planes:          planes,
	}
}

type Images []Image

func (i *Images) ToProto() []*pb.Image {
	var pbimages []*pb.Image
	for _, image := range *i {
		image_proto := image.ToProto()
		pbimages = append(pbimages, image_proto)
	}
	return pbimages
}
