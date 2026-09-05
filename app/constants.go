package app

import "github.com/daandejongen/projectile-trajectory/linalg"

var constants = struct {
	gravityAcceleration linalg.Vector2d
	airDensity          float64
}{
	gravityAcceleration: linalg.Vector2d{X: 0, Y: -9.81},
	airDensity:          1.2,
}
