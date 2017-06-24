package computation_status

import (
	"time"

	pb "github.com/stupschwartz/qubit/proto-gen/go/computations"
)

const (
	TableName = "computation_statuses"
)

const (
	ComputationStatusCreated   string = "created"
	ComputationStatusStarted          = "started"
	ComputationStatusCompleted        = "completed"
)

type ComputationStatus struct {
	Id            string    `db:"id"`
	Status        string    `db:"status"`
	ComputationId string    `db:"computation_id"`
	CreatedAt     time.Time `db:"created_at"`
}

type ComputationStatuses []ComputationStatus

func NewFromProto(pbComputationStatus *pb.ComputationStatus) ComputationStatus {
	return ComputationStatus{
		Id:            pbComputationStatus.GetId(),
		ComputationId: pbComputationStatus.GetComputationId(),
		Status:        pbComputationStatus.GetStatus(),
		CreatedAt:     time.Unix(pbComputationStatus.GetCreatedAt(), 0),
	}
}

func (cs *ComputationStatus) ToProto() *pb.ComputationStatus {
	return &pb.ComputationStatus{
		Id:            cs.Id,
		ComputationId: cs.ComputationId,
		Status:        cs.Status,
		CreatedAt:     cs.CreatedAt.Unix(),
	}
}

func (cs *ComputationStatus) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"computation_id": cs.ComputationId,
	}
}

func (cs *ComputationStatus) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"status": cs.Status,
	}
}

func (cs *ComputationStatus) ValidateCreate() error {
	return nil
}

func (cs *ComputationStatus) ValidateUpdate(newObj interface{}) error {
	//obj := newObj.(*ComputationStatus)
	return nil
}

func (cs *ComputationStatuses) ToProto() []*pb.ComputationStatus {
	var pbComputationStatuses []*pb.ComputationStatus
	for _, computationStatus := range *cs {
		pbComputationStatuses = append(pbComputationStatuses, computationStatus.ToProto())
	}
	return pbComputationStatuses
}
