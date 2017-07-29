package scene_event

import (
	"encoding/json"

	"github.com/pkg/errors"
	pb "github.com/stupschwartz/qubit/proto-gen/go/scene_events"
)

type SceneEventChange struct {
	Action  string                 `json:"action"` // create, update, delete
	Changes map[string]interface{} `json:"changes"`
	//OperatorId string                 `json:"operator_id"`
}

func (s *SceneEventChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SceneEventChange) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, s)
}

type SceneEvent struct {
	DownChangeData *SceneEventChange
	DownVersion    int32
	Id             string
	SceneId        string
	UpChangeData   *SceneEventChange
	UpVersion      int32
}

type SceneEvents []SceneEvent

// TODO: Return a pointer
func NewFromProto(pbSceneEvent *pb.SceneEvent) SceneEvent {
	var upChangeData SceneEventChange
	if err := json.Unmarshal(pbSceneEvent.GetUpChangeData(), &upChangeData); err != nil {

	}
	var downChangeData SceneEventChange
	if err := json.Unmarshal(pbSceneEvent.GetDownChangeData(), &downChangeData); err != nil {

	}
	return SceneEvent{
		DownChangeData: &downChangeData,
		DownVersion:    pbSceneEvent.GetDownVersion(),
		Id:             pbSceneEvent.GetId(),
		SceneId:        pbSceneEvent.GetSceneId(),
		UpChangeData:   &upChangeData,
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

func (s *SceneEvent) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{}
}

func (s *SceneEvent) ToProto() *pb.SceneEvent {
	pbDownChangeData, err := json.Marshal(s.DownChangeData)
	if err != nil {
		// meh
	}
	pbUpChangeData, err := json.Marshal(s.UpChangeData)
	if err != nil {
		// meh
	}
	return &pb.SceneEvent{
		DownChangeData: pbDownChangeData,
		DownVersion:    s.DownVersion,
		Id:             s.Id,
		SceneId:        s.SceneId,
		UpChangeData:   pbUpChangeData,
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
