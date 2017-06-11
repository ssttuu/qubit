package scene

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

type Scene struct {
	Id        string `json:"id" db:"id"`
	ProjectId string `json:"project_id" db:"project_id"`
	Name      string `json:"name" db:"name"`
}

func (s *Scene) ToProto() *pb.Scene {
	return &pb.Scene{
		Id:        s.Id,
		ProjectId: s.ProjectId,
		Name:      s.Name,
	}
}

func NewSceneFromProto(pbscene *pb.Scene) Scene {
	return Scene{
		Id:        pbscene.Id,
		ProjectId: pbscene.ProjectId,
		Name:      pbscene.Name,
	}
}

type Scenes []*Scene

func (s *Scenes) ToProto() []*pb.Scene {
	var pbscenes []*pb.Scene
	for _, scene := range *s {
		scene_proto := scene.ToProto()
		pbscenes = append(pbscenes, scene_proto)
	}
	return pbscenes
}
