package operator

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/parameter"
	"github.com/stupschwartz/qubit/core/image"
)


const Kind string = "Operator"

const (
	IMAGE string = "image"
	GEOMETRY string = "geometry"
)

type Operator struct {
	Id      string `json:"id" datastore:"id"`
	Name    string `json:"name" datastore:"name"`
	Context string `json:"context" datastore:"context"`
	Type    string `json:"type" datastore:"type"`
	Inputs  []string `json:"inputs" datastore:"inputs"`
	Outputs []string `json:"outputs" datastore:"outputs"`
}

func (o *Operator) ToProto() (*pb.Operator, error) {
	return &pb.Operator{Id: o.Id, Name: o.Name, Context: o.Context, Type: o.Type}, nil
}

func NewOperatorFromProto(pb_op *pb.Operator) Operator {
	return Operator{Id: fmt.Sprint(pb_op.Id), Name: pb_op.Name, Context: pb_op.Context, Type: pb_op.Type}
}

type Operators []*Operator

func (o *Operators) ToProto() ([]*pb.Operator, error) {
	var pb_ops []*pb.Operator
	for _, operator := range *o {
		operator_proto, err := operator.ToProto()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to convert operator to proto, %v", operator)
		}
		pb_ops = append(pb_ops, operator_proto)
	}

	return pb_ops, nil
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