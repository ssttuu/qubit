package image

import (
	"image"
	"image/color"
	"log"
	pb "github.com/stupschwartz/qubit/compute/protos/compute"
	//"fmt"
)


type Row struct {
	Data []float64
}

func (r *Row) Merge(otherRow *Row, startX int64) {
	// fmt.Printf("RowMerge: lens: %v %v", len(r.Data), len(otherRow.Data))
	r.Data = append(r.Data[:startX], otherRow.Data...)
}

func (r *Row) ToProto() *pb.Row {
	return &pb.Row{Data: r.Data}
}

func NewRowFromProto (rp *pb.Row) *Row {
	return &Row{Data: rp.GetData()}
}

type Component struct {
	Rows []*Row
}

func (c *Component) HasData() bool {
	return len(c.Rows) > 0
}

func (c *Component) At(row int64, column int64) float64 {
	return c.Rows[row].Data[column]
}

func (c *Component) Zero(width int64, height int64) {
	var y, x int64
	for y = 0; y < height; y++ {
		c.Rows = append(c.Rows, &Row{})

		for x = 0; x < width; x++ {
			// fmt.Printf("addValue:%v,%v=0.0", x, y)
			c.Rows[y].Data = append(c.Rows[y].Data, 0.0)
		}
	}

	// fmt.Printf("ROWS: %v", len(c.Rows))
}

func (c *Component) Merge(otherComponent *Component, startX int64, startY int64) {
	// fmt.Printf("ComponentMerge lens: %v, %v\n", len(c.Rows), len(otherComponent.Rows))


	for index, row := range otherComponent.Rows {
		rowIndex := startY + int64(index)
		// fmt.Printf("index: %v, startY: %v, rowIndex: %v\n", index, startY, rowIndex)

		c.Rows[rowIndex].Merge(row, startX)
	}
}

func (c *Component) ToProto() *pb.ImageComponent {
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

func (p *Plane) Merge(otherPlane *Plane, startX int64, startY int64) {
	// fmt.Printf("PlaneMerge: lens: %v, %v\n", len(p.Components), len(otherPlane.Components))
	for index, cp := range otherPlane.Components {
		// fmt.Printf("PlaneMerge: index: %v\n", index)
		p.Components[index].Merge(&cp, startX, startY)
	}
}

func NewRGBPlane(width int64, height int64, components []Component) Plane {
	return Plane{Width: width, Height: height, Labels: []string{"Red", "Green", "Blue"}, Components: components}
}

func NewRGBZeroPlane(width int64, height int64) Plane {
	components := []Component{}
	for i := 0; i < 3; i++ {
		c := Component{}
		c.Zero(width, height)
		components = append(components, c)
	}

	// fmt.Println("")
	// fmt.Println(len(components))
	// fmt.Println(components)
	// fmt.Println(len(components[0].Rows))

	return NewRGBPlane(width, height, components)
	return Plane{Width: width, Height: height, Labels: []string{"Red", "Green", "Blue"}, Components: components}
}

func NewRGBAPlane(width int64, height int64, components []Component) Plane {
	return Plane{Width: width, Height: height, Labels: []string{"Red", "Green", "Blue", "Alpha"}, Components: components}
}

func NewPlaneFromProto(ipp *pb.ImagePlane) Plane {
	return Plane{Width: ipp.GetWidth(), Height: ipp.GetHeight(), Labels: ipp.GetLabels(), Components: NewComponentsFromProtos(ipp.GetComponents())}
}

func (p *Plane) ToNRGBA() *image.NRGBA {
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

func (p *Plane) ToProto() *pb.ImagePlane {
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
