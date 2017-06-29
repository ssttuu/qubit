package parameter

import (
	"encoding/json"
)

type ParameterType interface{}

type Parameter struct {
	TypeName string
	Type     ParameterType
}

type ParameterSpec struct {
	Name      string
	Parameter *Parameter
}

type ParameterSpecs []ParameterSpec
type Parameters []Parameter

func (p *Parameter) GetBool() Bool {
	return p.Type.(Bool)
}

func (p *Parameter) GetInt64() Int64 {
	return p.Type.(Int64)
}

func (p *Parameter) GetFloat64() Float64 {
	return p.Type.(Float64)
}

func (p *Parameter) GetString() String {
	return p.Type.(String)
}

func (p *Parameter) GetEnum() Enum {
	return p.Type.(Enum)
}

func (p *Parameter) GetGroup() Group {
	return p.Type.(Group)
}

func (p *Parameter) GetMulti() Multi {
	return p.Type.(Multi)
}

func (p *Parameter) UnmarshalJSON(b []byte) error {
	type unmarshalParam struct {
		TypeName  string
		TypeBytes json.RawMessage `json:"Type"`
	}
	var unParam unmarshalParam
	if err := json.Unmarshal(b, &unParam); err != nil {
		return err
	}
	p.TypeName = unParam.TypeName
	switch unParam.TypeName {
	case "bool":
		var paramType Bool
		if err := json.Unmarshal(unParam.TypeBytes, &paramType); err != nil {
			return err
		}
		p.Type = paramType
	case "int64":
		var paramType Int64
		if err := json.Unmarshal(unParam.TypeBytes, &paramType); err != nil {
			return err
		}
		p.Type = paramType
	case "float64":
		var paramType Float64
		if err := json.Unmarshal(unParam.TypeBytes, &paramType); err != nil {
			return err
		}
		p.Type = paramType
	case "string":
		var paramType String
		if err := json.Unmarshal(unParam.TypeBytes, &paramType); err != nil {
			return err
		}
		p.Type = paramType
	case "enum":
		var paramType Enum
		if err := json.Unmarshal(unParam.TypeBytes, &paramType); err != nil {
			return err
		}
		p.Type = paramType
	case "group":
		var paramType Group
		if err := json.Unmarshal(unParam.TypeBytes, &paramType); err != nil {
			return err
		}
		p.Type = paramType
	case "multi":
		var paramType Multi
		if err := json.Unmarshal(unParam.TypeBytes, &paramType); err != nil {
			return err
		}
		p.Type = paramType
	}
	return nil
}

///////////////////////
// Parameter Helpers //
///////////////////////

func NewPosition2DParameter() *Parameter {
	return NewGroupParameter(
		"position2d",
		ParameterSpecs{
			{Name: "x", Parameter: NewFloat64Parameter(0.0)},
			{Name: "y", Parameter: NewFloat64Parameter(0.0)},
		},
	)
}

func NewRGBParameter() *Parameter {
	return NewGroupParameter(
		"rgb",
		ParameterSpecs{
			{Name: "r", Parameter: NewFloat64Parameter(0.0)},
			{Name: "g", Parameter: NewFloat64Parameter(0.0)},
			{Name: "b", Parameter: NewFloat64Parameter(0.0)},
		},
	)
}

func NewRGBAParameter() *Parameter {
	return NewGroupParameter(
		"rgba",
		ParameterSpecs{
			{Name: "r", Parameter: NewFloat64Parameter(0.0)},
			{Name: "g", Parameter: NewFloat64Parameter(0.0)},
			{Name: "b", Parameter: NewFloat64Parameter(0.0)},
			{Name: "a", Parameter: NewFloat64Parameter(0.0)},
		},
	)
}
