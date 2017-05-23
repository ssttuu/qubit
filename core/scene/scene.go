package scene

import pb "github.com/stupschwartz/qubit/server/protos/scenes"

type Scene struct {
	Id      int64 `json:"id" datastore:"id"`
	Name    string `json:"name" datastore:"name"`
	Type    string `json:"type" datastore:"type"`
}

func (s *Scene) ToProto() *pb.Scene {
	return &pb.Scene{Id: s.Id}
}

func NewSceneFromProto(pbscene *pb.Scene) Scene {
	return Scene{Id: pbscene.Id}
}

type Scenes []*Scene

func (s *Scenes) ToProto() *pb.ScenesList {
	var pbscenes []*pb.Scene
	for _, scene := range *s {
		pbscenes = append(pbscenes, scene.ToProto())
	}

	return &pb.ScenesList{Scenes:pbscenes}
}
