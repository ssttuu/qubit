package parameter

import (
	pb "github.com/stupschwartz/qubit/server/protos/parameters"
)



type Parameter struct {
	Id       string `json:"id"`
	Components Components `json:"components"`
}

func (p *Parameter) GetValueByIndex(index int) float64 {
	return p.Components[index].GetValue()
}

func (p *Parameter) SetValueByIndex(index int, value float64) {
	p.Components[index].SetValue(value)
}

func (p *Parameter) GetComponentById(id string) *Component {
	for _, component := range p.Components {
		if component.Id == id {
			return component
		}
	}

	return nil
}

func (p *Parameter) GetValueById(id string) float64 {
	return p.GetComponentById(id).GetValue()
}

func (p *Parameter) SetValueById(id string, value float64) {
	p.GetComponentById(id).SetValue(value)
}

func (p *Parameter) ToProto() *pb.Parameter {
	return &pb.Parameter{
		Id: p.Id,
		Components: p.Components.ToProto(),
	}
}

func NewParameterFromProto(pb_param *pb.Parameter) *Parameter {
	return &Parameter{Id: pb_param.Id, Components: NewComponentsFromProto(pb_param.Components)}
}

func NewFloatParameter(id string) *Parameter {
	return &Parameter{
		Id: id,
		Components: Components{
			&Component{Id: "float", Value: 0.0},
		},
	}
}

func NewColorParameter(id string) *Parameter {
	return &Parameter{
		Id: id,
		Components: Components{
			&Component{Id: "Red", Value: 0.0},
			&Component{Id: "Green", Value: 0.0},
			&Component{Id: "Blue", Value: 0.0},
		},
	}
}

type Parameters []*Parameter

func (p *Parameters) GetById(id string) *Parameter {
	for _, param := range *p {
		if param.Id == id {
			return param
		}
	}

	return nil
}

func (p *Parameters) SetById(id string, component string, value float64) {
	p.GetById(id).SetValueById(component, value)
}

func (p *Parameters) ToProto() []*pb.Parameter {
	var pb_params []*pb.Parameter
	for _, param := range *p {
		pb_params = append(pb_params, param.ToProto())
	}

	return pb_params
}
