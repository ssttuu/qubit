package computation

import (
	"github.com/stupschwartz/qubit/core/operator"
	pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

const (
	PubSubTopicID = "computations"
	TableName     = "computations"
)

type Computation struct {
	Id             string                       `db:"id"`
	RootOperatorId string                       `db:"root_operator_id"`
	OperatorMap    map[string]operator.Operator `db:"operator_map"`
	ResourceId     string                       `db:"resource_id"`
}

type Computations []Computation

func NewFromProto(pbComputation *pb.Computation) Computation {
	operatorMap := map[string]operator.Operator{}
	for key, op := range pbComputation.GetOperatorMap() {
		operatorMap[key] = operator.NewFromProto(op)
	}
	return Computation{
		Id:             pbComputation.GetId(),
		RootOperatorId: pbComputation.GetRootOperatorId(),
		OperatorMap:    operatorMap,
		ResourceId:     pbComputation.GetResourceId(),
	}
}

func (c *Computation) ToProto() *pb.Computation {
	opProtoMap := map[string]*operators_pb.Operator{}
	for key, op := range c.OperatorMap {
		opProtoMap[key] = op.ToProto()
	}
	return &pb.Computation{
		Id:             c.Id,
		RootOperatorId: c.RootOperatorId,
		OperatorMap:    opProtoMap,
		ResourceId:     c.ResourceId,
	}
}

func (c *Computation) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"root_operator_id": c.RootOperatorId,
		"operator_map":     c.OperatorMap,
	}
}

func (c *Computation) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"root_operator_id": c.RootOperatorId,
		"operator_map":     c.OperatorMap,
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
