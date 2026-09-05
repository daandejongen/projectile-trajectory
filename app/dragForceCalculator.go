package app

import "github.com/daandejongen/projectile-trajectory/linalg"

type dragForceCalculator interface {
	compute(velocity linalg.Vector2d) linalg.Vector2d
}

func newDragForceCalculator(dragType DragType, projectile projectile) dragForceCalculator {
	switch dragType {
	case NoDrag:
		return noDragForceCalculator{}
	case Quadratic:
		return newQuadraticDragForceCalculator(projectile)
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

func newQuadraticDragForceCalculator(projectile projectile) quadraticDragForceCalculator {
	return quadraticDragForceCalculator{
		constant: -0.5 * constants.airDensity * projectile.FrontalArea() * projectile.DragCoefficient(),
	}
}

func (calculator quadraticDragForceCalculator) compute(velocity linalg.Vector2d) linalg.Vector2d {
	return velocity.ComputeScalarMultiplication(calculator.constant * velocity.Length())
}
