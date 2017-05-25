package checkerboard

import (
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/params"
	"github.com/stupschwartz/qubit/core/operator"
)

const Name string = "CheckerBoard"

type CheckerBoard struct {

}

func (c *CheckerBoard) Process(inputs []image.Plane, p params.Parameters, startX int64, startY int64, endX int64, endY int64) image.Plane {
	sizeParam := p.GetByName("Size")
	sizeValue := int64(sizeParam.GetValue(0))

	width := endX - startX
	height := endY - startY

	redComponent := image.Component{Rows: make([]*image.Row, height)}
	greenComponent := image.Component{Rows: make([]*image.Row, height)}
	blueComponent := image.Component{Rows: make([]*image.Row, height)}

	// TODO: each row should be a goroutine
	var row, col int64
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

		redComponent.Rows[row] = &image.Row{Data: rowData}
		greenComponent.Rows[row] = &image.Row{Data: rowData}
		blueComponent.Rows[row] = &image.Row{Data: rowData}
	}

	return image.NewRGBPlane(width, height, []image.Component{redComponent, greenComponent, blueComponent})
}

var Params params.Parameters = params.Parameters{
	params.NewFloatParameter("Size"),
}

func init() {
	operator.RegisterOperation(Name, CheckerBoard{}, Params)
}
