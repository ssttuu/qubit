package constant

import (
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/core/operator"
)

const Name string = "Constant"

type Constant struct {}

func (c *Constant) Process(inputs []image.Plane, p params.Parameters, startX int32, startY int32, endX int32, endY int32) image.Plane {
	colorParam := p.GetByName("Color")

	redValue := colorParam.GetComponentValueByLabel("Red")
	greenValue := colorParam.GetComponentValueByLabel("Green")
	blueValue := colorParam.GetComponentValueByLabel("Blue")

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

	return image.NewPlane(width, height, []image.Channel{redComponent, greenComponent, blueComponent})

}

var Params params.Parameters = params.Parameters{
	params.NewColorParameter("Color"),
}

func init() {
	operator.RegisterOperation(Name, &Constant{}, Params)
}
