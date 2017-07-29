package scene

import (
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/core/geometry"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/parameter"
	pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

// TODO: Move to scene package?
type Operator struct {
	Context    string               `json:"context"` // 2d, 3d, etc.
	Id         string               `json:"id"`
	InputIds   []string             `json:"input_ids"`
	Name       string               `json:"name"`
	Parameters *parameter.Parameter `json:"parameters"`
	Type       string               `json:"type"`
}

func (o *Operator) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"id": o.Id,
	}
}

func (o *Operator) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"id": o.Id,
	}
}

func (o *Operator) ValidateCreate() error {
	return nil
}
func (o *Operator) ValidateUpdate(existingObject interface{}) error {
	return nil
}

func (o *Operator) ToProto() *pb.Operator {
	return &pb.Operator{
		Id:   o.Id,
		Name: o.Name,
	}
}

func NewOperatorFromProto(pbOperator *pb.Operator) Operator {
	return Operator{
		//Context: pbOperator.Context,
		Id: pbOperator.Id,
		//InputIds: pbOperator.InputIds,
		Name: pbOperator.Name,
		//Parameters: pbOperator.Parameters,
		//Type: pbOperator.Type,
	}
}

type Operators []Operator

func (o Operators) ToProto() []*pb.Operator {
	var pbOperators []*pb.Operator
	for _, op := range o {
		pbOperators = append(pbOperators, op.ToProto())
	}
	return pbOperators
}

type OperatorMap map[string]Operator

func (o OperatorMap) ToProto() map[string]*pb.Operator {
	var pbOperatorsMap map[string]*pb.Operator
	for opId, op := range o {
		pbOperatorsMap[opId] = op.ToProto()
	}
	return pbOperatorsMap
}

type RenderImageContext struct {
	BoundingBox *geometry.BoundingBox2D
	Inputs      []image.Plane
	Parameters  *parameter.Parameter
	Time        float64
}

type Operable interface {
	Process(renderContext *RenderImageContext) (*image.Plane, error)
}

var OperatorsRegistry = make(map[string]Operable)
var ParameterRootRegistry = make(map[string]parameter.ParameterSpecs)

func RegisterOperation(opType string, operation Operable, parameterRoot parameter.ParameterSpecs) {
	OperatorsRegistry[opType] = operation
	ParameterRootRegistry[opType] = parameterRoot
}

func GetOperation(opType string) (Operable, error) {
	if operable, ok := OperatorsRegistry[opType]; ok {
		return operable, nil
	}
	return nil, errors.Errorf("Operation does not exist, %v", opType)
}

const OperatorsTableName = "operators"

func CreateOperator(pgClient *sqlx.DB, sceneId string, operator *Operator) error {
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     pgClient,
		Object: operator,
		Table:  OperatorsTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to create Operator: %v", operator)
	}
	s, err := GetScene(pgClient, sceneId)
	if err != nil {
		return err
	}
	s.Operators[operator.Id] = operator
	err = UpdateScene(pgClient, s)
	if err != nil {
		return err
	}
	return err
}

func GetOperator(pgClient *sqlx.DB, operatorId string) (*Operator, error) {
	var o Operator
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    pgClient,
		Id:    operatorId,
		Out:   &o,
		Table: OperatorsTableName,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Operator by id: %v", operatorId)
	}
	return o, nil
}

func ListOperator(pgClient *sqlx.DB, sceneId string) (Operators, error) {
	var operatorList Operators
	err := apiutils.List(&apiutils.ListConfig{
		DB:    pgClient,
		Out:   &operatorList,
		Table: OperatorsTableName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get Operators")
	}
	return operatorList, nil
}

func UpdateOperator(pgClient *sqlx.DB, sceneId string, operator *Operator) error {
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        pgClient,
		Id:        operator.Id,
		NewObject: operator,
		OldObject: &Operator{},
		Table:     OperatorsTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to update Operator: %v", operator)
	}
	s, err := GetScene(pgClient, sceneId)
	if err != nil {
		return err
	}
	s.Operators[operator.Id] = operator
	err = UpdateScene(pgClient, s)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOperator(pgClient *sqlx.DB, sceneId string, operatorId string) error {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    pgClient,
		Id:    operatorId,
		Table: OperatorsTableName,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to delete Operator by id: %v", operatorId)
	}
	s, err := GetScene(pgClient, sceneId)
	if err != nil {
		return nil, err
	}
	delete(s.Operators, operatorId)
	err = UpdateScene(pgClient, s)
	if err != nil {
		return nil, err
	}
	return err
}
