package constant

import (
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/parameter"
)

const Name string = "Constant"

var ConstantParameters parameter.ParameterSpecs = parameter.ParameterSpecs{
	{Name: "color", Parameter: parameter.NewRGBParameter()},
}

type Constant struct{}

func (c *Constant) Process(imageContext *operator.RenderImageContext) (*image.Plane, error) {
	colorRoot := imageContext.Parameters.GetGroup()
	colorGroup := colorRoot.GetGroup("color")
	redParam := colorGroup.GetFloat64("red")
	greenParam := colorGroup.GetFloat64("green")
	blueParam := colorGroup.GetFloat64("blue")
	redValue := redParam.GetValue(0.0)
	greenValue := greenParam.GetValue(0.0)
	blueValue := blueParam.GetValue(0.0)
	width := imageContext.BoundingBox.EndX - imageContext.BoundingBox.StartX
	height := imageContext.BoundingBox.EndY - imageContext.BoundingBox.StartY
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
	operator.RegisterOperation(Name, &Constant{}, ConstantParameters)
}
