package constant

import (
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/parameter"
)

const Name string = "Constant"

var Params parameter.Parameter = parameter.NewGroupParameter(
	"root",
	parameter.ParameterMap{
		"color": parameter.NewRGBParameter(),
	},
	[]string{"color"},
)

type Constant struct{}

func (c *Constant) Process(inputs []*image.Plane, p parameter.Parameter, startX int32, startY int32, endX int32, endY int32) (*image.Plane, error) {
	colorParam := p.GetGroup().GetParameter("color").GetGroup()
	redValue := colorParam.GetParameter("red").GetFloat64().GetValue(0.0)
	greenValue := colorParam.GetParameter("green").GetFloat64().GetValue(0.0)
	blueValue := colorParam.GetParameter("blue").GetFloat64().GetValue(0.0)
	width := endX - startX
	height := endY - startY
	redComponent := image.Channel{Rows: make([]*image.Row, height)}
	greenComponent := image.Channel{Rows: make([]*image.Row, height)}
	blueComponent := image.Channel{Rows: make([]*image.Row, height)}
	// TODO: each row should be a goroutine
	var row, col int32
	for row = 0; row < height; row++ {
		for col = 0; col < width; col++ {
			redComponent.Rows[row].Data[col] = redValue
			greenComponent.Rows[row].Data[col] = greenValue
			blueComponent.Rows[row].Data[col] = blueValue
		}
	}
	return image.NewPlane(width, height, []image.Channel{redComponent, greenComponent, blueComponent}), nil

}

func init() {
	operator.RegisterOperation(Name, &Constant{}, Params)
}
