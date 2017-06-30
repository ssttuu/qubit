package parameter

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
		// TODO
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

func NewBoolParameter(defaultValue bool) *Parameter {
	return &Parameter{
		TypeName: "bool",
		Type: Bool{
			Default: defaultValue,
		},
	}
}
