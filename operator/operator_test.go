package operator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOperator(t *testing.T) {
}

func TestGetOperatorFromJson(t *testing.T) {
	request_bytes := []byte(`{"type": "Read", "name": "read1", "id": "ae62f1c3"}`)

	op := GetOperatorFromJson(request_bytes)

	assert.Equal(t, op.Name, "read1")
}
