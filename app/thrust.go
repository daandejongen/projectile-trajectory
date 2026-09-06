package app

import (
	"fmt"

	"github.com/daandejongen/projectile-trajectory/trigonometry"
)

type Thrust struct {
	Angle    trigonometry.Radians
	Force    float64
	Duration float64
}

func (thrust Thrust) string() string {
	return fmt.Sprintf("Angle=%f rad, force=%fN, duration=%fs", thrust.Angle, thrust.Force, thrust.Duration) 
}
