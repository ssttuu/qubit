package operator

import (
	pb "github.com/stupschwartz/qubit/server/protos/operators"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/core/image"
)


const (
	IMAGE string = "image"
	GEOMETRY string = "geometry"
)

type Operator struct {
	Id      string `json:"id" datastore:"id"`
	Name    string `json:"name" datastore:"name"`
	Context string `json:"kind" datastore:"kind"`
	Type    string `json:"type" datastore:"type"`
	Inputs  []string `json:"inputs" datastore:"inputs"`
	Outputs []string `json:"outputs" datastore:"outputs"`
}

func (o *Operator) ToProto() (*pb.Operator, error) {
	return &pb.Operator{Id: o.Id, Name: o.Name}, nil
}

func NewOperatorFromProto(pb_op *pb.Operator) Operator {
	return Operator{Id: fmt.Sprint(pb_op.Id), Name: pb_op.Name}
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
	Process(inputs []image.Plane, p params.Parameters, startX int32, startY int32, endX int32, endY int32) image.Plane
}

var OperatorsRegistry = make(map[string]Operable)
var ParametersRegistry = make(map[string]params.Parameters)

func RegisterOperation(opType string, operation Operable, parameters params.Parameters) {
	OperatorsRegistry[opType] = operation
	ParametersRegistry[opType] = parameters
}

func GetOperation(opType string) (Operable, error) {
	if operable, ok := OperatorsRegistry[opType]; ok {
		return operable, nil
	}
	return nil, errors.Errorf("Operation does not exist, %v", opType)
}

func GetParameters(opType string) (params.Parameters, error) {
	if parameters, ok := ParametersRegistry[opType]; ok {
		return parameters, nil
	}
	return nil, errors.Errorf("Parameters do not exist for operation type, %v", opType)
}