package operator

import (
	pb "github.com/stupschwartz/qubit/server/protos/operators"
	"fmt"
	"strconv"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/core/image"
)


type Type string
const (
	IMAGE Type = "image"
	GEOMETRY Type = "geometry"
)

type Operator struct {
	Id      string `json:"id" datastore:"id"`
	Name    string `json:"name" datastore:"name"`
	Type    string `json:"type" datastore:"type"`
	Inputs  []string `json:"inputs" datastore:"inputs"`
	Outputs []string `json:"outputs" datastore:"outputs"`
}

func (n *Operator) ToProto() (*pb.Operator, error) {
	i, err := strconv.ParseInt(n.Id, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert Id from string to int64, %v", n.Id)
	}
	return &pb.Operator{Id: i}, nil
}

func NewOperatorFromProto(pboperator *pb.Operator) Operator {
	return Operator{Id: fmt.Sprint(pboperator.Id)}
}

type Operators []*Operator

func (n *Operators) ToProto() ([]*pb.Operator, error) {
	var pboperators []*pb.Operator
	for _, operator := range *n {
		operator_proto, err := operator.ToProto()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to convert operator to proto, %v", operator)
		}
		pboperators = append(pboperators, operator_proto)
	}

	return pboperators, nil
}

// TODO: cleanup
type Operation func(inputs []image.Plane, p params.Parameters, startX int64, startY int64, endX int64, endY int64) image.Plane

var OperatorsRegistry = make(map[string]Operation)
var ParametersRegistry = make(map[string]params.Parameters)

func RegisterOperation(opType string, operation Operation, parameters params.Parameters) {
	OperatorsRegistry[opType] = operation
	ParametersRegistry[opType] = parameters
}

func GetOperation(opType string) Operation {
	return OperatorsRegistry[opType]
}

func GetParameters(opType string) params.Parameters {
	return ParametersRegistry[opType]
}