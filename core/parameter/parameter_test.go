package parameter

import (
	"encoding/json"
	"testing"
)

func TestConvertBoolToJson(t *testing.T) {
	b := NewBool(true)

	data, err := json.Marshal(b)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)

}

func TestConvertInt64ToJson(t *testing.T) {
	i64 := NewInt64(0)

	data, err := json.Marshal(i64)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)

}

func TestConvertEnumToJson(t *testing.T) {
	enum := NewEnum([]string{"one", "two", "three"}, "one")

	data, err := json.Marshal(enum)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)
}

func TestConvertPositionToJson(t *testing.T) {
	grp := NewGroupParameter("position",
		ParameterMap{
			"x": NewFloat64(1.0),
			"y": NewFloat64(1.0),
		},
		[]string{"x", "y"},
	)

	data, err := json.Marshal(grp)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)
}

func TestConvertColorToJson(t *testing.T) {
	pos := NewGroupParameter("color",
		ParameterMap{
			"r": NewFloat64(1.0),
			"g": NewFloat64(1.0),
			"b": NewFloat64(1.0),
		},
		[]string{"r", "g", "b"},
	)
	data, err := json.Marshal(pos)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)
}

func TestConvertColorAlphaToJson(t *testing.T) {
	pos := NewGroupParameter("colorAlpha",
		ParameterMap{
			"r": NewFloat64(1.0),
			"g": NewFloat64(1.0),
			"b": NewFloat64(1.0),
			"a": NewFloat64(1.0),
		},
		[]string{"r", "g", "b", "a"},
	)

	data, err := json.Marshal(pos)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)
}

func TestConvertMultiPointToJson(t *testing.T) {
	multi := NewMulti(
		NewGroupParameter("position",
			ParameterMap{
				"x": NewFloat64(1.0),
				"y": NewFloat64(1.0),
			},
			[]string{"x", "y"},
		),
		0,
		ParameterArray{},
	)

	data, err := json.Marshal(multi)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)
}

func TestConvertComplexGroupToJson(t *testing.T) {
	multi := NewGroupParameter("custom",
		ParameterMap{
			"name": NewString(""),
			"points": NewMulti(
				NewGroupParameter("position",
					ParameterMap{
						"x": NewFloat64(1.0),
						"y": NewFloat64(1.0),
					},
					[]string{"x", "y"},
				),
				0,
				ParameterArray{},
			),
		},
		[]string{"name", "points"},
	)

	data, err := json.Marshal(multi)
	if err != nil {
		t.Fail()
	}

	t.Log(string(data))

	var i64FromJson Parameter
	err = json.Unmarshal(data, &i64FromJson)
	if err != nil {
		t.Fail()
	}

	t.Log(i64FromJson)
}

func TestConvertFromJson(t *testing.T) {

}
