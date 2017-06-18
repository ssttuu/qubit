package parameter

import (
	"encoding/json"
)

/////////////////////
// Parameter Types //
/////////////////////

type ParameterType interface {
}

type BoolKeyFrame struct {
	Time  float64
	Value bool
	// TODO: will need slope and magnitude values for more fine grained control
	// TODO: interpolation function
}

type Bool struct {
	ParameterType
	Default bool
	// Keyframes are sorted by time
	KeyFrames  []BoolKeyFrame
	Expression string
}

func (b *Bool) GetValue(time float64) bool {
	if b.Expression != "" {

	}

	if len(b.KeyFrames) != 0 {
		var lowKeyFrame BoolKeyFrame
		//var highKeyFrame BoolKeyFrame
		for _, kf := range b.KeyFrames {

			if kf.Time > time {
				//highKeyFrame = kf
				break
			}

			lowKeyFrame = kf
		}

		// TODO: interpolate between keyframes to get value
		return lowKeyFrame.Value
	}

	return b.Default
}

func NewBoolParameter(Default bool) Parameter {
	return Parameter{
		TypeName: "bool",
		Type: Bool{
			Default: Default,
		},
	}
}

type Int64KeyFrame struct {
	Time  float64
	Value int64
	// TODO: will need slope and magnitude values for more fine grained control
	// TODO: interpolation function
}

type Int64 struct {
	ParameterType
	Default int64
	// Keyframes are sorted by time
	KeyFrames  []Int64KeyFrame
	Expression string
}

func (i *Int64) GetValue(time float64) int64 {
	if i.Expression != "" {

	}

	if len(i.KeyFrames) != 0 {
		var lowKeyFrame Int64KeyFrame
		//var highKeyFrame Int64KeyFrame
		for _, kf := range i.KeyFrames {

			if kf.Time > time {
				//highKeyFrame = kf
				break
			}

			lowKeyFrame = kf
		}

		// TODO: interpolate between keyframes to get value
		return lowKeyFrame.Value
	}

	return i.Default
}

func NewInt64Parameter(Default int64) Parameter {
	return Parameter{
		TypeName: "int64",
		Type: Int64{
			Default: Default,
		},
	}
}

type Float64KeyFrame struct {
	Time  float64
	Value float64
	// TODO: will need slope and magnitude values for more fine grained control
	// TODO: interpolation function
}

type Float64 struct {
	ParameterType
	Default float64
	// Keyframes are sorted by time
	KeyFrames  []Float64KeyFrame
	Expression string
}

func (f *Float64) GetValue(time float64) float64 {
	if f.Expression != "" {

	}

	if len(f.KeyFrames) != 0 {
		var lowKeyFrame Float64KeyFrame
		//var highKeyFrame Float64KeyFrame
		for _, kf := range f.KeyFrames {

			if kf.Time > time {
				//highKeyFrame = kf
				break
			}

			lowKeyFrame = kf
		}

		// TODO: interpolate between keyframes to get value
		return lowKeyFrame.Value
	}

	return f.Default
}

func NewFloat64Parameter(Default float64) Parameter {
	return Parameter{
		TypeName: "float64",
		Type: Float64{
			Default: Default,
		},
	}
}

type StringKeyFrame struct {
	Time  float64
	Value string
	// TODO: will need slope and magnitude values for more fine grained control
	// TODO: interpolation function
}

type String struct {
	ParameterType
	Default string
	// Keyframes are sorted by time
	KeyFrames  []StringKeyFrame
	Expression string
}

func (s *String) GetValue(time float64) string {
	if s.Expression != "" {

	}

	if len(s.KeyFrames) != 0 {
		var lowKeyFrame StringKeyFrame
		//var highKeyFrame StringKeyFrame
		for _, kf := range s.KeyFrames {

			if kf.Time > time {
				//highKeyFrame = kf
				break
			}

			lowKeyFrame = kf
		}

		// TODO: interpolate between keyframes to get value
		return lowKeyFrame.Value
	}

	return s.Default
}

func NewStringParameter(Default string) Parameter {
	return Parameter{
		TypeName: "string",
		Type: String{
			Default: Default,
		},
	}
}

type EnumKeyFrame struct {
	Time  float64
	Value string
	// TODO: will need slope and magnitude values for more fine grained control
	// TODO: interpolation function
}

type Enum struct {
	ParameterType
	Options []string
	Default string
	// Keyframes are sorted by time
	KeyFrames  []EnumKeyFrame
	Expression string
}

func (e *Enum) GetValue(time float64) string {
	if e.Expression != "" {

	}

	if len(e.KeyFrames) != 0 {
		var lowKeyFrame EnumKeyFrame
		//var highKeyFrame EnumKeyFrame
		for _, kf := range e.KeyFrames {

			if kf.Time > time {
				//highKeyFrame = kf
				break
			}

			lowKeyFrame = kf
		}

		// TODO: interpolate between keyframes to get value
		return lowKeyFrame.Value
	}

	return e.Default
}

func NewEnumParameter(Options []string, Default string) Parameter {
	return Parameter{
		TypeName: "float64",
		Type: Enum{
			Options: Options,
			Default: Default,
		},
	}
}

// A Multi Parameter represents a dynamically sized array of parameters.
type Multi struct {
	ParameterType
	Template   Parameter
	ArrayCount int32
	Array      ParameterArray
}

func (m *Multi) GetParameterByIndex(index int32) Parameter {
	return m.Array[index]
}

func NewMultiParameter(Template Parameter, ArrayCount int32, Array ParameterArray) Parameter {
	return Parameter{
		TypeName: "multi",
		Type: Multi{
			Template:   Template,
			ArrayCount: ArrayCount,
			Array:      Array,
		},
	}
}

// A Tuple Parameter represents a tuple of parameters
type Group struct {
	ParameterType
	GroupType string
	Children  ParameterMap
	Order     []string
}

func (g *Group) GetParameter(name string) Parameter {
	return g.Children[name]
}

func NewGroupParameter(GroupType string, Children ParameterMap, Order []string) Parameter {
	return Parameter{
		TypeName: "group",
		Type: Group{
			GroupType: GroupType,
			Children:  Children,
			Order:     Order,
		},
	}
}

///////////////
// Parameter //
///////////////

type Parameter struct {
	TypeName string
	Type     ParameterType
}

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

type ParameterMap map[string]Parameter
type ParameterArray []Parameter

///////////////////////
// Parameter Helpers //
///////////////////////

func NewRGBParameter() Parameter {
	return NewGroupParameter(
		"rgb",
		ParameterMap{
			"r": NewFloat64Parameter(0.0),
			"g": NewFloat64Parameter(0.0),
			"b": NewFloat64Parameter(0.0),
		},
		[]string{"r", "g", "b"},
	)
}

func NewRGBAParameter() Parameter {
	return NewGroupParameter(
		"rgba",
		ParameterMap{
			"r": NewFloat64Parameter(0.0),
			"g": NewFloat64Parameter(0.0),
			"b": NewFloat64Parameter(0.0),
			"a": NewFloat64Parameter(0.0),
		},
		[]string{"r", "g", "b", "a"},
	)
}
