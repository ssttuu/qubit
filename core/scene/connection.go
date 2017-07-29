package scene

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stupschwartz/qubit/applications/lib/apiutils"
	pb "github.com/stupschwartz/qubit/proto-gen/go/operators"
)

type Connection struct {
	Id         string `db:"id"`
	InputId    string `db:"input_id"`
	InputIndex int32  `db:"input_index"`
	OutputId   string `db:"output_id"`
	//OutputIndex int32 `db:"output_index"`
}

func (c *Connection) GetCreateData() map[string]interface{} {
	return map[string]interface{}{
		"id": c.Id,
	}
}

func (c *Connection) GetUpdateData() map[string]interface{} {
	return map[string]interface{}{
		"id": c.Id,
	}
}

func (c *Connection) ValidateCreate() error {
	return nil
}
func (c *Connection) ValidateUpdate(existingObject interface{}) error {
	return nil
}

func (o *Connection) ToProto() *pb.Connection {
	return &pb.Connection{
		Id:         o.Id,
		InputId:    o.InputId,
		InputIndex: o.InputIndex,
		OutputId:   o.OutputId,
		//OutputIndex: o.OutputIndex,
	}
}

func NewConnectionFromProto(pbConnection *pb.Connection) Connection {
	return Connection{
		Id:         pbConnection.Id,
		InputId:    pbConnection.InputId,
		InputIndex: pbConnection.InputIndex,
		OutputId:   pbConnection.OutputId,
		//OutputIndex: pbConnection.OutputIndex,
	}
}

type Connections []Connection

func (cs Connections) ToProto() []*pb.Connection {
	var pbConnections []*pb.Connection
	for _, c := range cs {
		pbConnections = append(pbConnections, c.ToProto())
	}
	return pbConnections
}

type ConnectionMap map[string]Connection

func (cs ConnectionMap) ToProto() map[string]*pb.Connection {
	var pbConnections map[string]*pb.Connection
	for id, c := range cs {
		pbConnections[id] = c.ToProto()
	}
	return pbConnections
}

// DB

const ConnectionsTableName = "connections"

// Create creates a new record in the connections table and adds a connection to the
// scene Connections map.
func CreateConnection(pgClient *sqlx.DB, sceneId string, connection *Connection) error {
	err := apiutils.Create(&apiutils.CreateConfig{
		DB:     pgClient,
		Object: connection,
		Table:  ConnectionsTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to create Connection: %v", connection)
	}
	s, err := GetScene(pgClient, sceneId)
	if err != nil {
		return err
	}
	s.Connections[connection.Id] = *connection
	err = UpdateScene(pgClient, s)
	if err != nil {
		return err
	}
	return err
}

func GetConnection(pgClient *sqlx.DB, connectionId string) (*Connection, error) {
	var c Connection
	err := apiutils.Get(&apiutils.GetConfig{
		DB:    pgClient,
		Id:    connectionId,
		Out:   &c,
		Table: ConnectionsTableName,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Connection by id: %v", connectionId)
	}
	return &c, nil
}

func ListConnection(pgClient *sqlx.DB, sceneId string) (Connections, error) {
	var connectionList Connections
	err := apiutils.List(&apiutils.ListConfig{
		DB:          pgClient,
		Out:         &connectionList,
		Table:       ConnectionsTableName,
		WhereClause: fmt.Sprintf("scene_id = %v", sceneId),
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get Connections")
	}
	return connectionList, nil
}

func UpdateConnection(pgClient *sqlx.DB, sceneId string, connection *Connection) error {
	err := apiutils.Update(&apiutils.UpdateConfig{
		DB:        pgClient,
		Id:        connection.Id,
		NewObject: connection,
		OldObject: &Connection{},
		Table:     ConnectionsTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to update Connection: %v", connection)
	}
	s, err := GetScene(pgClient, sceneId)
	if err != nil {
		return err
	}
	s.Connections[connection.Id] = *connection
	err = UpdateScene(pgClient, s)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConnection(pgClient *sqlx.DB, sceneId string, connectionId string) error {
	err := apiutils.Delete(&apiutils.DeleteConfig{
		DB:    pgClient,
		Id:    connectionId,
		Table: ConnectionsTableName,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to delete Connection by id: %v", connectionId)
	}
	s, err := GetScene(pgClient, sceneId)
	if err != nil {
		return err
	}
	delete(s.Connections, connectionId)
	err = UpdateScene(pgClient, s)
	if err != nil {
		return err
	}
	return err
}
