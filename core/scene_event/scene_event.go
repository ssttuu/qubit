package scene_event

import (
	"encoding/json"
	"errors"

	pb "github.com/stupschwartz/qubit/proto-gen/go/scene_events"
)

const TableName = "scene_events"

type SceneEventChange struct {
	// TODO: Enum
	Action     string                 `json:"action"` // create, update, delete
	Changes    map[string]interface{} `json:"changes"`
	OperatorId string                 `json:"operator_id"`
}

type SceneEvent struct {
	DownChangeData []byte
	DownVersion    int32
	Id             string
	SceneId        string
	UpChangeData   []byte
	UpVersion      int32
}

type SceneEvents []SceneEvent

// TODO: Return a pointer
func NewFromProto(pbSceneEvent *pb.SceneEvent) SceneEvent {
	return SceneEvent{
		DownChangeData: pbSceneEvent.GetDownChangeData(),
		DownVersion:    pbSceneEvent.GetDownVersion(),
		Id:             pbSceneEvent.GetId(),
		SceneId:        pbSceneEvent.GetSceneId(),
		UpChangeData:   pbSceneEvent.GetUpChangeData(),
		UpVersion:      pbSceneEvent.GetUpVersion(),
	}
}

func (s *SceneEvent) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"down_change_data": s.DownChangeData,
		"down_version":     s.DownVersion,
		"scene_id":         s.SceneId,
		"up_change_data":   s.UpChangeData,
		"up_version":       s.UpVersion,
	}
}

func (s *SceneEvent) GetDownSceneEventChange() (*SceneEventChange, error) {
	var sec SceneEventChange
	err := json.Unmarshal(s.DownChangeData, &sec)
	return &sec, err
}

func (s *SceneEvent) GetUpSceneEventChange() (*SceneEventChange, error) {
	var sec SceneEventChange
	err := json.Unmarshal(s.UpChangeData, &sec)
	return &sec, err
}

func (s *SceneEvent) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{}
}

func (s *SceneEvent) ToProto() *pb.SceneEvent {
	return &pb.SceneEvent{
		DownChangeData: s.DownChangeData,
		DownVersion:    s.DownVersion,
		Id:             s.Id,
		SceneId:        s.SceneId,
		UpChangeData:   s.UpChangeData,
		UpVersion:      s.UpVersion,
	}
}

func (s *SceneEvent) ValidateCreate() error {
	// TODO
	return nil
}

func (s *SceneEvent) ValidateUpdate(newObj interface{}) error {
	return errors.New("Update is not supported for scene events")
}

func (s *SceneEvents) ToProto() []*pb.SceneEvent {
	var pbscenes []*pb.SceneEvent
	for _, sceneEvent := range *s {
		sceneEventProto := sceneEvent.ToProto()
		pbscenes = append(pbscenes, sceneEventProto)
	}
	return pbscenes
}
