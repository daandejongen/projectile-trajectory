package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/daandejongen/projectile-trajectory/app"
	"github.com/daandejongen/projectile-trajectory/trigonometry"
	"github.com/spf13/cobra"
)

const outputDir string = "output"
const plotDir string = "plots"

func NewSimulateCommand() *cobra.Command {
	var dragTypeInput string
	var addThrust bool
	var makePlot bool
	command := &cobra.Command{
		Use:   "sim [angle-in-degrees] [initial-speed-in-meters-per-second]",
		Short: "Simulate the trajectory of a projectile.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			angleInDegrees, angleParsingErr := strconv.ParseFloat(args[0], 64)
			speed, speedParsingError := strconv.ParseFloat(args[1], 64)
			exitIfError(angleParsingErr)
			exitIfError(speedParsingError)

			drag, dragErr := translateDrag(dragTypeInput)
			exitIfError(dragErr)

			simulator := app.NewTrajectorySimulator().
				WithInitialAngle(trigonometry.RadiansFromDegrees(angleInDegrees)).
				WithInitialSpeed(speed).
				WithDragType(drag)
			if addThrust {
				simulator.WithThrust(app.Thrust{Angle: trigonometry.RadiansFromDegrees(45), Force: 100, Duration: 5})
			}

			simulator.Print(os.Stdout)
			trajectory, err := simulator.Simulate()
			exitIfError(err)

			landingPoint := trajectory.LandingPoint()
			fmt.Fprintf(os.Stdout, "Projectile landed at x = %f\n", landingPoint)
			fmt.Fprintf(os.Stdout, "Airtime: %f seconds\n", trajectory.AirTime())

			if makePlot {
				plotErr := app.Plot(trajectory, fmt.Sprintf("./%s/%s/simulated-projectile-trajectory-%s", outputDir, plotDir, time.Now().Format("20060102-150405")))
				exitIfError(plotErr)
			}
		},
	}
	command.Flags().StringVarP(&dragTypeInput, "drag", "D", "", "How drag is modeled proportional to the projectile's velocity: 'l' for linear or 'q' for quadratic. Leave empty for no drag.")
	command.Flags().BoolVarP(&addThrust, "thrust", "T", false, "Add a thrust of 100N in a constant 45 degree angle for the first 5 seconds.")
	command.Flags().BoolVarP(&makePlot, "plot", "P", false, fmt.Sprintf("Bool to indicate whether the trajectory should be plotted and saved to ./%s/%s/", outputDir, plotDir))
	return command
}

func translateDrag(input string) (app.DragType, error) {
	switch input {
	case "l":
		return app.Linear, nil
	case "q":
		return app.Quadratic, nil
	case "":
		return app.NoDrag, nil
	default:
		return app.NoDrag, errors.New("dragtype not supported")
	}
}
