package computation

import (
	"github.com/stupschwartz/qubit/core/operator"
	pb "github.com/stupschwartz/qubit/proto-gen/go/compute"
	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

var Kind string = "computation"
var PubsubTopicID string = "computations"

type Computation struct {
	Id             string                       `db:"id"`
	Status         string                       `db:"status"`
	RootOperatorId string                       `db:"root_operator_id"`
	OperatorMap    map[string]operator.Operator `db:"operator_map"`
	ResourceId     string                       `db:"resource_id"`
}

type Computations []Computation

func NewFromProto(pbcomputation *pb.Computation) Computation {
	operatorMap := map[string]operator.Operator{}
	for key, op := range pbcomputation.GetOperatorMap() {
		operatorMap[key] = operator.NewFromProto(op)
	}
	return Computation{
		Id:             pbcomputation.GetId(),
		Status:         pbcomputation.GetStatus(),
		RootOperatorId: pbcomputation.GetRootOperatorId(),
		OperatorMap:    operatorMap,
		ResourceId:     pbcomputation.GetResourceId(),
	}
}

func (o *Computation) ToProto() *pb.Computation {
	opProtoMap := map[string]*operators_pb.Operator{}
	for key, op := range o.OperatorMap {
		opProtoMap[key] = op.ToProto()
	}
	return &pb.Computation{
		Id:             o.Id,
		Status:         o.Status,
		RootOperatorId: o.RootOperatorId,
		OperatorMap:    opProtoMap,
		ResourceId:     o.ResourceId,
	}
}

func (o *Computation) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"root_operator_id": o.RootOperatorId,
		"operator_map":     o.OperatorMap,
	}
}

func (o *Computation) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"root_operator_id": o.RootOperatorId,
		"operator_map":     o.OperatorMap,
	}
}

func (o *Computation) ValidateCreate() error {
	return nil
}

func (o *Computation) ValidateUpdate(newObj interface{}) error {
	//org := newObj.(*Computation)
	return nil
}

func (o *Computations) ToProto() []*pb.Computation {
	var pbcomputations []*pb.Computation
	for _, computation := range *o {
		pbcomputations = append(pbcomputations, computation.ToProto())
	}
	return pbcomputations
}
