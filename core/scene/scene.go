package scene

import pb "github.com/stupschwartz/qubit/compute/protos/compute"

type Scene struct {
	Id string `json:"id" datastore:"id"`
	Version int `json:"version" datastore:"version"`
	Name string `json:"name" datastore:"name"`
	Type string `json:"type" datastore:"type"`
}

func (s *Scene) ToProto() *pb.Scene {
	return &pb.Scene{Id: s.Id}
}
