package scene

import (
	pb "github.com/stupschwartz/qubit/server/protos/scenes"
	"fmt"
	"github.com/pkg/errors"
)

type Scene struct {
	Id   string `json:"id" datastore:"id"`
	Name string `json:"name" datastore:"name"`
	Type string `json:"type" datastore:"type"`
}

func (s *Scene) ToProto() (*pb.Scene, error) {
	return &pb.Scene{Id: s.Id}, nil
}

func NewSceneFromProto(pbscene *pb.Scene) Scene {
	return Scene{Id: fmt.Sprint(pbscene.Id)}
}

type Scenes []*Scene

func (s *Scenes) ToProto() ([]*pb.Scene, error) {
	var pbscenes []*pb.Scene
	for _, scene := range *s {
		scene_proto, err := scene.ToProto()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to convert scene to proto, %v", scene)
		}
		pbscenes = append(pbscenes, scene_proto)
	}

	return pbscenes, nil
}
