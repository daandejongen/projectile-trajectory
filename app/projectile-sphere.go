package app

import (
	"fmt"
	"math"
)

type Sphere struct {
	radiusInCm           float64
	densityInGramsPerCm3 float64
}

func (sphere Sphere) FrontalAreaInCubicMeters() float64 {
	return math.Pi * sphere.radiusInCm * sphere.radiusInCm / 10000
}

func (sphere Sphere) MassInGrams() float64 {
	return sphere.densityInGramsPerCm3 * (4 / 3 * math.Pi * sphere.radiusInCm * sphere.radiusInCm * sphere.radiusInCm)
}

func (sphere Sphere) DragCoefficient() float64 {
	return 0.47
}

func (sphere Sphere) string() string {
	return fmt.Sprintf("Shape=sphere, radius=%fcm, density=%fg/cm^3", sphere.radiusInCm, sphere.densityInGramsPerCm3)
}
