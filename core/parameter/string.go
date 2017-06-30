package parameter

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
		// TODO
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

func NewStringParameter(defaultValue string) *Parameter {
	return &Parameter{
		TypeName: "string",
		Type: String{
			Default: defaultValue,
		},
	}
}
