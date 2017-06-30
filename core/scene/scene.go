package scene

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/operator"
	pb "github.com/stupschwartz/qubit/proto-gen/go/scenes"
)

const TableName = "scenes"

type Scene struct {
	Id           string `db:"id"`
	Name         string `db:"name"`
	OperatorData []byte `db:"operator_data"`
	ProjectId    string `db:"project_id"`
	Version      int32  `db:"version"`
}

type Scenes []Scene

// TODO: Return a pointer
func NewFromProto(pbScene *pb.Scene) Scene {
	return Scene{
		Id:           pbScene.GetId(),
		Name:         pbScene.GetName(),
		OperatorData: pbScene.GetOperatorData(),
		ProjectId:    pbScene.GetProjectId(),
		Version:      pbScene.GetVersion(),
	}
}

func (s *Scene) GetCreateData() map[string]interface{} {
	defaultOpData := operator.OperatorData{
		RootOperatorIds: []string{},
		OperatorMap:     map[string]operator.Operator{},
	}
	defaultOpDataJSON, err := json.Marshal(defaultOpData)
	if err != nil {
		panic(errors.Wrapf(err, "Failed to marshal defaultOpData: %v", defaultOpData))
	}
	return map[string]interface{}{
		"name":          s.Name,
		"operator_data": defaultOpDataJSON,
		"project_id":    s.ProjectId,
		"version":       1,
	}
}

func (s *Scene) GetOperators() (operator.OperatorData, error) {
	var opData operator.OperatorData
	err := json.Unmarshal(s.OperatorData, &opData)
	return opData, err
}

func (s *Scene) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"name": s.Name,
	}
}

func (s *Scene) SetOperators(opData operator.OperatorData) error {
	opsJSONData, err := json.Marshal(&opData)
	if err != nil {
		return err
	}
	s.OperatorData = opsJSONData
	return nil
}

func (s *Scene) ToProto() *pb.Scene {
	return &pb.Scene{
		Id:           s.Id,
		Name:         s.Name,
		OperatorData: s.OperatorData,
		ProjectId:    s.ProjectId,
		Version:      s.Version,
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
