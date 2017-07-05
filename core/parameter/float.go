package parameter

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
		// TODO
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

func NewFloat64Parameter(defaultValue float64) *Parameter {
	return &Parameter{
		TypeName: "float64",
		Type: Float64{
			Default: defaultValue,
		},
	}
}
