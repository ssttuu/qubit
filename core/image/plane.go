package image

import (
	"image"
	"image/color"

	pb "github.com/stupschwartz/qubit/proto-gen/go/images"
)

type Plane struct {
	Name       string
	Width      int32
	Height     int32
	Labels     map[string]string
	Channels   []Channel
}

func (p *Plane) Merge(otherPlane *Plane, startX int32, startY int32) {
	for index, cp := range otherPlane.Channels {
		p.Channels[index].Merge(&cp, startX, startY)
	}
}

func NewPlane(width int32, height int32, channels []Channel) *Plane {
	return &Plane{Width:width, Height:height, Channels: channels}
}

func NewRGBZeroPlane(width int32, height int32) *Plane {
	channels := []Channel{}
	for i := 0; i < 3; i++ {
		c := Channel{}
		c.Zero(width, height)
		channels = append(channels, c)
	}

	return NewPlane(width, height, channels)
}

func NewPlaneFromProto(pb_plane *pb.Plane) *Plane {
	return &Plane{Labels: pb_plane.GetLabels(), Channels: NewChannelsFromProtos(pb_plane.GetChannels())}
}

func (p *Plane) ToNRGBA() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, int(p.Width), int(p.Height)))

	var y, x int32
	for y = 0; y < p.Height; y++ {
		for x = 0; x < p.Width; x++ {
			img.Set(int(x), int(y), color.NRGBA{
				R: uint8(p.Channels[0].At(y, x) * 255),
				G: uint8(p.Channels[1].At(y, x) * 255),
				B: uint8(p.Channels[2].At(y, x) * 255),
				A: 255,
			})
		}
	}

	return img
}

func (p *Plane) ToProto() *pb.Plane {
	pbchannels := make([]*pb.Channel, len(p.Channels))

	for index, channel := range p.Channels {
		pbchannels[index] = channel.ToProto()
	}

	return &pb.Plane{Labels: p.Labels, Channels: pbchannels}
}