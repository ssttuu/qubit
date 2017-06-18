package image_sequence

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/image_sequences"
)

type ImageSequence struct {
	Id        string `json:"id" db:"id"`
	ProjectId string `json:"project_id" db:"project_id"`
	Name      string `json:"name" db:"name"`
}

type ImageSequences []ImageSequence

func NewFromProto(pbimage *pb.ImageSequence) ImageSequence {
	return ImageSequence{
		Id:        pbimage.Id,
		ProjectId: pbimage.ProjectId,
		Name:      pbimage.Name,
	}
}

func (i *ImageSequence) ToProto() *pb.ImageSequence {
	return &pb.ImageSequence{
		Id:        i.Id,
		ProjectId: i.ProjectId,
		Name:      i.Name,
	}
}

func (i *ImageSequence) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"project_id": i.ProjectId,
		"name":       i.Name,
	}
}

func (i *ImageSequence) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"name": i.Name,
	}
}

func (i *ImageSequence) ValidateCreate() error {
	return nil
}

func (i *ImageSequence) ValidateUpdate(newObj interface{}) error {
	//im := newObj.(*ImageSequence)
	return nil
}

func (i *ImageSequences) ToProto() []*pb.ImageSequence {
	var pbimages_sequences []*pb.ImageSequence
	for _, image_sequence := range *i {
		image_seq_proto := image_sequence.ToProto()
		pbimages_sequences = append(pbimages_sequences, image_seq_proto)
	}
	return pbimages_sequences
}
