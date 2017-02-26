package operator

import (
	"encoding/json"
)

type Operator struct {
	Name   string                 `json:"name"`
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params"`
	Inputs []string               `json:"inputs"`
}

type Operation func(op Operator) string

var Operators = make(map[string]Operation)

func RegisterOperation(opType string, operation Operation) {
	Operators[opType] = operation
}

func GetOperation(opType string) Operation {
	return Operators[opType]
}

func GetOperatorFromJson(jsonBytes []byte) *Operator {
	op := new(Operator)
	if err := json.Unmarshal(jsonBytes, op); err != nil {

	}

	return op
}
