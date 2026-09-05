package app

type IntegrationMethod string

const (
	ForwardEuler IntegrationMethod = "Forward Euler"
	SimplecticEuler = "Simplectic Euler"
)