package app

type projectile interface {
	FrontalArea() float64
	Mass() float64
	DragCoefficient() float64
	String() string
}

