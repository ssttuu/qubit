package operator

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stupschwartz/qubit/core/image"
	"github.com/gonum/matrix/mat64"
)

func TestConstantOperation(t *testing.T) {
	var inputs []image.Plane
	var width, height int = 256, 512

	constantParameters := GetParameters("Constant")
	constantParameters[0].Components[2].Value = 1.0

	imageOutput := ConstantOperation(inputs, constantParameters, width, height)

	assert.Equal(t, "Color", imageOutput.Label)
	assert.Equal(t, 3, len(imageOutput.Components))

	redComponent := imageOutput.Components[0]
	greenComponent := imageOutput.Components[1]
	blueComponent := imageOutput.Components[2]

	rows, columns := redComponent.Dims()
	assert.Equal(t, width, columns)
	assert.Equal(t, height, rows)

	assert.Equal(t, 0.0, mat64.Min(redComponent))
	assert.Equal(t, 0.0, mat64.Max(redComponent))

	assert.Equal(t, 0.0, mat64.Min(greenComponent))
	assert.Equal(t, 0.0, mat64.Max(greenComponent))

	assert.Equal(t, 1.0, mat64.Min(blueComponent))
	assert.Equal(t, 1.0, mat64.Max(blueComponent))
}
