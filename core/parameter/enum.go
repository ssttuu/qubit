package parameter

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
		// TODO
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

func NewEnumParameter(options []string, defaultValue string) *Parameter {
	return &Parameter{
		TypeName: "enum",
		Type: Enum{
			Options: options,
			Default: defaultValue,
		},
	}
}
