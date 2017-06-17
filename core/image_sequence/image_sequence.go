package image

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/image_sequences"
)

type ImageSequence struct {
	Id        string `json:"id" db:"id"`
	ProjectId string `json:"project_id" db:"project_id"`
	Name      string `json:"name" db:"name"`
}

func (i *ImageSequence) ToProto() *pb.ImageSequence {
	return &pb.ImageSequence{
		Id:        i.Id,
		ProjectId: i.ProjectId,
		Name:      i.Name,
	}
}

func NewFromProto(pbimage *pb.ImageSequence) ImageSequence {
	return ImageSequence{
		Id:        pbimage.Id,
		ProjectId: pbimage.ProjectId,
		Name:      pbimage.Name,
	}
}

type ImageSequences []ImageSequence

func (i *ImageSequences) ToProto() []*pb.ImageSequence {
	var pbimages_sequences []*pb.ImageSequence
	for _, image_sequence := range *i {
		image_seq_proto := image_sequence.ToProto()
		pbimages_sequences = append(pbimages_sequences, image_seq_proto)
	}
	return pbimages_sequences
}
