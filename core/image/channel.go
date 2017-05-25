package image

import (
	pb "github.com/stupschwartz/qubit/server/protos/images"
)

type Channel struct {
	Name string
	Rows []*Row
}

func (c *Channel) HasData() bool {
	return len(c.Rows) > 0
}

func (c *Channel) At(row int32, column int32) float64 {
	return c.Rows[row].Data[column]
}

func (c *Channel) Zero(width int32, height int32) {
	var y, x int32
	for y = 0; y < height; y++ {
		c.Rows = append(c.Rows, &Row{})

		for x = 0; x < width; x++ {
			// fmt.Printf("addValue:%v,%v=0.0", x, y)
			c.Rows[y].Data = append(c.Rows[y].Data, 0.0)
		}
	}

	// fmt.Printf("ROWS: %v", len(c.Rows))
}

func (c *Channel) Merge(otherChannel *Channel, startX int32, startY int32) {
	// fmt.Printf("ChannelMerge lens: %v, %v\n", len(c.Rows), len(otherChannel.Rows))


	for index, row := range otherChannel.Rows {
		rowIndex := startY + int32(index)
		// fmt.Printf("index: %v, startY: %v, rowIndex: %v\n", index, startY, rowIndex)

		c.Rows[rowIndex].Merge(row, startX)
	}
}

func (c *Channel) ToProto() *pb.Channel {
	rows := make([]*pb.Row, len(c.Rows))

	for index, row := range c.Rows {
		rows[index] = row.ToProto()
	}

	return &pb.Channel{Rows: rows}
}

func NewChannelFromProto(cp *pb.Channel) Channel {
	rows := make([]*Row, len(cp.Rows))

	for index, row := range cp.Rows {
		rows[index] = NewRowFromProto(row)
	}

	return Channel{Rows: rows}
}

func NewChannelsFromProtos(cps []*pb.Channel) []Channel {
	channels := make([]Channel, len(cps))

	for index, cp := range cps {
		channels[index] = NewChannelFromProto(cp)
	}

	return channels
}
