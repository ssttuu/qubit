package operators

import (
	"cloud.google.com/go/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/core/scene"
	"github.com/stupschwartz/qubit/core/scene_event"
	operators_pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

type Server struct {
	PostgresClient *sqlx.DB
	StorageClient  *storage.Client
}

// API

func Register(grpcServer *grpc.Server, postgresClient *sqlx.DB, storageClient *storage.Client) {
	operators_pb.RegisterOperatorsServer(grpcServer, &Server{
		PostgresClient: postgresClient,
		StorageClient:  storageClient,
	})
}

// Create adds a new CreateOperator event to the OperatorEvent table and updates the current snapshot.
func (s *Server) Create(ctx context.Context, in *operators_pb.CreateOperatorRequest) (*operators_pb.Operator, error) {
	newOperator := scene.NewOperatorFromProto(in.Operator)
	err := scene.CreateOperator(s.PostgresClient, in.GetSceneId(), newOperator)
	if err != nil {
		return nil, err
	}
	_, err = scene_event.Append(
		s.PostgresClient,
		newOperator.Id, // TODO: create unique id
		scene_event.SceneEventChange{
			Action: scene.OperatorCreated,
		},
		scene_event.SceneEventChange{
			Action: scene.OperatorUnCreated,
		},
	)
	if err != nil {
		return nil, err
	}
	return newOperator.ToProto(), nil
}

// Delete adds a DeleteOperator event to the OperatorEvent table.
func (s *Server) Delete(ctx context.Context, in *operators_pb.DeleteOperatorRequest) (*empty.Empty, error) {
	err := scene.DeleteOperator(s.PostgresClient, in.GetSceneId(), in.GetId())
	if err != nil {
		return nil, err
	}
	_, err = scene_event.Append(
		s.PostgresClient,
		in.GetSceneId(),
		scene_event.SceneEventChange{
			Action: scene.OperatorDeleted,
		},
		scene_event.SceneEventChange{
			Action: scene.OperatorUnDeleted,
		},
	)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// Get selects the operator by the given id from the Operators table.
func (s *Server) Get(ctx context.Context, in *operators_pb.GetOperatorRequest) (*operators_pb.Operator, error) {
	op, err := scene.GetOperator(s.PostgresClient, in.GetId())
	if err != nil {
		return nil, err
	}
	return op.ToProto(), nil
}

// List selects all operators from the Operators table.
func (s *Server) List(ctx context.Context, in *operators_pb.ListOperatorsRequest) (*operators_pb.ListOperatorsResponse, error) {
	operatorsList, err := scene.ListOperator(s.PostgresClient, in.GetSceneId())
	if err != nil {
		return nil, err
	}
	return &operators_pb.ListOperatorsResponse{Operators: operatorsList.ToProto()}, nil
}

// Rename creates a new rename event on the scene_event table and updates the latest operator object with the new name.
func (s *Server) Rename(ctx context.Context, in *operators_pb.RenameOperatorRequest) (*operators_pb.Operator, error) {
	operatorId := in.GetId()
	previousOperator, err := scene.GetOperator(s.PostgresClient, operatorId)
	if err != nil {
		return nil, err
	}
	// TODO: apply event logic in Operator object
	newOperator := previousOperator
	newOperator.Name = in.GetName()
	err = scene.UpdateOperator(s.PostgresClient, in.GetSceneId(), &newOperator)
	if err != nil {
		return nil, err
	}
	_, err = scene_event.Append(
		s.PostgresClient,
		in.GetSceneId(),
		scene_event.SceneEventChange{
			Action: scene.OperatorRenamed,
			Changes: map[string]interface{}{
				"name": in.GetName(),
			},
		},
		scene_event.SceneEventChange{
			Action: scene.OperatorUnRenamed,
			Changes: map[string]interface{}{
				"name": previousOperator.Name,
			},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create OperatorEvent")
	}
	return &empty.Empty{}, nil
}

func (s *Server) Connect(ctx context.Context, in *operators_pb.ConnectOperatorRequest) (*operators_pb.Connection, error) {
	newConnection := scene.NewConnectionFromProto(in.Connection)
	err := scene.CreateConnection(s.PostgresClient, in.GetSceneId(), newConnection)
	if err != nil {
		return nil, err
	}
	_, err = scene_event.Append(
		s.PostgresClient,
		in.GetSceneId(),
		scene_event.SceneEventChange{
			Action: scene.ConnectionCreated,
			Changes: map[string]interface{}{
				"id":          newConnection.Id,
				"input_id":    newConnection.InputId,
				"input_index": newConnection.InputIndex,
				"output_id":   newConnection.OutputId,
			},
		},
		scene_event.SceneEventChange{
			Action: scene.ConnectionUnCreated,
			Changes: map[string]interface{}{
				"id":          newConnection.Id,
				"input_id":    newConnection.InputId,
				"input_index": newConnection.InputIndex,
				"output_id":   newConnection.OutputId,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return newConnection.ToProto(), nil
}

func (s *Server) Disconnect(ctx context.Context, in *operators_pb.DisconnectOperatorRequest) (*empty.Empty, error) {
	sceneId := in.GetSceneId()
	connectionId := in.GetConnectionId()
	err := scene.DeleteConnection(s.PostgresClient, sceneId, connectionId)
	if err != nil {
		return nil, err
	}
	_, err = scene_event.Append(
		s.PostgresClient,
		sceneId,
		scene_event.SceneEventChange{
			Action: scene.ConnectionDeleted,
			Changes: map[string]interface{}{
				"id": connectionId,
			},
		},
		scene_event.SceneEventChange{
			Action: scene.ConnectionUnDeleted,
			Changes: map[string]interface{}{
				"id": connectionId,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// TODO: haven't figured out how to properly store / access parameters.  Not sure the data model is right yet.
func (s *Server) SetValue(ctx context.Context, in *operators_pb.SetValueRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *Server) SetKeyFrame(ctx context.Context, in *operators_pb.SetKeyFrameRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *Server) SetExpression(ctx context.Context, in *operators_pb.SetExpressionRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
