package scene

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

const TableName = "scenes"

type Scene struct {
	Id        string `json:"id" db:"id"`
	ProjectId string `json:"project_id" db:"project_id"`
	Name      string `json:"name" db:"name"`
}

type Scenes []Scene

func NewFromProto(pbscene *pb.Scene) Scene {
	return Scene{
		Id:        pbscene.GetId(),
		ProjectId: pbscene.GetProjectId(),
		Name:      pbscene.GetName(),
	}
}

func (s *Scene) ToProto() *pb.Scene {
	return &pb.Scene{
		Id:        s.Id,
		ProjectId: s.ProjectId,
		Name:      s.Name,
	}
}

func (s *Scene) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"project_id": s.ProjectId,
		"name":       s.Name,
	}
}

func (s *Scene) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"name": s.Name,
	}
}

func (s *Scene) ValidateCreate() error {
	return nil
}

func (s *Scene) ValidateUpdate(newObj interface{}) error {
	//scene := newObj.(*Scene)
	return nil
}

func (s *Scenes) ToProto() []*pb.Scene {
	var pbscenes []*pb.Scene
	for _, scene := range *s {
		scene_proto := scene.ToProto()
		pbscenes = append(pbscenes, scene_proto)
	}
	return pbscenes
}
