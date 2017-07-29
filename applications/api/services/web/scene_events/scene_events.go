package scene_events

import (
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	"github.com/stupschwartz/qubit/applications/lib/pgutils"
	"github.com/stupschwartz/qubit/core/scene"
	"github.com/stupschwartz/qubit/core/scene_event"
	scene_events_pb "github.com/stupschwartz/qubit/proto-gen/go/scene_events"
)

type Server struct {
	PostgresClient *sqlx.DB
}

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB) {
	scene_events_pb.RegisterSceneEventsServer(grpcServer, &Server{
		PostgresClient: postgresClient,
	})
}

func (s *Server) Create(ctx context.Context, in *scene_events_pb.CreateSceneEventRequest) (*scene_events_pb.SceneEvent, error) {
	newObject := scene_event.NewFromProto(in.SceneEvent)
	tx, err := s.PostgresClient.Beginx()
	if err != nil {
		log.Println(errors.Wrap(err, "Failed to begin transaction"))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	var sc scene.Scene
	err = pgutils.SelectByID(&pgutils.SelectConfig{
		Args:        []interface{}{newObject.DownVersion},
		ForClause:   "FOR UPDATE",
		Id:          newObject.SceneId,
		Out:         &sc,
		Table:       scene.ScenesTableName,
		Tx:          tx,
		WhereClause: "WHERE scene_version=$1",
	})
	if err != nil {
		log.Println(errors.Wrapf(err, "Failed to get scene with ID %v and version %v", newObject.SceneId, newObject.DownVersion))
		tx.Rollback()
		return nil, status.Error(codes.NotFound, "Not found")
	}
	ops, err := sc.GetOperators()
	if err != nil {
		log.Println(errors.Wrapf(err, "Failed to unmarshal operators for scene ID %v", newObject.SceneId))
		tx.Rollback()
		return nil, status.Error(codes.NotFound, "Not found")
	}
	log.Println(ops)
	// TODO: Validate change against scene
	// TODO: Apply change to scene
	log.Println(newObject.UpChangeData)
	// TODO: Set DownChangeData
	// Increment versions
	newObject.UpVersion = sc.Version + 1
	sc.Version = newObject.UpVersion
	err = apiutils.Create(&apiutils.CreateConfig{
		Object: &newObject,
		Table:  scene_event.TableName,
		Tx:     tx,
	})
	if err != nil {
		log.Println(errors.Wrap(err, "Failed to create scene event"))
		tx.Rollback()
		return nil, status.Error(codes.Internal, "Internal error")
	}
	operatorJSONData, err := json.Marshal(ops)
	if err != nil {
		log.Println(errors.Wrap(err, "Failed to marshal scene operators"))
		tx.Rollback()
		return nil, status.Error(codes.Internal, "Internal error")
	}
	err = pgutils.UpdateByID(&pgutils.UpdateConfig{
		Id:    sc.Id,
		Table: scene.ScenesTableName,
		Tx:    tx,
		Updates: map[string]interface{}{
			"operators": operatorJSONData,
			"version":   sc.Version,
		},
	})
	if err != nil {
		log.Println(errors.Wrap(err, "Failed to update scene"))
		tx.Rollback()
		return nil, status.Error(codes.Internal, "Internal error")
	}
	err = tx.Commit()
	if err != nil {
		log.Println(errors.Wrap(err, "Failed to commit transaction"))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return newObject.ToProto(), nil
}

func (s *Server) Get(ctx context.Context, in *scene_events_pb.GetSceneEventRequest) (*scene_events_pb.SceneEvent, error) {
	var obj scene_event.SceneEvent
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    s.PostgresClient,
		Id:    in.GetId(),
		Out:   &obj,
		Table: scene_event.TableName,
	})
	if err != nil {
		return nil, err
	}
	return obj.ToProto(), nil
}

func (s *Server) List(ctx context.Context, in *scene_events_pb.ListSceneEventsRequest) (*scene_events_pb.ListSceneEventsResponse, error) {
	var objectList scene_event.SceneEvents
	err := apiutils.List(&apiutils.ListConfig{
		DB:    s.PostgresClient,
		Out:   &objectList,
		Table: scene_event.TableName,
	})
	if err != nil {
		return nil, err
	}
	return &scene_events_pb.ListSceneEventsResponse{SceneEvents: objectList.ToProto()}, nil
}
