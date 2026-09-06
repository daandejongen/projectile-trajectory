package app

type IntegrationMethod string

const (
	NoIntegrationMethod        IntegrationMethod = "No Integration Method"
	ForwardEulerIntegration                      = "Forward Euler"
	SimplecticEulerIntegration                   = "Simplectic Euler"
)
