package nodes

import (
	"github.com/stupschwartz/qubit/core/node"
	"github.com/stupschwartz/qubit/server/env"
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	nodes_pb "github.com/stupschwartz/qubit/server/protos/nodes"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

const SceneKind string = "Scene"
const NodeKind string = "Node"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}


// Server implements `service Health`.
type Server struct {
	env *env.Env
}

func (s *Server) List(ctx context.Context, in *nodes_pb.ListNodesRequest) (*nodes_pb.NodesList, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)

	var nodes node.Nodes
	_, err := s.env.DatastoreClient.GetAll(ctx, datastore.NewQuery(NodeKind).Ancestor(sceneKey), &nodes)
	if err != nil {
		return nil, errors.Wrap(err, "Could not get all")
	}

	return nodes.ToProto(), nil
}

func (s *Server) Get(ctx context.Context, in *nodes_pb.GetNodeRequest) (*nodes_pb.Node, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)
	nodeKey := datastore.IDKey(NodeKind, in.NodeId, sceneKey)

	var existingNode node.Node
	if err := s.env.DatastoreClient.Get(ctx, nodeKey, &existingNode); err != nil {
		return nil, errors.Wrap(err, "Could not get datastore entity")
	}

	return existingNode.ToProto(), nil
}

func (s *Server) Create(ctx context.Context, in *nodes_pb.CreateNodeRequest) (*nodes_pb.Node, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)

	nodeId := r.Int63()
	nodeKey := datastore.IDKey(NodeKind, nodeId, sceneKey)

	newNode := node.NewNodeFromProto(in.Node)
	newNode.Id = nodeId

	if _, err := s.env.DatastoreClient.Put(ctx, nodeKey, &newNode); err != nil {
		return nil, errors.Wrapf(err, "Failed to put node %v", newNode.Id)
	}

	return newNode.ToProto(), nil
}

func (s *Server) Update(ctx context.Context, in *nodes_pb.UpdateNodeRequest) (*nodes_pb.Node, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)
	nodeKey := datastore.IDKey(NodeKind, in.NodeId, sceneKey)

	newNode := node.NewNodeFromProto(in.Node)

	_, err := s.env.DatastoreClient.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var existingNode node.Node
		if err := tx.Get(nodeKey, &existingNode); err != nil {
			return errors.Wrapf(err, "Failed to get node in tx %v", existingNode)
		}

		existingNode.Name = newNode.Name

		_, err := tx.Put(nodeKey, &existingNode)
		if err != nil {
			return errors.Wrapf(err, "Failed to put node in tx %v", existingNode)
		}

		newNode = existingNode

		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "Failed to update node")
	}

	return newNode.ToProto(), nil
}

func (s *Server) Delete(ctx context.Context, in *nodes_pb.DeleteNodeRequest) (*empty.Empty, error) {
	sceneKey := datastore.IDKey(SceneKind, in.SceneId, nil)
	nodeKey := datastore.IDKey(NodeKind, in.NodeId, sceneKey)

	if err := s.env.DatastoreClient.Delete(ctx, nodeKey); err != nil {
		return nil, errors.Wrapf(err, "Failed to delete node by id: %v", in.NodeId)
	}

	return &empty.Empty{}, nil
}

func newServer(e *env.Env) *Server {
	return &Server{
		env: e,
	}
}

func Register(server *grpc.Server, e *env.Env) {
	nodes_pb.RegisterNodesServer(server, newServer(e))
}
