package linalg

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const floatingPointTolerance = 1e-10

func TestNewUnitVectorWithAngleZero(t *testing.T) {
	vector := NewVectorFromLengthAndAngle(1, 0)
	assert.Less(t, vector.X - 1, floatingPointTolerance)
	assert.Less(t, vector.Y, floatingPointTolerance)
}

func TestNewUnitVectorWithAngleHalfPi(t *testing.T) {
	vector := NewVectorFromLengthAndAngle(1, 0.5*math.Pi)
	assert.Less(t, vector.X, floatingPointTolerance)
	assert.Less(t, vector.Y - 1, floatingPointTolerance)
}

func TestNewUnitVectorWithAnglePi(t *testing.T) {
	vector := NewVectorFromLengthAndAngle(1, math.Pi)
	assert.Less(t, vector.X + 1, floatingPointTolerance)
	assert.Less(t, vector.Y, floatingPointTolerance)
}

func TestNewUnitVectorWithAngleMinusHalfPi(t *testing.T) {
	vector := NewVectorFromLengthAndAngle(1, -0.5*math.Pi)
	assert.Less(t, vector.X, floatingPointTolerance)
	assert.Less(t, vector.Y + 1, floatingPointTolerance)
}