package scene

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
	"fmt"
	"github.com/pkg/errors"
)

const Kind string = "Scene"

type Scene struct {
	Id   string `json:"id" datastore:"id"`
	Name string `json:"name" datastore:"name"`
	Type string `json:"type" datastore:"type"`
}

func (s *Scene) ToProto() (*pb.Scene, error) {
	return &pb.Scene{Id: s.Id, Name: s.Name}, nil
}

func NewSceneFromProto(pbscene *pb.Scene) Scene {
	return Scene{Id: fmt.Sprint(pbscene.Id), Name: pbscene.Name}
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
