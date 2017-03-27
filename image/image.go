package image

import (
	"github.com/gonum/matrix/mat64"
	"fmt"
	"image"
	"image/color"
	"log"
)

type Component struct {
	*mat64.Dense
	Label string
}

func (c Component) String() string {
	return fmt.Sprintf("%v: %v", c.Label, mat64.Formatted(c, mat64.Prefix(" "), mat64.Squeeze()))
}

type Plane struct {
	Components      []Component
	Label string
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

type Frame struct {
	Planes      []Plane
}
