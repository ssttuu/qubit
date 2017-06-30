package roto

import (
	"fmt"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/parameter"
)

const Name string = "Roto"

var RotoParameters parameter.ParameterSpecs = parameter.ParameterSpecs{
	{Name: "lod", Parameter: parameter.NewFloat64Parameter(2.0)},
	{Name: "antialias", Parameter: parameter.NewGroupParameter(
		"antialias",
		parameter.ParameterSpecs{
			{Name: "x", Parameter: parameter.NewFloat64Parameter(0.0)},
			{Name: "y", Parameter: parameter.NewFloat64Parameter(0.0)},
		},
	)},
	{Name: "curves", Parameter: parameter.NewMultiParameter(
		parameter.NewGroupParameter(
			"curve",
			parameter.ParameterSpecs{
				{Name: "curve", Parameter: parameter.NewGroupParameter(
					"spline",
					parameter.ParameterSpecs{
						{Name: "splineType", Parameter: parameter.NewEnumParameter([]string{"polygon", "bezier"}, "bezier")},
						{Name: "shapeComposite", Parameter: parameter.NewEnumParameter([]string{"over", "under", "atop", "etc"}, "over")},
						{Name: "fill", Parameter: parameter.NewEnumParameter([]string{"closed", "open"}, "closed")},
						{Name: "thickness", Parameter: parameter.NewFloat64Parameter(0.02)},
						{Name: "feather", Parameter: parameter.NewBoolParameter(false)},
						{Name: "featherDropoff", Parameter: parameter.NewEnumParameter([]string{"linear", "gaussian"}, "gaussian")},
						//{Name: "featherRamp", Parameter: parameter.Ramp{}},
					},
				)},
				{Name: "points", Parameter: parameter.NewMultiParameter(
					parameter.NewGroupParameter(
						"point",
						parameter.ParameterSpecs{
							{Name: "position", Parameter: parameter.NewPosition2DParameter()},
							{Name: "tieSlopes", Parameter: parameter.NewGroupParameter(
								"tieSlopes",
								parameter.ParameterSpecs{
									{Name: "x", Parameter: parameter.NewFloat64Parameter(0.0)},
									{Name: "y", Parameter: parameter.NewFloat64Parameter(0.0)},
									{Name: "z", Parameter: parameter.NewFloat64Parameter(0.0)},
									{Name: "w", Parameter: parameter.NewFloat64Parameter(0.0)},
								},
							)},
							{Name: "thickness", Parameter: parameter.NewFloat64Parameter(1.0)},
						},
					),
					0,
					parameter.Parameters{},
				)},
			},
		),
		0,
		parameter.Parameters{},
	)},
}

type Roto struct{}

func (c *Roto) Process(renderContext *operator.RenderImageContext) (*image.Plane, error) {
	rotoRoot := renderContext.Parameters.GetGroup()
	curves := rotoRoot.GetMulti("curves")
	for curve := range curves.Iterator() {
		curveRoot := curve.GetGroup()
		points := curveRoot.GetMulti("points")
		for point := range points.Iterator() {
			pointRoot := point.GetGroup()
			position := pointRoot.GetGroup("position")
			x := position.GetFloat64("x")
			y := position.GetFloat64("y")
			xVal := x.GetValue(renderContext.Time)
			yVal := y.GetValue(renderContext.Time)
			fmt.Printf("x:%s y: %s", xVal, yVal)
		}
	}
	return nil, nil
}

func init() {
	operator.RegisterOperation(Name, &Roto{}, RotoParameters)
}
