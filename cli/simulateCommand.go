package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/daandejongen/projectile-trajectory/app"
	"github.com/daandejongen/projectile-trajectory/trigonometry"
	"github.com/spf13/cobra"
)

func NewSimulateCommand() *cobra.Command {
	var drag string
	var thrust bool
	var plotPath string
	command := &cobra.Command{
		Use:   "sim [angle-in-degrees] [initial-speed-in-meters-per-second]",
		Short: "Simulate the trajectory of a projectile.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			angleInDegrees, angleParsingErr := strconv.ParseFloat(args[0], 64)
			speed, speedParsingError := strconv.ParseFloat(args[1], 64)
			exitIfError(angleParsingErr)
			exitIfError(speedParsingError)

			drag, dragErr := translateDrag(drag)
			exitIfError(dragErr)

			simulator := app.NewTrajectorySimulator().
				WithInitialAngle(trigonometry.RadiansFromDegrees(angleInDegrees)).
				WithInitialSpeed(speed).
				WithDragType(drag)

			if thrust {
				simulator.WithThrust(app.Thrust{
					Angle:    trigonometry.RadiansFromDegrees(45),
					Force:    100,
					Duration: 5,
				})
			}

			simulator.Print(os.Stdout)
			trajectory, err := simulator.Simulate()

			exitIfError(err)
			landingPoint1, landingPoint2 := trajectory.PositionOnGroundHit()
			if landingPoint1 == landingPoint2 {
				fmt.Fprintf(os.Stdout, "Projectile landed at x = %f\n", landingPoint1)
			} else {
				fmt.Fprintf(os.Stdout, "Projectile landed between x = %f and x = %f\n", landingPoint1, landingPoint2)
			}
			fmt.Fprintf(os.Stdout, "Airtime: %f seconds\n", trajectory.AirTime())

			if plotPath != "" {
				plotErr := app.Plot(trajectory, plotPath)
				exitIfError(plotErr)
			}
		},
	}
	command.Flags().StringVar(&drag, "drag", "", "How drag is modeled proportional to the projectile's velocity: linear or quadratic. Leave empty for no drag.")
	command.Flags().BoolVar(&thrust, "thrust", false, "Add a thrust of 100N in a constant 45 degree angle for the first 5 seconds.")
	command.Flags().StringVar(&plotPath, "plot", "", "Path were plot should be stored. Omit if you don't want to plot.")
	return command
}

func translateDrag(input string) (app.DragType, error) {
	switch input {
	case "linear":
		return app.Linear, nil
	case "quadratic":
		return app.Quadratic, nil
	default:
		return app.NoDrag, errors.New("dragtype not supported")
	}
}
