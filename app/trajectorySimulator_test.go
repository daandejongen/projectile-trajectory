package app

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const floatingPointTolerance = 1e-5

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

func TestSimulationOfProjectileStraightInTheAirWithoutDrag(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(9.81).WithInitialAngle(0.5 * math.Pi).WithIntegrationMethod(SimplecticEuler)
	trajectory, err := simulator.Simulate()
	assert.NoError(t, err)
	assert.Less(t, math.Abs(trajectory.AirTime()-2), floatingPointTolerance)
}

func TestSimulationDoesNotExceedMaxTimeSteps(t *testing.T) {
	simulator := NewTrajectorySimulator().WithInitialSpeed(10000)
	simulator.maxTimeSteps = 3
	trajectory, _ := simulator.Simulate()
	assert.Equal(t, simulator.maxTimeSteps + 1, len(trajectory.Points))
}
