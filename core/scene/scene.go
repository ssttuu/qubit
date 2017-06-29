package scene

import (
	"encoding/json"

	"github.com/stupschwartz/qubit/core/operator"
	pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

const TableName = "scenes"

type DBScene struct {
	Id        string `db:"id"`
	ProjectId string `db:"project_id"`
	Version   string `db:"version"`
	Name      string `db:"name"`
	Operators []byte `db:"operators"`
}

type DBScenes []DBScene

type Scene struct {
	Id        string
	ProjectId string
	Version   string
	Name      string
	Operators operator.Operators
}

type Scenes []Scene

// TODO: Return a pointer
func NewFromProto(pbscene *pb.Scene) Scene {
	var ops operator.Operators
	// TODO: Handler error
	json.Unmarshal(pbscene.GetOperators(), &ops)
	return Scene{
		Id:        pbscene.GetId(),
		Name:      pbscene.GetName(),
		Operators: ops,
		ProjectId: pbscene.GetProjectId(),
		Version:   pbscene.GetVersion(),
	}
}

func (ds *DBScene) ToScene() *Scene {
	var ops operator.Operators
	// TODO: Handler error
	json.Unmarshal(ds.Operators, &ops)
	return &Scene{
		Id:        ds.Id,
		Name:      ds.Name,
		Operators: ops,
		ProjectId: ds.ProjectId,
		Version:   ds.Version,
	}
}

func (dss *DBScenes) ToScenes() Scenes {
	var scenes Scenes
	for _, ds := range *dss {
		scenes = append(scenes, *ds.ToScene())
	}
	return scenes
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

func (s *Scene) ToDBScene() *DBScene {
	// TODO: Handler error
	jsonData, _ := json.Marshal(s.Operators)
	return &DBScene{
		Id:        s.Id,
		Name:      s.Name,
		Operators: jsonData,
		ProjectId: s.ProjectId,
		Version:   s.Version,
	}
}

func (s *Scene) ToProto() *pb.Scene {
	// TODO: Handler error
	jsonData, _ := json.Marshal(s.Operators)
	return &pb.Scene{
		Id:        s.Id,
		Name:      s.Name,
		Operators: jsonData,
		ProjectId: s.ProjectId,
		Version:   s.Version,
	}
}

func (s *Scene) ValidateCreate() error {
	return nil
}

func (s *Scene) ValidateUpdate(newObj interface{}) error {
	//scene := newObj.(*Scene)
	return nil
}

func (s *Scenes) ToDBScenes() DBScenes {
	var dbScenes []DBScene
	for _, scene := range *s {
		dbScenes = append(dbScenes, *scene.ToDBScene())
	}
	return dbScenes
}

func (s *Scenes) ToProto() []*pb.Scene {
	var pbscenes []*pb.Scene
	for _, scene := range *s {
		scene_proto := scene.ToProto()
		pbscenes = append(pbscenes, scene_proto)
	}
	return pbscenes
}
