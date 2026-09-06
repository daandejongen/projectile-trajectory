package app

import "github.com/daandejongen/projectile-trajectory/linalg"

type generalThrustForceCalculator struct {
	force    linalg.Vector2d
	duration float64
}

func newThrustForceCalculator(thrust Thrust) thrustForceCalculator {
	return generalThrustForceCalculator{
		force:    linalg.NewVectorFromLengthAndAngle(thrust.Force, thrust.Angle),
		duration: thrust.Duration,
	}
}

func (calculator generalThrustForceCalculator) compute(velocity linalg.Vector2d, timePassed float64) linalg.Vector2d {
	if calculator.duration < timePassed {
		return linalg.Vector2d{X: 0, Y: 0}
	}
	return calculator.force
}
