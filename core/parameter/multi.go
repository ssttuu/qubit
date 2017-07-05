package parameter

type Multi struct {
	ParameterType
	Template   *Parameter
	ArrayCount int32
	Array      Parameters
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

func NewMultiParameter(template *Parameter, arrayCount int32, array Parameters) *Parameter {
	return &Parameter{
		TypeName: "multi",
		Type: Multi{
			Template:   template,
			ArrayCount: arrayCount,
			Array:      array,
		},
	}
}
