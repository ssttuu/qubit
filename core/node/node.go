package node

import (
	pb "github.com/stupschwartz/qubit/server/protos/nodes"
)

type Node struct {
	Id int64 `json:"id" datastore:"id"`
	Name string `json:"name" datastore:"name"`
	Type string `json:"type" datastore:"type"`
	Inputs []string `json:"inputs" datastore:"inputs"`
	Outputs []string `json:"outputs" datastore:"outputs"`
}

func (n *Node) ToProto() *pb.Node {
	return &pb.Node{Id: n.Id}
}

func NewNodeFromProto(pbnode *pb.Node) Node {
	return Node{Id: pbnode.Id}
}

type Nodes []*Node

func (n *Nodes) ToProto() *pb.NodesList {
	var pbnodes []*pb.Node
	for _, node := range *n {
		pbnodes = append(pbnodes, node.ToProto())
	}

	return &pb.NodesList{Nodes:pbnodes}
}
