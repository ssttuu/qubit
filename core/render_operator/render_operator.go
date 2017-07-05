package scene_event

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/render_operators"
)

const TableName = "render_operators"

type RenderOperator struct {
	Id           string
	SceneId      string
	SceneVersion int32
}

type RenderOperators []RenderOperator

// TODO: Return a pointer
func NewFromProto(pbRenderOperator *pb.RenderOperator) RenderOperator {
	return RenderOperator{
		Id:           pbRenderOperator.GetId(),
		SceneId:      pbRenderOperator.GetSceneId(),
		SceneVersion: pbRenderOperator.GetSceneVersion(),
	}
}

func (ro *RenderOperator) ToProto() *pb.RenderOperator {
	return &pb.RenderOperator{
		Id:           ro.Id,
		SceneId:      ro.SceneId,
		SceneVersion: ro.SceneVersion,
	}
}

func (ro *RenderOperators) ToProto() []*pb.RenderOperator {
	var pbscenes []*pb.RenderOperator
	for _, sceneEvent := range *ro {
		sceneEventProto := sceneEvent.ToProto()
		pbscenes = append(pbscenes, sceneEventProto)
	}
	return pbscenes
}
