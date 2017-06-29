package parameter

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
		// TODO
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

func NewInt64Parameter(defaultValue int64) *Parameter {
	return &Parameter{
		TypeName: "int64",
		Type: Int64{
			Default: defaultValue,
		},
	}
}
