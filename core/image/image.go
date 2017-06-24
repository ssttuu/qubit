package image

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/images"
)

const TableName = "images"

type Image struct {
	Id              string            `json:"id" db:"id"`
	ImageSequenceId string            `json:"image_sequence_id" db:"image_sequence_id"`
	Name            string            `json:"name" db:"name"`
	Width           int32             `json:"width" db:"width"`
	Height          int32             `json:"height" db:"height"`
	Labels          map[string]string `json:"labels" db:"labels"`
	Planes          []Plane           `json:"planes" db:"planes"`
}

type Images []Image

func NewFromProto(pbimage *pb.Image) Image {
	planes := make([]Plane, len(pbimage.Planes))
	for index, plane := range pbimage.Planes {
		planes[index] = *NewPlaneFromProto(plane)
	}
	return Image{
		Id:              pbimage.GetId(),
		ImageSequenceId: pbimage.GetImageSequenceId(),
		Name:            pbimage.GetName(),
		Width:           pbimage.GetWidth(),
		Height:          pbimage.GetHeight(),
		Labels:          pbimage.GetLabels(),
		Planes:          planes,
	}
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

func (i *Image) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"image_sequence_id": i.ImageSequenceId,
		"name":              i.Name,
		"width":             i.Width,
		"height":            i.Height,
		// TODO
	}
}

func (i *Image) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"name":   i.Name,
		"width":  i.Width,
		"height": i.Height,
		// TODO
	}
}

func (i *Image) ValidateCreate() error {
	return nil
}

func (i *Image) ValidateUpdate(newObj interface{}) error {
	//im := newObj.(*Image)
	return nil
}

func (i *Images) ToProto() []*pb.Image {
	var pbimages []*pb.Image
	for _, image := range *i {
		image_proto := image.ToProto()
		pbimages = append(pbimages, image_proto)
	}
	return pbimages
}
