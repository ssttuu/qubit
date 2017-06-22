package roto

import (
	"fmt"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/parameter"
)

const Name string = "Roto"

var ParameterRoot parameter.Parameter = parameter.NewGroupParameter(
	"root",
	parameter.ParameterMap{
		"lod": parameter.NewFloat64Parameter(2.0),
		"antialias": parameter.NewGroupParameter(
			"antialias",
			parameter.ParameterMap{
				"x": parameter.NewFloat64Parameter(0.0),
				"y": parameter.NewFloat64Parameter(0.0),
			},
			[]string{"x", "y"},
		),
		"curves": parameter.NewMultiParameter(
			parameter.NewGroupParameter(
				"curve",
				parameter.ParameterMap{
					"curve": parameter.NewGroupParameter(
						"spline",
						parameter.ParameterMap{
							"splineType":     parameter.NewEnumParameter([]string{"polygon", "bezier"}, "bezier"),
							"shapeComposite": parameter.NewEnumParameter([]string{"over", "under", "atop", "etc"}, "over"),
							"fill":           parameter.NewEnumParameter([]string{"closed", "open"}, "closed"),
							"thickness":      parameter.NewFloat64Parameter(0.02),
							"feather":        parameter.NewBoolParameter(false),
							"featherDropoff": parameter.NewEnumParameter([]string{"linear", "gaussian"}, "gaussian"),
							//"featherRamp": parameter.Ramp{},
						},
						[]string{"splineType", "shapeComposite", "fill", "thickness", "feather", "featherDropoff"},
					),
					"points": parameter.NewMultiParameter(
						parameter.NewGroupParameter(
							"point",
							parameter.ParameterMap{
								"position": parameter.NewPosition2DParameter(),
								"tieSlopes": parameter.NewGroupParameter(
									"tieSlopes",
									parameter.ParameterMap{
										"x": parameter.NewFloat64Parameter(0.0),
										"y": parameter.NewFloat64Parameter(0.0),
										"z": parameter.NewFloat64Parameter(0.0),
										"w": parameter.NewFloat64Parameter(0.0),
									},
									[]string{"x", "y", "z", "w"},
								),
								"thickness": parameter.NewFloat64Parameter(1.0),
							},
							[]string{"position", "tieSlopes", "thickness"},
						),
						0,
						parameter.ParameterArray{},
					),
				},
				[]string{"curve", "points"},
			),
			0,
			parameter.ParameterArray{},
		),
	},
	[]string{"root"},
)

type Roto struct{}

func (c *Roto) Process(renderContext *operator.RenderImageContext) (*image.Plane, error) {
	rotoRoot := renderContext.ParameterRoot.GetGroup()
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
	operator.RegisterOperation(Name, &Roto{}, ParameterRoot)
}
