package operator

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/stupschwartz/qubit/core/geometry"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/parameter"
	pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

const TableName = "operators"

type Operator struct {
	Id      string `json:"id" db:"id"`
	SceneId string `json:"scene_id" db:"scene_id"`
	// 2d, 3d, etc.
	Context string `json:"context" db:"context"`
	Type    string `json:"type" db:"type"`
	Name    string `json:"name" db:"name"`
	// Array of IDs of input operators
	Inputs        []string            `json:"inputs" db:"inputs"`
	ParameterRoot parameter.Parameter `json:"parameter_root" db:"parameter_root"`
}

type Operators []Operator

func NewFromProto(pb_op *pb.Operator) Operator {
	var parameterRoot parameter.Parameter
	_ = json.Unmarshal(pb_op.GetParameterRoot(), &parameterRoot)
	return Operator{
		Id:            pb_op.Id,
		SceneId:       pb_op.SceneId,
		Type:          pb_op.Type,
		Name:          pb_op.Name,
		Context:       pb_op.Context,
		ParameterRoot: parameterRoot,
	}
}

func (o *Operator) ToProto() *pb.Operator {
	// TODO: handle marshaling error
	parameterRootBytes, _ := json.Marshal(o.ParameterRoot)
	return &pb.Operator{
		Id:            o.Id,
		SceneId:       o.SceneId,
		Type:          o.Type,
		Name:          o.Name,
		Context:       o.Context,
		ParameterRoot: parameterRootBytes,
	}
}

func (o *Operator) GetCreateData() map[string]interface{} {
	// TODO: handle marshaling error
	parameterRootBytes, _ := json.Marshal(o.ParameterRoot)
	return map[string]interface{}{
		"scene_id":       o.SceneId,
		"context":        o.Context,
		"type":           o.Type,
		"name":           o.Name,
		"parameter_root": parameterRootBytes,
	}
}

func (o *Operator) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"context": o.Context,
		"type":    o.Type,
		"name":    o.Name,
	}
}

func (o *Operator) ValidateCreate() error {
	return nil
}

func (o *Operator) ValidateUpdate(newObj interface{}) error {
	//op := newObj.(*Operator)
	return nil
}

func (o *Operators) ToProto() []*pb.Operator {
	var pb_ops []*pb.Operator
	for _, operator := range *o {
		pb_ops = append(pb_ops, operator.ToProto())
	}
	return pb_ops
}

type RenderImageContext struct {
	Inputs        []image.Plane
	ParameterRoot *parameter.Parameter
	BoundingBox   *geometry.BoundingBox2D
	Time          float64
}

type Operable interface {
	Process(renderContext *RenderImageContext) (*image.Plane, error)
}

var OperatorsRegistry = make(map[string]Operable)
var ParameterRootRegistry = make(map[string]parameter.ParameterRoot)

func RegisterOperation(opType string, operation Operable, parameterRoot parameter.ParameterRoot) {
	OperatorsRegistry[opType] = operation
	ParameterRootRegistry[opType] = parameterRoot
}

func GetOperation(opType string) (Operable, error) {
	if operable, ok := OperatorsRegistry[opType]; ok {
		return operable, nil
	}
	return nil, errors.Errorf("Operation does not exist, %v", opType)
}

func GetParameterRoot(opType string) (parameter.ParameterRoot, error) {
	var p parameter.ParameterRoot
	if p, ok := ParameterRootRegistry[opType]; ok {
		return p, nil
	}
	return p, errors.Errorf("Parameters do not exist for operation type, %v", opType)
}
