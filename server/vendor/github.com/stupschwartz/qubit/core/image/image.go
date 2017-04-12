package image

import (
	"image"
	"image/color"
	"log"
	pb "github.com/stupschwartz/qubit/protos"
)


type Row struct {
	Data []float64
}

func (r Row) ToProto() *pb.Row {
	return &pb.Row{Data: r.Data}
}

func NewRowFromProto (rp *pb.Row) *Row {
	return &Row{Data: rp.GetData()}
}

type Component struct {
	Rows []*Row
}

func (c Component) At(row int64, column int64) float64 {
	return c.Rows[row].Data[column]
}

func (c Component) ToProto() *pb.ImageComponent {
	rows := make([]*pb.Row, len(c.Rows))

	for index, row := range c.Rows {
		rows[index] = row.ToProto()
	}

	return &pb.ImageComponent{Rows: rows}
}

func NewComponentFromProto(cp *pb.ImageComponent) Component {
	rows := make([]*Row, len(cp.Rows))

	for index, row := range cp.Rows {
		rows[index] = NewRowFromProto(row)
	}

	return Component{Rows: rows}
}

func NewComponentsFromProtos(cps []*pb.ImageComponent) []Component {
	components := make([]Component, len(cps))

	for index, cp := range cps {
		components[index] = NewComponentFromProto(cp)
	}

	return components
}

type Plane struct {
	Width      int64
	Height     int64
	Labels     []string
	Components []Component
}

func NewRGBPlane(width int64, height int64, components []Component) Plane {
	return Plane{Width: width, Height: height, Labels: []string{"Red", "Green", "Blue"}, Components: components}
}

func NewPlaneFromProto(ipp *pb.ImagePlane) Plane {
	return Plane{Width: ipp.GetWidth(), Height: ipp.GetHeight(), Labels: ipp.GetLabels(), Components: NewComponentsFromProtos(ipp.GetComponents())}
}

func (p Plane) ToNRGBA() *image.NRGBA {
	log.Printf("Width: %v, Height: %v", p.Width, p.Height)
	img := image.NewNRGBA(image.Rect(0, 0, int(p.Width), int(p.Height)))

	var y, x int64
	for y = 0; y < p.Height; y++ {
		for x = 0; x < p.Width; x++ {
			img.Set(int(x), int(y), color.NRGBA{
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

	return &pb.ImagePlane{Width: p.Width, Height: p.Height, Labels: p.Labels, Components: cps}
}

type Frame struct {
	Labels []string
	Planes []Plane
}
