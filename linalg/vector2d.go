package linalg

import (
	"math"

	"github.com/daandejongen/projectile-trajectory/trigonometry"
)

type Vector2d struct {
	X float64
	Y float64
}

func NewVectorFromLengthAndAngle(length float64, angle trigonometry.Radians) Vector2d {
	return Vector2d{
		X: math.Cos(float64(angle)) * length,
		Y: math.Sin(float64(angle)) * length,
	}
}

func (vector Vector2d) IsZero() bool {
	return vector.X == 0 && vector.Y == 0
}

func (vector Vector2d) Length() float64 {
	return math.Sqrt(vector.X*vector.X + vector.Y*vector.Y)
}

func (vector Vector2d) Unit() Vector2d {
	return vector.ComputeScalarMultiplication(1/vector.Length())
}

func (vector Vector2d) Angle() trigonometry.Radians {
	return trigonometry.Radians(math.Acos(vector.Unit().X))
}

func (vector Vector2d) Add(other Vector2d) Vector2d {
	return Vector2d{
		X: vector.X + other.X,
		Y: vector.Y + other.Y,
	}
}

func (vector Vector2d) ComputeScalarMultiplication(scalar float64) Vector2d {
	return Vector2d{
		X: scalar * vector.X,
		Y: scalar * vector.Y,
	}
}
