package app

import "github.com/daandejongen/projectile-trajectory/linalg"

const testToleranceForRoundingErrors = 1e-5
const airDensity float64 = 1.2

var gravityAcceleration linalg.Vector2d = linalg.Vector2d{X: 0, Y: -9.81}
