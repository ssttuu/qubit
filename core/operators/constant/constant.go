package constant

//import (
//	"github.com/stupschwartz/qubit/core/image"
//	"github.com/stupschwartz/qubit/core/params"
//)
//
//func ConstantOperation(inputs []image.Plane, p params.Parameters, startX int64, startY int64, endX int64, endY int64) image.Plane {
//	colorParam := p.GetByName("Color")
//
//	redValue := colorParam.GetComponentValueByLabel("Red")
//	greenValue := colorParam.GetComponentValueByLabel("Green")
//	blueValue := colorParam.GetComponentValueByLabel("Blue")
//
//	width := endX - startX
//	height := endY - startY
//
//	redComponent := image.Component{Rows: make([]*image.Row, height)}
//	greenComponent := image.Component{Rows: make([]*image.Row, height)}
//	blueComponent := image.Component{Rows: make([]*image.Row, height)}
//
//	// TODO: each row should be a goroutine
//	var row, col int64
//	for row = 0; row < height; row++ {
//		for col = 0; col < width; col++ {
//			redComponent.Rows[row].Data[col] = redValue
//			greenComponent.Rows[row].Data[col] = greenValue
//			blueComponent.Rows[row].Data[col] = blueValue
//		}
//	}
//
//	return image.NewRGBPlane(width, height, []image.Component{redComponent, greenComponent, blueComponent})
//
//}
//
//func init() {
//	parameters := make([]params.Parameter, 1)
//	parameters[0] = params.NewColorParameter("Color")
//
//	RegisterOperation("Constant", ConstantOperation, parameters)
//}
