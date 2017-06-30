package parameter

import (
	"encoding/json"
	"testing"
)

func TestConvertBoolToJson(t *testing.T) {
	b := NewBoolParameter(true)

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
	i64 := NewInt64Parameter(0)

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
	enum := NewEnumParameter([]string{"one", "two", "three"}, "one")

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
		ParameterSpecs{
			{Name: "x", Parameter: NewFloat64Parameter(1.0)},
			{Name: "y", Parameter: NewFloat64Parameter(1.0)},
		},
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
		ParameterSpecs{
			{Name: "r", Parameter: NewFloat64Parameter(1.0)},
			{Name: "g", Parameter: NewFloat64Parameter(1.0)},
			{Name: "b", Parameter: NewFloat64Parameter(1.0)},
		},
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
		ParameterSpecs{
			{Name: "r", Parameter: NewFloat64Parameter(1.0)},
			{Name: "g", Parameter: NewFloat64Parameter(1.0)},
			{Name: "b", Parameter: NewFloat64Parameter(1.0)},
			{Name: "a", Parameter: NewFloat64Parameter(1.0)},
		},
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
	multi := NewMultiParameter(
		NewGroupParameter("position",
			ParameterSpecs{
				{Name: "x", Parameter: NewFloat64Parameter(1.0)},
				{Name: "y", Parameter: NewFloat64Parameter(1.0)},
			},
		),
		0,
		Parameters{},
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
	multi := NewGroupParameter(
		"custom",
		ParameterSpecs{
			{Name: "name", Parameter: &String("")},
			{Name: "points", Parameter: NewMultiParameter(
				NewGroupParameter("position",
					ParameterSpecs{
						{Name: "x", Parameter: NewFloat64Parameter(1.0)},
						{Name: "y", Parameter: NewFloat64Parameter(1.0)},
					},
				),
				0,
				Parameters{},
			)},
		},
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
