package operator

import (
	"github.com/pkg/errors"

	"github.com/stupschwartz/qubit/core/geometry"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/parameter"
)

// TODO: Move to scene package?
type Operator struct {
	Context    string               `json:"context"` // 2d, 3d, etc.
	Id         string               `json:"id"`
	InputIds   string               `json:"input_ids"`
	Name       string               `json:"name"`
	Parameters *parameter.Parameter `json:"parameters"`
	Type       string               `json:"type"`
}

type OperatorData struct {
	RootOperatorIds []string            `json:"root_operator_ids"`
	OperatorMap     map[string]Operator `json:"operator_map"`
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
