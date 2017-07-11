package computation

import (
	"github.com/stupschwartz/qubit/core/geometry"
	"github.com/stupschwartz/qubit/core/scene"
	pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
)

const (
	TableName = "computations"
)

type Computation struct {
	Id          string
	Scene       scene.Scene
	OperatorId  string
	Time        float64
	BoundingBox *geometry.BoundingBox2D
	ResourceId  string
}

type Computations []Computation

func NewFromProto(pbComputation *pb.Computation) Computation {
	return Computation{
		Id:          pbComputation.GetId(),
		Scene:       scene.NewFromProto(pbComputation.GetScene()),
		OperatorId:  pbComputation.GetOperatorId(),
		Time:        pbComputation.GetTime(),
		BoundingBox: geometry.NewBoundingBoxFromProto(pbComputation.GetBoundingBox()),
		ResourceId:  pbComputation.GetResourceId(),
	}
}

func (c *Computation) ToProto() *pb.Computation {
	return &pb.Computation{
		Id:          c.Id,
		Scene:       c.Scene.ToProto(),
		OperatorId:  c.OperatorId,
		Time:        c.Time,
		BoundingBox: c.BoundingBox.ToProto(),
		ResourceId:  c.ResourceId,
	}
}

func (c *Computation) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"operator_id":    c.OperatorId,
		"time":           c.Time,
		"bounding_box2d": c.BoundingBox,
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
