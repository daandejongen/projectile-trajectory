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

func (trajectory Trajectory) PositionOnGroundHit() (float64, float64) {
	n := len(trajectory.Points)
	if trajectory.Points[n-1].Position.Y == 0 {
		return trajectory.Points[n-1].Position.X, trajectory.Points[n-1].Position.X
	}
	return trajectory.Points[n-1].Position.X, trajectory.Points[n-2].Position.X
}
