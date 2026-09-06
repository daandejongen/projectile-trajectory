package app

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/daandejongen/projectile-trajectory/linalg"
	"github.com/daandejongen/projectile-trajectory/trigonometry"
)

type TrajectorySimulator struct {
	projectile              projectile
	initialPosition         linalg.Vector2d
	initialAngle            trigonometry.Radians
	initialSpeed            float64
	dragType                DragType
	thrust                  Thrust
	timeStepInterval        float64
	integrationMethod       IntegrationMethod
	maxTimeSteps            int
	errorToleranceAtLanding float64
}

// dependency declarations, implemented elsewhere
type dragForceCalculator interface {
	compute(velocity linalg.Vector2d) linalg.Vector2d
}

type thrustForceCalculator interface {
	compute(velocity linalg.Vector2d, timePassed float64) linalg.Vector2d
}

// end dependencies

func NewTrajectorySimulator() *TrajectorySimulator {
	return &TrajectorySimulator{
		projectile:              Sphere{radius: 1, density: 1},
		initialPosition:         linalg.Vector2d{X: 0, Y: 0},
		initialAngle:            0.25 * math.Pi,
		initialSpeed:            1,
		dragType:                NoDrag,
		thrust:                  Thrust{Angle: 0, Force: 0, Duration: 0},
		timeStepInterval:        1e-5,
		integrationMethod:       ForwardEuler,
		maxTimeSteps:            1000000,
		errorToleranceAtLanding: 1e-5,
	}
}

func (simulator TrajectorySimulator) Simulate() (Trajectory, error) {
	initialVelocity := linalg.NewVectorFromLengthAndAngle(simulator.initialSpeed, simulator.initialAngle)
	if initialVelocity.Y <= 0 {
		return Trajectory{Points: []Point{{Position: simulator.initialPosition, Velocity: linalg.Vector2d{X: 0, Y: 0}}}, timeStepInterval: simulator.timeStepInterval}, nil
	}
	timeStepCount := 0
	projectileLanded := false
	points := make([]Point, simulator.maxTimeSteps+1)
	points[timeStepCount] = Point{
		Position: simulator.initialPosition,
		Velocity: linalg.NewVectorFromLengthAndAngle(simulator.initialSpeed, simulator.initialAngle),
	}

	dragForceCalculator := newDragForceCalculator(simulator.dragType, simulator.projectile)
	thrustForceCalculator := newThrustForceCalculator(simulator.thrust)

	for !projectileLanded {
		if timeStepCount == simulator.maxTimeSteps {
			return Trajectory{Points: points, timeStepInterval: simulator.timeStepInterval}, errors.New("Projectile did not land after max time step iterations was reached.")
		}

		currentVelocity := points[timeStepCount].Velocity
		dragForce := dragForceCalculator.compute(currentVelocity)
		thrustForce := thrustForceCalculator.compute(currentVelocity, simulator.timeStepInterval*float64(timeStepCount))
		acceleration := gravityAcceleration.
			Add(dragForce.ComputeScalarMultiplication(1 / simulator.projectile.Mass())).
			Add(thrustForce.ComputeScalarMultiplication(1 / simulator.projectile.Mass()))
		nextVelocity := currentVelocity.Add(acceleration.ComputeScalarMultiplication(simulator.timeStepInterval))

		var nextPosition linalg.Vector2d
		switch simulator.integrationMethod {
		case ForwardEuler:
			nextPosition = points[timeStepCount].Position.Add(currentVelocity.ComputeScalarMultiplication(simulator.timeStepInterval))
		case SimplecticEuler:
			nextPosition = points[timeStepCount].Position.Add(nextVelocity.ComputeScalarMultiplication(simulator.timeStepInterval))
		}

		timeStepCount++
		points[timeStepCount].Position = nextPosition
		points[timeStepCount].Velocity = nextVelocity
		projectileLanded = nextPosition.Y < simulator.errorToleranceAtLanding
	}

	return Trajectory{Points: points[:timeStepCount+1], timeStepInterval: simulator.timeStepInterval}, nil
}

func (simulator *TrajectorySimulator) WithInitialSpeed(speed float64) *TrajectorySimulator {
	simulator.initialSpeed = speed
	return simulator
}

func (simulator *TrajectorySimulator) WithInitialAngle(angle trigonometry.Radians) *TrajectorySimulator {
	simulator.initialAngle = angle
	return simulator
}

func (simulator *TrajectorySimulator) WithInitialPosition(position linalg.Vector2d) *TrajectorySimulator {
	simulator.initialPosition = position
	return simulator
}

func (simulator *TrajectorySimulator) WithDragType(dragType DragType) *TrajectorySimulator {
	simulator.dragType = dragType
	return simulator
}

func (simulator *TrajectorySimulator) WithIntegrationMethod(method IntegrationMethod) *TrajectorySimulator {
	simulator.integrationMethod = method
	return simulator
}

func (simulator *TrajectorySimulator) WithThrust(thrust Thrust) *TrajectorySimulator {
	simulator.thrust = thrust
	return simulator
}

func (simulator TrajectorySimulator) Print(writer io.Writer) {
	fmt.Fprintf(writer, "Simulator conditions\n")
	fmt.Fprintf(writer, "projectile:        %s\n", simulator.projectile.String())
	fmt.Fprintf(writer, "initial position:  %f\n", simulator.initialPosition)
	fmt.Fprintf(writer, "initial angle:     %f\n", simulator.initialAngle)
	fmt.Fprintf(writer, "initial speed:     %f\n", simulator.initialSpeed)
	fmt.Fprintf(writer, "thrust:            %s\n", simulator.thrust.string())
	fmt.Fprintf(writer, "timeStepInterval:  %f\n", simulator.timeStepInterval)
	fmt.Fprintf(writer, "drag type:         %s\n", simulator.dragType)
	fmt.Fprintf(writer, "integrationMethod: %s\n", simulator.integrationMethod)
	fmt.Fprintf(writer, "maxTimeSteps:      %d\n", simulator.maxTimeSteps)
	fmt.Fprintf(writer, "errTolerAtLanding: %f\n", simulator.errorToleranceAtLanding)
}
