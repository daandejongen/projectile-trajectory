package trigonometry

import "math"

type Radians float64

func RadiansFromDegrees(degrees float64) Radians {
	return Radians(degrees * math.Pi / 180)
}
