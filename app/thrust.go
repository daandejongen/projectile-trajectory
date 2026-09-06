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

func (thrust Thrust) String() string {
	return fmt.Sprintf("angle=%frad, force=%fN, duration=%fs", thrust.Angle, thrust.Force, thrust.Duration)
}
