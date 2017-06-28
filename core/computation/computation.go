package computation

import (
	"github.com/stupschwartz/qubit/core/geometry"
	pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
)

const (
	TableName = "computations"
)

type Computation struct {
	Id            string                  `db:"id"`
	OperatorKey   string                  `db:"operator_key"`
	Time          float32                 `db:"time"`
	BoundingBox2D *geometry.BoundingBox2D `db:"bounding_box2d"`
	ResourceId    string                  `db:"resource_id"`
}

type Computations []Computation

func NewFromProto(pbComputation *pb.Computation) Computation {
	return Computation{
		Id:            pbComputation.GetId(),
		OperatorKey:   pbComputation.GetOperatorKey(),
		Time:          pbComputation.GetTime(),
		BoundingBox2D: geometry.NewBoundingBoxFromProto(pbComputation.GetBoundingBox2D()),
		ResourceId:    pbComputation.GetResourceId(),
	}
}

func (c *Computation) ToProto() *pb.Computation {
	return &pb.Computation{
		Id:            c.Id,
		OperatorKey:   c.OperatorKey,
		Time:          c.Time,
		BoundingBox2D: c.BoundingBox2D.ToProto(),
		ResourceId:    c.ResourceId,
	}
}

func (c *Computation) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"operator_key":   c.OperatorKey,
		"time":           c.Time,
		"bounding_box2d": c.BoundingBox2D,
	}
}

func (c *Computation) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"resource_id": c.ResourceId,
	}
}

func (c *Computation) ValidateCreate() error {
	return nil
}

func (c *Computation) ValidateUpdate(newObj interface{}) error {
	//org := newObj.(*Computation)
	return nil
}

func (c *Computations) ToProto() []*pb.Computation {
	var pbComputations []*pb.Computation
	for _, computation := range *c {
		pbComputations = append(pbComputations, computation.ToProto())
	}
	return pbComputations
}
