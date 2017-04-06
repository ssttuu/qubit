package operator

import (
	"github.com/stupschwartz/qubit/image"
	"github.com/stupschwartz/qubit/params"
	"github.com/gonum/blas/blas64"
)


func MultiplyOperation(inputs []image.Plane, p params.Parameters, width int, height int) image.Plane {
	multiplyByParam := p.GetByName("MultiplyBy")
	multiplyByValue := multiplyByParam.GetValue(0)

	inputImage := inputs[0]


	for componentIndex := range inputImage.Components {
		inputComponent := inputImage.Components[componentIndex]
		newMatrixData := make([]float64, width * height)
		for pixelIndex, pixelValue := range inputComponent.RawMatrix().Data {
			newMatrixData[pixelIndex] = pixelValue * multiplyByValue
		}
		inputImage.Components[componentIndex].SetRawMatrix(blas64.General{Rows: height, Cols: width, Stride: width, Data:newMatrixData})
	}

	return inputImage
}

func init() {
	parameters := make([]params.Parameter, 1)
	parameters[0] = params.NewFloatParameter("MultiplyBy")

	RegisterOperation("Multiply", MultiplyOperation, parameters)
}
