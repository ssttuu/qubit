package checkerboard

import (
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/core/operator"
)

const Name string = "CheckerBoard"

type CheckerBoard struct {

}

func (c *CheckerBoard) Process(inputs []image.Plane, p params.Parameters, startX int32, startY int32, endX int32, endY int32) image.Plane {
	sizeParam := p.GetByName("Size")
	sizeValue := int32(sizeParam.GetValue(0))

	width := endX - startX
	height := endY - startY

	redChannel := image.Channel{Rows: make([]*image.Row, height)}
	greenChannel := image.Channel{Rows: make([]*image.Row, height)}
	blueChannel := image.Channel{Rows: make([]*image.Row, height)}

	// TODO: each row should be a goroutine
	var row, col int32
	for row = 0; row < height; row++ {
		rowData := make([]float64, width)
		for col = 0; col < width; col++ {
			checkerBoardRow := (row + startY) / sizeValue
			checkerBoardColumn := (col + startX) / sizeValue

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

	return image.NewPlane(width, height, []image.Channel{redChannel, greenChannel, blueChannel})
}

var Params params.Parameters = params.Parameters{
	params.NewFloatParameter("Size"),
}

func init() {
	operator.RegisterOperation(Name, &CheckerBoard{}, Params)
}
