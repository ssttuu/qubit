package operator

import (
	"github.com/pkg/errors"

	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/parameter"
	pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

const (
	IMAGE    string = "image"
	GEOMETRY string = "geometry"
)

type Operator struct {
	Id      string `json:"id" db:"id"`
	SceneId string `json:"scene_id" db:"scene_id"`
	Type    string `json:"type" db:"type"`
	Name    string `json:"name" db:"name"`
	// 2d, 3d, etc.
	Context string `json:"context" db:"context"`
	// Array of IDs of input operators
	Inputs     []string             `json:"inputs" db:"inputs"`
	Parameters parameter.Parameters `json:"parameters" db:"parameters"`
}

func (o *Operator) ToProto() *pb.Operator {
	return &pb.Operator{
		Id:      o.Id,
		SceneId: o.SceneId,
		Type:    o.Type,
		Name:    o.Name,
		Context: o.Context,
	}
}

func NewOperatorFromProto(pb_op *pb.Operator) Operator {
	return Operator{
		Id:      pb_op.Id,
		SceneId: pb_op.SceneId,
		Type:    pb_op.Type,
		Name:    pb_op.Name,
		Context: pb_op.Context,
	}
}

type Operators []*Operator

func (o *Operators) ToProto() []*pb.Operator {
	var pb_ops []*pb.Operator
	for _, operator := range *o {
		pb_ops = append(pb_ops, operator.ToProto())
	}
	return pb_ops
}

type Operable interface {
	Process(inputs []*image.Plane, p parameter.Parameters, startX int32, startY int32, endX int32, endY int32) (*image.Plane, error)
}

var OperatorsRegistry = make(map[string]Operable)
var ParametersRegistry = make(map[string]parameter.Parameters)

func RegisterOperation(opType string, operation Operable, parameters parameter.Parameters) {
	OperatorsRegistry[opType] = operation
	ParametersRegistry[opType] = parameters
}

func GetOperation(opType string) (Operable, error) {
	if operable, ok := OperatorsRegistry[opType]; ok {
		return operable, nil
	}
	return nil, errors.Errorf("Operation does not exist, %v", opType)
}

func GetParameters(opType string) (parameter.Parameters, error) {
	if parameters, ok := ParametersRegistry[opType]; ok {
		return parameters, nil
	}
	return nil, errors.Errorf("Parameters do not exist for operation type, %v", opType)
}
