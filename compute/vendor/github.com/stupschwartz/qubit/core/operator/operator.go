package operator

import (
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/core/image"
)

type Operator struct {
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params"`
	Inputs []string               `json:"inputs"`
}

type Operation func(inputs []image.Plane, p params.Parameters, width int64, height int64) image.Plane

var Operators = make(map[string]Operation)
var Parameters = make(map[string]params.Parameters)

func RegisterOperation(opType string, operation Operation, parameters params.Parameters) {
	Operators[opType] = operation
	Parameters[opType] = parameters
}

func GetOperation(opType string) Operation {
	return Operators[opType]
}

func GetParameters(opType string) params.Parameters {
	return Parameters[opType]
}
