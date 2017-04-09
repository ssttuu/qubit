package params

import pb "github.com/stupschwartz/qubit/protos"

type Component struct {
	Label string `json:"label"`
	Value float64 `json:"value"`
}

func (c *Component) GetValue() float64 {
	return c.Value
}

func (c *Component) SetValue(value float64) {
	c.Value = value
}

func (c *Component) ToProto() *pb.Component {
	return &pb.Component{Label: c.Label, Value: c.Value}
}

type Parameter struct {
	Name       string `json:"name"`
	Components []Component `json:"components"`
}

func (p *Parameter) GetComponentByLabel(label string) *Component {
	for _, component := range p.Components {
		if component.Label == label {
			return &component
		}
	}

	return nil
}

func (p *Parameter) GetComponentValueByLabel(label string) float64 {
	return p.GetComponentByLabel(label).Value
}

func (p *Parameter) GetValue(index int) float64 {
	return p.Components[index].GetValue()
}

func (p *Parameter) SetValue(index int, value float64) {
	p.Components[index].SetValue(value)
}

func (p *Parameter) GetComponentByName(name string) *Component {
	for _, component := range p.Components {
		if component.Label == name {
			return &component
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
	cps := make([]*pb.Component, len(p.Components))

	for index, component := range p.Components {
		cps[index] = component.ToProto()
	}

	return &pb.Parameter{Name: p.Name, Components: cps}
}

func NewFloatParameter(name string) Parameter {
	p := Parameter{Name: name}
	floatComponent := Component{Label: "float", Value: 0.0}
	p.Components = append(p.Components, floatComponent)
	return p
}

func NewColorParameter(name string) Parameter {
	p := Parameter{Name: name}

	p.Components = append(p.Components, Component{Label: "Red", Value: 0.0})
	p.Components = append(p.Components, Component{Label: "Green", Value: 0.0})
	p.Components = append(p.Components, Component{Label: "Blue", Value: 0.0})

	return p
}

type Parameters []Parameter

func (p Parameters) GetByName(name string) *Parameter {
	for _, param := range p {
		if param.Name == name {
			return &param
		}
	}

	return nil
}

func (p Parameters) SetByName(name string, component string, value float64) {
	p.GetByName(name).SetValueByName(component, value)
}

func (p Parameters) ToProto() []*pb.Parameter {
	pps := make([]*pb.Parameter, len(p))

	for index, param := range p {
		pps[index] = param.ToProto()
	}

	return pps
}