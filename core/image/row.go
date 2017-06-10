package image

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/images"
)

type Row struct {
	Data []float64
}

func (r *Row) Merge(otherRow *Row, startX int32) {
	r.Data = append(r.Data[:startX], otherRow.Data...)
}

func (r *Row) ToProto() *pb.Row {
	return &pb.Row{Data: r.Data}
}

func NewRowFromProto(rp *pb.Row) *Row {
	return &Row{Data: rp.GetData()}
}
