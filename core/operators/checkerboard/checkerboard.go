package checkerboard

import (
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/operator"
	"github.com/stupschwartz/qubit/core/parameter"
)

const Name string = "CheckerBoard"

var ParameterRoot parameter.Parameter = parameter.NewGroupParameter(
	"root",
	parameter.ParameterMap{
		"size": parameter.NewFloat64Parameter(256.0),
	},
	[]string{"size"},
)

type CheckerBoard struct{}

func (c *CheckerBoard) Process(imageContext *operator.RenderImageContext) (*image.Plane, error) {
	rootGroup := imageContext.ParameterRoot.GetGroup()
	sizeFl64 := rootGroup.GetFloat64("size")
	sizeValue := int32(sizeFl64.GetValue(0.0))
	width := imageContext.BoundingBox.EndX - imageContext.BoundingBox.StartX
	height := imageContext.BoundingBox.EndY - imageContext.BoundingBox.StartY
	redChannel := image.Channel{Rows: make([]*image.Row, height)}
	greenChannel := image.Channel{Rows: make([]*image.Row, height)}
	blueChannel := image.Channel{Rows: make([]*image.Row, height)}
	// TODO: each row should be a goroutine
	var row, col int32
	for row = 0; row < height; row++ {
		rowData := make([]float64, width)
		for col = 0; col < width; col++ {
			checkerBoardRow := (row + imageContext.BoundingBox.StartY) / sizeValue
			checkerBoardColumn := (col + imageContext.BoundingBox.StartX) / sizeValue
			value := 0.0
			if ((checkerBoardRow + checkerBoardColumn) % 2) == 0 {
				value = 1.0
			}
			rowData[col] = value
		}
		redChannel.Rows[row] = &image.Row{Data: rowData}
		greenChannel.Rows[row] = &image.Row{Data: rowData}
		blueChannel.Rows[row] = &image.Row{Data: rowData}
	}
	return image.NewPlane(width, height, []image.Channel{redChannel, greenChannel, blueChannel}), nil
}

func init() {
	operator.RegisterOperation(Name, &CheckerBoard{}, ParameterRoot)
}
