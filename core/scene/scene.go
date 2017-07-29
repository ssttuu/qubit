package scene

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

const ScenesTableName = "scenes"

// TODO: obviously
type Scene struct {
	Id          string        `db:"id"`
	Name        string        `db:"name"`
	Operators   OperatorMap   `db:"operators"`
	Connections ConnectionMap `db:"connections"`
	ProjectId   string        `db:"project_id"`
	Version     int32         `db:"version"`
}

type Scenes []Scene

// TODO: Return a pointer
func NewFromProto(pbScene *pb.Scene) Scene {
	return Scene{
		Id:        pbScene.GetId(),
		Name:      pbScene.GetName(),
		Operators: pbScene.GetOperators(),
		ProjectId: pbScene.GetProjectId(),
		Version:   pbScene.GetVersion(),
	}
}

func (s *Scene) ApplyEvent(eventName string, eventData map[string]interface{}) {

}

func (s *Scene) GetCreateData() map[string]interface{} {
	// TODO: Store as Proto instead of JSON?
	defaultOpDataJSON, err := json.Marshal(s.Operators)
	if err != nil {
		// TODO: handle/return error
		panic(errors.Wrap(err, "Failed to marshal operators"))
	}
	return map[string]interface{}{
		"name":       s.Name,
		"operators":  defaultOpDataJSON,
		"project_id": s.ProjectId,
		"version":    1,
	}
}

func (s *Scene) GetOperators() (OperatorMap, error) {
	var opData OperatorMap
	err := json.Unmarshal(s.Operators, &opData)
	return opData, err
}

func (s *Scene) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"name": s.Name,
	}
}

func (s *Scene) SetOperators(opData OperatorMap) error {
	opsJSONData, err := json.Marshal(&opData)
	if err != nil {
		return err
	}
	s.Operators = opsJSONData
	return nil
}

func (s *Scene) ToProto() *pb.Scene {
	return &pb.Scene{
		Id:          s.Id,
		Name:        s.Name,
		Operators:   s.Operators,
		Connections: s.Connections,
		ProjectId:   s.ProjectId,
		Version:     s.Version,
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
		sceneProto := scene.ToProto()
		pbscenes = append(pbscenes, sceneProto)
	}
	return pbscenes
}

// DB

func CreateScene(pgClient *sqlx.DB, scene *Scene) error {
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     pgClient,
		Object: scene,
		Table:  ScenesTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to create Scene: %v", scene)
	}
	return err
}

func GetScene(pgClient *sqlx.DB, sceneId string) (*Scene, error) {
	var s Scene
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    pgClient,
		Id:    sceneId,
		Out:   &s,
		Table: ScenesTableName,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Scene by id: %v", sceneId)
	}
	return s, nil
}

func ListScene(pgClient *sqlx.DB) (Scenes, error) {
	var sceneList Scenes
	err := apiutils.List(&apiutils.ListConfig{
		DB:    pgClient,
		Out:   &sceneList,
		Table: ScenesTableName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get Scenes")
	}
	return sceneList, nil
}

func UpdateScene(pgClient *sqlx.DB, scene *Scene) error {
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        pgClient,
		Id:        scene.Id,
		NewObject: scene,
		OldObject: &Scene{},
		Table:     ScenesTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to update Scene: %v", scene)
	}
	return nil
}

func DeleteScene(pgClient *sqlx.DB, sceneId string) error {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    pgClient,
		Id:    sceneId,
		Table: ScenesTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to delete Scene by id: %v", sceneId)
	}
	return err
}
