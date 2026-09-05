package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAirTimeIsComputedCorrectly(t *testing.T) {
	trajectory := Trajectory{
		Points:           make([]Point, 10),
		timeStepInterval: 1,
	}

	assert.Equal(t, trajectory.AirTime(), float64(9))
}
