package image

import (
	"github.com/gonum/matrix/mat64"
	"fmt"
	"image"
	"image/color"
	"log"
	pb "github.com/stupschwartz/qubit/protos"
)

type Component struct {
	*mat64.Dense
	Label string
}

func (c Component) String() string {
	return fmt.Sprintf("%v: %v", c.Label, mat64.Formatted(c, mat64.Prefix(" "), mat64.Squeeze()))
}

func (c Component) ToProto() *pb.ImageComponent {
	matrix := c.RawMatrix()
	return &pb.ImageComponent{Label: c.Label, Width: int32(matrix.Cols), Height: int32(matrix.Rows), Data: matrix.Data}
}

func NewComponentFromProto(cp *pb.ImageComponent) Component {
	return Component{Label: cp.GetLabel(), Dense: mat64.NewDense(int(cp.GetHeight()), int(cp.GetWidth()), cp.GetData())}
}

func NewComponentsFromProtos(cps []*pb.ImageComponent) []Component {
	components := make([]Component, len(cps))

	for index, cp := range cps {
		components[index] = NewComponentFromProto(cp)
	}

	return components
}

type Plane struct {
	Components []Component
	Label      string
}

func NewPlaneFromProto(imagePlaneProto *pb.ImagePlane) Plane {
	p := Plane{Label: imagePlaneProto.GetLabel(), Components: NewComponentsFromProtos(imagePlaneProto.GetComponents())}
	return p
}

func (p Plane) ToNRGBA() *image.NRGBA {
	height, width := p.Components[0].Caps()
	log.Printf("Width: %v, Height: %v", width, height)
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8(p.Components[0].At(y, x) * 255),
				G: uint8(p.Components[1].At(y, x) * 255),
				B: uint8(p.Components[2].At(y, x) * 255),
				A: 255,
			})
		}
	}

	return img
}

func (p Plane) ToProto() *pb.ImagePlane {
	cps := make([]*pb.ImageComponent, len(p.Components))

	for index, component := range p.Components {
		cps[index] = component.ToProto()
	}

	return &pb.ImagePlane{Label: p.Label, Components: cps}
}

type Frame struct {
	Planes []Plane
}
