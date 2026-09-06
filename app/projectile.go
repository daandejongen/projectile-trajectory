package app

type projectile interface {
	FrontalAreaInCubicMeters() float64
	MassInGrams() float64
	DragCoefficient() float64
	string() string
}
