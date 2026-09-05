package app

import (
	"fmt"
	"math"
)

type Sphere struct {
	radius  float64
	density float64
}

func (sphere Sphere) FrontalArea() float64 {
	return math.Pi * sphere.radius * sphere.radius
}

func (sphere Sphere) Mass() float64 {
	return 4 / 3 * math.Pi * sphere.radius * sphere.radius * sphere.radius
}

func (sphere Sphere) DragCoefficient() float64 {
	return 0.47
}

func (sphere Sphere) String() string {
	return fmt.Sprintf("Shape=sphere, radius=%fcm, density=%fg/cm^3", sphere.radius, sphere.density) 
}