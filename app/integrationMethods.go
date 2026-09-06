package app

type IntegrationMethod string

const (
	ForwardEulerIntegration    IntegrationMethod = "Forward Euler"
	SimplecticEulerIntegration                   = "Simplectic Euler"
)
