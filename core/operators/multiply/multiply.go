package multiply

//import (
//	"github.com/stupschwartz/qubit/core/image"
//	"github.com/stupschwartz/qubit/core/params"
//)
//
//
//func MultiplyOperation(inputs []image.Plane, p params.Parameters, startX int64, startY int64, endX int64, endY int64) image.Plane {
//	multiplyByParam := p.GetByName("MultiplyBy")
//	multiplyByValue := multiplyByParam.GetValue(0)
//
//	inputImage := inputs[0]
//
//	for componentIndex, inputComponent := range inputImage.Components {
//		for rowIndex, row := range inputComponent.Rows {
//			for pixelIndex, pixelValue := range row.Data {
//				inputImage.Components[componentIndex].Rows[rowIndex].Data[pixelIndex] = pixelValue * multiplyByValue
//			}
//		}
//	}
//
//	return inputImage
//}
//
//func init() {
//	parameters := make([]params.Parameter, 1)
//	parameters[0] = params.NewFloatParameter("MultiplyBy")
//
//	RegisterOperation("Multiply", MultiplyOperation, parameters)
//}
