package app

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicSphereHasCorrectFrontalAreaInCubicMeters(t *testing.T) {
	sphere := Sphere{radiusInCm: 10, densityInGramsPerCm3: 10}
	expected := math.Pi * 100 / 10000
	assert.Less(t, sphere.FrontalAreaInCubicMeters()-expected, testToleranceForRoundingErrors)
}

func TestBasicSphereHasCorrectMass(t *testing.T) {
	sphere := Sphere{radiusInCm: 3, densityInGramsPerCm3: 7.8}
	expected := 4 / 3  * math.Pi * 27 * 7.8
	mass := sphere.MassInGrams()
	assert.Less(t, mass-expected, testToleranceForRoundingErrors)
}
