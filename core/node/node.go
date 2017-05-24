package node

import (
	pb "github.com/stupschwartz/qubit/server/protos/nodes"
	"fmt"
	"strconv"
	"github.com/pkg/errors"
)

type Node struct {
	Id      string `json:"id" datastore:"id"`
	Name    string `json:"name" datastore:"name"`
	Type    string `json:"type" datastore:"type"`
	Inputs  []string `json:"inputs" datastore:"inputs"`
	Outputs []string `json:"outputs" datastore:"outputs"`
}

func (n *Node) ToProto() (*pb.Node, error) {
	i, err := strconv.ParseInt(n.Id, 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert Id from string to int64, %v", n.Id)
	}
	return &pb.Node{Id: i}, nil
}

func NewNodeFromProto(pbnode *pb.Node) Node {
	return Node{Id: fmt.Sprint(pbnode.Id)}
}

type Nodes []*Node

func (n *Nodes) ToProto() (*pb.NodesList, error) {
	var pbnodes []*pb.Node
	for _, node := range *n {
		node_proto, err := node.ToProto()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to convert node to proto, %v", node)
		}
		pbnodes = append(pbnodes, node_proto)
	}

	return &pb.NodesList{Nodes:pbnodes}, nil
}
