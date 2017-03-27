package operator

import (
	"github.com/stupschwartz/qubit/image"
	"github.com/stupschwartz/qubit/params"
	"github.com/gonum/matrix/mat64"
)

func CheckerBoardOperation(inputs []image.Plane, p params.Parameters, width int, height int) image.Plane {
	sizeParam := params.GetByName(p, "Size")

	sizeValue := int(sizeParam.GetValue())

	var totalPixels int = width * height
	redPixels := make([]float64, totalPixels)
	greenPixels := make([]float64, totalPixels)
	bluePixels := make([]float64, totalPixels)

	for index := range redPixels {
		row := index / width
		column := index % width

		checkerBoardRow := row / sizeValue
		checkerBoardColumn := column / sizeValue

		value := 0.0
		if ((checkerBoardRow + checkerBoardColumn) % 2) == 0 {
			value = 1.0
		}

		redPixels[index] = value
		greenPixels[index] = value
		bluePixels[index] = value
	}

	redMatrix := mat64.NewDense(height, width, redPixels)
	greenMatrix := mat64.NewDense(height, width, greenPixels)
	blueMatrix := mat64.NewDense(height, width, bluePixels)

	redComponent := image.Component{Dense: redMatrix, Label: "Red"}
	greenComponent := image.Component{Dense: greenMatrix, Label: "Green"}
	blueComponent := image.Component{Dense: blueMatrix, Label: "Blue"}

	components := make([]image.Component, 3)
	components[0] = redComponent
	components[1] = greenComponent
	components[2] = blueComponent

	return image.Plane{
		Components: components,
		Label: "Color",
	}
}

func init() {
	parameters := make([]params.Parameter, 1)
	parameters[0] = params.NewFloatParameter("Size")

	RegisterOperation("CheckerBoard", CheckerBoardOperation, parameters)
}
