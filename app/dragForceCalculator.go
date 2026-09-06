package app

import "github.com/daandejongen/projectile-trajectory/linalg"

func newDragForceCalculator(dragType DragType, projectile projectile) dragForceCalculator {
	switch dragType {
	case NoDrag:
		return noDragForceCalculator{}
	case QuadraticDrag:
		return quadraticDragForceCalculator{constant: -0.5 * airDensity * projectile.FrontalAreaInCubicMeters() * projectile.DragCoefficient()}
	default:
		return noDragForceCalculator{}
	}
}

type noDragForceCalculator struct{}

func (calculator noDragForceCalculator) compute(velocity linalg.Vector2d) linalg.Vector2d {
	return linalg.Vector2d{X: 0, Y: 0}
}

type quadraticDragForceCalculator struct {
	constant float64
}

func (calculator quadraticDragForceCalculator) compute(velocity linalg.Vector2d) linalg.Vector2d {
	return velocity.ComputeScalarMultiplication(calculator.constant * velocity.Length())
}
