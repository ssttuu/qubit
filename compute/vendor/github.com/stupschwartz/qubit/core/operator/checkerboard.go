package operator

import (
	"github.com/stupschwartz/qubit/core/image"
	"github.com/stupschwartz/qubit/core/params"
)

func CheckerBoardOperation(inputs []image.Plane, p params.Parameters, width int64, height int64) image.Plane {
	sizeParam := p.GetByName("Size")
	sizeValue := int64(sizeParam.GetValue(0))

	redComponent := image.Component{Rows: make([]*image.Row, height)}
	greenComponent := image.Component{Rows: make([]*image.Row, height)}
	blueComponent := image.Component{Rows: make([]*image.Row, height)}

	// TODO: each row should be a goroutine
	var row, col int64
	for row = 0; row < height; row++ {
		rowData := make([]float64, width)
		for col = 0; col < width; col++ {
			checkerBoardRow := row / sizeValue
			checkerBoardColumn := col / sizeValue

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

func init() {
	parameters := make([]params.Parameter, 1)
	parameters[0] = params.NewFloatParameter("Size")

	RegisterOperation("CheckerBoard", CheckerBoardOperation, parameters)
}
