package app

import "github.com/daandejongen/projectile-trajectory/linalg"

type Trajectory struct {
	Points           []Point
	timeStepInterval float64
}

type Point struct {
	Position linalg.Vector2d
	Velocity linalg.Vector2d
}

func (trajectory Trajectory) AirTime() float64 {
	return float64(len(trajectory.Points)-1) * trajectory.timeStepInterval
}

func (trajectory Trajectory) LandingPoint() float64 {
	return trajectory.Points[len(trajectory.Points)-1].Position.X
}
