package app

import (
	"fmt"
	"io"
)

func WriteCsv(writer io.Writer, trajectory Trajectory) {
	writer.Write([]byte("t,x,y,velocityX,velocityY"))
	writer.Write([]byte("\n"))
	for i, point := range trajectory.Points {
		fmt.Fprintf(writer, "%f,%f,%f,%f,%f", float64(i)*trajectory.timeStepInterval, point.Position.X, point.Position.Y, point.Velocity.X, point.Velocity.Y)
		writer.Write([]byte("\n"))
	}
}
