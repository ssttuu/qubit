package params

type Component struct {
	Label string `json:"label"`
	Value float64 `json:"value"`
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

func (p *Parameter) GetValue() float64 {
	return p.Components[0].Value
}

func NewFloatParameter(name string) Parameter {
	p := Parameter{Name: name}

	floatComponent := Component{Label: "float", Value: 0.0}

	p.Components = append(p.Components, floatComponent)

	return p
}

func NewColorParameter(name string) Parameter {
	p := Parameter{Name: name}

	red := Component{Label: "Red", Value: 0.0}
	green := Component{Label: "Green", Value: 0.0}
	blue := Component{Label: "Blue", Value: 0.0}

	p.Components = append(p.Components, red)
	p.Components = append(p.Components, green)
	p.Components = append(p.Components, blue)

	return p
}

type Parameters []Parameter

func GetByName(p Parameters, name string) *Parameter {
	for _, param := range p {
		if param.Name == name {
			return &param
		}
	}

	return nil
}
