package scene

import (
	pb "github.com/stupschwartz/qubit/server/protos/scenes"
	"strconv"
	"fmt"
	"github.com/pkg/errors"
)

type Scene struct {
	Id   string `json:"id" datastore:"id"`
	Name string `json:"name" datastore:"name"`
	Type string `json:"type" datastore:"type"`
}

func (s *Scene) ToProto() (*pb.Scene, error) {
	i, err := strconv.ParseInt(s.Id, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert Id from string to int64: %v", s.Id)
	}
	return &pb.Scene{Id: i}, nil
}

func NewSceneFromProto(pbscene *pb.Scene) Scene {
	return Scene{Id: fmt.Sprintf("%d", pbscene.Id)}
}

type Scenes []*Scene

func (s *Scenes) ToProto() (*pb.ScenesList, error) {
	var pbscenes []*pb.Scene
	for _, scene := range *s {
		scene_proto, err := scene.ToProto()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to convert scene to proto, %v", scene)
		}
		pbscenes = append(pbscenes, scene_proto)
	}

	return &pb.ScenesList{Scenes:pbscenes}, nil
}
