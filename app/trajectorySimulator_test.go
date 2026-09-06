package app

import (
	"math"
	"testing"

	"github.com/daandejongen/projectile-trajectory/linalg"
	"github.com/stretchr/testify/assert"
)

func TestSimulationDoesNotExceedMaxTimeSteps(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(10000)
	simulator.maxTimeSteps = 3
	trajectory, _ := simulator.Simulate()
	assert.Equal(t, simulator.maxTimeSteps+1, len(trajectory.Points))
}

func TestSimulation_WithZeroVelocity_YieldsTrajectory_WithZeroAirTime(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(0)
	trajectory, err := simulator.Simulate()
	assert.NoError(t, err)
	assert.Equal(t, float64(0), trajectory.AirTime())
}

func TestSimulation_WithZeroAngle_YieldsTrajectory_WithZeroAirTime(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialAngle(0)
	trajectory, err := simulator.Simulate()
	assert.NoError(t, err)
	assert.Equal(t, float64(0), trajectory.AirTime())
}

func TestSimulation_YieldsTrajectoryWhereFirstPoint_EqualsTheInitialCondition(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialAngle(0.5 * math.Pi).WithInitialPosition(linalg.Vector2d{X: 2, Y: 1}).WithInitialSpeed(50)
	trajectory, _ := simulator.Simulate()
	firstPoint := trajectory.Points[0]
	assert.Equal(t, float64(2), firstPoint.Position.X)
	assert.Equal(t, float64(1), firstPoint.Position.Y)
	assert.Less(t, firstPoint.Velocity.X, testToleranceForRoundingErrors)
	assert.Less(t, firstPoint.Velocity.Y-50, testToleranceForRoundingErrors)
}

func TestTrajectoryWithoutDrag_WithInitalSpeedEqualToGravity_Spends2SecondsInTheAir(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(9.81).WithInitialAngle(0.5 * math.Pi).WithIntegrationMethod(SimplecticEulerIntegration)
	trajectory, err := simulator.Simulate()
	assert.NoError(t, err)
	assert.Less(t, math.Abs(trajectory.AirTime()-2), testToleranceForRoundingErrors)
}

func TestTrajectoryWithoutDrag_HasMoreAirTimeThan_TrajectoryWithDrag(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(10)
	trajectoryWithoutDrag, _ := simulator.Simulate()

	simulator.WithDragType(QuadraticDrag)
	trajectoryWithDrag, _ := simulator.Simulate()

	assert.Greater(t, trajectoryWithoutDrag, trajectoryWithDrag)
}

func TestTrajectoryWithoutThrust_HasLessAirTimeThan_TrajectoryWithThrust(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(50).WithThrust(Thrust{Duration: 0})
	trajectoryWithoutThrust, _ := simulator.Simulate()

	simulator.WithThrust(Thrust{Angle: 45, Force: 50, Duration: 5})
	trajectoryWithThrust, _ := simulator.Simulate()

	assert.Less(t, trajectoryWithoutThrust.AirTime(), trajectoryWithThrust.AirTime())
}

func TestTrajectoryWithThrust(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(50).WithThrust(Thrust{Duration: 1, Force: 30, Angle: 45})
	trajectory, _ := simulator.Simulate()

	assert.Equal(t, trajectory.timeStepInterval, 1)
}
