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

func NewBoolParameter(defaultValue bool) Parameter {
	return Parameter{
		TypeName: "bool",
		Type: Bool{
			Default: defaultValue,
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

func NewInt64Parameter(defaultValue int64) Parameter {
	return Parameter{
		TypeName: "int64",
		Type: Int64{
			Default: defaultValue,
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

func NewFloat64Parameter(defaultValue float64) Parameter {
	return Parameter{
		TypeName: "float64",
		Type: Float64{
			Default: defaultValue,
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

func NewStringParameter(defaultValue string) Parameter {
	return Parameter{
		TypeName: "string",
		Type: String{
			Default: defaultValue,
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

func NewEnumParameter(options []string, defaultValue string) Parameter {
	return Parameter{
		TypeName: "enum",
		Type: Enum{
			Options: options,
			Default: defaultValue,
		},
	}
}

type Multi struct {
	ParameterType
	Template   Parameter
	ArrayCount int32
	Array      ParameterArray
}

func (m *Multi) Iterator() <-chan Parameter {
	ch := make(chan Parameter)
	go func() {
		for _, p := range m.Array {
			ch <- p
		}
		close(ch)
	}()
	return ch
}

func (m *Multi) GetParameterByIndex(index int32) Parameter {
	return m.Array[index]
}

func NewMultiParameter(template Parameter, arrayCount int32, array ParameterArray) Parameter {
	return Parameter{
		TypeName: "multi",
		Type: Multi{
			Template:   template,
			ArrayCount: arrayCount,
			Array:      array,
		},
	}
}

type Group struct {
	ParameterType
	GroupType string
	Children  ParameterMap
	Order     []string
}

func (g *Group) GetParameter(name string) *Parameter {
	p := g.Children[name]
	return &p
}

func (g *Group) GetBool(name string) Bool {
	return g.GetParameter(name).GetBool()
}

func (g *Group) GetInt64(name string) Int64 {
	return g.GetParameter(name).GetInt64()
}

func (g *Group) GetFloat64(name string) Float64 {
	return g.GetParameter(name).GetFloat64()
}

func (g *Group) GetString(name string) String {
	return g.GetParameter(name).GetString()
}

func (g *Group) GetEnum(name string) Enum {
	return g.GetParameter(name).GetEnum()
}

func (g *Group) GetGroup(name string) Group {
	return g.GetParameter(name).GetGroup()
}

func (g *Group) GetMulti(name string) Multi {
	return g.GetParameter(name).GetMulti()
}

func NewGroupParameter(groupType string, children ParameterMap, order []string) Parameter {
	return Parameter{
		TypeName: "group",
		Type: Group{
			GroupType: groupType,
			Children:  children,
			Order:     order,
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

func NewPosition2DParameter() Parameter {
	return NewGroupParameter(
		"position2d",
		ParameterMap{
			"x": NewFloat64Parameter(0.0),
			"y": NewFloat64Parameter(0.0),
		},
		[]string{"x", "y"},
	)
}

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
