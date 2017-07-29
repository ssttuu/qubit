package scene_event

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/applications/lib/pgutils"
)

const TableName = "scene_events"

func getLastSceneEvent(pgClient *sqlx.DB, sceneId string) (*SceneEvent, error) {
	var sceneEvent SceneEvent
	err := pgutils.Select(&pgutils.SelectConfig{
		DB:   pgClient,
		Args: []interface{}{sceneId},
		// Strictly limit columns for performance
		Columns: []string{
			"up_version",
		},
		Limit:         1,
		OrderByClause: "ORDER BY up_version",
		Out:           &[]SceneEvent{sceneEvent},
		Table:         TableName,
		WhereClause:   "WHERE scene_id=$1",
	})
	if err != nil {
		return nil, err
	}
	return &sceneEvent, nil
}

func Append(pgClient *sqlx.DB, sceneId string, upChangeData *SceneEventChange, downChangeData *SceneEventChange) (*SceneEvent, error) {
	previousSceneEvent, err := getLastSceneEvent(pgClient, sceneId)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get last SceneEvent")
	}
	newSceneEvent := SceneEvent{
		DownVersion:    previousSceneEvent.UpVersion,
		DownChangeData: downChangeData,
		SceneId:        sceneId,
		UpChangeData:   upChangeData,
		UpVersion:      previousSceneEvent.UpVersion + 1,
	}
	err = apiutils.Create(&apiutils.CreateConfig{
		DB:     pgClient,
		Object: &newSceneEvent,
		Table:  TableName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create SceneEvent")
	}
	return &newSceneEvent, nil
}
