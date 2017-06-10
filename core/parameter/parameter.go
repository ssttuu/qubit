package parameter

import (
	pb "github.com/stupschwartz/qubit/proto-gen/go/parameters"
)

type Parameter struct {
	Id         int64      `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Components Components `json:"components" db:"components"`
}

func (p *Parameter) GetValueByIndex(index int) float64 {
	return p.Components[index].GetValue()
}

func (p *Parameter) SetValueByIndex(index int, value float64) {
	p.Components[index].SetValue(value)
}

func (p *Parameter) GetComponentByName(name string) *Component {
	for _, component := range p.Components {
		if component.Name == name {
			return component
		}
	}
	return nil
}

func (p *Parameter) GetValueByName(name string) float64 {
	return p.GetComponentByName(name).GetValue()
}

func (p *Parameter) SetValueByName(name string, value float64) {
	p.GetComponentByName(name).SetValue(value)
}

func (p *Parameter) ToProto() *pb.Parameter {
	return &pb.Parameter{
		Id:         p.Id,
		Components: p.Components.ToProto(),
	}
}

func NewParameterFromProto(pb_param *pb.Parameter) *Parameter {
	return &Parameter{
		Id:         pb_param.Id,
		Components: NewComponentsFromProto(pb_param.Components),
	}
}

func NewFloatParameter(name string) *Parameter {
	return &Parameter{
		Name: name,
		Components: Components{
			&Component{Name: "float", Value: 0.0},
		},
	}
}

func NewColorParameter(id int64) *Parameter {
	return &Parameter{
		Id: id,
		Components: Components{
			&Component{Name: "Red", Value: 0.0},
			&Component{Name: "Green", Value: 0.0},
			&Component{Name: "Blue", Value: 0.0},
		},
	}
}

type Parameters []*Parameter

func (p *Parameters) GetById(id int64) *Parameter {
	for _, param := range *p {
		if param.Id == id {
			return param
		}
	}
	return nil
}

func (p *Parameters) GetByName(name string) *Parameter {
	for _, param := range *p {
		if param.Name == name {
			return param
		}
	}
	return nil
}

func (p *Parameters) SetById(id int64, component string, value float64) {
	p.GetById(id).SetValueByName(component, value)
}

func (p *Parameters) ToProto() []*pb.Parameter {
	var pb_params []*pb.Parameter
	for _, param := range *p {
		pb_params = append(pb_params, param.ToProto())
	}
	return pb_params
}

func NewParametersFromProto(pb_params []*pb.Parameter) Parameters {
	var params Parameters
	for _, pb_param := range pb_params {
		params = append(params, NewParameterFromProto(pb_param))
	}
	return params
}
