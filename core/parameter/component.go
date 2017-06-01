package parameter

import pb "github.com/stupschwartz/qubit/server/protos/parameters"

type Component struct {
	Id    string `json:"id"`
	Value float64 `json:"value"`
}

func (c *Component) GetValue() float64 {
	return c.Value
}

func (c *Component) SetValue(value float64) {
	c.Value = value
}

func (c *Component) ToProto() *pb.Component {
	return &pb.Component{Id: c.Id, Value: c.Value}
}

func NewComponentFromProto(pb_cp *pb.Component) *Component {
	return &Component{Id: pb_cp.Id, Value: pb_cp.Value}
}

type Components []*Component

func (c *Components) ToProto() []*pb.Component {
	var pb_components []*pb.Component
	for _, component := range *c {
		pb_components = append(pb_components, component.ToProto())
	}

	return pb_components
}

func NewComponentsFromProto(pb_cps []*pb.Component) Components {
	var cps Components
	for _, cp := range pb_cps {
		cps = append(cps, NewComponentFromProto(cp))
	}
	return cps
}
